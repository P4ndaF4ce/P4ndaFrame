/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/54L1M/CraftyPanda/cmd/flags"
	"github.com/54L1M/CraftyPanda/cmd/program"
	spinner "github.com/54L1M/CraftyPanda/cmd/ui"
	"github.com/54L1M/CraftyPanda/cmd/utils"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/spf13/cobra"
)

func init() {
	var flagProject flags.Project
	var flagsDatabase flags.DataBase

	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().VarP(&flagProject, "project", "p", fmt.Sprintf("Project type. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	createCmd.MarkPersistentFlagRequired("project")

	createCmd.PersistentFlags().StringP("name", "n", "", "Name of project to create")
	createCmd.MarkPersistentFlagRequired("name")

	createCmd.PersistentFlags().VarP(&flagsDatabase, "database", "d", fmt.Sprintf("Database type. Allowed values: %s", strings.Join(flags.AllowedDataBaseTypes, ", ")))

}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		flagName := cmd.Flag("name").Value.String()
		if flagName != "" && !utils.ValidateModuleName(flagName) {
			err = fmt.Errorf("'%s' is not a valid module name. Please choose a different name", flagName)
			cobra.CheckErr(err)
		}

		rootDirName := utils.GetRootDir(flagName)
		if rootDirName != "" && doesDirectoryExistAndIsNotEmpty(rootDirName) {
			err = fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", rootDirName)
			cobra.CheckErr(err)
		}

		flagProject := flags.Project(cmd.Flag("project").Value.String())
		flagDataBase := flags.DataBase(cmd.Flag("database").Value.String())

		if flagProject == "django" || flagProject == "drf" {
			if flagDataBase == "" {
				err = fmt.Errorf("upon choosing django or drf as project, database should be chosen as well. available values are: %s", strings.Join(flags.AllowedDataBaseTypes, ", "))
				cobra.CheckErr(err)
			}
		}

		project := &program.Program{
			Name:         flagName,
			ProjectType:  flagProject,
			DataBaseType: flagDataBase,
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(err)
		}
		project.AbsolutePath = currentWorkingDir
		spinner := tea.NewProgram(spinner.InitialModelNew())

		// add synchronization to wait for spinner to finish
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := spinner.Run(); err != nil {
				cobra.CheckErr(err)
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("The program encountered an unexpected issue and had to exit. The error was:", r)
				if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
					log.Printf("Problem releasing terminal: %v", releaseErr)
				}
			}
		}()

		err = project.CreateDjangoProject()
		if err != nil {
			log.Printf("problem creating files for project. %v", err)
			cobra.CheckErr(err)
		}
	},
}

func doesDirectoryExistAndIsNotEmpty(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Printf("could not read directory: %v", err)
			cobra.CheckErr(err)
		}
		if len(dirEntries) > 0 {
			return true
		}
	}
	return false
}
