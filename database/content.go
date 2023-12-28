package database

import (
	"encoding/json"
)

type Content struct {
	TextContent `prefix:"text_"`
	Attachments `prefix:"attachments_"`
}

type TextContent struct {
	HTML     string
	Plain    string
	Markdown string
}

type Attachments struct {
	// Images contain URLs to images attached to a post.
	//
	// Images should be deduplicated before saving. If there are multiple resolutions
	// of image available, choose the highest resolution (at least for now).
	Images []string
	// Other are attachments other than images.
	Other json.RawMessage
	// many to many
	Tags []Tag
}
