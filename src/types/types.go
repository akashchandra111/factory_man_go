package types

type Employee struct {
	EmpId      int64
	FirstName  string
	LastName   string
	DOB        string
	MobileNo   string
	EmpType    string
	SubEmpType string
	IsWorking  bool
	IsSalaried bool
	Salary     int64
	Advance    int64
	JobId      int64
}

type Salary struct {
	EmpId        int64
	SalaryAmount int64
	DateOfCredit string
}

type AdvanceSalary struct {
	EmpId         int64
	AdvanceAmount int64
	AdvanceDate   string
}

type Job struct {
	JobId          int64
	JobType        string
	JobSubType     string
	JobSubTypeName string
	IsRcOrStone    byte
}

type PersonJob struct {
	EmpId int64
	JobId int64
}

type MakingJob struct {
	EmpId          int64
	JobId          int64
	TotalItemsMade int64
	MakingDate     string
	PaidAmount     int64
}
