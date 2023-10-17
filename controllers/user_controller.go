package controllers

import (
	"net/http"

	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/helpers"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/models"

	"github.com/gin-gonic/gin"
	"github.com/saepudinasep/task-5-pbi-btpns-AsepSaepudin/database"
)

func RegisterUser(c *gin.Context) {
	// Implementasi registrasi pengguna
	// Mendeklarasikan struct untuk menerima data dari pengguna
	var userInput models.UserDTO

	// Binding JSON data yang dikirimkan oleh pengguna ke struct userInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Membuat hash dari kata sandi yang diberikan oleh pengguna
	hashedPassword, err := helpers.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal meng-hash kata sandi"})
		return
	}

	// Membuat objek User berdasarkan data yang diterima
	newUser := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	// Simpan pengguna ke database
	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal melakukan registrasi"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil registered"})
}

func LoginUser(c *gin.Context) {
	// Implementasi login pengguna dan pembuatan token JWT
	// Mendeklarasikan struct untuk menerima data dari pengguna
	var loginInput models.UserDTO

	// Binding JSON data yang dikirimkan oleh pengguna ke struct loginInput
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Temukan pengguna berdasarkan email
	var user models.User
	if err := database.DB.Where("email = ?", loginInput.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Memeriksa kata sandi
	if err := helpers.CheckPasswordHash(loginInput.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Kata sandi salah"})
		return
	}

	// Membuat token JWT
	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	// Memberikan respons dengan token JWT
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func UpdateUser(c *gin.Context) {
	// Implementasi pembaruan data pengguna
	// Mendapatkan ID pengguna dari konteks (setelah melewati middleware otentikasi)
	userID, _ := c.Get("userID")

	// Mendapatkan data pembaruan dari JSON yang dikirimkan oleh pengguna
	var updateInput models.UserDTO
	if err := c.ShouldBindJSON(&updateInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Temukan pengguna yang akan diperbarui
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pengguna tidak ditemukan"})
		return
	}

	// Memperbarui data pengguna
	user.Username = updateInput.Username
	user.Email = updateInput.Email

	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal update user"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil updated"})
}

func DeleteUser(c *gin.Context) {
	// Implementasi penghapusan pengguna
	userID := c.Param("userId")

	// Temukan pengguna yang akan dihapus
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User tidak ditemukan"})
		return
	}

	// Hanya pengguna dengan ID yang sesuai atau admin yang dapat menghapus pengguna
	// Implementasikan logika otorisasi sesuai kebutuhan

	// Menghapus pengguna dari database
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus user"})
		return
	}

	// Memberikan respons sukses
	c.JSON(http.StatusOK, gin.H{"message": "User berhasil deleted"})
}
