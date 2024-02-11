package utils

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
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

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("HTTP Error: Resource not found")
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return nil, fmt.Errorf("HTTP Error: Something went wrong on server")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("IO Error: unable to read response body: %s", err.Error())
	}

	return bodyBytes, nil
}

func PrintToTable(header []string, data [][]string) {
	// The tables are formatted to look similar to how it looks in say `kubectl get deployments`
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header) // The header of the table
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // The data in the table
	table.Render()         // Render the table
}

func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}