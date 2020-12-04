package core

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
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

func DirBrute(dir *Dir) {
	Banner()
	fmt.Printf(`
----------Target----------
Target: %s
Wordlist: %s
--------Req Config--------
User-Agent: %s
Timeout: %d
Codes: %s
--------------------------

`, dir.url, dir.wordlist, dir.useragent, dir.timeout, dir.code)

	// Variables declaration
	var wg sync.WaitGroup
	var line string
	results := make(chan *Request)

	// Open Wordlist
	file, err := os.Open(dir.wordlist)
	if err != nil {
		Error(fmt.Sprintf("Error opening wordlist: %s", err.Error()))
		return
	}

	// Close on end
	defer file.Close()

	// Create I/O Reader
	reader := bufio.NewReader(file)

	Info("Scanning...\n")

	for {
		// Read line
		line, err = reader.ReadString('\n')
		line = strings.ReplaceAll(line, "\n", "")
		line = strings.ReplaceAll(line, "\r", "")

		// Handle read error
		if err != nil && err != io.EOF {
			break
		}

		// Format URl
		url := fmt.Sprintf("%s/%s", dir.url, line)

		// Make a new Request
		go func() {
			wg.Add(1)
			results <- NewRequest(url, dir, &wg)
		}()

		// Sleep
		time.Sleep(time.Duration(dir.interval) * time.Millisecond)
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
