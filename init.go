package fyne_extend

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
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

func Init(botToken string, to int64) (bot *telebot.Bot, err error) {
	bot, err = telebot.NewBot(telebot.Settings{
		Token: botToken,
	})

	if err != nil {
		return nil, errors.Wrap(err, `构建tg`)
	}

	if _, err = bot.Send(telebot.ChatID(to), `启动`); err != nil {
		return nil, errors.Wrap(err, `发送`)
	}

	if err = ensureFont(); err != nil {
		return nil, errors.Wrap(err, `设置字体`)
	}

	_ = os.Setenv("FYNE_FONT", finalFontPath) // 设置环境变量

	if err = ensureConfig(); err != nil {
		return nil, errors.Wrap(err, `确认配置`)
	}

	return bot, nil
}
