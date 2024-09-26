package general

import _ "embed"

var GeneralFiles = map[string][]byte{"README.md": ReadMeTemplate(), ".gitignore": GitIgnoreTemplate(), "Taskfile.yml": TaskFileTemplate()}

//go:embed files/README.md.tmpl
var readMeTemplate []byte

//go:embed files/.gitignore.tmpl
var gitIgnoreTemplate []byte

//go:embed files/Taskfile.yml.tmpl
var taskFileTemplate []byte

func ReadMeTemplate() []byte {
	return readMeTemplate
}

func GitIgnoreTemplate() []byte {
	return gitIgnoreTemplate
}

func TaskFileTemplate() []byte {
	return taskFileTemplate
}
