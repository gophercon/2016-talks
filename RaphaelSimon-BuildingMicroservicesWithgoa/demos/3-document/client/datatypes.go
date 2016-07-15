//************************************************************************//
// User Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/raphael/gophercon2016/demos/3-document/design
// --out=$(GOPATH)/src/github.com/raphael/gophercon2016/demos/3-document
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import "net/http"

// Bottle media type.
type Bottle struct {
	// Unique bottle ID
	ID int `json:"ID" xml:"ID" form:"ID"`
	// Name of bottle
	Name string `json:"name" xml:"name" form:"name"`
	// Rating of bottle
	Rating int `json:"rating" xml:"rating" form:"rating"`
	// Vintage of bottle
	Vintage int `json:"vintage" xml:"vintage" form:"vintage"`
}

// DecodeBottle decodes the Bottle instance encoded in resp body.
func (c *Client) DecodeBottle(resp *http.Response) (*Bottle, error) {
	var decoded Bottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
