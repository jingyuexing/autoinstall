package check

import (
	"autoinstall/config"
	"embed"
)

//go:embed meta/*
var _configFiles embed.FS

func Analyze(folder string) []config.Items {
	files, err := _configFiles.ReadDir("meta")
	if err != nil {
		return make([]config.Items, 0)
	}
	techs := make([]config.Items, 0)
	fingerprint := make([]config.Items, 0)
	for _, file := range files {
		path := "meta/" + file.Name()
		fingerPrintPreset := config.LoadConfig(path, _configFiles)
		fingerprint = append(fingerprint, fingerPrintPreset.Preset...)
	}

	for _, item := range fingerprint {
		hit := 0
		for _, fingerFiles := range item.DepndciesFile {
			if CheckFileExists(folder, fingerFiles) {
				hit++
			}
		}
		if hit > 0 {
			techs = append(techs, item)
		}
	}
	return techs
}
