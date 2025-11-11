package main

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	"github.com/henyxia/satisfactory-megafactory/models"
	_ui "github.com/henyxia/satisfactory-megafactory/ui"
)

type config struct {
	DB *gorm.DB
}

func (cfg *config) migrateDB() error {
	cfg.DB.AutoMigrate(
		models.Floor{},
		models.ProductionLine{},
		models.Port{},
		models.Building{},
		models.Item{},
		models.Recipe{},
		models.RecipeInput{},
		models.RecipeOutput{},
	)

	cfg.setData()

	return nil
}

//go:embed web
var webFS embed.FS

func main() {
	// setup logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Satisfactory Megafactory starting")

	cfg := &config{}

	// initialize db
	db, err := gorm.Open(sqlite.Open("megafactory.db"), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to initialize database")
	}
	cfg.DB = db

	// migrate db
	err = cfg.migrateDB()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to migrate database")
	}

	// setup ui controller
	ui := _ui.New(db)

	// setup http server
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)

		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.RequestURI).
			Int("status", c.Writer.Status()).
			Str("ip", c.ClientIP()).
			Dur("latency", latency).
			Msg("http request")
	})
	router.Use(func(c *gin.Context) { // template error handling
		c.Next()
		err := c.Errors.Last()
		if err != nil {
			log.Error().Err(err).Msg("html template rendering failed")
		}
	})

	// custom functions
	router.SetFuncMap(template.FuncMap{
		"mul":   _ui.Mul,
		"add":   _ui.Add,
		"add3":  _ui.Add3,
		"add4":  _ui.Add4,
		"add5":  _ui.Add5,
		"minus": _ui.Minus,
		"div":   _ui.Div,
	})

	// web pages
	router.LoadHTMLGlob("web/pages/*")

	// prepare subfs
	assetFS, _ := fs.Sub(webFS, "web/assets")

	// load static
	router.StaticFS("/assets", http.FS(assetFS))

	router.GET("/", ui.Index)
	router.GET("/floor/:id", ui.FloorView)
	router.GET("/floor/:id/plan", ui.FloorPlan)
	router.GET("/item", ui.ItemList)
	router.GET("/recipes", ui.RecipeList)
	router.Run()
}
