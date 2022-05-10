package notification

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/mail"
	"time"

	"github.com/phk13/host-monitor/pkg/utils"
	log "github.com/sirupsen/logrus"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

type GmailNotificationService struct {
	mailSrv  *gmail.Service
	mailAddr string
	test     bool
}

func (srv *GmailNotificationService) SendNotification(hostIP string, timestamp time.Time) {
	// If test flag is true, only log and return
	if srv.test {
		log.Infof("Test notification for %s completed.", hostIP)
		return
	}
	// Compose message
	var message gmail.Message

	messageStr := []byte(fmt.Sprintf(
		"From: %s\r\n"+
			"To: %s\r\n"+
			"Subject: Host %s is down\r\n\r\n"+
			"Host %s has been detected down at %s", srv.mailAddr, srv.mailAddr, hostIP, hostIP, timestamp.String()))

	// Place messageStr into message.Raw in base64 encoded format
	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	// Send the message
	_, err := srv.mailSrv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Errorf("Error sending notification: %v", err)
	} else {
		log.Debugf("Notification email for %s sent correctly.", hostIP)
	}
}

func NewMailNotificationService(mailAddr string, test bool) *GmailNotificationService {
	if _, err := mail.ParseAddress(mailAddr); err != nil {
		log.Fatal("Invalid email")
	}

	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := utils.GetOAuthClientClient(config)

	srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	return &GmailNotificationService{
		mailSrv:  srv,
		mailAddr: mailAddr,
		test:     test,
	}
}
