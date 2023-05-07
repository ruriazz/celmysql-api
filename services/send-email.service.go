package services

import (
	"context"
	"database/sql"
	"net/smtp"

	"github.com/go-playground/validator/v10"
)

type SendEmailService struct {
	DB       *sql.DB
	Validate *validator.Validate
}

func NewSendEmailService(DB *sql.DB, validate *validator.Validate) ISendEmailService {
	return &SendEmailService{
		DB:       DB,
		Validate: validate,
	}
}

func (service *SendEmailService) SendEmail(ctx context.Context) {
	from := "ariviasimoes@gmail.com"
	password := "Serverh5n783"

	toEmailAddress := "weslyaioria@gmail.com"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "Subject: This is the subject of the mail\n"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		panic(err)
	}
	// return mapping.ToSatuanResponses(satuans)
}
