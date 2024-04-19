package main

import (
	"CRUDAPI/internal/app"
	"CRUDAPI/internal/app/confi"
	// "net/http"

	// "github.com/gin-gonic/gin"
	//"CRUDAPI/internal/app/initializers"
	//"CRUDAPI/internal/app/controllers"
	//"net/http"
	//"log"
	//"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

var (
	appConf *confi.AppConfig
)

func init() {
	// initializers.LoadEnvVariables()

	// initializers.LoadEnvVariables()
	appConf = confi.NewConfig()
	//initializers.ConnectToDB()

}
func main() {
	app.StartApplication(appConf)
	// engine := html.New("./views", ".tmpl")
	// // Create app
	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })
	// router := mux.NewRouter()
	// router.HandleFunc("/posts", handlePosts).Methods("GET")

	// corsHandler := handlers.CORS(
	//     handlers.AllowedHeaders([]string{"Content-Type"}),
	//     handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
	//     handlers.AllowedOrigins([]string{"http://localhost:3001"}),
	// )

	// http.ListenAndServe(":3000", corsHandler(router))
	// http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
	// })

	//  r := gin.Default()

	// r.Use(func(c *gin.Context) {
	// 	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	// 	if c.Request.Method == http.MethodOptions {
	// 		c.AbortWithStatus(http.StatusOK)
	// 		return
	// 	}

	// 	c.Next()
	//  })

	// r.POST("/posts",controllers.PostsCreate)
	// 	r.PUT("/posts/:id", controllers.PostUpdate)

	// 	r.GET("/posts", controllers.PostIndex)
	// 	r.GET("/posts/:id", controllers.PostShow)

	// 	r.DELETE("/posts/:id", controllers.PostDelete)

	// r.Run()
}

// func handlePosts(w http.ResponseWriter, r *http.Request) {
// 	// Your handler logic for /posts
// 	// Example: w.Write([]byte("Hello, this is the /posts endpoint"))
// }
