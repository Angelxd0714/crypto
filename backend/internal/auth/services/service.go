package services

import (
	"gin-quickstart/internal/auth/dto"
	"gin-quickstart/internal/database"
	"gin-quickstart/internal/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (string, error)
}

type authService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) AuthService {
	return &authService{DB: db}
}

func (s *authService) Register(req dto.RegisterRequest) error {
	var existingUser database.User
	if err := s.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	newUser := database.User{
		ID:       uuid.New(),
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {
	var user database.User
	if err := s.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		return "", err
	}

	if utils.ComparePassword(req.Password, user.Password) != nil {
		return "", nil
	}

	token, err := utils.GenerateJWTSecret(&user)
	if err != nil {
		return "", err
	}

	return token, nil
}
