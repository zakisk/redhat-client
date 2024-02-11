package utils

import (
	"bytes"
	"fmt"
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

func MakeRequest(method, url string, body *bytes.Buffer, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("Unable to make request object: %s", err.Error())
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s", err.Error())
	}

	return resp, nil
}
