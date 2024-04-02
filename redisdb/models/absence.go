package db_models

// Absence (LongAbsence) Отсутствия сотрудника/Длительные отсутствия сотрудника
type Absence struct {
	//Идентификатор сотрудника
	EmployeeId string `json:"employee_id"`
	// Причина отсутствия: отпуск, болезнь или иные причины / Причина длительного отсутствия: декретный отпуск, иное
	Cause string `json:"cause"`
	// Дата начала отсутствия
	StartedDt string `json:"started_dt"`
	// Дата завершения отсутствия
	EndedDt string `json:"ended_dt"`
}
