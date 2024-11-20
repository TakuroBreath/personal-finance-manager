package service

import (
	"errors"
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
	"github.com/TakuroBreath/personal-finance-manager/internal/repository"
)

type UserServiceImpl interface {
	CreateUser(request UserCreateRequest) (uint, error)
	GetUserByEmail(email string) (UserResponse, error)
	LoginUser(request LoginRequest) (*AuthResponse, error)
	UpdateUser(userID uint, updateData UserUpdateRequest) error
	DeleteUser(userID uint) error
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type UserUpdateRequest struct {
	Nickname    *string `json:"username,omitempty"`
	NewPassword *string `json:"password,omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthResponse struct {
	Token    string       `json:"token"`
	UserInfo UserResponse `json:"user_info"`
}

type UserService struct {
	repo       *repository.Repository
	jwtService *JWTService
}

func NewUserService(repo *repository.Repository, jwtService *JWTService) *UserService {
	return &UserService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *UserService) CreateUser(user UserCreateRequest) (uint, error) {
	if _, err := s.repo.GetUserByEmail(user.Email); err == nil {
		return 0, errors.New("user with this email already exists")
	}

	// Create user model
	newUser := &models.User{
		Email:    user.Email,
		Password: user.Password,
		Nickname: user.Username,
	}

	return s.repo.CreateUser(*newUser)
}

func (s *UserService) GetUserByEmail(email string) (UserResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:       user.ID,
		Username: user.Nickname,
		Email:    user.Email,
	}, nil
}

func (s *UserService) LoginUser(request LoginRequest) (*AuthResponse, error) {
	user, err := s.repo.GetUserByEmail(request.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := user.ComparePassword(request.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := s.jwtService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		UserInfo: UserResponse{
			ID:       user.ID,
			Username: user.Nickname,
			Email:    user.Email,
		},
	}, nil
}

func (s *UserService) UpdateUser(userID uint, updateData UserUpdateRequest) error {
	// Fetch existing user
	user, err := s.repo.GetUserByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	// Update nickname if provided
	if updateData.Nickname != nil {
		user.Nickname = *updateData.Nickname
	}

	// Update password if provided
	if updateData.NewPassword != nil {
		user.Password = *updateData.NewPassword
	}

	return s.repo.UpdateUser(*user)
}

func (s *UserService) DeleteUser(userID uint) error {
	return s.repo.DeleteUser(userID)
}
