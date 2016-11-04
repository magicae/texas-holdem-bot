package config

import (
	"github.com/magicae/telegram-bot"
)

type BotConfig struct {
	ID            int
	Token         string
	Username      string
	GetMoneyBase  int64
	GetMoneyBonus int64
	RaiseButtons  [][]*bot.KeyboardButton
}

var Bot *BotConfig = &BotConfig{
	Token:         "IMPORTANT:SET_YOUR_TOKEN_HERE",
	GetMoneyBase:  500,
	GetMoneyBonus: 9500,
	RaiseButtons: [][]*bot.KeyboardButton{
		[]*bot.KeyboardButton{
			&bot.KeyboardButton{Text: "100"},
			&bot.KeyboardButton{Text: "200"},
			&bot.KeyboardButton{Text: "300"},
			&bot.KeyboardButton{Text: "500"},
		},
		[]*bot.KeyboardButton{
			&bot.KeyboardButton{Text: "700"},
			&bot.KeyboardButton{Text: "1000"},
			&bot.KeyboardButton{Text: "1500"},
			&bot.KeyboardButton{Text: "2000"},
		},
		[]*bot.KeyboardButton{
			&bot.KeyboardButton{Text: "2500"},
			&bot.KeyboardButton{Text: "3000"},
			&bot.KeyboardButton{Text: "4000"},
			&bot.KeyboardButton{Text: "/allin"},
		},
	},
}
