//************************************************************************//
// API "cellar": Application Resource Href Factories
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/raphael/gophercon2016/demos/2-implement/design
// --out=$(GOPATH)/src/github.com/raphael/gophercon2016/demos/2-implement
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "fmt"

// BottleHref returns the resource href.
func BottleHref(id interface{}) string {
	return fmt.Sprintf("/bottles/%v", id)
}
