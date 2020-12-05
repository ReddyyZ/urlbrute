package brute

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/ReddyyZ/urlbrute/core"
	"github.com/bogdanovich/dns_resolver"
)

func Resolve(url string, domain string, resolver *dns_resolver.DnsResolver, wg *sync.WaitGroup, verbose_level int) {
	if verbose_level == 2 {
		core.Info(fmt.Sprintf("Testing: %s", url))
	}

	// Resolve DNS
	ip, err := resolver.LookupHost(url)
	if err != nil {
		// if verbose_level == 2 {
		// 	fmt.Println("Erro:", url, "-", err.Error())
		// }
		wg.Done()
		return
	}

	// Print DNS founded
	core.DNSFound(url, domain, ip[0].String(), verbose_level)
	wg.Done()
}

func DNSBrute(dns *core.DNS) {
	if dns.Verbose != -1 {
		core.Banner()
		fmt.Printf(`
----------Target----------
Target: %s
Wordlist: %s
--------DNS Config--------
Interval: %dms
Retry: %d
DNSServer: %s
--------------------------

`, dns.Domain, dns.Wordlist, dns.Interval, dns.Retry, dns.DNSServer)
	}

	// Create resolver
	resolver := dns_resolver.New(dns.DNSServer)
	resolver.RetryTimes = dns.Retry

	// Declare variables
	var wg sync.WaitGroup
	var line string

	// Open wordlist
	fd, err := os.Open(dns.Wordlist)
	if err != nil {
		core.Error(fmt.Sprintf("Error opening wordlist: %s", err.Error()))
		return
	}
	defer fd.Close()

	// Create I/O Reader
	reader := bufio.NewReader(fd)

	if dns.Verbose != -1 {
		core.Info("Scanning...\n")
	}

	for {
		// Read a line
		line, err = reader.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\r", "")

		// Handle read error
		if err != nil && err != io.EOF || line == "" {
			break
		}

		// Format URL
		url := fmt.Sprintf("%s.%s", line, dns.Domain)

		// Resolve DNS func
		go func() {
			wg.Add(1)
			Resolve(url, line, resolver, &wg, dns.Verbose)
		}()

		// Sleep
		time.Sleep(time.Duration(dns.Interval) * time.Millisecond)
	}

	// Wait for threads to finish
	go func() {
		wg.Wait()
	}()
}
