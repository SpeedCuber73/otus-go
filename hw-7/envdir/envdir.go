package envdir

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func parseEnvDir(envPath string) error {
	files, err := ioutil.ReadDir(envPath)
	if err != nil {
		return err
	}

	for _, fileInfo := range files {
		name := fileInfo.Name()
		val, err := ioutil.ReadFile(path.Join(envPath, name))
		if err != nil {
			return err
		}
		os.Setenv(name, string(val))
	}

	return nil
}

// Run starts command with envPath environment
func Run(envPath string, command []string) (string, error) {
	err := parseEnvDir(envPath)
	if err != nil {
		return "", err
	}

	out, err := exec.Command(command[0], command[1:]...).Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}
