package database

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/oklog/ulid"
	"gorm.io/gorm"
)

type Model struct {
	gorm.Model
	ULID string
}

func (m *Model) BeforeCreate(*gorm.DB) error {
	newULID, err := ulid.New(uint64(time.Now().UnixNano()), rand.Reader)
	if err != nil {
		return fmt.Errorf("could not generate ulid: %w", err)
	}
	m.ULID = newULID.String()
	return nil
}
