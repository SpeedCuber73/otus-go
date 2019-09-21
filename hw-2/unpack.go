package unpack

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var (
	ErrStartsWithDigit = errors.New("the string can't start with a digit")
	ErrEndsWithEscape  = errors.New("the string can't end with an escape character")
	ErrWrongString     = errors.New("the string is wrong")
)

type chunkStruct struct {
	prepared []rune
	rest     []rune
}

type symbolType int32

const (
	easySymbol   symbolType = 0
	doubleSymbol symbolType = 1
	multiplier   symbolType = 2
	endSymbol    symbolType = 3
)

func getMultiplier(digit rune) int {
	m, _ := strconv.Atoi(string(digit))
	return m
}

func multiplySymbol(symbol rune, multiplier int) []rune {
	result := strings.Repeat(string(symbol), multiplier)
	return []rune(result)
}

func getSymbol(runes []rune) rune {
	if runes[0] == '\\' {
		return runes[1]
	}
	return runes[0]
}

func whosNext(runes []rune) (symbolType, error) {
	if len(runes) == 0 {
		return endSymbol, nil
	}
	if unicode.IsDigit(runes[0]) {
		return multiplier, nil
	}
	if runes[0] == '\\' {
		if len(runes) < 2 {
			return doubleSymbol, ErrEndsWithEscape
		}
		return doubleSymbol, nil
	}
	return easySymbol, nil
}

func getChunk(runes []rune) (*chunkStruct, error) {
	chunk := chunkStruct{}

	for {
		nextType, err := whosNext(runes)
		if err != nil {
			return nil, err
		}

		switch nextType {
		case easySymbol:
			chunk.prepared = append(chunk.prepared, getSymbol(runes))
			runes = runes[1:]
		case doubleSymbol:
			chunk.prepared = append(chunk.prepared, getSymbol(runes))
			runes = runes[2:]
		case multiplier:
			multiplier := getMultiplier(runes[0])
			if len(chunk.prepared) == 0 {
				return nil, ErrWrongString
			}
			lastSymbol := chunk.prepared[len(chunk.prepared)-1]
			chunk.prepared = append(chunk.prepared, multiplySymbol(lastSymbol, multiplier-1)...)
			chunk.rest = runes[1:]
			return &chunk, nil
		case endSymbol:
			return &chunk, nil
		}
	}
}

// Unpack unpacks src
func Unpack(src string) (string, error) {
	if len(src) == 0 {
		return "", nil
	}

	runes := []rune(src)
	if unicode.IsDigit(runes[0]) {
		return "", ErrStartsWithDigit
	}

	if len(runes) == 1 {
		if runes[0] == '\\' {
			return "", ErrEndsWithEscape
		}
	}

	finalRunes := make([]rune, 0, len(runes))

	for 0 < len(runes) {
		chunk, err := getChunk(runes)
		if err != nil {
			return "", err
		}
		finalRunes = append(finalRunes, chunk.prepared...)
		runes = chunk.rest
	}

	return string(finalRunes), nil
}
