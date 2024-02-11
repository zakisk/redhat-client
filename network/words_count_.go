package network

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/zakisk/redhat-client/utils"
	"github.com/zakisk/redhat-server/models"
)

func (nc *NetworkCaller) GetWordsCount() (*models.WordCountResponse, error) {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "9254"
	}

	url := fmt.Sprintf("http://localhost:%s/count_all_words", port)
	bodyBytes, err := utils.MakeRequest(http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to call api: %s", err.Error())
	}

	response := &models.WordCountResponse{}
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
