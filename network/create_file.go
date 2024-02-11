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
	if !utils.IsFileExist(fileName) {
		return nil, fmt.Errorf("There is no such file `%s`\n", fileName)
	}

	response, err := nc.ChecksumFile(fileName)
	if err != nil {
		return nil, err
	}

	exists := response.Metadata["checksum_exists"].(bool)
	mFileName := response.Metadata["file_name"].(string)
	if exists {
		return nil, fmt.Errorf("File is already existed on server with name `%s`\n", mFileName)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %s\n", err.Error())
	}
	defer file.Close()

	writer := multipart.NewWriter(nc.body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to prepare multipart file: %s\n", err.Error())
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, fmt.Errorf("Unable to copy multipart file: %s\n", err.Error())
	}

	writer.Close()

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9254"
	}

	url := fmt.Sprintf("http://localhost:%s/upload_file", port)
	bodyBytes, err := utils.MakeRequest(
		http.MethodPost, url, nc.body,
		map[string]string{"Content-Type": writer.FormDataContentType()})
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s\n", err.Error())
	}

	response = &models.Response{}
	err = json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, fmt.Errorf("IO Error: unable to unmarshal response: %s\n", err.Error())
	}

	if !response.Success {
		return nil, fmt.Errorf("Server Error: %s\n", response.Message)
	} else {
		return response, nil
	}
}
