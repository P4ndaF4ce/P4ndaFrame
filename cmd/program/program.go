package program

import "github.com/54L1M/CraftyPanda/cmd/flags"

type Program struct {
	Name         string
	ProjectType  flags.Project
	DataBaseType flags.DataBase
}
