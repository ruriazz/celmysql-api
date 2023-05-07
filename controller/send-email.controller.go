package controller

import (
	"net/smtp"

	"github.com/celmysql-api/services"
	"github.com/gin-gonic/gin"
)

type SendEmailController struct {
	SendEmailService services.ISendEmailService
}

func NewSendEmailController(satuanService services.ISendEmailService) ISendEmailController {
	return &SendEmailController{
		SendEmailService: satuanService,
	}
}

// Send Email
// @Summary List existing Send Email
// @Description Get all the existing Send Email
// @Tags Send Email
// @Accept  json
// @Produce  json
// @Success 200 {object} common.DefaultResponse	"ok"
// @Security ApiKeyAuth
// @Router /send-email [post]
func (controller *SendEmailController) SendEmail(c *gin.Context) {
	from := "adm.teknis2020@gmail.com"
	password := "ovzwexfbcoklgyni"

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
}
