package fyne_extend

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/pkg/errors"
)

type Config struct {
}

var (
	configPath      = `config.json`
	finalConfigPath = filepath.Join(os.Getenv(`HOME`), configPath)
)

func SaveConfig(ctx context.Context, build func(ctx2 context.Context) (conf any, err error), win fyne.Window) {
	var (
		err  error
		conf any
	)

	if conf, err = build(ctx); err != nil {
		dialog.ShowError(err, win)
		return
	}

	if err = saveConfig(conf); err != nil {
		dialog.ShowError(err, win)
		return
	}

	dialog.ShowInformation(`jira`, `配置更新`, win)

	win.Hide()
}

func saveConfig(conf any) error {
	file, err := os.OpenFile(finalConfigPath, os.O_RDWR, os.ModePerm)

	if err != nil {
		return errors.Wrap(err, `打开文件`)
	}

	defer func() {
		_ = file.Close()
	}()

	if err = json.NewEncoder(file).Encode(conf); err != nil {
		return errors.Wrap(err, `写入配置`)
	}

	return nil
}

func ensureConfig() error {
	var (
		err  error
		file *os.File
	)

	if _, err = os.Stat(finalConfigPath); err == nil {
		return nil
	}

	if !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if file, err = os.Create(finalConfigPath); err != nil {
		return errors.Wrap(err, `创建配置文件`)
	}

	_, _ = file.WriteString(`{}`)

	_ = file.Close()

	return nil
}
