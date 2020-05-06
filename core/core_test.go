package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexOf(t *testing.T) {
	values := []string{
		"PwDW1H7X", "bAUx72As", "/Hw0+pjR", "nVBO+e8X", "COvZVrN1",
		"ISA9nv9N", "spGrpKz9", "Xp/Kr03I", "dE733qOP", "dNPoSOOX",
	}

	assert.Equal(t, 0, IndexOf(values, "PwDW1H7X"))
	assert.Equal(t, 9, IndexOf(values, "dNPoSOOX"))
	assert.Equal(t, 7, IndexOf(values, "Xp/Kr03I"))
	assert.Equal(t, -1, IndexOf(values, "Kr03I"))
}

func TestSendEmail(t *testing.T) {
	smtp := GetSMTP()
	smtp.Init(SMTPConfig{
		Username: "thanh.vinh@hotmail.com",
		Password: "",
		Server:   "smtp.office365.com",
		Port:     587,
	})

	// assert.NoError(t, smtp.Send("thanh.vinh@hotmail.com", "thanh.vinh@hotmail.com", "Hi", "You"))
}
