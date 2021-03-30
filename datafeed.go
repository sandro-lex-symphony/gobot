package gobot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CreateDatafeedResult struct {
	Id string `json:"id"`
}

type Datafeed struct {
	Events []Event `json:"events`
	AckId  string  `"json:"ackId"`
}

type Event struct {
	Id        string    `json:"id"`
	MessageId string    `json:"messageId"`
	Timestamp int       `json:"timestamp"`
	Type      string    `json:"type"`
	Initiator Initiator `json:"initiator`
	Payload   Payload   `json:"payload"`
}

type Initiator struct {
	User User `json:"user"`
}

type User struct {
	UserId      int    `json:"userId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Username    string `json:"username"`
}

type Payload struct {
	MessageSent MessageSent `json:"messageSent"`
}

type MessageSent struct {
	Message Message `json:"message"`
}

type Message struct {
	MessageId          string `json:"messageId"`
	Timestamp          int    `json:"timestamp"`
	Message            string `json:"message`
	Data               string `json:"data"`
	User               User   `json:"user"`
	Stream             Stream `json:"stream"`
	ExternalRecipients bool   `json:"externalrecipients"`
	OriginalFormat     string `json:"originalFormat"`
	Sid                string `json:"sid"`
}

type Stream struct {
	StreamId   string `json:"streamId"`
	StreamType string `json:"streamType"`
}

func ListDatafeeds() ([]CreateDatafeedResult, error) {
	var result []CreateDatafeedResult
	url := fmt.Sprintf("https://%s:%d/agent/v5/datafeeds", Conf.AgentHost, Conf.AgentPort)
	resp := RequestGet(url)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return result, fmt.Errorf("list Datafeed failed %s", resp.Status)
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return result, err
	}
	resp.Body.Close()

	return result, nil
}

func CreateDatafeed() (string, error) {
	url := fmt.Sprintf("https://%s:%d/agent/v5/datafeeds", Conf.AgentHost, Conf.AgentPort)
	resp := RequestPost(url, nil)

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return "", fmt.Errorf("create Datafeed failed %s", resp.Status)
	}

	var result CreateDatafeedResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return "", err
	}
	resp.Body.Close()
	return result.Id, nil
}

func ReadDatafeed(dfId, ackId string) (Datafeed, error) {
	var result Datafeed
	url := fmt.Sprintf("https://%s:%d/agent/v5/datafeeds/%s/read", Conf.AgentHost, Conf.AgentPort, dfId)

	str := fmt.Sprintf("{\"ackId\": \"%s\"}", ackId)
	var jsonStr = []byte(str)
	resp := RequestPostLP(url, jsonStr)

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return result, fmt.Errorf("read Datafeed failed %s", resp.Status)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	resp.Body.Close() //  must close
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// fmt.Println("XXXXXXXXXXXXXXXXXX")
	// fmt.Println(string(bodyBytes))

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return result, err
	}
	resp.Body.Close()

	return result, nil
}

func DeleteDatafeed(dfId string) error {
	// https://YOUR-AGENT-URL.symphony.com/agent/v5/datafeeds/:datafeedId
	url := fmt.Sprintf("https://%s:%d/agent/v5/datafeeds/%s", Conf.AgentHost, Conf.AgentPort, dfId)
	resp := RequestDelete(url)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("delete Datafeed failed %s", resp.Status)
	}
	resp.Body.Close()
	return nil
}
