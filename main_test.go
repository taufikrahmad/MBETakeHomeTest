package main

import (
	"MBETakeHomeTest/dto"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetUpRouter() *gin.Engine{
	router := gin.Default()
	return router
}

func TestRegister(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/register", authController.Register)
	userTestRegister := dto.RegisterDTO{
		Name: "test",
		Email: "test2@gmail.com",
		Password: "jakarta@1",
	}

	jsonValue, _ := json.Marshal(userTestRegister)
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestDuplicateRegister(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/register", authController.Register)
	userTestRegister := dto.RegisterDTO{
		Name: "test",
		Email: "test2@gmail.com",
		Password: "jakarta@1",
	}

	jsonValue, _ := json.Marshal(userTestRegister)
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}