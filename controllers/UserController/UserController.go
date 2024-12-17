package UserController

import (
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"rest_api_golang/models"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "Gagal melakukan hashing password"})
		return
	}
	user.Password = string(hashedPassword) // Simpan hash ke dalam struct

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"msg":  "berhasil registrasi !"})

}

func GetUserWithProducts(c *gin.Context) {
	var user models.User
	id := c.Param("id") // Ambil ID user dari URL

	// Query dengan Preload untuk memuat relasi Products
	if err := models.DB.Preload("Products").First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "User tidak ditemukan"})
		return
	}

	// Kembalikan respons JSON
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
