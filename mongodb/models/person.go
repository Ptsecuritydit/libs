package models

type Person struct {
	//Идентификатор сотрудника
	PersonId string `bson:"person_id"`
	//ЛогиеИд в домене `mivanov`
	DomainId string `bson:"login_id"`
	//электронная почта
	Email string `bson:"email"`
	//телефон в 1С
	Phone string `bson:"phone"`
	//ТелеграмЛогин @ivanov
	Telegram string `bson:"telegram"`
	//Телеграм Ид
	TelegramId string `bson:"telegram_id"`
}
