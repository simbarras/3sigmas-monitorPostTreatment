package api

import (
	"github.com/gin-gonic/gin"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api/storage"
)

type Server struct {
	Store *storage.PostgresStore
}

func SetRoutes(app *gin.Engine, prefix string, store *storage.PostgresStore) {
	s := Server{
		Store: store,
	}

	app.GET(prefix+"/action", s.getAction)
	app.POST(prefix+"/action", postAction)
	app.PUT(prefix+"/action", putAction)
	app.DELETE(prefix+"/action", deleteAction)
	app.POST(prefix+"/trigger/bucket/:bucketName", postTriggerBucket)
	app.POST(prefix+"/trigger/action/:actionID", postTriggerAction)
	app.GET(prefix+"/bucket", getBucket)
	app.GET(prefix+"/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
}

// getAction Get all actions
func (s *Server) getAction(ctx *gin.Context) {
	res := s.Store.GetAllActions()
	ctx.JSON(200, res)
}

func postAction(ctx *gin.Context) {
	ctx.String(500, "POST action not implemented yet")
}

func putAction(ctx *gin.Context) {
	ctx.String(500, "PUT action not implemented yet")
}

func deleteAction(ctx *gin.Context) {
	ctx.String(500, "DELETE action not implemented yet")
}

func postTriggerBucket(ctx *gin.Context) {
	ctx.String(500, "POST trigger bucket not implemented yet")
}

func postTriggerAction(ctx *gin.Context) {
	ctx.String(500, "POST trigger action not implemented yet")
}

func getBucket(ctx *gin.Context) {
	ctx.String(500, "GET bucket not implemented yet")
}
