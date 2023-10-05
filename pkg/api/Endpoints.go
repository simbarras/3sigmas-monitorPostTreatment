package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	worker *Worker
}

func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func SetRoutes(app *gin.Engine, prefix string, worker *Worker) {
	s := Server{
		worker: worker,
	}

	app.GET(prefix+"/action", s.getAction)
	app.POST(prefix+"/action", postAction)
	app.PUT(prefix+"/action", putAction)
	app.DELETE(prefix+"/action", deleteAction)
	app.POST(prefix+"/trigger/bucket/:bucketName", postTriggerBucket)
	app.POST(prefix+"/trigger/action/:actionID", postTriggerAction)
	app.GET(prefix+"/bucket", s.getBucket)
	app.GET(prefix+"/function", s.getFunction)
	app.GET(prefix+"/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
}

// getAction Get all actions
func (s *Server) getAction(ctx *gin.Context) {
	res := s.worker.GetActions()
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

func (s *Server) getBucket(ctx *gin.Context) {
	ctx.JSON(200, s.worker.GetBuckets())
}

func (s *Server) getFunction(ctx *gin.Context) {
	ctx.JSON(200, s.worker.GetFunctions())
}
