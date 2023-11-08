Telegram Bot for Tashkeel Arabic Text

This is a Telegram Bot written in Golang that adds Tashkeel to Arabic text to make it visually appealing. This README will guide you through the steps required to install and configure the bot.

## Requirements

- Golang
- `gopkg.in/telebot.v3` package
- Sudo ID (owner ID)
- Bot token

## Installation

1. Install Golang following the official guide: https://golang.org/doc/install.

2. Install the `gopkg.in/telebot.v3` package using the following command:

```shell
go get gopkg.in/telebot.v3
```

## Configuration

Before running the bot, you need to configure the bot token and the sudo ID. Follow the steps below:
0. Write in terminal:
```shell
go mod init bot
```
1. Open the `main.go` file in a text editor.

2. Locate the following lines:

```go
const (
    token   = "YOUR_BOT_TOKEN"
    sudo  = 144444444
)
```

3. Replace `YOUR_BOT_TOKEN` with the token obtained from BotFather when you created your bot.

4. Replace `144444444` with your Telegram user ID. You can get your ID by messaging the `@userinfobot` on Telegram.


## Bot Admin Panel

This bot includes an admin panel to control its functionality. To access send /admin.

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
