//************************************************************************//
// API "cellar": Application User Types
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

import "github.com/goadesign/goa"

// BottlePayload is the type used to create bottles
type bottlePayload struct {
	// Name of bottle
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	// Rating of bottle
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty" form:"rating,omitempty"`
	// Vintage of bottle
	Vintage *int `json:"vintage,omitempty" xml:"vintage,omitempty" form:"vintage,omitempty"`
}

// Validate validates the bottlePayload type instance.
func (ut *bottlePayload) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "vintage"))
	}
	if ut.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rating"))
	}

	if ut.Name != nil {
		if len(*ut.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, *ut.Name, len(*ut.Name), 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 1, true))
		}
	}
	if ut.Rating != nil {
		if *ut.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, *ut.Rating, 5, false))
		}
	}
	if ut.Vintage != nil {
		if *ut.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, *ut.Vintage, 1900, true))
		}
	}
	return
}

// Publicize creates BottlePayload from bottlePayload
func (ut *bottlePayload) Publicize() *BottlePayload {
	var pub BottlePayload
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Rating != nil {
		pub.Rating = *ut.Rating
	}
	if ut.Vintage != nil {
		pub.Vintage = *ut.Vintage
	}
	return &pub
}

// BottlePayload is the type used to create bottles
type BottlePayload struct {
	// Name of bottle
	Name string `json:"name" xml:"name" form:"name"`
	// Rating of bottle
	Rating int `json:"rating" xml:"rating" form:"rating"`
	// Vintage of bottle
	Vintage int `json:"vintage" xml:"vintage" form:"vintage"`
}

// Validate validates the BottlePayload type instance.
func (ut *BottlePayload) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if len(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, ut.Name, len(ut.Name), 1, true))
	}
	if ut.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, ut.Rating, 1, true))
	}
	if ut.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.rating`, ut.Rating, 5, false))
	}
	if ut.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.vintage`, ut.Vintage, 1900, true))
	}
	return
}
