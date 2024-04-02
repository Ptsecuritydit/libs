package db_models

// EmployeePosition Список должностей сотрудников
type EmployeePosition struct {
	//Идентификатор связи
	Id string `json:"id"`
	//Идентификатор 1С должности
	PositionId string `json:"position_id"`
	//Идентификатор 1С сотрудника
	EmployeeId string `json:"employee_id"`
	//Основная должность
	IsMain bool `json:"is_main"`
	//Дата принятия на должность
	ReceiptDt string `json:"receipt_dt"`
	//Дата увольнения с должности
	DismissalDt string `json:"dismissal_dt"`
	//Идентификатор 1С подразделения
	DepartmentId string `json:"department_id"`
}
