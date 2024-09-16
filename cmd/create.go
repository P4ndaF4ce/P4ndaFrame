/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/54L1M/CraftyPanda/cmd/flags"
	"github.com/54L1M/CraftyPanda/cmd/utils"
	"github.com/spf13/cobra"
)

func init() {
	var flagProject flags.Project

	rootCmd.AddCommand(createCmd)

	createCmd.PersistentFlags().VarP(&flagProject, "project", "p", fmt.Sprintf("Project type. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	createCmd.MarkPersistentFlagRequired("project")

	createCmd.PersistentFlags().StringP("name", "n", "", "Name of project to create")
	createCmd.MarkPersistentFlagRequired("name")

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
		if flagProject != "" {
			fmt.Println(flagProject)
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
