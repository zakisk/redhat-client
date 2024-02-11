package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/briandowns/spinner"
)

func CreateDefaultSpinner(suffix string, finalMsg string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)

	s.Suffix = " " + suffix
	s.FinalMSG = finalMsg + "\n"
	return s
}

func MakeRequest(method, url string, body *bytes.Buffer, headers map[string]string) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		reqBody = body
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("Unable to make request object: %s", err.Error())
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP Error: Status Code: %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("IO Error: unable to read response body: %s", err.Error())
	}

	return bodyBytes, nil
}
