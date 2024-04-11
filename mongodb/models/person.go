package models

type Person struct {
	//Идентификатор сотрудника
	PersonId string `bson:"person_id"`
	//ЛогиеИд в домене `mivanov`
	DomainId string `bson:"login_id"`
	//электронная почта
	Email string `bson:"email"`
	//Телефон в 1С
	Phone string `bson:"phone"`
	//ТелеграмЛогин @ivanov
	Telegram string `bson:"telegram"`
	//Телеграм Ид
	TelegramId string `bson:"telegram_id"`
	//Код ключа
	PassId string `bson:"pass_id"`
}
