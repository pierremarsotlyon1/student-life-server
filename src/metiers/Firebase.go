package metiers

import (
	"github.com/NaySoftware/go-fcm"
	"errors"
)

type FirebaseMetier struct {}

const (
	serverKey = "AIzaSyAaifblESuQNaUzQ9YzacFYTLgi0APvEi4"
)

func (*FirebaseMetier) SendMessageToTokens (idsToken []string, data map[string]string) error {
	c := fcm.NewFcmClient(serverKey)
	c.NewFcmRegIdsMsg(idsToken, data)

	if _, err := c.Send(); err != nil {
		return errors.New("Erreur lors de l'envoi de la notification")
	}

	return nil
}
