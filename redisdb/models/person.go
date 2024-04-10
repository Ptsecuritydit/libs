package db_models

// Person Данные физического лица
type Person struct {
	//Идентификатор 1с физического лица
	Id string `json:"id"`
	//Имя
	FirstName string `json:"first_name"`
	//Отчество
	MiddleName string `json:"middle_name"`
	//Фамилия
	LastName string `json:"last_name"`
	//Пол [MALE, FEMALE]
	Sex string `json:"sex"`
	//Дата рождения"
	DateOfBirth string `json:"dob"`
	//Город рабочего места
	City string `json:"city"`
	//номер карты сотрудника
	PassId string `json:"pass_id"`
	//LoginAD доменный идентификатор сотрудника
	DomainId string `json:"login_id"`
	//почта физического лица
	Email string `json:"email"`
	//Телефон
	Phone string `json:"phone"`
	//Телеграм
	TelegramId string `json:"telegram_id"`
	//Списов id Сотрудников
	Employees []string `json:"employees"`
}
