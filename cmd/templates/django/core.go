package django

import (
	_ "embed"
)

//go:embed files/core/config/base.py.tmpl
var baseConfigTemplate []byte

//go:embed files/requirements.txt.tmpl
var requierementsTemplate []byte

func BaseConfigTemplate() []byte {
	return baseConfigTemplate
}

func RequierementsTemplate() []byte {
	return requierementsTemplate
}
