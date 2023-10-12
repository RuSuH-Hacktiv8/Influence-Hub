package notification

import (
	"influence-hub-notification/models"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SuccessRegisterEmail(user *models.User) error {
	from := mail.NewEmail("Influencer Hub", "mhandyalfurqon@gmail.com")
	subject := "Konfirmasi Pendaftaran"
	to := mail.NewEmail(user.Username, user.Email)
	plainTextContent := "Terima kasih telah mendaftar!"
	htmlContent, err := os.ReadFile("success_register.html")
	if err != nil {
		return err
	}
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, string(htmlContent))

	client := sendgrid.NewSendClient("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") // Ganti dengan API key SendGrid kamu
	_, err = client.Send(message)

	return err
}

func SuccessApplyCampaign(user *models.User) error {
	from := mail.NewEmail("Influencer Hub", "mhandyalfurqon@gmail.com")
	subject := "Konfirmasi apply Campaign"
	to := mail.NewEmail(user.Username, user.Email)
	plainTextContent := "Terima kasih telah mendaftar!"
	htmlContent, err := os.ReadFile("success_apply_campaign.html")
	if err != nil {
		return err
	}
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, string(htmlContent))

	client := sendgrid.NewSendClient("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")
	_, err = client.Send(message)

	return err
}
