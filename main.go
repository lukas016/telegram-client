package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	exitOK  = 0
	exitErr = 1
	// a few messages were succeeded sent
	exitPartialErr = 2
)

// Create a bot with token
func login(token string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(token)
}

// Parse chat ID from command line arguments
func parseChatID(args []string) (chatIDs []int64, err error) {
	chatIDs = make([]int64, len(args))
	err = nil

	for i := 0; i < len(args); i++ {
		chatIDs[i], err = strconv.ParseInt(args[i], 10, 64)
		if err != nil {
			return
		}
	}

	return
}

// Create a message handler to write content of stdout (each line) to someone
func messageWriter(bot *tgbotapi.BotAPI, chatIDs []int64) int {
	var msgTxt string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msgTxt += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		return exitErr
	}

	exitCode := exitOK
	errCode := exitErr
	for _, chatID := range chatIDs {
		msg := tgbotapi.NewMessage(chatID, msgTxt)
		msg.ParseMode = tgbotapi.ModeHTML

		_, err := bot.Send(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error sending message to %d from stdin: '%s'\n", chatID, err.Error())
			exitCode = errCode
		} else {
			errCode = exitPartialErr
			if exitCode != exitOK {
				exitCode = errCode
			}
		}
	}

	return exitCode
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing chat IDs")
		os.Exit(exitErr)
	}

	token := os.Getenv("TLGCLI_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "You need to set TLGCLI_TOKEN in your env")
		os.Exit(exitErr)
	}

	bot, err := login(token)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitErr)
	}

	bot.Debug = false

	ids, err := parseChatID(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(exitErr)
	}

	os.Exit(messageWriter(bot, ids))
}
