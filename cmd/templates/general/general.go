package general

import _ "embed"

var GeneralFiles = map[string][]byte{"README.md": ReadMeTemplate(), ".gitignore": GitIgnoreTemplate()}

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
