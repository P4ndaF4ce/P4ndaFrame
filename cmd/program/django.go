package program

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/54L1M/CraftyPanda/cmd/flags"
	"github.com/54L1M/CraftyPanda/cmd/templates/django"
	"github.com/54L1M/CraftyPanda/cmd/utils"
	"github.com/spf13/cobra"
)

type Program struct {
	Name         string
	ProjectType  flags.Project
	DataBaseType flags.DataBase
	AbsolutePath string
}

const (
	coreConfigPath = "core/config"
	core           = "core"
	root           = "/"
)

func (p *Program) CreateDjangoProject() error {
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		if err := os.MkdirAll(p.AbsolutePath, 0o754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}
	p.Name = strings.TrimSpace(p.Name)

	projectPath := filepath.Join(p.AbsolutePath, utils.GetRootDir(p.Name))

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := os.MkdirAll(projectPath, 0o751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	err := p.CreatePath(coreConfigPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		cobra.CheckErr(err)
	}

	err = utils.CreateDunderInits(projectPath, coreConfigPath, core)
	if err != nil {
		log.Printf("Error creating __init__.py files: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileFromTemplate(root, projectPath, "requirements.txt", django.RequierementsTemplate())
	if err != nil {
		log.Printf("Error creating requierements.txt file: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileFromTemplate(coreConfigPath, projectPath, "base.py", django.BaseConfigTemplate())
	if err != nil {
		log.Printf("Error creating base.py file: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileFromTemplate(coreConfigPath, projectPath, "dev.py", django.DevConfigTemplate())
	if err != nil {
		log.Printf("Error creating dev.py file: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = p.CreateFileFromTemplate(coreConfigPath, projectPath, "prod.py", django.ProdConfigTemplate())
	if err != nil {
		log.Printf("Error creating prod.py file: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = utils.CreatePyEnv(p.Name, projectPath)
	if err != nil {
		log.Printf("Error creating virtualenv: %s", err)
		cobra.CheckErr(err)
		return err
	}

	err = utils.InstallRequirements(p.Name, projectPath)
	if err != nil {
		log.Printf("Error installing requirements: %s", err)
		cobra.CheckErr(err)
		return err
	}

	return nil
}
