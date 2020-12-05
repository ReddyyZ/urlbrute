package core

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
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

func NewRequest(url string, dir *Dir, wg *sync.WaitGroup) *Request {
	codes, _ := sliceAtoi(strings.Split(dir.Code, ",")) // Create int array with codes to show on result

	client := &http.Client{
		Timeout: time.Duration(dir.Timeout) * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil) // Create request
	if err != nil {
		Error(fmt.Sprintf("Error making request: %s", err.Error()))
		(*wg).Done()

		DirFound(strings.Replace(url, dir.Url, "", 1), 404, codes) // Print founded dir

		return &Request{
			url:        url,
			useragent:  dir.Useragent,
			statuscode: 404,
		}
	}

	req.Header.Set("User-Agent", dir.Useragent) // Set user-agent

	res, err := client.Do(req) // Make request

	if err != nil {
		Error(fmt.Sprintf("Error making request: %s", err.Error()))
		fmt.Println(err)
		(*wg).Done()

		DirFound(strings.Replace(url, dir.Url, "", 1), 404, codes) // Print founded dir
		return &Request{
			url:        url,
			useragent:  dir.Useragent,
			statuscode: 404,
		}
	}

	(*wg).Done()

	DirFound(strings.Replace(url, dir.Url, "", 1), res.StatusCode, codes) // Print founded dir
	return &Request{
		url:        url,
		useragent:  dir.Useragent,
		statuscode: res.StatusCode,
		body:       res.Body,
	}
}
