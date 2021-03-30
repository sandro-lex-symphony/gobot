package gobot

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func GetStreamId(p Payload) string {
	return p.MessageSent.Message.Stream.StreamId
}

func GetText(p Payload) string {
	r := strings.NewReader(p.MessageSent.Message.Message)
	z := html.NewTokenizer(r)
	var text string
	for {
		tt := z.Next()
		token := z.Token()

		switch {
		case tt == html.ErrorToken:
			return text
		case tt == html.TextToken:
			text += token.Data
		}
	}
}

func IsMentioned(text string) bool {
	if strings.Contains(text, Conf.BotUsername) {
		return true
	}
	return false
}

func GetCommand(text string) string {
	words := strings.Fields(text)
	size := len(words)
	for i, word := range words {
		//fmt.Println(i, " => ", word)
		if word == "@"+Conf.BotUsername {
			fmt.Println("Found mention")
			if i < size-1 {
				return words[i+1]
			}
		}
	}
	return ""
}
