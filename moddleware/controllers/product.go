package controllers

import (
	"moddleware/database"
	"moddleware/helpers"
	"moddleware/model"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	// ambil token dan data dari autentifikasi
	userData := c.MustGet("userData").(jwt.MapClaims)

	// model yang ingin diubah
	product := model.Product{}

	// ambil map data berdasarkan generate token
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	// memberikan user ID sebagai User ID product
	product.UsergoID = userID

	err := db.Debug().Create(&product).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, product)

}

func UpdateProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	// ambil token dan data dari autentifikasi
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))

	// model yang ingin diubah
	product := model.Product{}

	// ambil map data berdasarkan generate token
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	// memberikan user ID sebagai User ID product
	product.UsergoID = userID

	// product ID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id=?", productId).Updates(model.Product{
		Title:       product.Title,
		Description: product.Description,
	}).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func ReadProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	// ambil token dan data dari autentifikasi
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))

	// model yang ingin diubah
	product := model.Product{}

	// ambil map data berdasarkan generate token
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	// memberikan user ID sebagai User ID product
	product.UsergoID = userID

	// product ID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id=?", productId).Find(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	// ambil token dan data dari autentifikasi
	userData := c.MustGet("userData").(jwt.MapClaims)

	productId, _ := strconv.Atoi(c.Param("productId"))

	// model yang ingin diubah
	product := model.Product{}

	// ambil map data berdasarkan generate token
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&product)
	} else {
		c.ShouldBind(&product)
	}

	// memberikan user ID sebagai User ID product
	product.UsergoID = userID

	// product ID
	product.ID = uint(productId)

	err := db.Model(&product).Where("id=?", productId).Delete(&product).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "bad request",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}
