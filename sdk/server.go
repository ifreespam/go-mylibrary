package sdk

import "github.com/ifreespam/go-mylibrary/v2/internal/logger"

// DoSomething is doing something :3
func DoSomething() {
	logger.GetLogger("server").Log("Just do something here!!111")
}

// DoMoreSomething should have a comment
func DoMoreSomething() {
	logger.GetLogger("server").Log("Just do something here!!111 12321312")
}
