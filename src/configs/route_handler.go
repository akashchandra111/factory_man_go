package configs

import (
	"log"
	"strconv"
	"time"

	"github.com/akashchandra111/factory_man/src/types"

	"github.com/gofiber/fiber/v2"
)


func RegisterUserRoute(c *fiber.Ctx) error {
	user := types.Employee{}

	if err := c.BodyParser(user); err != nil {
		return err
	}
	return c.JSON(user)
}

func GetUserRoute(c *fiber.Ctx) error {
	userInfo := types.LoginInfo{}

	if err := c.BodyParser(&userInfo); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	emp := GetUser(userInfo)
	return c.JSON(emp)
}

func GetAllJobsRoute(c *fiber.Ctx) error {
	return c.JSON(GetAllJobs())
}

func GetEmployeesByJobIdRoute(c *fiber.Ctx) error {
	jobId, err := c.ParamsInt("job")

	if err != nil {
		log.Println("Error parsing job id")
		return err
	}

	return c.JSON(GetEmployeesByJobId(int64(jobId)))
}

func InsertUserRoute(c *fiber.Ctx) error {
	emp := types.Employee{}

	if err := c.BodyParser(&emp); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	return c.SendString(strconv.Itoa(int(InsertUser(&emp))))
}

func InsertJobRoute(c *fiber.Ctx) error {
	job := types.Job{}

	if err := c.BodyParser(&job); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	return c.SendString(strconv.Itoa(int(InsertJob(&job))))
}

func InsertAdvanceRoute(c *fiber.Ctx) error {
	advance := types.AdvanceSalary{}

	if err := c.BodyParser(&advance); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	return c.SendString(strconv.Itoa(int(InsertAdvance(&advance))))
}

func InsertSalaryRoute(c *fiber.Ctx) error {
	salary := types.Salary{}

	if err := c.BodyParser(&salary); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	return c.SendString(strconv.Itoa(int(InsertSalary(&salary))))
}

func InsertMakingJobRoute(c *fiber.Ctx) error {
	makingJob := types.MakingJob{}

	if err := c.BodyParser(&makingJob); err != nil {
		log.Println("Not able to unmarshal")
		return err
	}

	return c.SendString(strconv.Itoa(int(InsertMakingJob(&makingJob))))
}

func GetSalaryByEmpIdRoute(c *fiber.Ctx) error {
	empId, err := c.ParamsInt("empId")

	if err != nil {
		log.Println("Error parsing emp id")
		return err
	}

	return c.JSON(GetSalaryByEmpId(int64(empId)))
}

func GetAdvanceSalaryByEmpIdRoute(c *fiber.Ctx) error {
	empId, err := c.ParamsInt("empId")

	if err != nil {
		log.Println("Error parsing emp id")
		return err
	}

	return c.JSON(GetAdvanceSalaryByEmpId(int64(empId)))
}

func GetMakingJobByEmpIdRoute(c *fiber.Ctx) error {
	empId, err := c.ParamsInt("empId")

	if err != nil {
		log.Println("Error parsing emp id")
		return err
	}

	return c.JSON(GetMakingJobByEmpId(int64(empId)))
}

func GetMakingJobByJobIdRoute(c *fiber.Ctx) error {
	jobId, err := c.ParamsInt("empId")

	if err != nil {
		log.Println("Error parsing job id")
		return err
	}

	return c.JSON(GetMakingJobByJobId(int64(jobId)))
}

func GetMakingJobByDateRoute(c *fiber.Ctx) error {
	date := c.Params("date")
	dt, err := time.Parse("2006-01-02T15:04:05.000Z", date)

	if err != nil {
		log.Println("Error parsing date")
		return err
	}

	return c.JSON(GetMakingJobByDate(dt))
}
