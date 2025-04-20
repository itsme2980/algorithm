package loadbalancer

import (
	"time"
)

func Retry(operation func() error, retries int, retryInterval time.Duration) error {
	var err error
	for i := 0; i < retries; i++ {
		err = operation()
		if err == nil {
			return nil
		}
		time.Sleep(retryInterval)
	}
	return err
}
