package check

import (
	"net"
	"strconv"
	"time"
)

func IsValidAppHost(host string) bool {
	// Check if the app host is valid
	if ip := net.ParseIP(host); ip != nil {
		return true
	}
	return false
}

func IsValidAppPort(port string) bool {
	// Check if the app port is valid
	if port, err := strconv.Atoi(port); err == nil {
		if port >= 0 && port <= 65535 {
			return true
		}
	}
	return false
}

func IsValidLogLevel(level string) bool {
	// Check if the log level is valid
	switch level {
	case "debug", "info", "warn", "error", "dpanic", "panic", "fatal":
		return true
	}
	return false
}

func IsValidTimeRange(start, end time.Time) bool {
	// Check if the time range is valid
	if start.Before(end) {
		return true
	}
	return false
}
