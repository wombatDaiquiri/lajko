package hejto

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/wombatDaiquiri/lajko/database"
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
	defer resp.Body.Close()

	respB, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//	fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\nbody: %s\n\n\n\n\n\n\n\n\n\n\n\n\n", respB)
	var postResp postResponse
	err = json.Unmarshal(respB, &postResp)
	if err != nil {
		fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\nbody: %s\n\n\n\n\n\n\n\n\n\n\n\n\n", respB)
		return nil, err
	}

	return postResp.DatabasePosts(), nil
}
