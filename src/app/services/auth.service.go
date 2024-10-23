package services

import (
	"errors"

	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/dal"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/models"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/app/types"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils/jwt"
	"github.com/bancodebogota/bbog-dig-pl-go-mngr-template/src/utils/password"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(ctx *fiber.Ctx) error {
	b := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	u := &types.UserResponse{}

	err := dal.FindUserByEmail(u, b.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(u.Password, b.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	t, errToken := jwt.Generate(&jwt.TokenPayload{
		ID: u.ID,
	})

	if errToken != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error generating token")
	}

	return ctx.JSON(&types.AuthResponse{
		User: u,
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}

func Signup(ctx *fiber.Ctx) error {
	b := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		return err
	}

	err := dal.FindUserByEmail(&struct{ ID string }{}, b.Email).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &models.User{
		Name:     b.Name,
		Password: password.Generate(b.Password),
		Email:    b.Email,
	}

	if err := dal.CreateUser(user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	t, errToken := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	if errToken != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Error generating token")
	}

	return ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: t,
		},
	})
}
