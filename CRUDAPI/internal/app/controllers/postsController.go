// package controllers

// import (
// 	"CRUDAPI/internal/app/confi"
// 	"CRUDAPI/internal/app/initializers"
// 	"CRUDAPI/internal/app/models"
// 	"CRUDAPI/internal/app/services"

// 	"github.com/gin-gonic/gin"
// )
// type PostController struct {
// 	AppConfig *confi.AppConfig
// 	Service   services.PostServices
// }
// func  PostsCreate(c *gin.Context) {
// 	//Get data of request body
// 	//create a post
// 	//return it
// 	// var body struct {
// 	// 	Body  string `json:""`
// 	// 	Title string
// 	// }
// 	c.Bind(&models.Post1)
// 	post := models.Post{Title: models.Post1.Title, Body: models.Post1.Body}
// 	result := confi.DB.Create(&post)
// 	if result.Error != nil {
// 		c.Status(400)
// 		return
// 	}
// 	c.JSON(200, result)
// }

// func  PostIndex(c *gin.Context) {
// 	//Get post and respond
// 	var posts []models.Post
// 	confi.DB.Find(&posts)
// 	//respond
// 	c.JSON(200, gin.H{
// 		"Post": posts,
// 	})
// }
// func PostShow(c *gin.Context) {
// 	//Get id of URL
// 	id := c.Param("id")

// 	var post models.Post
// 	initializers.Db.Find(&post, id)
// 	//respond
// 	c.JSON(200, gin.H{
// 		"Post": post,
// 	})
// }
// func PostUpdate(c *gin.Context) {
// 	//Get the id of URL
// 	id := c.Param("id")
// 	//GEt data of request body
// 	// var body struct {
// 	// 	Body  string
// 	// 	Title string
// 	// }
// 	c.Bind(&models.Post1)
// 	//find post to update
// 	var post models.Post
// 	initializers.Db.Find(&post, id)
// 	//update
// 	initializers.Db.Model(&post).Updates(models.Post{Title: models.Post1.Title, Body: models.Post1.Body})
// 	//response
// 	c.JSON(200, gin.H{
// 		"Post": post,
// 	})
// }

// func PostDelete(c *gin.Context) {
// 	//Get the id
// 	id := c.Param("id")
// 	//Delete The Post
// 	initializers.Db.Delete(&models.Post{}, id)
// 	//Respond
// 	c.JSON(200, id)
// }

package controllers

import (
	"CRUDAPI/internal/app/models"
	"CRUDAPI/internal/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	Service services.PostService
}

func NewPostController(service services.PostService) *PostController {
	return &PostController{
		Service: service,
	}
}

func (pc *PostController) PostsCreate(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if err := pc.Service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (pc *PostController) PostIndex(c *gin.Context) {
	posts, err := pc.Service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (pc *PostController) PostShow(c *gin.Context) {
	id := c.Param("id")

	post, err := pc.Service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := pc.Service.UpdatePost(id, &post); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) PostDelete(c *gin.Context) {
	id := c.Param("id")

	if err := pc.Service.DeletePost(id); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "Post deleted successfully")
}
