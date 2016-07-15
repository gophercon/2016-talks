package design

import . "github.com/goadesign/goa/design"
import . "github.com/goadesign/goa/design/apidsl"
import . "cellar/design/public" // HL

var CellarPayload = Type("BottlePayload", func() {
	Attribute("bottles", CollectionOf(BottlePayload), "Cellar bottles", func() { // HL
		MinLength(1)
	})
})
