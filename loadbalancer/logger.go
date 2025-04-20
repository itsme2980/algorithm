package loadbalancer

import (
	"log"
)

func LogInfo(format string, args ...interface{}) {
	log.Printf("[INFO] "+format, args...)
}

func LogWarn(format string, args ...interface{}) {
	log.Printf("[WARN] "+format, args...)
}
