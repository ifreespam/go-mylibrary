package sdk

import (
	"github.com/ifreespam/go-mylibrary/internal/logger"
	"strconv"
)

// Client is a main struct of the client
type Client struct {
	host string
	port int
}

// NewClient is a factory of clients
func NewClient(host string, port int) *Client {
	logger.GetLogger("client").Log("HOST: " + host + ", PORT: " + strconv.Itoa(port))
	return &Client{
		host: host,
		port: port,
	}
}
