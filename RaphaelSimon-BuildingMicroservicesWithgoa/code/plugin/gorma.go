package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("MyAPIStorage", func() {
	Store("postgres", gorma.Postgres, func() {
		Model("User", func() {
			BuildsFrom(func() { Payload("user", "create") })
			RendersTo(User)
			Field("id", gorma.Integer, func() {
				PrimaryKey()
				Description("This is the User Model PK field")
			})
		})
	})
})
