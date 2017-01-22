// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/EncryptAny/AditumAllAPI/design
// --out=$(GOPATH)/src/github.com/EncryptAny/AditumAllAPI
// --version=v1.1.0-dirty
//
// API "aditum-all": ap Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"bytes"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// CreateApPath computes a request path to the create action of ap.
func CreateApPath() string {

	return fmt.Sprintf("/ap")
}

// Create an AP
func (c *Client) CreateAp(ctx context.Context, path string, payload *NewAP, contentType string) (*http.Response, error) {
	req, err := c.NewCreateApRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateApRequest create the request corresponding to the create action endpoint of the ap resource.
func (c *Client) NewCreateApRequest(ctx context.Context, path string, payload *NewAP, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType != "*/*" {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// DownvoteApPath computes a request path to the downvote action of ap.
func DownvoteApPath(apID int) string {
	param0 := strconv.Itoa(apID)

	return fmt.Sprintf("/ap/%s/upvote", param0)
}

// Downvote a particular AP
func (c *Client) DownvoteAp(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDownvoteApRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDownvoteApRequest create the request corresponding to the downvote action endpoint of the ap resource.
func (c *Client) NewDownvoteApRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpvoteApPath computes a request path to the upvote action of ap.
func UpvoteApPath(apID int) string {
	param0 := strconv.Itoa(apID)

	return fmt.Sprintf("/ap/%s/downvote", param0)
}

// Upvote a particular AP
func (c *Client) UpvoteAp(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewUpvoteApRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpvoteApRequest create the request corresponding to the upvote action endpoint of the ap resource.
func (c *Client) NewUpvoteApRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
