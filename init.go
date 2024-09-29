package fyne_extend

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	finalFontPath = filepath.Join(os.Getenv(`HOME`), `.fyneFont.ttf`)
)

func ensureFont() error {
	file, err := os.Create(finalFontPath)

	if err != nil {
		return errors.Wrap(err, `create`)
	}

	defer func() {
		_ = file.Close()
	}()

	_, _ = file.Write(miSans)

	return nil
}

func init() {
	if err := ensureFont(); err != nil {
		panic(`设置字体` + err.Error())
	}

	_ = os.Setenv("FYNE_FONT", finalFontPath) // 设置环境变量
}
