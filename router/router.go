package router

import "github.com/gin-gonic/gin"

func Initialize(){
	// Initialize routes
	router := gin.Default()
	
	initializeRoutes(router)

	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

