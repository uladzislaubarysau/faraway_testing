package client

import (
	"errors"
	"log"
	"net"

	"github.com/uladzislaubarysau/faraway_testing/internal/helpers"
)

type Client struct {
	address string
}

func New(address string) *Client {
	return &Client{address}
}

func (c *Client) GetWow() (string, error) {
	conn, err := net.Dial("tcp", c.address)
	if err != nil {
		return "", err
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("can't close connection:", err)
		}
	}()

	if err = helpers.WriteMessage(conn, []byte("Hello there")); err != nil {
		return "", err
	}

	challenge, err := helpers.ReadMessage(conn)
	if err != nil {
		return "", err
	}

	response := helpers.Solve(challenge)

	if err = helpers.WriteMessage(conn, response); err != nil {
		return "", err
	}

	wow, err := helpers.ReadMessage(conn)
	if err != nil {
		return "", err
	}

	if err == nil && len(wow) == 0 {
		err = errors.New("empty response")
	}

	return string(wow), err
}
