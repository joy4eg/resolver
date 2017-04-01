package resolver

import (
	"time"
)

// Resolver configuration
type Config struct {
	Servers  []string      // servers to use
	Timeout  time.Duration // a cumulative timeout for dial, write and read, defaults to 2s
	Attempts int           // Total number of resolving attempts
}

// Default resolver configuration
var (
	ConfigDefault = Config{
		Servers: []string{
			"208.67.220.220", // OpenDNS
			"8.8.4.4",        // Google DNS
			"64.6.65.6",      // Verisign DNS
		},
		Timeout:  2 * time.Second,
		Attempts: 2,
	}
)

var (
	config = ConfigDefault
)

// SetConfig sets the resolver configuration
func SetConfig(c Config) {
	config = c
}
