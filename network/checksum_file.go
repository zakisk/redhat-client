package network

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/zakisk/redhat-client/utils"
	"github.com/zakisk/redhat-server/models"
	"golang.org/x/crypto/blake2b"
)

func (nc *NetworkCaller) ChecksumFile(fileName string) (*models.Response, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Unable to open file: %s\n", err.Error())
	}
	defer file.Close()

	checksum, err := getChecksum(file)
	if err != nil {
		return nil, err
	}

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

	url := fmt.Sprintf("http://localhost:%s/file_exists?checksum=%s", port, checksum)
	bodyBytes, err := utils.MakeRequest(
		http.MethodGet, url, nc.body,
		map[string]string{"Content-Type": writer.FormDataContentType()})
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s\n", err.Error())
	}

	response := &models.Response{}
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

func getChecksum(file *os.File) (string, error) {
	hasher, _ := blake2b.New256(nil)
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("Failed to copy file content\nerror: %s\n", err.Error())
	}

	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash), nil
}
