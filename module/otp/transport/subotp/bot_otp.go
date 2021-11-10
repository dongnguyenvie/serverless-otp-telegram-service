package subotp

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sv-otp/common"
	"github.com/sv-otp/component/appctx"

	otpbiz "github.com/sv-otp/module/otp/biz"
	"github.com/sv-otp/module/otp/command"
)

func BotOtp(appCtx appctx.AppContext, bot *tgbotapi.BotAPI) {

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	webhookProvider := appCtx.WebhookProvider()

	submitOtpBiz := otpbiz.NewSubmitOtpBiz(webhookProvider)

	updates, _ := bot.GetUpdatesChan(u)
	var msg tgbotapi.MessageConfig

	go func() {
		defer common.AppRecover()

		for update := range updates {

			commandCtx := command.NewCommandContext(update, bot, appCtx)

			var cmd command.Command

			if update.CallbackQuery != nil {
				bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "đang xử lý"))

				switch {
				case command.IsCallbackGetOtpCommand(&update):
					cmd = command.NewCallbackGetOtpCommand(commandCtx)

				}
				if cmd == nil {
					continue
				}

				cmd.Execute(&msg)
				bot.Send(msg)
			}

			if update.Message != nil {
				switch {
				case command.IsMirrorCommandCase(&update):
					cmd = command.NewMirrorCommand(commandCtx)
				case command.IsStartCommandCase(&update):
					cmd = command.NewStartCommand(commandCtx)
				case command.IsGetOtpCommand(&update):
					cmd = command.NewGetOtpCommand(commandCtx)
				case command.IsConfirmAndGenOTP(&update): // Get contact from user
					cmd = command.NewConfirmAndGenOTPCommand(commandCtx, submitOtpBiz)
				case command.IsConfirmFail(&update):
					cmd = command.NewReceiveContactFailCommand(commandCtx)
				default:
					cmd = command.NewMirrorCommand(commandCtx)
				}

				if cmd == nil {
					continue
				}

				cmd.Execute(&msg)

				bot.Send(msg)
			}

		}
	}()
}
