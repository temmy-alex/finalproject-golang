package controllers

import (
	"encoding/json"
	"final-project/database"
	"final-project/helpers"
	"final-project/models"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ListSocialMedia(c *gin.Context) {
	var (
		socialmedias []models.SocialMedia
	)

	db := database.GetDB()

	err := db.Preload("User").Find(&socialmedias).Error

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": &socialmedias,
	})
}

func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.UserID = userID

	err := db.Debug().Create(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"created_at":       SocialMedia.CreatedAt,
	})
}

func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType

	SocialMedia := models.SocialMedia{}
	NewSocialMedia := models.SocialMedia{}

	id := c.Param("socialMediaId")

	err := db.First(&SocialMedia, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&NewSocialMedia)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Model(&SocialMedia).Updates(NewSocialMedia).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":               SocialMedia.ID,
		"name":             SocialMedia.Name,
		"social_media_url": SocialMedia.SocialMediaUrl,
		"user_id":          SocialMedia.UserID,
		"updated_at":       SocialMedia.UpdatedAt,
	})
}

func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	SocialMedia := models.SocialMedia{}

	id := c.Param("socialMediaId")

	err := db.First(&SocialMedia, id).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err = db.Delete(&SocialMedia).Error
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	})
}
