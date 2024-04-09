package db_models

type EmployeesHrPartner struct {
	EmployeeID string      `redis:"employee_id"`
	HrPartners []HrPartner `redis:"hr_partner"`
}

type HrPartner struct {
	HrPartnerId string `redis:"hr_id"`
}
