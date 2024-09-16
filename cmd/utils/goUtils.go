package utils

func InitGoMod(projectName string, appDir string) error {
	if err := ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		return err
	}

	return nil
}

func GoGetPackage(appDir string, packages []string) error {
	for _, packageName := range packages {
		if err := ExecuteCmd("go",
			[]string{"get", "-u", packageName},
			appDir); err != nil {
			return err
		}
	}

	return nil
}

func GoFmt(appDir string) error {
	if err := ExecuteCmd("gofmt",
		[]string{"-s", "-w", "."},
		appDir); err != nil {
		return err
	}

	return nil
}

func GoTidy(appDir string) error {
	err := ExecuteCmd("go", []string{"mod", "tidy"}, appDir)
	if err != nil {
		return err
	}
	return nil
}
