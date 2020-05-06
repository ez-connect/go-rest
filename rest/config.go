package rest

import "time"

type ServerConfig struct {
	Host          string        `yaml:"host"`
	Port          int           `yaml:"port"`
	BodyLimit     string        `yaml:"bodyLimit"`
	Gzip          bool          `yaml:"gzip"`
	GzipLevel     int           `yaml:"gzipLevel"`
	CacheCapacity int           `yaml:"cacheCapacity"`
	CacheTTL      time.Duration `yaml:"cacheTTL"`
}
