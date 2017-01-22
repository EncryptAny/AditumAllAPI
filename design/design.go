package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("aditum-all", func() {
	Title("AditumAll API")
	Description("The Universal Accessibility API")
	Scheme("http")
	Host("066727de.ngrok.io")
})

var votes = Type("Votes", func() {
	Description("Describes the number of upvotes and downvotes on an item.")
	Attribute("upvotes", Integer)
	Attribute("downvotes", Integer)
	Required("upvotes", "downvotes")
})

// AP //
var apEnum = func() {
	Enum("parking", "door", "ramp")
}

var apAddType = Type("NewAP", func() {
	Description("Type for creating a new AP")
	Attribute("venueID", Integer, "Associated Venue ID")
	Attribute("type", String, "Type of AP", apEnum)
	Attribute("lat", Number, "Latitude of AP")
	Attribute("long", Number, "Longitude of AP")
	Required("type", "lat", "long")
})

var apMedia = MediaType("ap", func() {
	Description("Describes an Accessiblity Point in full")
	Attributes(func() {
		Attribute("id", Integer, "Unique AP ID")
		Attribute("venueID", Integer, "Associated Venue")
		Attribute("type", String, "Type of Accessiblity Point", apEnum)
		Attribute("upvotes", Integer, "Number of Upvotes for given AP")
		Attribute("downvotes", Integer, "Number of Downvotes for given AP")
		Attribute("lat", Number, "Latitude of given AP")
		Attribute("long", Number, "Longitude of given AP")
		Required("id", "type", "lat", "long")
	})
	View("default", func() {
		Attribute("id")
		Attribute("venueID")
		Attribute("type")
		Attribute("upvotes")
		Attribute("downvotes")
		Attribute("lat")
		Attribute("long")
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

var _ = Resource("ap", func() {
	BasePath("/ap")

	Action("create", func() {
		Description("Create an AP")
		Routing(POST("/"))
		Payload(apAddType)
		Response(OK, apMedia)
		Response(BadRequest)
	})

	Action("upvote", func() {
		Description("Upvote a particular AP")
		Routing(POST("/:apID/downvote"))
		Params(func() {
			Param("apID", Integer, "AP ID")
		})
		Response(OK, votes)
		Response(NotFound)
	})
	Action("downvote", func() {
		Description("Downvote a particular AP")
		Routing(POST("/:apID/upvote"))
		Params(func() {
			Param("apID", Integer, "AP ID")
		})
		Response(OK, votes)
		Response(NotFound)
	})
})

// Venue //

var venueMedia = MediaType("venue", func() {
	Description("Fully describes a venue")
	Attributes(func() {
		Attribute("id", Integer, "Unique Venue ID")
		Attribute("name", String, "Name of Venue")
		Attribute("aps", ArrayOf(apMedia), "Accessibility Points of the given location.")
		Attribute("epiPen", aiType, "Is filled in if venue is reported to have epi-pens on premisis")
		Attribute("aed", aiType, "Exists if venue is reported to contain an Automatic Electronic Defibulator (AED).")
		Attribute("bathroom", aiType, "Exists if venue is reported to contain an accessible bathroom")
		Attribute("lat", Number, "Latitude of Venue")
		Attribute("long", Number, "Longitude of Venue")
		Attribute("radius", Integer, "Distance in meters for the notification distance")
		Required("id", "name")
	})
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("lat")
		Attribute("long")
		Attribute("radius")
	})
	View("detailed", func() {
		Attribute("id")
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

var venueAddType = Type("NewVenue", func() {
	Description("Type for creating a new venue")
	Attribute("name", String, "Name of Venue")
	Attribute("lat", Number, "Latitude of Venue")
	Attribute("long", Number, "Longitude of Venue")
	Attribute("radius", Number,
		"Radius in meters of the distance in which the geofence is to activate.",
		func() {
			Minimum(20)
			Maximum(200)
		},
	)
	Required("name", "lat", "long", "radius")
})

var aiEnum = func() {
	Enum("epiPen", "bathroom", "aed")
}

var _ = Resource("venue", func() {
	BasePath("/venue")

	Action("all", func() {
		Description("Get all venues")
		Routing(GET("/"))
		Response(OK, ArrayOf(venueMedia))
	})

	Action("show", func() {
		Description("Get venue by ID")
		Routing(GET("/:venueID"))
		Params(func() {
			Param("venueID", Integer, "Venue ID")
		})
		Response(OK, venueMedia)
		Response(NotFound)
	})

	Action("new", func() {
		Description("Create a new venue")
		Routing(POST("/"))
		Payload(venueAddType)
		Response(OK, venueMedia)
		Response(BadRequest)
	})

	Action("upvoteAI", func() {
		Description("Upvote an existing AI for a venue")
		Routing(POST("/:venueID/:aiType/upvote"))
		Params(func() {
			Param("venueID", Integer, "Venue ID")
			Param("aiType", String, "Type of AI", aiEnum)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("downvoteAI", func() {
		Description("Downvote an existing AI for a venue")
		Routing(POST("/:venueID/:aiType/downvote"))
		Params(func() {
			Param("venueID", Integer, "Venue ID")
			Param("aiType", String, "Type of AI", aiEnum)
		})
		Response(OK)
		Response(NotFound)
	})

	Action("newAI", func() {
		Description("Create a new AI for a venue")
		Routing(POST("/:venueID/:aiType"))
		Params(func() {
			Param("venueID", Integer, "Venue ID")
			Param("aiType", String, "Type of AI to Create", aiEnum)
		})
		Response(Created, aiType)
		Response(Conflict)
	})
})

// AI //

var aiType = Type("ai", func() {
	Description("Accessiblity Information for an item")
	Attribute("upvotes", Integer, "Number of Upvotes for given AI")
	Attribute("downvotes", Integer, "Number of Downvotes for given AI")
	Required("upvotes", "downvotes")
})

// var _ = Resource("ai", func() {
// 	BasePath("/ai")
//
// 	Action("create", func() {
// 		Description("Create an AP")
// 		Routing(POST("/:apType"))
// 		Params(func() {
// 			Param("apType", String, "Type of AP Point", func() {
// 				Enum("parking", "door", "ramp")
// 			})
// 		})
// 		Payload(apAddType)
// 		Response(OK, apMedia)
// 		Response(BadRequest)
// 	})
//
// 	Action("upvote", func() {
// 		Description("Upvote a particular AP")
// 		Routing(POST("/:apID/downvote"))
// 		Params(func() {
// 			Param("apID", Integer, "AP ID")
// 		})
// 		Response(OK, votes)
// 		Response(NotFound)
// 	})
// 	Action("upvote", func() {
// 		Description("Downvote a particular AP")
// 		Routing(POST("/:apID/upvote"))
// 		Params(func() {
// 			Param("apID", Integer, "AP ID")
// 		})
// 		Response(OK, votes)
// 		Response(NotFound)
// 	})
// })
