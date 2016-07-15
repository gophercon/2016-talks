package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"

var _ = API("cellar", func() {
	Description("The wine review service")
	Host("localhost:8080")
})

var BottlePayload = Type("BottlePayload", func() {
	Description("BottlePayload is the type used to create bottles")

	Attribute("name", String, "Name of bottle", func() {
		MinLength(1)
	})
	Attribute("vintage", Integer, "Vintage of bottle", func() {
		Minimum(1900)
	})
	Attribute("rating", Integer, "Rating of bottle", func() {
		Minimum(1)
		Maximum(5)
	})
	Required("name", "vintage", "rating")
})

var BottleMedia = MediaType("application/vnd.gophercon.goa.bottle", func() {
	TypeName("bottle")
	Reference(BottlePayload)

	Attributes(func() {
		Attribute("ID", Integer, "Unique bottle ID")
		Attribute("name")
		Attribute("vintage")
		Attribute("rating")
		Required("ID", "name", "vintage", "rating")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("vintage")
		Attribute("rating")
	})
})

var _ = Resource("bottle", func() {
	Description("A wine bottle")
	BasePath("/bottles")

	Action("create", func() {
		Description("creates a bottle")
		Routing(POST("/"))
		Payload(BottlePayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a bottle")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, BottleMedia)
	})
})
