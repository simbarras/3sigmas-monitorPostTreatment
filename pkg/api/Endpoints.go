package api

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"log"
	"sort"
)

type Server struct {
	worker *Worker
}

func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE")

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
	app.POST(prefix+"/action", s.postAction)
	app.DELETE(prefix+"/action/:id", s.deleteAction)
	app.POST(prefix+"/trigger/bucket/:name", s.postTriggerBucket)
	app.POST(prefix+"/trigger/action/:id", s.postTriggerAction)
	app.GET(prefix+"/bucket", s.getBucket)
	app.GET(prefix+"/function", s.getFunction)
	app.GET(prefix+"/health", func(ctx *gin.Context) {
		ctx.String(200, "OK")
	})
}

// getAction Get all actions
func (s *Server) getAction(ctx *gin.Context) {
	res := s.worker.GetActions()
	sort.Slice(res, func(i, j int) bool {
		if res[i].BucketName != res[j].BucketName {
			return res[i].BucketName < res[j].BucketName
		}
		if res[i].EquationName != res[j].EquationName {
			return res[i].EquationName < res[j].EquationName
		}
		return res[i].ID.String() < res[j].ID.String()
	})
	ctx.JSON(200, res)
}

func (s *Server) postAction(ctx *gin.Context) {
	var actionRequest data.AddActionRequest
	err := ctx.BindJSON(&actionRequest)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Error while parsing request: %s\nResult: %v\n", err, actionRequest)
	}
	action := data.ToAction(actionRequest)
	if action.ID == uuid.Nil {
		log.Printf("Adding action for bucket %s\n", action.BucketName)
		s.worker.AddAction(action)
	} else {
		log.Printf("Updating action %s\n", action.ID)
		s.worker.UpdateAction(action)
	}
	s.getAction(ctx)
}

func (s *Server) deleteAction(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Deleting action %s\n", id)
	s.worker.DeleteAction(id)
	s.getAction(ctx)
}

func (s *Server) postTriggerBucket(ctx *gin.Context) {
	name := ctx.Param("name")
	log.Printf("Triggering bucket %s\n", name)
	err := s.worker.TriggerBucket(name)
	if err != nil {
		sentry.CaptureException(err)
		ctx.String(500, err.Error())
		log.Printf("Error while triggering bucket %s: %s\n", name, err)
	}
	ctx.String(200, "OK")
}

func (s *Server) postTriggerAction(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Printf("Triggering action %s\n", id)
	err := s.worker.TriggerAction(id)
	if err != nil {
		sentry.CaptureException(err)
		ctx.String(500, err.Error())
		log.Printf("Error while triggering action %s: %s\n", id, err)
	}
	ctx.String(200, "OK")
}

func (s *Server) getBucket(ctx *gin.Context) {
	ctx.JSON(200, s.worker.GetBuckets())
}

func (s *Server) getFunction(ctx *gin.Context) {
	ctx.JSON(200, s.worker.GetFunctions())
}
