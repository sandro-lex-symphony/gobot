package gobot

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func RequestGet(url string) *http.Response {
	timeout := time.Duration(time.Duration(Conf.HTTPTimeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("GET", url, nil)
	request.Header.Set("sessionToken", SessionToken)
	request.Header.Set("keyManagerToken", KeyManagerToken)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func RequestDelete(url string) *http.Response {
	timeout := time.Duration(time.Duration(Conf.HTTPTimeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	request, err := http.NewRequest("DELETE", url, nil)
	request.Header.Set("sessionToken", SessionToken)
	request.Header.Set("keyManagerToken", KeyManagerToken)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func genPost(url string, data []byte, client http.Client) *http.Response {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("sessionToken", SessionToken)
	request.Header.Set("keyManagerToken", KeyManagerToken)

	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func RequestPost(url string, data []byte) *http.Response {
	timeout := time.Duration(time.Duration(Conf.HTTPTimeout) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return genPost(url, data, client)
}

func RequestPostLP(url string, data []byte) *http.Response {
	timeout := time.Duration(time.Duration(Conf.HTTPTimeoutLP) * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	return genPost(url, data, client)
}
