package main

import (
	"fmt"
	"log"
	"time"

	"vector-tile/src/config"
	"vector-tile/src/controller"

	"vector-tile/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/logrusorgru/aurora"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	fmt.Println("Using timezone:", aurora.Green(time.Now().Location().String()))
}

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
// @description E.g. Bearer Your.Token
func main() {
	appPort := ":" + config.APP_PORT
	if !config.DEBUG {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println(aurora.Red("DEBUG"))
	}
	router := gin.Default()

	router.Use(gzip.Gzip(gzip.BestSpeed))
	router.Use(cors.Default())

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.StaticFile("/rapidoc.html", "docs/rapidoc.html")

	apiV1 := router.Group(basePath)
	apiV1.GET("/ping", func(c *gin.Context) { c.JSON(200, gin.H{"message": "Imusegipo"}) })

	//* ------------------------ REGISTER CRUD CONTROLLER ------------------------ */
	controller.NewMvtController(apiV1)

	//? API DOC
	log.Println(aurora.Green(fmt.Sprintf("http://localhost%s/swagger/index.html", appPort)))
	log.Println(aurora.Yellow(fmt.Sprintf("http://localhost%s/rapidoc.html\n", appPort)))

	log.Fatalln(router.Run(appPort))
}
