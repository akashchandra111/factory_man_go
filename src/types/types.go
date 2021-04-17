package types

import "time"

type Employee struct {
	EmpId      int64     `json: "EmpId"`
	FirstName  string    `json: "FirstName"`
	LastName   string    `json: "LastName"`
	DOB        time.Time `json: "DOB"`
	MobileNo   string    `json: "MobileNo"`
	Password   string    `json: "Password"`
	EmpType    string    `json: "EmpType"`
	SubEmpType string    `json: "SubEmpType"`
	IsWorking  string    `json: "IsWorking"`
	IsSalaried string    `json: "IsSalaried`
	Salary     int64     `json: "Salary"`
	Advance    int64     `json: "Advance"`
	JobId      int64     `json: "JobId"`
}

type Salary struct {
	EmpId        int64     `json: "EmpId"`
	SalaryAmount int64     `json: "SalaryAmount"`
	DateOfCredit time.Time `json: "DateOfCredit"`
}

type AdvanceSalary struct {
	EmpId         int64     `json: "EmpId"`
	AdvanceAmount int64     `json: "AdvanceAmount"`
	AdvanceDate   time.Time `json: "AdvanceDate"`
}

type Job struct {
	JobId          int64  `json: "JobId"`
	JobType        string `json: "JobType"`
	JobSubType     string `json: "JobSubType"`
	JobSubTypeName string `json: "JobSubTypeName"`
	IsRcOrStone    string `json: "IsRcOrStone"`
}

type PersonJob struct {
	EmpId int64 `json: "EmpId"`
	JobId int64 `json: "JobId"`
}

type MakingJob struct {
	EmpId          int64     `json: "EmpId"`
	JobId          int64     `json: "JobId"`
	TotalItemsMade int64     `json: "TotalItemsMade"`
	MakingDate     time.Time `json: "MakingDate"`
	PaidAmount     int64     `json: "PaidAmount"`
}

type LoginInfo struct {
	MobileNo string `json: "MobileNo"`
	Password string `json: "Password"`
}
