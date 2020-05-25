package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendMail(t *testing.T) {
	GetSMTP().Init(SMTPConfig{
		Username: "",
		Password: "",
		Server:   "smtp.office365.com",
		Port:     587,
	})

	assert.NoError(t, GetSMTP().Send(
		"recipient@hotmail.com",
		"Hello",
		"How are you?",
		"customer@example.com",
		[]string{"a@example.com", "b@example.com"},
	))
}
