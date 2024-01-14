package hejto

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/wombatDaiquiri/lajko/database"
	"github.com/wombatDaiquiri/lajko/ee"
)

type Client struct{}

func (c *Client) Posts(ctx context.Context, pagination PostPagination) ([]database.Post, error) {
	req, err := http.NewRequest(http.MethodGet, "https://api.hejto.pl/posts?"+pagination.Query(), nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer ee.CloseHTTPResponse(resp)

	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var postResp postResponse
	err = json.Unmarshal(respB, &postResp)
	if err != nil {
		return nil, err
	}

	return postResp.DatabasePosts(), nil
}
