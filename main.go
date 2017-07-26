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
		whiteSquaresText,
		blackSquaresText,
		parenthesisText,
		frakturText,
		scriptText,
		doubleStruckText,
		monospaceText,
		underlineText,
		overlineText,
		strikethroughText,
		dottedText,
		triangledText,
		deniedText,
		brailleText :=
			vaporwave(update.InlineQuery.Query),
			zalgo(update.InlineQuery.Query),
			ideographs(update.InlineQuery.Query),
			whiteCircles(update.InlineQuery.Query),
			blackCircles(update.InlineQuery.Query),
			whiteSquares(update.InlineQuery.Query),
			blackSquares(update.InlineQuery.Query),
			parenthesis(update.InlineQuery.Query),
			fraktur(update.InlineQuery.Query),
			script(update.InlineQuery.Query),
			doubleStruck(update.InlineQuery.Query),
			monospace(update.InlineQuery.Query),
			underline(update.InlineQuery.Query),
			overline(update.InlineQuery.Query),
			strikethrough(update.InlineQuery.Query),
			dotted(update.InlineQuery.Query),
			triangled(update.InlineQuery.Query),
			denied(update.InlineQuery.Query),
			braille(update.InlineQuery.Query)

		inlineConf := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			IsPersonal:    true,
			CacheTime:     0,
			Results: []interface{}{
				NewInlineQueryResultArticleDesc( "1", "ＶＡＰＯＲＷＡＶＥ", vaporwaveText, vaporwaveText),
				NewInlineQueryResultArticleDesc( "2", "Z͉̩̖̝͗aͩl̵͖̯̰̠͎̘ͣͭͥ͋ͅḡ̒̏ő̫ͣ͋ͅ", zalgoText, zalgoText),
				NewInlineQueryResultArticleDesc( "3", "工刀モ口ム尺丹ㄗ卄ち", ideographsText, ideographsText),
				NewInlineQueryResultArticleDesc( "4", "Ⓦⓗⓘⓣⓔ ⓒⓘⓡⓒⓛⓔⓢ", whiteCirclesText, whiteCirclesText),
				NewInlineQueryResultArticleDesc( "5", "🅑🅛🅐🅒🅚 🅒🅘🅡🅒🅛🅔🅢", blackCirclesText, blackCirclesText),
				NewInlineQueryResultArticleDesc( "6", "🅆🄷🄸🅃🄴 🅂🅀🅄🄰🅁🄴🅂", whiteSquaresText, whiteSquaresText),
				NewInlineQueryResultArticleDesc( "7", "🆆🅷🅸🆃🅴 🆂🆀🆄🅰🆁🅴🆂", blackSquaresText, blackSquaresText),
				NewInlineQueryResultArticleDesc( "8", "🄟⒜⒭⒠⒩⒯⒣⒠⒮⒤⒮", parenthesisText, parenthesisText),
				NewInlineQueryResultArticleDesc( "9", "𝕱𝖗𝖆𝖐𝖙𝖚𝖗", frakturText, frakturText),
				NewInlineQueryResultArticleDesc("10", "𝓢𝓬𝓻𝓲𝓹𝓽", scriptText, scriptText),
				NewInlineQueryResultArticleDesc("11", "𝔻𝕠𝕦𝕓𝕝𝕖-𝕤𝕥𝕣𝕦𝕔𝕜", doubleStruckText, doubleStruckText),
				NewInlineQueryResultArticleDesc("12", "𝙼𝚘𝚗𝚘𝚜𝚙𝚊𝚌𝚎", monospaceText, monospaceText),
				NewInlineQueryResultArticleDesc("13", "U̲n̲d̲e̲r̲l̲i̲n̲e̲", underlineText, underlineText),
				NewInlineQueryResultArticleDesc("14", "O̅v̅e̅r̅l̅i̅n̅e̅", overlineText, overlineText),
				NewInlineQueryResultArticleDesc("15", "S̶t̶r̶i̶k̶e̶t̶h̶r̶o̶u̶g̶h̶", strikethroughText, strikethroughText),
				NewInlineQueryResultArticleDesc("16", "Ḋȯṫṫėḋ", dottedText, dottedText),
				NewInlineQueryResultArticleDesc("17", "T⃤r⃤i⃤a⃤n⃤g⃤l⃤e⃤d⃤", triangledText, triangledText),
				NewInlineQueryResultArticleDesc("18", "D⃠e⃠n⃠i⃠e⃠d⃠", deniedText, deniedText),
				NewInlineQueryResultArticleDesc("19", "⠠⠃⠗⠁⠊⠇⠇⠑", brailleText, brailleText),
			},
		}

		if _, err := bot.AnswerInlineQuery(inlineConf); err != nil {
			log.Println(err)
		}
	}
}
