package program

import (
	"log"
	"os"
	"path/filepath"
	"text/template"
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

func (p *Program) CreateFilesFromMap(pathToCreate string, projectPath string, files map[string][]byte) error {

	for fileName, fileTemplate := range files {
		err := p.CreateFileFromTemplate(pathToCreate, projectPath, fileName, fileTemplate)
		if err != nil {
			return nil
		}
	}
	return nil
}
