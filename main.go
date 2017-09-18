package main

import (
	"net/http"

	"github.com/DeanThompson/ginpprof"
	"github.com/Verifier/emailVerifier/constants"
	"github.com/Verifier/emailVerifier/handlers"
	"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/gin"
	newrelic "github.com/newrelic/go-agent"
	"path/filepath"
	"os"
	"fmt"
)

const (
	apiBaseURL = "/email-verifier/api/1.0"
)

func main() {

	appKey := constants.Env.NewrelicAppKey   //newrelic app key
	appName := constants.Env.NewrelicAppName //newrelic app name

	router := gin.New() // Create a gin router without any default middleware

	// Add go-agent middleware to send profiling info to newrelic
	router.Use(newrelicMiddleware(appKey, appName))
	// Add gorelic middleware to send extra metrics to newrelic
	gorelic.InitNewrelicAgent(appKey, appName, false)
	router.Use(gorelic.Handler)
	// Add recovery (crash-free) middleware
	router.Use(gin.Recovery())
	// Add logger middleware only when running in debug mode (GIN_MODE env variable)
	if gin.Mode() == gin.DebugMode {
		router.Use(gin.Logger())
	}

	//Loading html files
	//router.LoadHTMLGlob(filepath.Join(os.Getenv("TMPL_DIR"), "*"))
	fmt.Println(filepath.Join(filepath.Join(os.Getenv("TEMPL_DIR"), "css")))
	router.LoadHTMLGlob(filepath.Join(os.Getenv("TEMPL_DIR"), "*.html"))
	router.Static("/css", filepath.Join(filepath.Join(os.Getenv("TEMPL_DIR"), "css")))
	router.Static("/images", filepath.Join(filepath.Join(os.Getenv("TEMPL_DIR"), "images")))
	router.Static("/js", filepath.Join(filepath.Join(os.Getenv("TEMPL_DIR"), "js")))
	router.Static("/fonts", filepath.Join(filepath.Join(os.Getenv("TEMPL_DIR"), "fonts")))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//emailVerifierClient api to verify list of emails
	emailVerifierClient := router.Group(apiBaseURL + "/emails")
	{
		emailVerifierClient.POST("/verify", handlers.VerifyEmails)
		emailVerifierClient.GET( "/index.html", handlers.GetIndexPage)
	}


	// Wrapping with pprof if running in debug mode
	if gin.Mode() == gin.DebugMode {
		ginpprof.Wrapper(router)
	}

	router.Run(":" + constants.Env.Port)
}

func newrelicMiddleware(appKey string, appName string) gin.HandlerFunc {

	if appName == "" || appKey == "" {
		return func(c *gin.Context) {}
	}

	config := newrelic.NewConfig(appName, appKey)
	app, _ := newrelic.NewApplication(config)

	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path, c.Writer, c.Request)
		defer txn.End()
		c.Next()
	}
}
