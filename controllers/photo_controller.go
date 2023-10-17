package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/database"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/models"
)

func CreatePhoto(c *gin.Context) {
	// Implementasi pembuatan foto
	// Mendapatkan ID pengguna dari konteks (setelah melewati middleware otentikasi)
	userID, _ := c.Get("userID")

	// Mendapatkan data foto dari JSON yang dikirimkan oleh pengguna
	var photoInput models.PhotoDTO
	if err := c.ShouldBindJSON(&photoInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat objek foto baru
	newPhoto := models.Photo{
		Title:    photoInput.Title,
		Caption:  photoInput.Caption,
		PhotoURL: photoInput.PhotoURL,
		UserID:   userID.(uint), // Mengambil ID pengguna dari konteks
	}

	// Menyimpan foto ke database
	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create photo"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Photo created successfully"})
}

func GetPhotos(c *gin.Context) {
	// Implementasi pengambilan daftar foto
	// Mengambil daftar foto dari database
	var photos []models.Photo
	if err := database.DB.Find(&photos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve photos"})
		return
	}

	// Memberikan respons dengan daftar foto
	c.JSON(http.StatusOK, photos)
}

func UpdatePhoto(c *gin.Context) {
	// Implementasi pembaruan data foto
	// Mendapatkan ID foto dari parameter rute
	photoID := c.Param("photoId")

	// Mendapatkan data pembaruan dari JSON yang dikirimkan oleh pengguna
	var updateInput models.PhotoDTO
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Temukan foto yang akan diperbarui
	var photo models.Photo
	if err := database.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Memperbarui data foto
	photo.Title = updateInput.Title
	photo.Caption = updateInput.Caption
	photo.PhotoURL = updateInput.PhotoURL

	if err := database.DB.Save(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update photo"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Photo updated successfully"})
}

func DeletePhoto(c *gin.Context) {
	// Implementasi penghapusan foto
	// Mendapatkan ID foto dari parameter rute
	photoID := c.Param("photoId")

	// Temukan foto yang akan dihapus
	var photo models.Photo
	if err := database.DB.Where("id = ?", photoID).First(&photo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Photo not found"})
		return
	}

	// Mendapatkan ID pengguna dari konteks (setelah melewati middleware otentikasi)
	userID, _ := c.Get("userID")

	// Mendapatkan peran pengguna (misalnya, admin atau pengguna biasa) dari database
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Memeriksa otorisasi: hanya pemilik foto yang dapat menghapus
	if photo.UserID != userID.(uint) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Menghapus foto dari database
	if err := database.DB.Delete(&photo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete photo"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted successfully"})
}
