package django

import (
	_ "embed"
)

var ConfigFiles = map[string][]byte{"base.py": BaseConfigTemplate(), "dev.py": DevConfigTemplate(), "prod.py": ProdConfigTemplate()}

var CoreFiles = map[string][]byte{"asgi.py": AsgiTemplate(), "wsgi.py": WsgiTemplate(), "urls.py": UrlsTemplate()}

var RootFiles = map[string][]byte{"requirements.txt": RequierementsTemplate(), "manage.py": ManagePyTemplate()}

var DjangoTemplateFiles = map[string][]byte{"base.html": BaseTemplate(), "footer.html": FooterTemplate(), "scripts.html": ScriptsTemplate(),
	"head.html": HeadTemplate()}

//go:embed files/core/config/base.py.tmpl
var baseConfigTemplate []byte

//go:embed files/core/config/dev.py.tmpl
var devConfigTemplate []byte

//go:embed files/core/config/prod.py.tmpl
var prodConfigTemplate []byte

//go:embed files/core/asgi.py.tmpl
var asgiTemplate []byte

//go:embed files/core/wsgi.py.tmpl
var wsgiTemplate []byte

//go:embed files/core/urls.py.tmpl
var urlsTemplate []byte

//go:embed files/manage.py.tmpl
var managePyTemplate []byte

//go:embed files/requirements.txt.tmpl
var requierementsTemplate []byte

//go:embed files/templates/base/base.html.tmpl
var baseTemplate []byte

//go:embed files/templates/base/head.html.tmpl
var headTemplate []byte

//go:embed files/templates/base/footer.html.tmpl
var footerTemplate []byte

//go:embed files/templates/base/scripts.html.tmpl
var scriptsTemplate []byte

func BaseConfigTemplate() []byte {
	return baseConfigTemplate
}

func DevConfigTemplate() []byte {
	return devConfigTemplate
}

func AsgiTemplate() []byte {
	return asgiTemplate
}

func WsgiTemplate() []byte {
	return wsgiTemplate
}

func UrlsTemplate() []byte {
	return urlsTemplate
}

func ProdConfigTemplate() []byte {
	return prodConfigTemplate
}

func RequierementsTemplate() []byte {
	return requierementsTemplate
}

func ManagePyTemplate() []byte {
	return managePyTemplate
}

func BaseTemplate() []byte {
	return baseTemplate
}
func HeadTemplate() []byte {
	return headTemplate
}
func FooterTemplate() []byte {
	return footerTemplate
}
func ScriptsTemplate() []byte {
	return scriptsTemplate
}
