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
	sequence := make([]int, n)
	sequence[0] = 1
	if n > 1 {
		sequence[1] = 1
		for i := 2; i < n; i++ {
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
	}

	// Convert the slice of integers to a string
	return fmt.Sprint(sequence)
}

func init() {
	fmt.Println("Starting bot...")
}
