package general

import _ "embed"

var GeneralFiles = map[string][]byte{"README.md": ReadMeTemplate(), ".gitignore": GitIgnoreTemplate(), "Taskfile.yml": TaskFileTemplate(),
	"Dockerfile": DockerfileTemplate(), "docker-compose.yml": DockerComposeTemplate()}

//go:embed files/README.md.tmpl
var readMeTemplate []byte

//go:embed files/.gitignore.tmpl
var gitIgnoreTemplate []byte

//go:embed files/Taskfile.yml.tmpl
var taskFileTemplate []byte

//go:embed files/Dockerfile.tmpl
var dockerfileTemplate []byte

//go:embed files/docker-compose.yml.tmpl
var dockerComposeTemplate []byte

func ReadMeTemplate() []byte {
	return readMeTemplate
}

func GitIgnoreTemplate() []byte {
	return gitIgnoreTemplate
}

func TaskFileTemplate() []byte {
	return taskFileTemplate
}

func DockerfileTemplate() []byte {
	return dockerfileTemplate
}

func DockerComposeTemplate() []byte {
	return dockerComposeTemplate
}
