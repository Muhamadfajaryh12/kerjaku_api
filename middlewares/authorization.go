package middlewares

import (
	"kerjaku/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware() fiber.Handler{
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message":"Unauthorized"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message":"Invalid token"})
		}

		authHeader = parts[1]

		claims, err := utils.VerifyToken(authHeader)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message":"Invalid token"})
		}

		c.Locals("user_id",claims["user_id"])
		return c.Next()
	}
}