package core

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Request struct {
	url        string
	useragent  string
	headers    string
	statuscode int
	body       io.ReadCloser
}

func NewRequest(url string, dir *Dir, wg *sync.WaitGroup) *Request {
	codes, _ := sliceAtoi(strings.Split(dir.code, ",")) // Create int array with codes to show on result

	client := &http.Client{
		Timeout: time.Duration(dir.timeout) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil) // Create request
	if err != nil {
		Error(fmt.Sprintf("Error making request: %s", err.Error()))
		(*wg).Done()

		DirFound(strings.Replace(url, dir.url, "", 1), 404, codes) // Print founded dir

		return &Request{
			url:        url,
			useragent:  dir.useragent,
			statuscode: 404,
		}
	}

	req.Header.Set("User-Agent", dir.useragent) // Set user-agent

	res, err := client.Do(req) // Make request

	if err != nil {
		Error(fmt.Sprintf("Error making request: %s", err.Error()))
		fmt.Println(err)
		(*wg).Done()

		DirFound(strings.Replace(url, dir.url, "", 1), 404, codes) // Print founded dir
		return &Request{
			url:        url,
			useragent:  dir.useragent,
			statuscode: 404,
		}
	}

	(*wg).Done()

	DirFound(strings.Replace(url, dir.url, "", 1), res.StatusCode, codes) // Print founded dir
	return &Request{
		url:        url,
		useragent:  dir.useragent,
		statuscode: res.StatusCode,
		body:       res.Body,
	}
}
