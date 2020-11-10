package v1

import (
	"fmt"
	"net/http"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/cosmos/atlas/config"
	"github.com/cosmos/atlas/server/models"
)

func (r *Router) sendEmailConfirmation(name, email, confirmURL string) error {
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(name, email))
	p.SetDynamicTemplateData("confirmURL", confirmURL)

	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("Atlas", "contact@atlas.cosmos.network"))
	m.SetTemplateID("d-93c251d606024aefaa2d35c2a445bb30")
	m.AddPersonalizations(p)

	r.logger.Debug().Msg("sending email confirmation")
	return r.sendEmail(m)
}

func (r *Router) sendOwnerInvitation(acceptURL, invitedBy string, invitee models.User, module models.Module) error {
	p := mail.NewPersonalization()
	p.AddTos(mail.NewEmail(invitee.Name, invitee.Email.String))
	p.SetDynamicTemplateData("acceptURL", acceptURL)
	p.SetDynamicTemplateData("invitedBy", invitedBy)
	p.SetDynamicTemplateData("module", module.Name)
	p.SetDynamicTemplateData("team", module.Team)

	m := mail.NewV3Mail()
	m.SetFrom(mail.NewEmail("Atlas", "contact@atlas.cosmos.network"))
	m.SetTemplateID("d-48f6c83a2072495fb359ceb600e87d18")
	m.AddPersonalizations(p)

	r.logger.Debug().Msg("sending module owner invitation")
	return r.sendEmail(m)
}

func (r *Router) sendEmail(msg *mail.SGMailV3) error {
	apiKey := r.cfg.String(config.SendGridAPIKey)
	if apiKey == "" {
		r.logger.Warn().Msg("cannot send email; sendgrid api key is empty")
		return nil
	}

	client := sendgrid.NewSendClient(apiKey)
	resp, err := client.Send(msg)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("failed to send email; unexpected status code (%d != %d): %s", resp.StatusCode, http.StatusAccepted, resp.Body)
	}

	return nil
}
