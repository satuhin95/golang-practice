package controller

import (
	"react-go-fiber-jwt/database"
	"react-go-fiber-jwt/models"

	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Hello(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
func Register(c *fiber.Ctx) error {
	var data map[string]string
	db := database.DBConn
	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]),14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	return c.JSON(user)
}

func Login(c *fiber.Ctx) error{
	var data map[string]string

	db := database.DBConn

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	var user models.User

	db.Where("email = ?", data["email"]).First(&user)

	if user.ID == 0 { 
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil{
		 c.Status(fiber.StatusBadRequest)
		 return c.JSON(fiber.Map{ "message": "Incorrect password",})

	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),            //issuer contains the ID of the user.
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //Adds time to the token i.e. 24 hours.
	})
	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name : "jwt",
		Value : token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.JSON(fiber.Map{"message":"Success"})

}

func User(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey),nil
	})
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "UnAuthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DBConn.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error{
	cookie := fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message":"Logout Successful",
	})
}