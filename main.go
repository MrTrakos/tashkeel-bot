// Writed By @trakoss on Telegram
// Open-Source Project, 2023.

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

func tashkill(text string) string {
	api := "http://www.7koko.com/api/tashkil/index.php"
	data := strings.NewReader(`textArabic=` + text)
	req, err := http.NewRequest("POST", api, data)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Origin", "http://www.7koko.com")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Referer", "http://www.7koko.com/apps/tashkil/index.php")
	req.Header.Set("Cookie", "_tccl_visitor=a0ba1532-cfb6-5bb6-80a7-e6c8ce65dc95; _tccl_visit=a0ba1532-cfb6-5bb6-80a7-e6c8ce65dc95; _ga_X6H4SPSJ35=GS1.1.1699441538.1.0.1699441538.0.0.0; _ga=GA1.1.451612605.1699441538; _ga_WNHWFV5ER8=GS1.1.1699441538.1.0.1699441538.0.0.0")

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	println(response.StatusCode)
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	response_ := (strings.TrimSpace(string(responseBody)))
	return response_

}

const (
	token  = "YOUR_BOT_TOKEN"
	sudoID = 144444444
)

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:   token,
		Updates: 1,
		Poller:  &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)

		return
	}
	users := []int{}
	bot.Handle("/start", func(msg telebot.Context) error {
		user_id, userFirstName := msg.Sender().ID, msg.Sender().FirstName
		lastUpdateTime := time.Now().Unix() - msg.Message().Unixtime // getting different between 2 Unix's.
		if lastUpdateTime > 60 {
			// if the update was 60 seconds old, bot will not reply to.
			return nil
		}
		if slices.Contains(users, int(user_id)) == false {
			// this gonna save user ids, without duplicate.
			users = append(users, int(user_id))
		}
		xstartMessage := strings.Builder{}
		xstartMessage.WriteString(
			fmt.Sprintf("Hello %s!\nYou can make your Arabic Text more beauty. with this Bot!\nNow just send your text..", userFirstName),
		)
		bot.Reply(msg.Message(), xstartMessage.String())
		return nil
	})
	bot.Handle(telebot.OnText, func(msg telebot.Context) error {
		userFirstName := msg.Sender().FirstName
		lastUpdateTime := time.Now().Unix() - msg.Message().Unixtime // getting different between 2 Unix's.
		if lastUpdateTime > 60 {
			// if the update was 60 seconds old, bot will not reply to.
			return nil
		}
		if strings.Contains(msg.Text(), "/brodcast") { // brodcasting Command
			if msg.Sender().ID != int64(sudo) {
				bot.Reply(msg.Message(), "ACCESS_DENIED")
				return nil
			}
			FinalText := strings.ReplaceAll(msg.Text(), "/brodcast ", "")
			bot.Reply(msg.Message(), "Well done, I'll complete this.")
			t, f := 0, 0
			for id, ids := range users {
				println(id, ids)
				Xsend, err := bot.Send(&telebot.Chat{ID: int64(ids)}, FinalText)
				if err != nil {
					fmt.Println(err)
					f++
				} else {
					t++
				}

				if t%10 == 0 {
					Xsend.Time().Unix()
					time.Sleep(3000)
				}
			}
			bot.Reply(msg.Message(), fmt.Sprintf("brodcast is complete, which is sent for %d, and failed to %d", t, f))
			return nil
		}
		GetResponseOf := tashkill(msg.Text())
		SuccessMessage := strings.Builder{}
		SuccessMessage.WriteString(
			fmt.Sprintf("Well %v.. Your Text is ready!\n<code>%s</code>", userFirstName, GetResponseOf),
		)
		SP := telebot.SendOptions{ParseMode: telebot.ModeHTML} // SP stands for SendOptions
		bot.Send(&telebot.Chat{ID: msg.Sender().ID}, SuccessMessage.String(), &SP)
		return nil // skipping the errors.
	})
	// admins commands
	bot.Handle(".admin", func(msg telebot.Context) error {
		if msg.Sender().ID != int64(sudo) {
			bot.Reply(msg.Message(), "ACCESS_DENIED")
			return nil
		}
		SudoMessage := strings.Builder{}
		SudoMessage.WriteString("Welcome To Admin Panel! Here you can see your accessble commands!\n")
		SudoMessage.WriteString("- <code>/brodcast TEXT </code> .\n")
		SudoMessage.WriteString("- <code>/stats</code>\n")
		SudoMessage.WriteString("This was your commands.")
		SP := telebot.SendOptions{ParseMode: telebot.ModeHTML}
		bot.Send(&telebot.Chat{ID: msg.Sender().ID}, SudoMessage.String(), &SP)
		return nil
	})
	// stats command for Sudo
	bot.Handle("/stats", func(msg telebot.Context) error {
		if msg.Sender().ID != int64(sudo) {
			bot.Reply(msg.Message(), "ACCESS_DENIED")
			return nil
		}
		bot.Reply(msg.Message(), fmt.Sprintf("Your bot audience %d Member.", len(users)))
		return nil
	})

	bot.Start()
}
