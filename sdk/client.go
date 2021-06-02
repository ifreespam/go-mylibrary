package sdk

import (
	"github.com/ifreespam/go-mylibrary/internal/logger"
	"strconv"
)

// Client is a main struct of the client
type Client struct {
	host string
	port int

	log *logger.CustomLogger
}

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
