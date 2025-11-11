package server

import (
	"embed"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	DB     *gorm.DB
	Log    *zerolog.Logger
}

func New(logger *zerolog.Logger, embedFS *embed.FS) (*Server, error) {
	s := &Server{
		Log: logger,
	}

	// initialize db
	db, err := gorm.Open(sqlite.Open("megafactory.db"), &gorm.Config{})
	if err != nil {
		s.Log.Fatal().Err(err).Msg("unable to initialize database")
		return nil, err
	}
	s.DB = db

	// migrate db
	err = s.migrateDB()
	if err != nil {
		s.Log.Fatal().Err(err).Msg("unable to migrate database")
		return nil, err
	}

	// setup http server
	s.router = gin.New()
	s.router.Use(gin.Recovery())

	// http log middleware
	s.router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		s.Log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status", c.Writer.Status()).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Msg("http request")
	})

	// http template error handling middleware
	s.router.Use(func(c *gin.Context) { // template error handling
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			s.Log.Error().Err(err).Msg("html template rendering failed")
		}
	})

	s.customTemplateFunctions()
	s.loadHTMLFromEmbedFS(embedFS, "web", "shards/*", "pages/*")
	s.setupRoutes()

	return s, nil
}

func (s *Server) Run() {
	s.router.Run()
}
