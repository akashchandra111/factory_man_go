package main

import (
	"os"

	"github.com/akashchandra111/factory_man/src/configs" // Configs

	"github.com/gofiber/fiber/v2" // Fiber import
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()      // Intializing fiber app
	err := configs.InitDB() // Intializing DB

	if err != nil {
		os.Stderr.WriteString("Failed to init database, exiting")
		panic(err)
	}

	app.Use(
		cors.New(
			cors.Config	{
				AllowOrigins: "*",
				AllowHeaders: "",
			},
		),
	)

	app.Get("/app/jobs", configs.GetAllJobsRoute)
	app.Get("/app/get/users/job_id/:job", configs.GetEmployeesByJobIdRoute)
	app.Get("/app/get/salary/emp_id/:empId", configs.GetSalaryByEmpIdRoute)
	app.Get("/app/get/advance/emp_id/:empId", configs.GetAdvanceSalaryByEmpIdRoute)
	app.Get("/app/get/making_job/emp_id/:empId", configs.GetMakingJobByEmpIdRoute)
	app.Get("/app/get/making_job/job_id/:jobId", configs.GetMakingJobByJobIdRoute)
	app.Get("/app/get/making_job/date/:date", configs.GetMakingJobByDateRoute)

	app.Post("/app/get/user", configs.GetUserRoute)

	app.Post("/app/salary/add", configs.InsertSalaryRoute)
	app.Post("/app/advance/add", configs.InsertAdvanceRoute)
	app.Post("/app/making/add", configs.InsertMakingJobRoute)
	app.Post("/app/job/add", configs.InsertJobRoute)
	app.Post("/app/user/add", configs.RegisterUserRoute)

	app.Listen(":8911")
}
