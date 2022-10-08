package config

import (
	"errors"
	"os"
	"os/exec"
)

var (
	defaultCalibreDB = "/opt/calibre/bin/calibredb"
	defaultConvert   = "/opt/calibre/bin/ebook-convert"
)

func init() {
	if _, err := os.Stat(defaultCalibreDB); errors.Is(err, os.ErrNotExist) {
		calibreDB, err := exec.LookPath("calibredb")
		if err == nil {
			defaultCalibreDB = calibreDB
		}
	}

	if _, err := os.Stat(defaultConvert); errors.Is(err, os.ErrNotExist) {
		ebookConvert, err := exec.LookPath("ebook-convert")
		if err == nil {
			defaultConvert = ebookConvert
		}
	}
}
