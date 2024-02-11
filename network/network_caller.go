package network

import (
	"bytes"
	"net/http"
)

type NetworkCaller struct {
	client *http.Client
	body   *bytes.Buffer
}

func NewNetworkCaller(client *http.Client, body *bytes.Buffer) *NetworkCaller {
	return &NetworkCaller{client: client, body: body}
}
