package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	gf "github.com/shareed2k/goth_fiber"
	"github.com/wambugucoder/simple-go-service/configs"
	"github.com/wambugucoder/simple-go-service/models"
	"github.com/wambugucoder/simple-go-service/repository"
	"log"
)

type BaseInput struct {
	Email    string `json:"email" validate:"required,email" `
	Password string `json:"password" validate:"required,min=6,max=12"`
}

type RegisterInput struct {
	BaseInput
	Username string `json:"username" validate:"required,min=6,max=12"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateInput(input RegisterInput) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()

			errors = append(errors, &element)

		}
	}
	return errors

}

func AddUser(ctx *fiber.Ctx) error {
	userdetails := new(RegisterInput)
	//PARSE
	err := ctx.BodyParser(userdetails)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	errors := ValidateInput(*userdetails)
	if errors != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": errors,
		})
	}
	if repository.DoesEmailExist(userdetails.Email) {
		return ctx.Status(404).JSON(fiber.Map{
			"error":   true,
			"message": "Email Already Exists",
		})
	}
	userinfo := &models.User{
		Username: userdetails.Username,
		Email:    userdetails.Email,
		Password: userdetails.Password,
	}
	results := repository.SaveUser(userinfo)
	return ctx.JSON(fiber.Map{
		"error":   false,
		"message": results,
	})
}

func GoogleAuth() {
	goth.UseProviders(
		google.New(configs.ExtractEnvKey("GOOGLE_CLIENT_ID"), configs.ExtractEnvKey("GOOGLE_CLIENT_SECRET"),
			"http://localhost:3000/api/v1/auth/google/callback", "email", "profile"),
	)
}

func BeginGoogleOauth(ctx *fiber.Ctx) error {
	err := gf.BeginAuthHandler(ctx)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func CompleteOauth(ctx *fiber.Ctx) error {
	user, err := gf.CompleteUserAuth(ctx)
	if err != nil {
		ctx.Redirect("http://127.0.0.1:4000/api/v1/error=" + err.Error())
		return err
	}
	ctx.Redirect("http://127.0.0.1:3000/api/v1/token=" + user.Name)
	return nil
}
