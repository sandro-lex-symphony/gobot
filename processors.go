package gobot

import (
	"fmt"

	snykTool "github.com/sandro-lex-symphony/snykTools"
)

func Hello(sid, data string) {
	err := SendMessage(sid, "<card accent='tempo-bg-color--blue' iconSrc=''><body><h1>filler</h1></body></card>")
	Fatal(err)
}

func Experiment(sid, data string) {
	cmd := GetCommand(data)
	err := SendMessage(sid, cmd)
	Fatal(err)

	fmt.Println("KJFLKJLD")
	result, err := snykTool.GetOrgs()
	Fatal(err)
	var msg string
	msg = "<card  accent='tempo-bg-color--blue' iconSrc=''><body>"
	for _, org := range result.Orgs {
		msg += "<p><b>" + org.Name + "</b> </p>"
	}
	msg += "</body></card>"
	fmt.Println(msg)
	err = SendMessage(sid, msg)
	Fatal(err)
}
