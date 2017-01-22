// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/EncryptAny/AditumAllAPI/design
// --out=$(GOPATH)/src/github.com/EncryptAny/AditumAllAPI
// --version=v1.1.0-dirty
//
// API "aditum-all": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import "github.com/goadesign/goa"

// Type for creating a new AP
type newAP struct {
	// Latitude of AP
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// Longitude of AP
	Long *float64 `form:"long,omitempty" json:"long,omitempty" xml:"long,omitempty"`
	// Type of AP
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// Associated Venue ID
	VenueID *int `form:"venueID,omitempty" json:"venueID,omitempty" xml:"venueID,omitempty"`
}

// Validate validates the newAP type instance.
func (ut *newAP) Validate() (err error) {
	if ut.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if ut.Lat == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "lat"))
	}
	if ut.Long == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "long"))
	}
	if ut.Type != nil {
		if !(*ut.Type == "parking" || *ut.Type == "door" || *ut.Type == "ramp") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, *ut.Type, []interface{}{"parking", "door", "ramp"}))
		}
	}
	return
}

// Publicize creates NewAP from newAP
func (ut *newAP) Publicize() *NewAP {
	var pub NewAP
	if ut.Lat != nil {
		pub.Lat = *ut.Lat
	}
	if ut.Long != nil {
		pub.Long = *ut.Long
	}
	if ut.Type != nil {
		pub.Type = *ut.Type
	}
	if ut.VenueID != nil {
		pub.VenueID = ut.VenueID
	}
	return &pub
}

// Type for creating a new AP
type NewAP struct {
	// Latitude of AP
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// Longitude of AP
	Long float64 `form:"long" json:"long" xml:"long"`
	// Type of AP
	Type string `form:"type" json:"type" xml:"type"`
	// Associated Venue ID
	VenueID *int `form:"venueID,omitempty" json:"venueID,omitempty" xml:"venueID,omitempty"`
}

// Validate validates the NewAP type instance.
func (ut *NewAP) Validate() (err error) {
	if ut.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	if !(ut.Type == "parking" || ut.Type == "door" || ut.Type == "ramp") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.type`, ut.Type, []interface{}{"parking", "door", "ramp"}))
	}
	return
}

// Type for creating a new venue
type newVenue struct {
	// Latitude of Venue
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// Longitude of Venue
	Long *float64 `form:"long,omitempty" json:"long,omitempty" xml:"long,omitempty"`
	// Name of Venue
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Radius in meters of the distance in which the geofence is to activate.
	Radius *float64 `form:"radius,omitempty" json:"radius,omitempty" xml:"radius,omitempty"`
}

// Validate validates the newVenue type instance.
func (ut *newVenue) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if ut.Lat == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "lat"))
	}
	if ut.Long == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "long"))
	}
	if ut.Radius == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "radius"))
	}
	if ut.Radius != nil {
		if *ut.Radius < 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.radius`, *ut.Radius, 20, true))
		}
	}
	if ut.Radius != nil {
		if *ut.Radius > 200 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.radius`, *ut.Radius, 200, false))
		}
	}
	return
}

// Publicize creates NewVenue from newVenue
func (ut *newVenue) Publicize() *NewVenue {
	var pub NewVenue
	if ut.Lat != nil {
		pub.Lat = *ut.Lat
	}
	if ut.Long != nil {
		pub.Long = *ut.Long
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Radius != nil {
		pub.Radius = *ut.Radius
	}
	return &pub
}

// Type for creating a new venue
type NewVenue struct {
	// Latitude of Venue
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// Longitude of Venue
	Long float64 `form:"long" json:"long" xml:"long"`
	// Name of Venue
	Name string `form:"name" json:"name" xml:"name"`
	// Radius in meters of the distance in which the geofence is to activate.
	Radius float64 `form:"radius" json:"radius" xml:"radius"`
}

// Validate validates the NewVenue type instance.
func (ut *NewVenue) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if ut.Radius < 20 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.radius`, ut.Radius, 20, true))
	}
	if ut.Radius > 200 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.radius`, ut.Radius, 200, false))
	}
	return
}

// Describes the number of upvotes and downvotes on an item.
type votes struct {
	Downvotes *int `form:"downvotes,omitempty" json:"downvotes,omitempty" xml:"downvotes,omitempty"`
	Upvotes   *int `form:"upvotes,omitempty" json:"upvotes,omitempty" xml:"upvotes,omitempty"`
}

// Validate validates the votes type instance.
func (ut *votes) Validate() (err error) {
	if ut.Upvotes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "upvotes"))
	}
	if ut.Downvotes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "downvotes"))
	}
	return
}

// Publicize creates Votes from votes
func (ut *votes) Publicize() *Votes {
	var pub Votes
	if ut.Downvotes != nil {
		pub.Downvotes = *ut.Downvotes
	}
	if ut.Upvotes != nil {
		pub.Upvotes = *ut.Upvotes
	}
	return &pub
}

// Describes the number of upvotes and downvotes on an item.
type Votes struct {
	Downvotes int `form:"downvotes" json:"downvotes" xml:"downvotes"`
	Upvotes   int `form:"upvotes" json:"upvotes" xml:"upvotes"`
}

// Validate validates the Votes type instance.
func (ut *Votes) Validate() (err error) {

	return
}

// Accessiblity Information for an item
type ai struct {
	// Number of Downvotes for given AI
	Downvotes *int `form:"downvotes,omitempty" json:"downvotes,omitempty" xml:"downvotes,omitempty"`
	// Number of Upvotes for given AI
	Upvotes *int `form:"upvotes,omitempty" json:"upvotes,omitempty" xml:"upvotes,omitempty"`
}

// Validate validates the ai type instance.
func (ut *ai) Validate() (err error) {
	if ut.Upvotes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "upvotes"))
	}
	if ut.Downvotes == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "downvotes"))
	}
	return
}

// Publicize creates Ai from ai
func (ut *ai) Publicize() *Ai {
	var pub Ai
	if ut.Downvotes != nil {
		pub.Downvotes = *ut.Downvotes
	}
	if ut.Upvotes != nil {
		pub.Upvotes = *ut.Upvotes
	}
	return &pub
}

// Accessiblity Information for an item
type Ai struct {
	// Number of Downvotes for given AI
	Downvotes int `form:"downvotes" json:"downvotes" xml:"downvotes"`
	// Number of Upvotes for given AI
	Upvotes int `form:"upvotes" json:"upvotes" xml:"upvotes"`
}

// Validate validates the Ai type instance.
func (ut *Ai) Validate() (err error) {

	return
}