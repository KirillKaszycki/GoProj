package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	botToken := "7355285986:AAELCwI0dUhpVPSoN7kkEVutuf9E_WZib7I"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()
	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)
			messageText := update.Message.Text

			n, err := strconv.Atoi(messageText)
			if err != nil {
				_, _ = bot.SendMessage(
					tu.Message(chatID, "Please enter a valid number"),
				)
				continue
			}

			fibSequence := fibonacci(n)

			_, _ = bot.SendMessage(
				tu.Message(chatID, fmt.Sprintf("Fibonacci sequence up to %d: %s", n, fibSequence)),
			)
		}
	}
}

func fibonacci(n int) string {
	if n == 0 {
		return "0"
	}

	sequence := []int{0, 1}
	for i := 2; i < n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return fmt.Sprint(sequence)
}

func init() {
	fmt.Println("Starting bot...")
}
