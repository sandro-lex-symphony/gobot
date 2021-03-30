package gobot

import (
	"fmt"
	"log"
)

type TokenResult struct {
	Token string
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	JWT             string
	SessionToken    string
	KeyManagerToken string
)

func Init() {
	err := ParseConf()
	Fatal(err)

	key := Conf.BotPrivateKeyPath + Conf.BotPrivateKeyName
	JWT = GenAuthJWT(key, Conf.BotUsername)

	url := fmt.Sprintf("https://%s:%d/login/pubkey/authenticate", Conf.SessionAuthHost, Conf.SessionAuthPort)
	SessionToken = Auth(JWT, url)

	url = fmt.Sprintf("https://%s:%d/relay/pubkey/authenticate", Conf.KeyAuthHost, Conf.KeyAuthPort)
	KeyManagerToken = Auth(JWT, url)
}

type Processor func(string, string)

func Loop(f Processor) {
	// create datafeed
	// loop reading datafeed
	df, err := CreateDatafeed()
	Fatal(err)
	var ackId string

	for {
		datafeed, err := ReadDatafeed(df, ackId)
		Fatal(err)
		//fmt.Printf("%v\n", datafeed)
		ackId = datafeed.AckId
		var data string
		if len(datafeed.Events) > 0 {
			// process message
			sid := GetStreamId(datafeed.Events[0].Payload)
			// check if @mentioned
			for _, e := range datafeed.Events {
				if e.Type == "MESSAGESENT" {
					data = GetText(e.Payload)
				}
				//if e.Type == "USERJOINEDROOM"
			}
			if IsMentioned(data) {
				//SendMessage(sid, "XXXX")
				f(sid, data)
			}
		}

	}
}
