package db_models

// Position Список должностей
type Position struct {
	//Идентификатор 1С должности
	Id string `json:"id"`
	//Наименование должности
	Title string `json:"title"`
}
