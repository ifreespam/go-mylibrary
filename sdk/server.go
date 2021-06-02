package sdk

import "github.com/ifreespam/go-mylibrary/internal/logger"

// DoSomething is doing something :3
func DoSomething() {
	logger.GetLogger("server").Log("Just do something here!!111")
}
