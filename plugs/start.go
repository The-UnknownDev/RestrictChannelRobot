package plugs

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Start(bot *gotgbot.Bot, ctx *ext.Context) error {
	var MSG string = `
**Hey User, I am a Highly Advanced Channel Spam Detector Bot.* π

π I can ban that channels which

spams your chat!

π I am simple, reliable and secure bot

*π₯ Powered By - @TheByteBots*

*πCreator - @TheUnknownDev*





	`
	if ctx.EffectiveChat.Type != "private" {
		ctx.EffectiveMessage.Reply(
			bot,
			"π Bot is working (β’βΏβ’) β’ 

If you're facing any issue then contact here - @ByteBotsSupport",
			&gotgbot.SendMessageOpts{ParseMode: "markdown"},
		)
	} else {
		ctx.EffectiveMessage.Reply(
			bot,
			MSG,
			&gotgbot.SendMessageOpts{
				ParseMode: "markdown",
				ReplyMarkup: gotgbot.InlineKeyboardMarkup{
					InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
						{Text: "My Source Code", Url: "github.com/The-UnknownDev/RestrictChannelRobot"},
					}},
				},
			},
		)
	}
	return ext.EndGroups
}
