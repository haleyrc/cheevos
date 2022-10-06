package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	C *http.Client
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpResponse struct{}

func (c *Client) SignUp(ctx context.Context, username, password string) error {
	req := SignUpRequest{
		Username: username,
		Password: password,
	}

	var resp SignUpResponse
	if err := c.Post(ctx, "/auth/SignUp", req, &resp); err != nil {
		return fmt.Errorf("sign up failed: %w", err)
	}

	return nil
}

func (c *Client) Post(ctx context.Context, path string, req interface{}, res interface{}) error {
	url := "http://localhost:8080" + path

	body, err := marshal(req)
	if err != nil {
		return fmt.Errorf("post failed: %w", err)
	}

	resp, err := c.C.Post(url, "application/json", body)
	if err != nil {
		return fmt.Errorf("post failed: %w", err)
	}
	defer resp.Body.Close()

	if err := unmarshal(resp.Body, res); err != nil {
		return fmt.Errorf("post failed: %w", err)
	}

	return nil
}

func marshal(data interface{}) (*bytes.Reader, error) {
	bs, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal failed: %w", err)
	}

	return bytes.NewReader(bs), nil
}

func unmarshal(r io.Reader, dest interface{}) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	if err := json.Unmarshal(bs, dest); err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	return nil
}
