package db_models

type EmployeesHrBp struct {
	EmployeeId string `json:"employee_id"`
	HrBp       []HrBp `json:"hr_bp"`
}

type HrBp struct {
	HrBpId string `json:"hr_id"`
}
