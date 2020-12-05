package cli

import (
	"fmt"

	"github.com/ReddyyZ/urlbrute/brute"
	"github.com/ReddyyZ/urlbrute/core"
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
		AddFlag("code,c", "Filter results by status codes", commando.String, "200,204,301,302,307,401,403").    // Optional
		AddFlag("timeout,t", "Request timeout", commando.Int, 10).                                              // Optional
		AddFlag("useragent,a", "Set User-Agent", commando.String, fmt.Sprintf("urlbrute/%s", config.Version_)). // Optional
		AddFlag("interval,i", "Interval between requests in ms", commando.Int, 300).                            // Optional
		AddFlag("extension,x", "Add extensions to end of each request", commando.String, "no_extension").       // Optional
		AddFlag("url,u", "URL to scan", commando.String, nil).                                                  // Required
		AddFlag("wordlist,w", "Wordlist to test", commando.String, nil).                                        // Required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			url, _ := flags["url"].GetString()
			wordlist, _ := flags["wordlist"].GetString()
			code, _ := flags["code"].GetString()
			timeout, _ := flags["timeout"].GetInt()
			useragent, _ := flags["useragent"].GetString()
			interval, _ := flags["interval"].GetInt()
			extension, _ := flags["extension"].GetString()

			dir := core.NewDir(url, wordlist, useragent, code, timeout, interval, extension)
			brute.DirBrute(dir)
		})

	commando.
		Register("dns").
		SetDescription("Scan for subdomains").
		AddFlag("dnsserver,s", "DNS Servers to resolve", commando.String, "8.8.8.8,8.8.4.4").
		AddFlag("retry,r", "Retry times", commando.Int, 5).
		AddFlag("interval,i", "Interval between requests in ms", commando.Int, 300).
		AddFlag("quiet,q", "Show only domain found", commando.Bool, false).
		AddFlag("ip,a", "Show IP address of domain", commando.Bool, false).
		AddFlag("verbose,v", "Verbose level", commando.Bool, false).
		AddFlag("domain,d", "Domain to scan", commando.String, nil).     // Required
		AddFlag("wordlist,w", "Wordlist to test", commando.String, nil). // Required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			domain, _ := flags["domain"].GetString()
			wordlist, _ := flags["wordlist"].GetString()
			interval, _ := flags["interval"].GetInt()
			retry, _ := flags["retry"].GetInt()
			dnsserver, _ := flags["dnsserver"].GetString()

			quiet, _ := flags["quiet"].GetBool()
			ip, _ := flags["ip"].GetBool()
			verbose, _ := flags["verbose"].GetBool()
			var verboseLevel int

			if quiet {
				verboseLevel = -1
			} else if ip {
				verboseLevel = 1
			} else if verbose {
				verboseLevel = 2
			} else {
				verboseLevel = 0
			}

			dns := core.NewDNS(domain, wordlist, interval, retry, dnsserver, verboseLevel)
			brute.DNSBrute(dns)
		})

	commando.Parse(nil)
}
