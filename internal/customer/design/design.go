package design

import (
	"goa.design/goa/v3/dsl"
)

var _ = dsl.API("jumia", func() {
	dsl.Title("jumia API")
	dsl.Description(`Jumia Application API, You can read and test  the APIs from this documentation`)
	dsl.Version("1.0")
	dsl.HTTP(func() {
		dsl.Consumes("application/json", "application/xml")
		dsl.Path("/customers/v1")
	})
})

// // JWTAuth defines a security scheme that uses JWT tokens.
// var JWTAuth = dsl.JWTSecurity("jwt", func() {
// 	dsl.Description(`Secures endpoint by requiring a valid JWT token retrieved via the signin endpoint.".`)

// })
