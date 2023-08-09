package entity

import (
	"crypto/rand"
	"encoding/base32"
	"io"
	"strings"
	"time"
)

type Session struct {
	ID        string    `db:"id"`
	UserID    int       `db:"userId"`
	ExpiresAt time.Time `db:"expiresAt"`
}

func (s *Session) SetSessionID() error {
	sidByte := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, sidByte)
	if err != nil {
		return err
	}
	sessionID := strings.TrimRight(base32.StdEncoding.EncodeToString(sidByte), "=")
	s.ID = sessionID

	return nil
}
