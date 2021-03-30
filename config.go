package gobot

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
)

type Config struct {
	PodURL            string `json:"podURL"`
	SessionAuthHost   string `json:"sessionAuthHost"`
	SessionAuthPort   int    `json:"sessionAuthPort"`
	KeyAuthHost       string `json:"keyAuthHost"`
	KeyAuthPort       int    `json:"keyAuthPort"`
	AgentHost         string `json:"AgentHost"`
	AgentPort         int    `json:"AgentPort"`
	BotPrivateKeyPath string `json:"botPrivateKeyPath"`
	BotPrivateKeyName string `json:"botPrivateKeyName"`
	BotUsername       string `json:"botUsername"`
	HTTPTimeout       int    `json:"HTTPTimeout"`
	HTTPTimeoutLP     int    `json:"HTTPTimeoutLP"`
}

var Conf Config

func ParseConf() error {
	//var config Config

	jsonFile, err := os.Open("conf/config.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &Conf)

	err = checkConf(Conf)
	return err
}

func checkConf(c Config) error {
	// check PodURL
	if c.PodURL == "" {
		return errors.New("empty podURL")
	}
	// parse URL
	_, err := url.ParseRequestURI(c.PodURL)
	if err != nil {
		return errors.New("invalid PodURL")
	}

	// check sessionAuth
	if !validHost(c.SessionAuthHost) {
		return errors.New("invalid sesion auth host")
	}
	if !validPort(c.SessionAuthPort) {
		return errors.New("invalid sesion auth port")
	}

	// check key Manager
	if !validHost(c.KeyAuthHost) {
		return errors.New("invalid Key Manager host")
	}
	if !validPort(c.KeyAuthPort) {
		return errors.New("invalid Key Manager port")
	}

	// check Agent
	if !validHost(c.AgentHost) {
		return errors.New("invalid Agent host")
	}
	if !validPort(c.AgentPort) {
		return errors.New("invalid Agent port")
	}

	// check private key
	if c.BotPrivateKeyName == "" || c.BotPrivateKeyPath == "" {
		return errors.New("Invalid Private key")
	}

	// check botname
	if c.BotUsername == "" {
		return errors.New("Invalid bot Username")
	}

	// defaul http timeout
	if c.HTTPTimeout == 0 {
		c.HTTPTimeout = 5
	}

	// defaul Long Poll http timeout
	if c.HTTPTimeoutLP == 0 {
		c.HTTPTimeoutLP = 30
	}

	return nil
}

func validHost(host string) bool {
	host = strings.Trim(host, " ")

	re, _ := regexp.Compile(`^(([a-zA-Z0-9]|[a-zA-Z0-9][a-zA-Z0-9\-]*[a-zA-Z0-9])\.)*([A-Za-z0-9]|[A-Za-z0-9][A-Za-z0-9\-]*[A-Za-z0-9])$`)
	if re.MatchString(host) {
		return true
	}
	return false
}

func validPort(p int) bool {
	if p > 1 && p < 65536 {
		return true
	}
	return false
}
