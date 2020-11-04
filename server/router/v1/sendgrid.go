package v1

import (
	"fmt"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/cosmos/atlas/config"
)

func (r *Router) sendEmailConfirmation(name, email, confirmURL string) error {
	apiKey := r.cfg.String(config.SendGridAPIKey)
	if apiKey == "" {
		r.logger.Warn().Msg("cannot send confirmation email; sendgrid api key is empty")
		return nil
	}

	r.logger.Debug().Msg("sending email confirmation")

	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(name, email))
	p.SetDynamicTemplateData("confirmURL", confirmURL)

	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("Atlas", "contact@atlas.cosmos.network"))
	m.SetTemplateID("d-93c251d606024aefaa2d35c2a445bb30")
	m.AddPersonalizations(p)

	client := sendgrid.NewSendClient(apiKey)
	resp, err := client.Send(m)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send email; unexpected status code (%d): %s", resp.StatusCode, resp.Body)
	}

	return nil
}
