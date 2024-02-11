package network

import (
	"bytes"
)

type NetworkCaller struct {
	body   *bytes.Buffer
}

func NewNetworkCaller(body *bytes.Buffer) *NetworkCaller {
	return &NetworkCaller{body: body}
}
