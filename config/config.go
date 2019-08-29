package config

type cfg struct {
	templateDir, homeID, aboutID string
}

func Default(home, about string) cfg {
	return cfg{"./html/*", home, about}
}

func (c cfg) TemplateDir() string {
	return c.templateDir
}

func (c cfg) HomeID() string {
	return c.homeID
}

func (c cfg) AboutID() string {
	return c.aboutID
}
