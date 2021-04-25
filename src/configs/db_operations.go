package configs

import (
	"database/sql"
	"os"
	"time"

	"github.com/akashchandra111/factory_man/src/types"
	_ "github.com/mattn/go-sqlite3"
)

const (
	selectUserByMobile     string = "SELECT EMP_ID, FIRST_NAME, LAST_NAME, DOB, MOBILE_NO, PASSWORD, EMP_TYPE, SUB_EMP_TYPE, IS_WORKING, IS_SALARIED, SALARY, ADVANCE, JOB_ID FROM EMPLOYEE WHERE MOBILE_NO=? AND PASSWORD=?"
	selectJobById          string = "SELECT JOB_ID, JOB_TYPE, JOB_SUBTYPE, JOB_SUBTYPE_NAME, IS_RC_OR_STONE FROM JOB WHERE JOB_ID=?"
	selectAllJobs          string = "SELECT JOB_ID, JOB_TYPE, JOB_SUBTYPE, JOB_SUBTYPE_NAME, IS_RC_OR_STONE FROM JOB"
	selectEmpsByJobId      string = "SELECT EMP_ID, FIRST_NAME, LAST_NAME, DOB, MOBILE_NO, EMP_TYPE, SUB_EMP_TYPE, IS_WORKING, IS_SALARIED, SALARY, ADVANCE, JOB_ID FROM EMPLOYEE WHERE JOB_ID=?"
	selectSalaryByEmpId    string = "SELECT EMP_ID, SALARY_AMOUNT, DATE_OF_CREDIT FROM SALARY WHERE EMP_ID=?"
	selectAdvanceByEmpId   string = "SELECT EMP_ID, ADVANCE_AMOUNT, ADVANCE_DATE FROM ADVANCE_SALARY WHERE EMP_ID=?"
	selectMakingJobByEmpId string = "SELECT EMP_ID, JOB_ID, TOTAL_ITEMS_MADE, MAKING_DATE, PAID_AMOUNT FROM MAKING_JOB WHERE EMP_ID=?"
	selectMakingJobByJobId string = "SELECT EMP_ID, JOB_ID, TOTAL_ITEMS_MADE, MAKING_DATE, PAID_AMOUNT FROM MAKING_JOB WHERE JOB_ID=?"
	selectMakingJobByDate  string = "SELECT EMP_ID, JOB_ID, TOTAL_ITEMS_MADE, MAKING_DATE, PAID_AMOUNT FROM MAKING_JOB WHERE MAKING_DATE=?"

	insertUser          string = "INSERT INTO EMPLOYEE(FIRST_NAME, LAST_NAME, DOB, MOBILE_NO, PASSWORD, EMP_TYPE, SUB_EMP_TYPE, IS_WORKING, IS_SALARIED, SALARY, ADVANCE, JOB_ID) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)"
	insertJob           string = "INSERT INTO JOB(JOB_TYPE, JOB_SUBTYPE, JOB_SUBTYPE_NAME, IS_RC_OR_STONE) VALUES (?,?,?,?)"
	insertSalary        string = "INSERT INTO SALARY(EMP_ID, SALARY_AMOUNT, DATE_OF_CREDIT) VALUES (?,?,?)"
	insertAdvanceSalary string = "INSERT INTO ADVANCE_SALARY(EMP_ID, ADVANCE_AMOUNT, ADVANCE_DATE) VALUES (?,?,?)"
	insertMakingJob     string = "INSERT INTO MAKING_JOB(EMP_ID, JOB_ID, TOTAL_ITEMS_MADE, MAKING_DATE, PAID_AMOUNT) VALUES (?,?,?,?,?)"
)

var dbRef *sql.DB = nil

func InitDB() error {
	db, err := sql.Open("sqlite3", "./factory_man.db")
	dbRef = db
	checkErrAndWrite(err, "Error, not able to connect to db")
	return err
}

func checkErrAndWrite(err error, msg string) {
	if err != nil {
		os.Stderr.WriteString(msg)
		panic(err)
	}
}

func GetUser(userInfo types.LoginInfo) types.Employee {
	var emp types.Employee = types.Employee{}

	rows, err := dbRef.Query(selectUserByMobile, userInfo.MobileNo, userInfo.Password)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		err := rows.Scan(&emp.EmpId, &emp.FirstName, &emp.LastName, &emp.DOB, &emp.MobileNo, &emp.Password,
			&emp.EmpType, &emp.SubEmpType, &emp.IsWorking, &emp.IsSalaried, &emp.Salary, &emp.Advance, &emp.JobId)
		checkErrAndWrite(err, "Error occured while scanning the data from row")
	}

	return emp
}

func GetJob(id int64) types.Job {
	var job types.Job = types.Job{}

	rows, err := dbRef.Query(selectJobById, id)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		err := rows.Scan(&job.JobId, &job.JobType, &job.JobSubType, &job.JobSubTypeName, &job.IsRcOrStone)
		checkErrAndWrite(err, "Error occured while scanning the data from row")
	}

	return job
}

func GetAllJobs() []types.Job {
	var jobs []types.Job = []types.Job{}

	rows, err := dbRef.Query(selectAllJobs)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		job := types.Job{}
		err := rows.Scan(&job.JobId, &job.JobType, &job.JobSubType, &job.JobSubTypeName, &job.IsRcOrStone)
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		jobs = append(jobs, job)
	}

	return jobs
}

func GetEmployeesByJobId(jobId int64) []types.Employee {
	var emps []types.Employee = []types.Employee{}

	rows, err := dbRef.Query(selectEmpsByJobId, jobId)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		emp := types.Employee{}
		err := rows.Scan(&emp.EmpId, &emp.FirstName, &emp.LastName, &emp.DOB, &emp.MobileNo,
			&emp.EmpType, &emp.SubEmpType, &emp.IsWorking, &emp.IsSalaried, &emp.Salary, &emp.Advance, &emp.JobId)
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		emps = append(emps, emp)
	}

	return emps
}

func InsertUser(emp *types.Employee) int64 {
	stmt, err := dbRef.Prepare(insertUser)
	checkErrAndWrite(err, "Error occured while preparing the SQL statement for insertion")
	res, err := stmt.Exec(emp.FirstName, emp.LastName, emp.DOB, emp.MobileNo, emp.Password, emp.EmpType, emp.SubEmpType, emp.IsWorking, emp.IsSalaried, emp.Salary, emp.Advance, emp.JobId)
	checkErrAndWrite(err, "Error occured while executing the SQL statement for insertion")

	rowsAffected, err := res.RowsAffected()
	checkErrAndWrite(err, "Error occured while checking the affected rows")
	return rowsAffected
}

func InsertJob(job *types.Job) int64 {
	stmt, err := dbRef.Prepare(insertJob)
	checkErrAndWrite(err, "Error occured while preparing the SQL statement for insertion")
	res, err := stmt.Exec(job.JobType, job.JobSubType, job.JobSubTypeName, job.IsRcOrStone)
	checkErrAndWrite(err, "Error occured while executing the SQL statement for insertion")

	rowsAffected, err := res.RowsAffected()
	checkErrAndWrite(err, "Error occured while checking the affected rows")
	return rowsAffected
}

func InsertSalary(salary *types.Salary) int64 {
	stmt, err := dbRef.Prepare(insertSalary)
	checkErrAndWrite(err, "Error occured while preparing the SQL statement for insertion")
	res, err := stmt.Exec(salary.EmpId, salary.SalaryAmount, salary.DateOfCredit)
	checkErrAndWrite(err, "Error occured while executing the SQL statement for insertion")

	rowsAffected, err := res.RowsAffected()
	checkErrAndWrite(err, "Error occured while checking the affected rows")
	return rowsAffected
}

func InsertAdvance(advance *types.AdvanceSalary) int64 {
	stmt, err := dbRef.Prepare(insertAdvanceSalary)
	checkErrAndWrite(err, "Error occured while preparing the SQL statement for insertion")
	res, err := stmt.Exec(advance.EmpId, advance.AdvanceAmount, advance.AdvanceDate)
	checkErrAndWrite(err, "Error occured while executing the SQL statement for insertion")

	rowsAffected, err := res.RowsAffected()
	checkErrAndWrite(err, "Error occured while checking the affected rows")
	return rowsAffected
}

func InsertMakingJob(making *types.MakingJob) int64 {
	stmt, err := dbRef.Prepare(insertMakingJob)
	checkErrAndWrite(err, "Error occured while preparing the SQL statement for insertion")
	res, err := stmt.Exec(making.EmpId, making.JobId, making.TotalItemsMade, making.MakingDate, making.PaidAmount)
	checkErrAndWrite(err, "Error occured while executing the SQL statement for insertion")

	rowsAffected, err := res.RowsAffected()
	checkErrAndWrite(err, "Error occured while checking the affected rows")
	return rowsAffected
}

func GetSalaryByEmpId(empId int64) []types.Salary {
	var salaries []types.Salary = []types.Salary{}

	rows, err := dbRef.Query(selectSalaryByEmpId, empId)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		salary := types.Salary{}
		err := rows.Scan()
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		salaries = append(salaries, salary)
	}

	return salaries
}

func GetAdvanceSalaryByEmpId(empId int64) []types.AdvanceSalary {
	var advances []types.AdvanceSalary = []types.AdvanceSalary{}
	rows, err := dbRef.Query(selectAdvanceByEmpId, empId)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		advance := types.AdvanceSalary{}
		err := rows.Scan()
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		advances = append(advances, advance)
	}

	return advances
}

func GetMakingJobByEmpId(empId int64) []types.MakingJob {
	var makingJobs []types.MakingJob = []types.MakingJob{}
	rows, err := dbRef.Query(selectMakingJobByEmpId, empId)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		makingJob := types.MakingJob{}
		err := rows.Scan()
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		makingJobs = append(makingJobs, makingJob)
	}

	return makingJobs
}

func GetMakingJobByJobId(jobId int64) []types.MakingJob {
	var makingJobs []types.MakingJob = []types.MakingJob{}
	rows, err := dbRef.Query(selectMakingJobByJobId, jobId)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		makingJob := types.MakingJob{}
		err := rows.Scan()
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		makingJobs = append(makingJobs, makingJob)
	}

	return makingJobs
}

func GetMakingJobByDate(date time.Time) []types.MakingJob {
	var makingJobs []types.MakingJob = []types.MakingJob{}
	rows, err := dbRef.Query(selectMakingJobByDate, date)
	checkErrAndWrite(err, "Error occured while executing query for getting user details")

	for rows.Next() {
		makingJob := types.MakingJob{}
		err := rows.Scan()
		checkErrAndWrite(err, "Error occured while scanning the data from row")
		makingJobs = append(makingJobs, makingJob)
	}

	return makingJobs
}
