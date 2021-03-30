package gobot

import (
	"fmt"
	"net/http"
)

// create message with stream id
// https://YOUR-AGENT-URL.symphony.com/agent/v4/stream/:sid/message/create

func SendMessage(sid, msg string) error {
	url := fmt.Sprintf("https://%s:%d/agent/v4/stream/%s/message/create", Conf.AgentHost, Conf.AgentPort, sid)
	str := fmt.Sprintf("{\"message\": \"<messageML>%s</messageML>\"}", msg)
	var jsonStr = []byte(str)
	resp := RequestPost(url, jsonStr)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("read Datafeed failed %s", resp.Status)
	}
	return nil
}
