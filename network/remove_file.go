package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zakisk/redhat-client/utils"
	"github.com/zakisk/redhat-server/models"
)

func (nc *NetworkCaller) RemoveFile(fileName string) (*models.Response, error) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9254"
	}

	url := fmt.Sprintf("http://localhost:%s/remove_file?file_name=%s", port, fileName)
	bodyBytes, err := utils.MakeRequest(http.MethodDelete, url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s", err.Error())
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
