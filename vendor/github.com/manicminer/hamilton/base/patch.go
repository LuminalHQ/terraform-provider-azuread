package base

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

// PatchHttpRequestInput configures a PATCH request.
type PatchHttpRequestInput struct {
	Body             []byte
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc
	Uri              Uri
}

// GetValidStatusCodes returns a []int of status codes considered valid for a PATCH request.
func (i PatchHttpRequestInput) GetValidStatusCodes() []int {
	return i.ValidStatusCodes
}

// GetValidStatusFunc returns a function used to evaluate whether the response to a PATCH request is considered valid.
func (i PatchHttpRequestInput) GetValidStatusFunc() ValidStatusFunc {
	return i.ValidStatusFunc
}

// Patch performs a PATCH request.
func (c Client) Patch(ctx context.Context, input PatchHttpRequestInput) (*http.Response, int, error) {
	var status int
	url, err := c.buildUri(input.Uri)
	if err != nil {
		return nil, status, fmt.Errorf("unable to make request: %v", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, url, bytes.NewBuffer(input.Body))
	if err != nil {
		return nil, status, err
	}
	resp, status, err := c.performRequest(req, input)
	if err != nil {
		return nil, status, err
	}
	return resp, status, nil
}
