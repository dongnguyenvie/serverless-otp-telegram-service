package command

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sv-otp/component/appctx"
)

type Command interface {
	Execute(msg *tgbotapi.MessageConfig) *tgbotapi.MessageConfig
}

type commandCtx struct {
	update tgbotapi.Update
	bot    *tgbotapi.BotAPI
	appCtx appctx.AppContext
}

func NewCommandContext(update tgbotapi.Update, bot *tgbotapi.BotAPI, appCtx appctx.AppContext) *commandCtx {
	return &commandCtx{update: update, bot: bot, appCtx: appCtx}
}
