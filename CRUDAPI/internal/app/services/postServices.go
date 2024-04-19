// package services
// import (
//
//	"CRUDAPI/internal/app/confi"
//	"CRUDAPI/internal/app/models"
//
// )
//
//	type PostServices struct {
//		AppConfig *confi.AppConfig
//		CartData  []models.Post
//	}
//
// // func(p *PostServices) GetAllPosts()([]models.Post,error){
// // 	//db:=confi.DB.Find(&p.CartData)
// // 	var db=confi.DB.Find(&p.CartData)
// // 	//return db,nil
// // }
package services

import (
	"CRUDAPI/internal/app/models"
	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(post *models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostByID(id string) (*models.Post, error)
	UpdatePost(id string, post *models.Post) error
	DeletePost(id string) error
}

type postService struct {
	DB *gorm.DB
}

func NewPostService(db *gorm.DB) PostService {
	return &postService{
		DB: db,
	}
}

func (ps *postService) CreatePost(post *models.Post) error {
	return ps.DB.Create(post).Error
}

func (ps *postService) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := ps.DB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (ps *postService) GetPostByID(id string) (*models.Post, error) {
	var post models.Post
	if err := ps.DB.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (ps *postService) UpdatePost(id string, post *models.Post) error {
	return ps.DB.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error
}

func (ps *postService) DeletePost(id string) error {
	return ps.DB.Delete(&models.Post{}, id).Error
}
