package configs

import (
	"github.com/akashchandra111/factory_man/src/types"

	"github.com/gofiber/fiber/v2"
)

func GetName(c *fiber.Ctx) error {

	sal := types.Salary{
		EmpId:        10,
		SalaryAmount: 25000,
		DateOfCredit: "Monday",
	}

	return c.JSON(sal)
}
