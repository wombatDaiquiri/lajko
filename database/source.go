package database

import (
	"encoding/json"
	"time"
)

// SourcePortal marks which portal was the post found on.
type SourcePortal string

const (
	SourcePortalHejto SourcePortal = "hejto"
)

// Source contains data about the original source of the post.
type Source struct {
	Source          SourcePortal
	SourceData      json.RawMessage
	SourceID        string
	SourceURL       string
	SourceCreatedAt time.Time
}
