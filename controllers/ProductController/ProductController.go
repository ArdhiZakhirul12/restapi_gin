package ProductController

import (
	"encoding/json"
	"net/http"
	"rest_api_golang/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{
		"products": products,
		"msg":      "berhasil menampilkan list produk !"})

}
func Show(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Preload("User").First(&product, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"msg": "Data tidak ditemukan:"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"product": product,
		"msg":     "berhasil menampilkan data produk dengan id " + id + "!"})
}

func Create(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	models.DB.Create(&product)
	c.JSON(http.StatusOK, gin.H{"product": product, "msg": "Berhasil menambahkan data produk !"})
}
func Update(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "tidak dapat mengupdate data !"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "data berhasil diupdate !"})
}
func Delete(c *gin.Context) {
	var product models.Product
	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "tidak dapat menghapus data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "berhasil menghapus data "})
}
