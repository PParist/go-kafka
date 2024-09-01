package controller

import "github.com/gofiber/fiber/v2"

type AccountController interface {
	OpenAccount(c *fiber.Ctx) error
	Deposit(c *fiber.Ctx) error
	Withdraw(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}
