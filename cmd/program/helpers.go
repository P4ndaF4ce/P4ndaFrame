package program

import (
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/54L1M/CraftyPanda/cmd/templates/general"
)

func (p *Program) CreatePath(pathToCreate string, projectPath string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0o751)
		if err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}
	}
	return nil
}

func (p *Program) CreateFileFromTemplate(pathToCreate string, projectPath string, fileName string, fileTmpl []byte) error {

	createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	createdTemplate := template.Must(template.New(fileName).Parse(string(fileTmpl)))
	err = createdTemplate.Execute(createdFile, p)
	if err != nil {
		return err
	}

	return nil
}

func (p *Program) CreateGeneralFiles(pathToCreate string, projectPath string) error {
	generalFiles := map[string][]byte{"README.md": general.ReadMeTemplate(), ".gitignore": general.GitIgnoreTemplate()}

	for fileName, fileTemplate := range generalFiles {
		createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
		if err != nil {
			return err
		}
		defer createdFile.Close()
		createdTemplate := template.Must(template.New(fileName).Parse(string(fileTemplate)))
		err = createdTemplate.Execute(createdFile, p)
		if err != nil {
			return err
		}
	}

	return nil
}
