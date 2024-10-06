package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetStudents(ctx echo.Context) error {
	var students []Student

	// Ambil semua data student dari database
	if err := db.Find(&students).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "gagal mengambil data students",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan semua students",
		"data":    students,
	})
}

func GetStudent(ctx echo.Context) error {
	studentRequestByID := new(StudentRequestByID)
	if err := ctx.Bind(studentRequestByID); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	id, err := uuid.Parse(studentRequestByID.ID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "format uuid salah",
			"data":    nil,
		})
	}

	// Cari student berdasarkan ID di database
	var student Student
	if err := db.First(&student, "id = ?", id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "student tidak ditemukan",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "berhasil mendapatkan student",
		"data":    student,
	})
}

// handler.go
func CreateStudent(ctx echo.Context) error {
	studentRequest := new(StudentRequest)
	if err := ctx.Bind(studentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}
	student := Student{
		ID:          uuid.New(),
		Name:        studentRequest.Nama,
		Email:       studentRequest.SuratElektronik,
		PhoneNumber: studentRequest.NoHP,
		Address:     studentRequest.Alamat,
		GPA:         studentRequest.IPK,
		IsGraduate:  studentRequest.IsGraduate,
	}

	// Simpan student ke database
	if err := db.Create(&student).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "gagal membuat student",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "student berhasil dibuat",
		"data":    student,
	})
}

func UpdateStudent(ctx echo.Context) error {
	studentRequest := new(StudentRequest)
	if err := ctx.Bind(studentRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"data":    nil,
		})
	}

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "format uuid salah",
			"data":    nil,
		})
	}

	// Cari student berdasarkan ID di database
	var student Student
	if err := db.First(&student, "id = ?", id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "student tidak ditemukan",
			"data":    nil,
		})
	}

	// Update data student
	student.Name = studentRequest.Nama
	student.Email = studentRequest.SuratElektronik
	student.PhoneNumber = studentRequest.NoHP
	student.Address = studentRequest.Alamat
	student.GPA = studentRequest.IPK
	student.IsGraduate = studentRequest.IsGraduate

	if err := db.Save(&student).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "gagal mengupdate student",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "student berhasil diupdate",
		"data":    student,
	})
}

func DeleteStudent(ctx echo.Context) error {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "format uuid salah",
			"data":    nil,
		})
	}

	// Cari student berdasarkan ID di database
	var student Student
	if err := db.First(&student, "id = ?", id).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "student tidak ditemukan",
			"data":    nil,
		})
	}

	// Hapus student
	if err := db.Delete(&student).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "gagal menghapus student",
			"data":    nil,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "student berhasil dihapus",
	})
}
