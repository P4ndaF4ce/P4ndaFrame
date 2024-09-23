package utils

import (
	"os"
	"path/filepath"
	"text/template"
)

func CreatePyEnv(projectName string, appDir string) error {
	if err := ExecuteCmd("python3",
		[]string{"-m", "venv", "env"},
		appDir); err != nil {
		return err
	}

	return nil
}

func InstallRequirements(projectName string, appDir string) error {
	if err := ExecuteCmd("pip",
		[]string{"install", "-r", "requirements.txt"},
		appDir); err != nil {
		return err
	}
	return nil
}

func CreateDunderInits(projectPath string, pathToCreate ...string) error {

	for _, endpath := range pathToCreate {
		createdFile, err := os.Create(filepath.Join(projectPath, endpath, "__init__.py"))
		if err != nil {
			return err
		}

		defer createdFile.Close()

		createdTemplate := template.Must(template.New("__init__.py").Parse(""))
		err = createdTemplate.Execute(createdFile, "")
		if err != nil {
			return err
		}
	}

	return nil
}
