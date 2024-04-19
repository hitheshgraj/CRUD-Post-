package app

import (
	"CRUDAPI/internal/app/controllers"
	"CRUDAPI/internal/app/services"

	//"CRUDAPI/internal/app/initializers"
	"CRUDAPI/internal/app/confi"
	// "CRUDAPI/internal/app/midelware"
	"CRUDAPI/internal/app/midelware"
	"log"

	//"github.com/gin-gonic/gin"
)

// func RunURL() {
// 	log.Println("Initialise all handlers")
// 	api := gin.Default()
// 	log.Println("Initialise all handlers")
// 	cartController := controller.NewController(app)

// 	router.GET("/ping", middleware.BasicAuth(app), func(ctx *gin.Context) {
// 		ctx.JSON(200, "ping")
// 	})
// 	api.POST("/posts", controllers.PostsCreate)
// 	api.PUT("/posts/:id", controllers.PostUpdate)

// 	api.GET("/posts", controllers.PostIndex)
// 	api.GET("/posts/:id", controllers.PostShow)

// 	api.DELETE("/posts/:id", controllers.PostDelete)

//		api.Run()
//	}
func RunURL(app *confi.AppConfig) {
	log.Println("Initialise all handlers")
	//cartController := controller.NewController(app)
	postService := services.NewPostService(confi.DB)

	// Initialize controllers
	postController := controllers.NewPostController(postService)
	// // Enable CORS middleware
	// router.Use(middleware.CORSMiddleware())
	// router.GET("/ping", middleware.BasicAuth(app), func(ctx *gin.Context) {
	// 	ctx.JSON(200, "ping")
	// })
	// api := gin.Default()
	api := router.Group("/v1")
	api.Use(middleware.BasicAuth(app))
	api.POST("/posts", postController.PostsCreate)
	api.PUT("/posts/:id", postController.PostUpdate)

	api.GET("/posts", postController.PostIndex)
	api.GET("/posts/:id", postController.PostShow)

	api.DELETE("/posts/:id", postController.PostDelete)
	//api.Run()

}
