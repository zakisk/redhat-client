package network

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/zakisk/redhat-client/utils"
	"github.com/zakisk/redhat-server/models"
)

func (nc *NetworkCaller) CreateFile(fileName string) (*models.Response, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %s", err.Error())
	}
	defer file.Close()

	writer := multipart.NewWriter(nc.body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to prepare multipart file: %s", err.Error())
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("Unable to copy multipart file: %s", err.Error())
	}

	writer.Close()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9254"
	}

	url := fmt.Sprintf("http://localhost:%s/upload_file", port)
	resp, err := utils.MakeRequest(
		http.MethodPost, url, nc.body,
		map[string]string{"Content-Type": writer.FormDataContentType()})
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

	response := &models.Response{}
	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, fmt.Errorf("IO Error: unable to unmarshal response: %s", err.Error())
	}

	if !response.Success {
		return nil, fmt.Errorf("Server Error: %s", response.Message)
	} else {
		return response, nil
	}
}
