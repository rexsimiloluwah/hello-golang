package models

import (
	"fmt"
	"time"

	u "../utils"
	"github.com/jinzhu/gorm"
)

type Post struct {
	ID          uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title       string     `gorm:"type varchar(300);not null" json:"title"`
	Category    string     `gorm:"type varchar(200);not null" json:"category"`
	Description string     `gorm:"type varchar(500)" json:"description"`
	Content     string     `gorm:"type:MEDIUMTEXT" json:"content"`
	UserId      uint       `gorm:"foreignkey:UserIdForUser" json:"user_id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

// Validate incoming Post data
func (post *Post) Validate() (map[string]interface{}, bool) {
	if len(post.Title) > 300 || len(post.Title) < 10 {
		return u.Message(false, "Title must be between 10 to 300 characters"), false
	}

	if len(post.Description) > 500 {
		return u.Message(false, "Description must not be more than 500 characters"), false
	}

	if len(post.Category) == 0 {
		return u.Message(false, "Category is required."), false
	}

	if len(post.Content) == 0 {
		return u.Message(false, "Content is required."), false
	}

	return u.Message(false, "Validation successful."), true
}

// Create a new Post
func (post *Post) CreatePost() map[string]interface{} {
	resp, ok := post.Validate()
	if !ok {
		return resp
	}

	Db().Create(post)
	response := u.Message(true, "New post successfully created.")
	response["post"] = post
	return response
}

// Get all posts
func GetPosts() map[string]interface{} {
	posts := make([]*Post, 0)
	var response map[string]interface{}
	err := Db().Table("posts").Find(&posts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response = u.Message(false, "No posts found")
			response["posts"] = nil
			return response
		}
	}

	response = u.Message(true, fmt.Sprintf("Successfully fetched %d post(s).", len(posts)))
	response["posts"] = posts
	return response
}

// Get post by post id
func GetPostById(post_id uint) map[string]interface{} {
	post := &Post{}
	var response map[string]interface{}
	err := Db().Table("posts").Where("id = ?", post_id).First(post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response = u.Message(false, "No post found with that id.")
			response["post"] = nil
			return response
		}
	}

	response = u.Message(true, "Post successfully fetched")
	response["post"] = post
	return response
}

// Get User posts
func GetUserPosts(user_id uint) map[string]interface{} {
	userPosts := make([]*Post, 0)
	var response map[string]interface{}
	err := Db().Table("posts").Where("user_id = ?", user_id).Find(&userPosts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response = u.Message(false, "No posts found for this user.")
			response["posts"] = nil
			return response
		}
	}

	response = u.Message(true, fmt.Sprintf("Successfully fetched %d post(s)", len(userPosts)))
	response["posts"] = userPosts
	return response
}

// Update post
func UpdatePost(post_id uint, fields map[string]interface{}) map[string]interface{} {
	var response map[string]interface{}
	err := db.First(&Post{}, post_id).Error
	if err == gorm.ErrRecordNotFound {
		response = u.Message(false, "Post does not exist.")
		return response
	}

	err = db.Table("posts").Where("id = ?", post_id).Updates(fields).Error
	if err != nil {
		response = u.Message(false, "An error occurred.")
		return response
	}

	response = u.Message(true, "Post successfully updated.")
	return response
}

// Delete post
func DeletePost(post_id uint) map[string]interface{} {
	var response map[string]interface{}
	err := db.First(&Post{}, post_id).Error
	if err == gorm.ErrRecordNotFound {
		response = u.Message(false, "Post does not exist.")
		return response
	}
	err = db.Delete(&Post{}, post_id).Error
	if err != nil {
		response = u.Message(false, "An error occurred.")
		return response
	}

	response = u.Message(true, "Post successfully deleted.")
	return response
}
