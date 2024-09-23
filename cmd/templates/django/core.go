package django

import (
	_ "embed"
)

//go:embed files/core/config/base.py.tmpl
var baseConfigTemplate []byte

//go:embed files/core/config/dev.py.tmpl
var devConfigTemplate []byte

//go:embed files/core/config/prod.py.tmpl
var prodConfigTemplate []byte

//go:embed files/requirements.txt.tmpl
var requierementsTemplate []byte

func BaseConfigTemplate() []byte {
	return baseConfigTemplate
}

func DevConfigTemplate() []byte {
	return devConfigTemplate
}

func ProdConfigTemplate() []byte {
	return prodConfigTemplate
}
func RequierementsTemplate() []byte {
	return requierementsTemplate
}
