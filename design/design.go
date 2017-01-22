package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("aditum-all", func() {
	Title("AditumAll API")
	Description("The Universal Accessibility API")
	Scheme("http")
	Host("127.0.0.1:8080")
})

var _ = Resource("venue", func() {
	BasePath("/venue")
	DefaultMedia(venueMedia)

	Action("show", func() {
		Description("Get venue by ID")
		Routing(GET("/:venueID"))
		Params(func() {
			Param("venueID", Integer, "Venue ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var venueMedia = MediaType("application/json", func() {
	Description("Fully describes a venue")
	Attributes(func() {
		Attribute("id", Integer, "Unique Venue ID")
		Attribute("name", String, "Name of Venue")
		Attribute("aps", ArrayOf(apMedia), "Accessibility Points of the given location.")
		Attribute("epiPen", aiMedia, "Is filled in if venue is reported to have epi-pens on premisis")
		Attribute("aed", aiMedia, "Exists if venue is reported to contain an Automatic Electronic Defibulator (AED).")
		Attribute("lat", Number, "Latitude of Venue")
		Attribute("long", Number, "Longitude of Venue")
		Required("id", "name")
	})
	View("detailed", func() {
		Attribute("id") // "id" and "name" must be media type attributes
		Attribute("name")
		Attribute("aps", func() {
			View("venue")
		})
		Attribute("epiPen")
		Attribute("aed")
		Attribute("lat")
		Attribute("long")
	})
})

var aiMedia = Type("application/json", func() {
	Description("Accessiblity Information for an item")
	Attributes(func() {
		Attribute("upvotes", Integer, "Number of Upvotes for given AI")
		Attribute("downvotes", Integer, "Number of Downvotes for given AI")
		Required("upvotes", "downvotes")
	})
	View("default", func() {
		Attribute("upvotes")
		Attribute("downvotes")
	})
})

var apMedia = MediaType("application/json", func() {
	Description("Describes an Accessiblity Point in full")
	Attributes(func() {
		Attribute("id", Integer, "Unique AP ID")
		Attribute("venueID", Integer, "Associated Venue")
		Attribute("type", func() {
			Enum("parking", "door", "ramp")
		})
		Attribute("upvotes", Integer, "Number of Upvotes for given AP")
		Attribute("downvotes", Integer, "Number of Downvotes for given AP")
		Attribute("lat", Number, "Latitude of given AP")
		Attribute("long", Number, "Longitude of given AP")
		Required("id", "type", "lat", "long")
	})
	View("venue", func() {
		Attribute("id")
		Attribute("type")
		Attribute("upvotes")
		Attribute("downvotes")
		Attribute("lat")
		Attribute("long")
	})
})
