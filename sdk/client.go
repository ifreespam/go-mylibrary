package sdk

import (
	"github.com/ifreespam/go-mylibrary/v2/internal/logger"
	"strconv"
)

// Client is a main struct of the client
type Client struct {
	host string
	port int

	log *logger.CustomLogger
}

// PrintInfo should have a comment
func (c *Client) PrintInfo() {
	c.log.Log("HOST: " + c.host + ", PORT: " + strconv.Itoa(c.port))
}

// NewClient is a factory of clients
func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
		log:  logger.GetLogger("client"),
	}
}
