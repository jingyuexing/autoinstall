package check

import (
	"autoinstall/config"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CheckFileExists(folderPath string, fileName string) bool {
	if _, err := os.Stat(folderPath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return false
	}

	for _, file := range files {
		if file.Name() == fileName {
			return true
		}
	}

	return false
}

func CheckWorkspace() {
	config := config.LoadConfig("", _configFiles)
	if config == nil {
		return
	}
	for _, item := range config.Preset {
		for _, dependcies := range item.DepndciesFile {
			_, err := os.Stat(dependcies)
			if err == nil {
				AutoInstall(item.Commander)
				return
			} else {
				continue
			}
		}
	}

}

func CheckFile(filename string) bool {
	dir, err := os.Getwd()
	if err != nil {
		return false
	}
	path := filepath.Join(dir, filename)
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func AutoInstall(cmmd string) {
	commandAndArgs := strings.Split(cmmd, " ")
	cmd := exec.Command(commandAndArgs[0], commandAndArgs[1:]...)
	output, err := cmd.Output()
	if err != nil {
		return
	}
}
