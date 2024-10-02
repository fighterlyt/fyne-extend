package fyne_extend

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ncruces/go-sqlite3/gormlite"
	"github.com/pkg/errors"
	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
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

func Init(botToken string, to int64, needSend bool, dbName string) (bot *telebot.Bot, db *gorm.DB, err error) {
	bot, err = telebot.NewBot(telebot.Settings{
		Token: botToken,
	})

	if err != nil {
		return nil, nil, errors.Wrap(err, `构建tg`)
	}

	send := func(data string, needSend bool) {
		if needSend {
			_, _ = bot.Send(telebot.ChatID(to), data)
		}
	}

	if _, err = bot.Send(telebot.ChatID(to), `启动`); err != nil {
		return nil, nil, errors.Wrap(err, `发送`)
	}

	defer func() {
		if x := recover(); x != nil {
			_, _ = bot.Send(telebot.ChatID(to), fmt.Sprintf(`%v`, x))
		}
	}()

	send(`ensureFont`, needSend)

	if err = ensureFont(); err != nil {
		return bot, nil, errors.Wrap(err, `设置字体`)
	}

	send(`ensureFont`, needSend)

	_ = os.Setenv("FYNE_FONT", finalFontPath) // 设置环境变量

	if err = ensureConfig(); err != nil {
		return bot, nil, errors.Wrap(err, `确认配置`)
	}

	send(`ensureConfig`, needSend)

	if dbName != `` {
		if db, err = initDB(dbName); err != nil {
			return bot, nil, errors.Wrap(err, `init DB`)
		}
	}

	return bot, db, nil
}

func initDB(name string) (db *gorm.DB, err error) {
	db, err = gorm.Open(gormlite.Open(filepath.Join(GetBasePath(), name+`.db`)), &gorm.Config{})

	return db, err
}
