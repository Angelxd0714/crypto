package auth

import (
	"gin-quickstart/internal/auth/dto"
	"gin-quickstart/internal/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{Service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user     body    dto.RegisterRequest     true        "User info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /auth/register [post]
func (ac *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.Service.Register(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   user     body    dto.LoginRequest     true        "User credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.Service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
