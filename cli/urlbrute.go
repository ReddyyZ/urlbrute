package cli

import (
	"fmt"
	"urlbrute/core"

	"github.com/thatisuday/commando"
)

func Run(config *core.Config) {
	commando.
		SetExecutableName("urlbrute").
		SetVersion(config.Version).
		SetDescription(config.Description)

	commando.
		Register("dir").
		SetDescription("Scan for diretories on website").
		AddFlag("url,u", "URL to scan", commando.String, nil).                                                  // Required
		AddFlag("wordlist,w", "Wordlist to test", commando.String, nil).                                        // Required
		AddFlag("code,c", "Filter results by status codes", commando.String, "200,204,301,302,307,401,403").    // Optional
		AddFlag("timeout,t", "Request timeout", commando.Int, 10).                                              // Optional
		AddFlag("useragent,a", "Set User-Agent", commando.String, fmt.Sprintf("urlbrute/%s", config.Version_)). // Optional
		AddFlag("interval,i", "Interval between requests in ms", commando.Int, 300).                            // Optional
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			url, _ := flags["url"].GetString()
			wordlist, _ := flags["wordlist"].GetString()
			code, _ := flags["code"].GetString()
			timeout, _ := flags["timeout"].GetInt()
			useragent, _ := flags["useragent"].GetString()
			interval, _ := flags["interval"].GetInt()

			dir := core.NewDir(url, wordlist, useragent, code, timeout, interval)
			core.DirBrute(dir)
		})

	commando.Parse(nil)
}
