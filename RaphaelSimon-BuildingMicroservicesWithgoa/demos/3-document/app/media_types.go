//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/raphael/gophercon2016/demos/3-document/design
// --out=$(GOPATH)/src/github.com/raphael/gophercon2016/demos/3-document
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// Bottle media type.
//
// Identifier: application/vnd.gophercon.goa.bottle
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

// Validate validates the Bottle media type instance.
func (mt *Bottle) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if len(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, len(mt.Name), 1, true))
	}
	if mt.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 1, true))
	}
	if mt.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, mt.Rating, 5, false))
	}
	if mt.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, mt.Vintage, 1900, true))
	}
	return
}
