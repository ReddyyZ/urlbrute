package brute

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/ReddyyZ/urlbrute/core"
)

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}

func DirBrute(dir *core.Dir) {
	core.Banner()
	fmt.Printf(`
----------Target----------
Target: %s
Wordlist: %s
--------Req Config--------
User-Agent: %s
Timeout: %d
Interval: %dms
Codes: %s
--------------------------

`, dir.Url, dir.Wordlist, dir.Useragent, dir.Timeout, dir.Interval, dir.Code)

	// Variables declaration
	var wg sync.WaitGroup
	var line string
	results := make(chan *core.Request)

	// Open Wordlist
	file, err := os.Open(dir.Wordlist)
	if err != nil {
		core.Error(fmt.Sprintf("Error opening wordlist: %s", err.Error()))
		return
	}

	// Close on end
	defer file.Close()

	// Create I/O Reader
	reader := bufio.NewReader(file)

	core.Info("Scanning...\n")

	for {
		// Read line
		line, err = reader.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\r", "")

		// Handle read error
		if err != nil && err != io.EOF || line == "" {
			break
		}

		// Format URl
		url := fmt.Sprintf("%s/%s", dir.Url, line)

		// Make a new Request
		go func() {
			wg.Add(1)
			results <- core.NewRequest(url, dir, &wg)
		}()

		// Make requests with defined extensions
		if dir.Extensions != "no_extension" {
			time.Sleep(time.Duration(dir.Interval) * time.Millisecond)

			extensions := strings.Split(dir.Extensions, ",")

			for _, ext := range extensions {
				go func() {
					wg.Add(1)
					results <- core.NewRequest(fmt.Sprintf("%s.%s", url, ext), dir, &wg)
				}()

				time.Sleep(time.Duration(dir.Interval) * time.Millisecond)
			}
		}

		// Sleep
		time.Sleep(time.Duration(dir.Interval) * time.Millisecond)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	// for req := range results {
	// 	if _, i := Find(codes, req.statuscode); i {
	// 		DirFound(strings.Replace(req.url, dir.url, "", 1), req.statuscode)
	// 	}
	// }
}
