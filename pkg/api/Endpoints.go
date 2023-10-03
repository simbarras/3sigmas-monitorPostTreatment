package api

import "github.com/gin-gonic/gin"

func SetRoutes(app *gin.Engine, prefix string) {
	app.GET(prefix+"/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
	app.GET(prefix+"/action", func(ctx *gin.Context) {
		ctx.String(500, "GET action not implemented yet")
	})
	app.POST(prefix+"/action", func(ctx *gin.Context) {
		ctx.String(500, "POST action not implemented yet")
	})
	app.PUT(prefix+"/action", func(ctx *gin.Context) {
		ctx.String(500, "PUT action not implemented yet")
	})
	app.DELETE(prefix+"/action", func(ctx *gin.Context) {
		ctx.String(500, "DELETE action not implemented yet")
	})
	app.POST(prefix+"/trigger/bucket/:bucketName", func(ctx *gin.Context) {
		ctx.String(500, "POST trigger bucket not implemented yet")
	})
	app.POST(prefix+"/trigger/action/:actionID", func(ctx *gin.Context) {
		ctx.String(500, "POST trigger action not implemented yet")
	})
	app.GET(prefix+"/bucket", func(ctx *gin.Context) {
		ctx.String(500, "GET bucket not implemented yet")
	})
}
