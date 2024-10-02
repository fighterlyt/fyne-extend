package fyne_extend

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
)

var (
	finalFontPath = filepath.Join(GetBasePath(), `.fyneFont.ttf`)
)

func ensureFont() error {
	if runtime.GOOS == `android` {
		finalFontPath = `/system/fonts/DroidSans.ttf`

		return nil
	}

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

func Init(botToken string, to int64, needSend bool) (bot *telebot.Bot, err error) {
	bot, err = telebot.NewBot(telebot.Settings{
		Token: botToken,
	})

	if err != nil {
		return nil, errors.Wrap(err, `构建tg`)
	}

	send := func(data string, needSend bool) {
		if needSend {
			_, _ = bot.Send(telebot.ChatID(to), data)
		}
	}

	if _, err = bot.Send(telebot.ChatID(to), `启动`); err != nil {
		return nil, errors.Wrap(err, `发送`)
	}

	defer func() {
		if x := recover(); x != nil {
			_, _ = bot.Send(telebot.ChatID(to), fmt.Sprintf(`%v`, x))
		}
	}()

	send(`ensureFont`, needSend)

	if err = ensureFont(); err != nil {
		return bot, errors.Wrap(err, `设置字体`)
	}

	send(`ensureFont`, needSend)

	_ = os.Setenv("FYNE_FONT", finalFontPath) // 设置环境变量

	if err = ensureConfig(); err != nil {
		return bot, errors.Wrap(err, `确认配置`)
	}

	send(`ensureConfig`, needSend)

	return bot, nil
}
