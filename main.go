package main

import (
	"log"
	"time"
	"math/rand"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func NewInlineQueryResultArticleDesc(id string, title string, messageText string, description string) tgbotapi.InlineQueryResultArticle {
	newArticle := tgbotapi.NewInlineQueryResultArticle(id, title, messageText)
	newArticle.Description = description
	return newArticle
}

func main() {
	// Initialize global pseudo random generator
	rand.Seed(time.Now().Unix())

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.InlineQuery == nil || len(update.InlineQuery.Query) == 0 { // if no inline query, ignore it
			continue
		}

		vaporwaveText,
		zalgoText,
		ideographsText,
		whiteCirclesText,
		blackCirclesText,
		parenthesisText :=
			vaporwave(update.InlineQuery.Query),
			zalgo(update.InlineQuery.Query),
			ideographs(update.InlineQuery.Query),
			whiteCircles(update.InlineQuery.Query),
			blackCircles(update.InlineQuery.Query),
			parenthesis(update.InlineQuery.Query)

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results: []interface{}{
				NewInlineQueryResultArticleDesc("1", "ＶＡＰＯＲＷＡＶＥ", vaporwaveText, vaporwaveText),
				NewInlineQueryResultArticleDesc("2", "Z͉̩̖̝͗aͩl̵͖̯̰̠͎̘ͣͭͥ͋ͅḡ̒̏ő̫ͣ͋ͅ", zalgoText, zalgoText),
				NewInlineQueryResultArticleDesc("3", "工刀モ口ム尺丹ㄗ卄ち", ideographsText, ideographsText),
				NewInlineQueryResultArticleDesc("4", "Ⓦⓗⓘⓣⓔ ⓒⓘⓡⓒⓛⓔⓢ", whiteCirclesText, whiteCirclesText),
				NewInlineQueryResultArticleDesc("5", "🅑🅛🅐🅒🅚 🅒🅘🅡🅒🅛🅔🅢", blackCirclesText, blackCirclesText),
				NewInlineQueryResultArticleDesc("6", "🄟⒜⒭⒠⒩⒯⒣⒠⒮⒤⒮", parenthesisText, parenthesisText),
			},
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}
