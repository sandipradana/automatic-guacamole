package service

import (
	"automatic-guacamole/internal/model"
	"automatic-guacamole/internal/repository"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetAll() ([]model.User, error)
	GetByID(id uint64) (*model.User, error)
	Create(User *model.User) error
	Update(User *model.User) error
	Delete(id uint64) error
	Login(UserLogin model.UserLogin) (*string, error)
}

type userService struct {
	userRepo repository.UserRepository
	db       *gorm.DB
}

func NewUserService(db *gorm.DB, repo repository.UserRepository) UserService {
	return &userService{db: db, userRepo: repo}
}

func (s *userService) GetAll() ([]model.User, error) {
	return s.userRepo.GetAll(s.db)
}

func (s *userService) GetByID(id uint64) (*model.User, error) {
	return s.userRepo.GetByID(s.db, id)
}

func (s *userService) Create(User *model.User) error {
	passwordEncoded, _ := bcrypt.GenerateFromPassword([]byte(User.Password), 14)
	User.Password = string(passwordEncoded)
	return s.userRepo.Create(s.db, User)
}

func (s *userService) Update(User *model.User) error {
	return s.userRepo.Update(s.db, User)
}

func (s *userService) Delete(id uint64) error {
	return s.userRepo.Delete(s.db, id)
}

func (s *userService) Login(UserLogin model.UserLogin) (*string, error) {

	User, err := s.userRepo.GetUserByEmail(s.db, UserLogin.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(UserLogin.Password)); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["obj"] = User.ID
	claims["sub"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to generate JWT token")
	}

	return &tokenString, nil
}
