package general

import _ "embed"

//go:embed files/README.md.tmpl
var readMeTemplate []byte

//go:embed files/.gitignore.tmpl
var gitIgnoreTemplate []byte

func ReadMeTemplate() []byte {
	return readMeTemplate
}

func GitIgnoreTemplate() []byte {
	return gitIgnoreTemplate
}
