package main_test

import (
	"autoinstall/check"
	"os"
	"testing"
)

func TestAnalyze(t *testing.T) {
	path,err := os.Getwd()
	if err != nil {
		return
	}
	check.Analyze(path)
}