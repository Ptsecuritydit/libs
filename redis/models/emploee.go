package db_models

// Employee Описание объекта сотрудника
type Employee struct {
	//Идентификатор сотрудника
	Id string `json:"id"`
	//Идентификатор физического лица
	Person string `json:"person"`
}
