package flags

import (
	"fmt"
	"strings"
)

type Project string

const (
	Django  Project = "django"
	DRF     Project = "drf"
	PyTorch Project = "pytorch"
	Go      Project = "go"
	GoCLI   Project = "go-cli"
)

var AllowedProjectTypes = []string{string(Django), string(DRF), string(PyTorch), string(Go), string(GoCLI)}

func (p Project) String() string {
	return string(p)
}

func (p *Project) Type() string {
	return "Project"
}

func (p *Project) Set(value string) error {
	for _, project := range AllowedProjectTypes {
		if project == value {
			*p = Project(value)
			return nil
		}
	}
	return fmt.Errorf("projects to use. allowed values: %s", strings.Join(AllowedProjectTypes, ", "))
}
