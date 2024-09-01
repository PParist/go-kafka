package controller

import (
	"producer/commands"
	"producer/services"

	"github.com/gofiber/fiber/v2"
)

type accountController struct {
	accountService services.AccountService
}

func NewAccountController(accountService services.AccountService) AccountController {
	return &accountController{accountService: accountService}
}

func (con *accountController) OpenAccount(c *fiber.Ctx) error {
	commands := commands.OpenAccountCommand{}

	if err := c.BodyParser(&commands); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := con.accountService.OpenAccount(commands); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": "open account success"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "open account success"})
}

func (con *accountController) Deposit(c *fiber.Ctx) error {
	commands := commands.DepositFunCommand{}

	if err := c.BodyParser(&commands); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := con.accountService.Deposit(commands); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "Deposit success"})
}

func (con *accountController) Withdraw(c *fiber.Ctx) error {
	commands := commands.WithdrawFunCommand{}

	if err := c.BodyParser(&commands); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := con.accountService.Withdraw(commands); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "open withdraw success"})
}

func (con *accountController) CloseAccount(c *fiber.Ctx) error {
	commands := commands.CloseAccountCommand{}

	if err := c.BodyParser(&commands); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := con.accountService.CloseAccount(commands); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"Message": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "close account success"})
}
