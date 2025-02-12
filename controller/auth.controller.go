package controller

import (
	"github.com/anjarmath/01_golang_api_sederhana/model"
	"github.com/anjarmath/01_golang_api_sederhana/repository"
	"github.com/anjarmath/01_golang_api_sederhana/utils"
	"github.com/gofiber/fiber/v2"
)

// Field Untuk request
type (
	ReqisterReq struct {
		Username string `json:"username"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	LoginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

type authController struct {
	userRepository repository.UserRepository
}

func NewAuthController(repo repository.UserRepository) authController {
	return authController{
		userRepository: repo,
	}
}

func (ac authController) Register(c *fiber.Ctx) error {
	userReq := new(ReqisterReq)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Mohon isikan data dengan benar",
		})
	}

	valid := validateUser(userReq)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: "Mohon isikan data dengan benar",
		})
	}

	user := model.User{
		Username: userReq.Username,
		Name:     userReq.Name,
		Password: userReq.Password,
	}

	if err := user.HashPassword(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Terjadi Kesalahan",
		})
	}

	if err := ac.userRepository.AddUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: "Terjadi Kesalahan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Pendaftaran sukses",
	})

}

func (ac authController) Login(c *fiber.Ctx) error {
	userReq := new(LoginReq)

	if err := c.BodyParser(userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "Mohon isikan data dengan benar",
		})
	}

	user, err := ac.userRepository.GetUserByUsername(userReq.Username)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Error{
			Message: "Pengguna tidak terdaftar!",
		})
	}

	valid := user.CheckPassword(userReq.Password)
	if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Error{
			Message: "Password salah!",
		})
	}

	if err := utils.GenerateSession(c, *user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Error{
			Message: err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login sukses",
		"user":    user,
	})
}

func validateUser(user *ReqisterReq) bool {
	return !(user.Username == "" || user.Name == "" || user.Password == "")
}
