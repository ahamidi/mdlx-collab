package main

import (
	"fmt"
	"log"
	"net/http"
	"net/mail"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/forms"
	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

const (
	TenantHost = "mycorp.na1.my.cloud"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	thisTenant := Tenant{
		Host:    TenantHost,
		Company: "MyCorp",
	}

	// join handler
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/join/:token", func(c echo.Context) error {
			token := c.PathParam("token")

			// verify token for this user
			inviteRecord, err := app.Dao().FindFirstRecordByData("invites", "token", token)
			if err != nil {
				log.Printf("error finding invite record: %v", err)
				return err
			}

			// if valid, set user to verified, mark invite as accepted
			inviteRecord.Set("accepted", true)
			err = app.Dao().Save(inviteRecord)
			if err != nil {
				return err
			}

			// expand record
			if errs := app.Dao().ExpandRecord(inviteRecord, []string{"invited_user"}, nil); len(errs) > 0 {
				return fmt.Errorf("failed to expand: %v", errs)
			}
			userRecord := inviteRecord.ExpandedOne("invited_user")
			err = userRecord.SetVerified(true)
			if err != nil {
				return err
			}
			err = app.Dao().Save(userRecord)

			return c.JSON(http.StatusOK, map[string]string{"message": "successfully joined " + thisTenant.Company})
		} /* optional middlewares */)

		return nil
	})

	// after user create hook
	app.OnRecordAfterCreateRequest("users").Add(func(e *core.RecordCreateEvent) error {
		// create invite record
		collection, err := app.Dao().FindCollectionByNameOrId("invites")
		if err != nil {
			return err
		}

		record := models.NewRecord(collection)
		form := forms.NewRecordUpsert(app, record)

		inviterUserID := apis.RequestInfo(e.HttpContext).AuthRecord.GetId()
		invitedUserID := e.Record.GetId()
		// generated verification token
		token := uuid.New().String()
		form.LoadData(map[string]any{
			"inviter":      inviterUserID,
			"invited_user": invitedUserID,
			"token":        token,
			"accepted":     false,
		})

		// validate and submit
		if err := form.Submit(); err != nil {
			return err
		}

		// send invite email
		message, err := newInviteMessage(e.Record.GetString("name"), e.Record.Email(), thisTenant, token)
		if err != nil {
			return err
		}

		return app.NewMailClient().Send(message)
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

type Tenant struct {
	Host    string `json:"host"`
	Company string `json:"company"`
}

func newInviteMessage(name string, email string, tenant Tenant, token string) (*mailer.Message, error) {
	subject := "You've been invited to join " + tenant.Company
	body := "Hi " + name + ",\n\n" +
		"You've been invited to join " + tenant.Company + " on Meroxa.\n\n" +
		"Click the link below to get started:\n\n" +
		"https://" + tenant.Host + "/join/" + token + "\n\n" +
		"Thanks,\n" +
		"Team Meroxa"

	message := &mailer.Message{
		From: mail.Address{
			Address: email,
			Name:    name,
		},
		To:      []mail.Address{{Address: email}},
		Subject: subject,
		HTML:    body,
	}

	return message, nil
}
