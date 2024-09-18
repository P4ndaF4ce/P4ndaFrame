package flags

import (
	"fmt"
	"strings"
)

type DataBase string

const (
	Postgres DataBase = "postgres"
	Mysql    DataBase = "mysql"
	Sqlite   DataBase = "sqlite"
)

var AllowedDataBaseTypes = []string{string(Postgres), string(Mysql), string(Sqlite)}

func (db DataBase) String() string {
	return string(db)
}

func (db *DataBase) Type() string {
	return "Project"
}

func (db *DataBase) Set(value string) error {
	for _, database := range AllowedDataBaseTypes {
		if database == value {
			*db = DataBase(value)
			return nil
		}
	}
	return fmt.Errorf("projects to use. allowed values: %s", strings.Join(AllowedProjectTypes, ", "))
}
