package controllers

import (
	"encoding/json"
	"fmt"
	"go-restapi-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {
	var product []models.Product

	models.DB.Find(&product)

	c.JSON(http.StatusOK, gin.H{
		"products": product,
	})
}

func Show(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := models.DB.First(&product, id).Error; err != nil{
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "Data Tidak Ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func Create(c *gin.Context){
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil{
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error in field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}

func Update(c *gin.Context){
	var product models.Product
	id := c.Param("id")

	if err := c.ShouldBindJSON(&product); err != nil{
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors){
			errorMessage := fmt.Sprintf("Error in field %s condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	if models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected == 0{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data Tidak Ditemukkan", 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil Di Update",
	})
}

func Delete(c *gin.Context){
	var product models.Product
	
	var input struct{
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&product, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Data Tidak Ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Berhasil Dihapus",
	})
}