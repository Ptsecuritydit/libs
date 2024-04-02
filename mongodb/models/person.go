package models

type Person struct {
	PersonId   string `bson:"person_id"`
	DomainId   string `bson:"login_id"`
	Email      string `bson:"email"`
	Phone      string `bson:"phone"`
	Telegram   string `bson:"telegram"`
	TelegramId string `bson:"telegram_id"`
}
