package utils

import (
	"bytes"
	"os/exec"
	"regexp"
	"strings"
)

func ExecuteCmd(name string, args []string, dir string) error {
	command := exec.Command(name, args...)
	command.Dir = dir
	var out bytes.Buffer
	command.Stdout = &out
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}

func CheckGitConfig(key string) (bool, error) {
	cmd := exec.Command("git", "config", "--get", key)
	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			// The command failed to run.
			if exitError.ExitCode() == 1 {
				// The 'git config --get' command returns 1 if the key was not found.
				return false, nil
			}
		}
		// Some other error occurred.
		return false, err
	}
	// The command ran successfully, so the key is set.
	return true, nil
}

// ValidateModuleName returns true if it's a valid module name.
// It allows any number of / and . characters in between.
func ValidateModuleName(moduleName string) bool {
	matched, _ := regexp.Match("^[a-zA-Z0-9_-]+(?:[\\/.][a-zA-Z0-9_-]+)*$", []byte(moduleName))
	return matched
}

// GetRootDir returns the project directory name from the module path.
// Returns the last token by splitting the moduleName with /
func GetRootDir(moduleName string) string {
	tokens := strings.Split(moduleName, "/")
	return tokens[len(tokens)-1]
}
