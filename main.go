package main

import (
	"first/controllers"
	"first/db"
	"net/http"
	"os"

	docs "first/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

// @title DB service
// @version 1.0
// @description Service for reports and ner results.

// @host localhost:7559
// @BasePath /api
func main() {
	db.Init()

	r := gin.Default()
	r.GET("/", homePage)

	docs.SwaggerInfo.BasePath = "/"
	v1 := r.Group("/")
	{
		eg := v1.Group("/api")
		{
			eg.GET("/reports", controllers.ReportsGetList)
			eg.POST("/reports", controllers.ReportsAddOne)
			eg.GET("/reports/:id", controllers.ReportsGetOne)
			eg.PATCH("/reports/:id", controllers.ReportsEditOne)
			eg.DELETE("/reports/:id", controllers.ReportsDeleteOne)
			eg.POST("/reports_upload", controllers.ReportsAddAll)

			eg.GET("/ner_results", controllers.NerResultsGetList)
			eg.POST("/ner_results", controllers.NerResultsAddOne)
			eg.GET("/ner_results/:id", controllers.NerResultsGetOne)
			eg.PATCH("/ner_results/:id", controllers.NerResultsEditOne)
			eg.DELETE("/ner_results/:id", controllers.NerResultsDeleteOne)
			eg.POST("/ner_results_upload", controllers.NerResultsAddAll)
		}
	}
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	godotenv.Load()
	port := os.Getenv("DB_SERVICE_PORT")
	r.Run(":" + port)
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "DB service is available.")
}
