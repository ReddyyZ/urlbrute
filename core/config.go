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
	Url        string
	Wordlist   string
	Useragent  string
	Code       string
	Timeout    int
	Interval   int
	Extensions string
}

type DNS struct {
	Domain    string
	Wordlist  string
	Interval  int
	Retry     int
	DNSServer []string
	Verbose   int
}

func NewConfig(version string, description string) *Config {
	return &Config{
		Version:     fmt.Sprintf("v%s", version),
		Version_:    version,
		Description: description,
	}
}

func NewDir(url string, wordlist string, useragent string, code string, timeout int, interval int, extensions string) *Dir {
	if !strings.Contains(url, "http") {
		url = fmt.Sprintf("http://%s", url)
	}

	return &Dir{
		Url:        url,
		Wordlist:   wordlist,
		Useragent:  useragent,
		Code:       code,
		Timeout:    timeout,
		Interval:   interval,
		Extensions: extensions,
	}
}

func NewDNS(domain string, wordlist string, interval int, retry int, dnsserver string, verbose int) *DNS {
	return &DNS{
		Domain:    domain,
		Wordlist:  wordlist,
		Interval:  interval,
		Retry:     retry,
		DNSServer: strings.Split(dnsserver, ","),
		Verbose:   verbose,
	}
}
