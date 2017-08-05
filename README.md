# Dabbext Bot
A Telegram inline bot to generate funny text effects

## Usage
To use it within Telegram, just type ```@DabbextBot <your message>``` and the bot will answer with the list of possible text transformations.

## Currently supported effects
- Vaporwave
- Zalgo
- Ideographs
- White circled text
- Black circled text
- White squared text
- Black squared text
- Parenthesized text
- Fraktur
- Script
- Double-struck
- Monospace
- Underline
- Overline
- Strikethrough
- Dotted
- Triangled
- Denied
- Braille

## Build and run
Implying you have already Go set up and configured

- Clone the repository: ```git clone https://github.com/Dabolus/dabbext.git && cd dabbext```
- Install the Go bindings for the Telegram Bot API (you have to do it manually because for now I'm not using a package manager): ```github.com/go-telegram-bot-api/telegram-bot-api```
- Clone the file ```tokens.go.tmpl``` to ```tokens.go``` (or rename the original) and replace ```<INSERT YOUR BOT TOKEN HERE>``` with your own bot token (you can get it from [@BotFather](https://t.me/BotFather)
- Finally, build the project and run it: ```go build && dabbext```

_N.B. Remember to set the bot as an inline bot by issuing the command_ ```/setinline``` _to BotFather_