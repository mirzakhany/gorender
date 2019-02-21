package storage

import "time"

// KVStorage interface
type KVStorage interface {
	Init() error
	Reset()
	IncIntKey(key string, value int64) error
	DecIntKey(key string, value int64) error
	SetIntKey(key string, value int64, exp time.Duration) error
	GetIntKey(key string) (int64, error)
	SetString(key string, value string, exp time.Duration) error
	GetString(key string) (string, error)
	RemoveKeys(keys ...string) error
}

// ConfStatus status package settings
type ConfStatus struct {
	Engine  string         `json:"engine"`
	Redis   SectionRedis   `json:"redis"`
	BoltDB  SectionBoltDB  `json:"boltdb"`
	LevelDB SectionLevelDB `json:"leveldb"`
}

// SectionRedis is sub section of config.
type SectionRedis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

// SectionBoltDB is sub section of config.
type SectionBoltDB struct {
	Path   string `json:"path"`
	Bucket string `json:"bucket"`
}

// SectionLevelDB is sub section of config.
type SectionLevelDB struct {
	Path string `json:"path"`
}
