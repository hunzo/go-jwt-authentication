package controller

import (
	"api/database"
	"api/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Home(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message":  "GO JWT-LOGIN Example",
		"register": "POST /register, payload: name,email,password",
		"login":    "POST /login, payload: email,password",
		"user":     "GET /user",
		"logout":   "POST /logout",
	})
}

const SecretKey = "secretkey"

func Register(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 1)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	if rs := database.DB.Create(&user); rs.Error != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"message": rs.Error,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
	})

}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	// if user := database.DB.Where("email=?", data["email"]).First(&user); user.Error != nil {
	// 	return c.JSON(fiber.Map{
	// 		"message": user.Error,
	// 	})
	// }

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthorize",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}

func User(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	// fmt.Printf("%v", cookie)

	token, err := jwt.ParseWithClaims(
		cookie,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
	)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, _ := token.Claims.(*jwt.StandardClaims)
	// claims, ok := token.Claims.(*jwt.StandardClaims)

	// if !ok {
	// 	fmt.Printf("%v", ok)
	// }

	// fmt.Printf("%v", claims)

	var user models.User

	database.DB.Where("id=?", claims.Issuer).First(&user)

	return c.JSON(user)

}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
