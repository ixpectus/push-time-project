package user_online

import (
	"net/http"

	"github.com/ixpectus/push-time-project/pkg/helpers"
)

type Client struct {
	c *http.Client
}

func New() Client {
	return Client{
		c: http.DefaultClient,
	}
}

func (g *Client) Get(userID int64) (error, *helpers.UserActivity) {
	// тут как будто бы настоящий клиент
	return nil, &helpers.UserActivity{}
}
