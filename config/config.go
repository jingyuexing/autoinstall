package config

import (
	"embed"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	json "encoding/json"

	yaml "gopkg.in/yaml.v3"
)

// // go:embed meta
var meta embed.FS

type Items struct {
	Type          string   `json:"type" yaml:"type"`
	Language      string   `json:"language" yaml:"language"`
	DepndciesFile []string `json:"dependcies" yaml:"dependcies"`
	Commander     string   `json:"command" yaml:"command"`
}

type Configure struct {
	Preset []Items `json:"preset" yaml:"preset"`
}

type File struct {
	Dir  string
	Name string
	Ext  string
}

func LoadJSONBytes(data []byte) *Configure {
	config_ := new(Configure)
	err := json.Unmarshal(data, config_)
	if err != nil {
		return nil
	}
	return config_
}

func LoadYAMLBytes(data []byte) *Configure {
	config_ := new(Configure)
	err := yaml.Unmarshal(data, config_)
	if err != nil {
		return nil
	}
	return config_
}

func LoadConfig(filename string, fs embed.FS) *Configure {
	file := ParserFileName(filename)
	data, err := fs.ReadFile(filename)
	if err != nil {
		return nil
	}
	switch file.Ext {
	case ".json":
		return LoadJSONBytes(data)
	case ".yml", ".yaml":
		return LoadYAMLBytes(data)
	}
	return nil
}
func LoadJSON(filename string) *Configure {
	data, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("file is not exist!")
		}
		switch err {
		case os.ErrPermission:
			fmt.Printf("ErrPermission")
		}
		return nil
	}
	var config *Configure
	err = json.Unmarshal(data, config)
	if err != nil {
		if errors.Is(err, &json.SyntaxError{}) {
			fmt.Printf("json syntax error %s", err)
		}
		return nil
	}
	return config
}

func LoadYAML(filename string) *Configure {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil
	}
	var config *Configure
	err = yaml.Unmarshal(data, config)
	if err != nil {
		fmt.Printf("%s", err)
	}

	return config

}

func ParserFileName(filename string) *File {
	return &File{
		Dir:  filepath.Dir(filename),
		Name: filepath.Base(filename),
		Ext:  path.Ext(filename),
	}
}
