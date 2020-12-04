package core

import (
	"fmt"
	"strings"
)

type Config struct {
	Version     string
	Version_    string
	Description string
}

type Dir struct {
	url       string
	wordlist  string
	useragent string
	code      string
	timeout   int
	interval  int
}

func NewConfig(version string, description string) *Config {
	return &Config{
		Version:     fmt.Sprintf("v%s", version),
		Version_:    version,
		Description: description,
	}
}

func NewDir(url string, wordlist string, useragent string, code string, timeout int, interval int) *Dir {
	if !strings.Contains(url, "http") {
		url = fmt.Sprintf("http://%s", url)
	}

	return &Dir{
		url:       url,
		wordlist:  wordlist,
		useragent: useragent,
		code:      code,
		timeout:   timeout,
		interval:  interval,
	}
}
