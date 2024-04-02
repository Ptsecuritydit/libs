package db_models

// Department Описание объекта подразделение (отдел, управление, группа)
type Department struct {
	//Идентификатор 1С подразделения
	Id string `json:"id"`
	//Наименование подразделения
	Title string `json:"title"`
	//дентификатор вышестоящего подразделения
	ParentId string `json:"parent_id"`
	//Идентификатор руководителя подразделения
	ManagerId string `json:"manager_id"`
	//Идентификатор должности руководителя подразделения
	ManagerPositionId string `json:"manager_position_id"`
	//Признак удаления элемента
	IsRemoved bool `json:"is_removed"`
}
