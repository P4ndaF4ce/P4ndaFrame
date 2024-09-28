package general

import _ "embed"

var GeneralFiles = map[string][]byte{"README.md": ReadMeTemplate(), ".gitignore": GitIgnoreTemplate(), "Taskfile.yml": TaskFileTemplate(),
	"Dockerfile.prod": DockerfileProdTemplate(), "Dockerfile.dev": DockerfileDevTemplate(), "docker-compose.yml": DockerComposeTemplate()}

//go:embed files/README.md.tmpl
var readMeTemplate []byte

//go:embed files/.gitignore.tmpl
var gitIgnoreTemplate []byte

//go:embed files/Taskfile.yml.tmpl
var taskFileTemplate []byte

//go:embed files/Dockerfile.prod.tmpl
var dockerfileProdTemplate []byte

//go:embed files/Dockerfile.dev.tmpl
var dockerfileDevTemplate []byte

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

func DockerfileProdTemplate() []byte {
	return dockerfileProdTemplate
}

func DockerfileDevTemplate() []byte {
	return dockerfileDevTemplate
}

func DockerComposeTemplate() []byte {
	return dockerComposeTemplate
}
