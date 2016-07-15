package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// JWT defines a security scheme using JWT.  The scheme uses the "Authorization" header to lookup
// the token.  It also defines then scope "api".
var JWT = JWTSecurity("jwt", func() { // HL
	Header("Authorization") // HL
}) // HL

// Resource jwt uses the JWTSecurity security scheme.
var _ = Resource("jwt", func() {
	Description("This resource uses JWT to secure its endpoints")
	DefaultMedia(SuccessMedia)

	Security(JWT) // HL
	// ...
	// OMIT
	Action("signin", func() {
		Description("Creates a valid JWT")
		Security(SigninBasicAuth)
		Routing(POST("/jwt/signin"))
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(Unauthorized)
	})

	Action("secure", func() {
		Description("This action is secured with the jwt scheme")
		Routing(GET("/jwt"))
		Params(func() {
			Param("fail", Boolean, "Force auth failure via JWT validation middleware")
		})
		Response(OK)
		Response(Unauthorized)
	})

	Action("unsecure", func() {
		Description("This action does not require auth")
		Routing(GET("/jwt/unsecure"))
		NoSecurity() // Override the need for auth
		Response(OK)
	})
})
