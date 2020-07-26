package utils

import "time"

/**
 * local config struct
 */
type LocalConfig struct {
	Name             string    `json:"name"`
	ServerAddr       string    `json:"server_addr"`
	LoggerLevel      *string   `json:"log_level"`
	VortexPort       *int      `json:"vortex_port"`
	DatabaseHost     *string   `json:"database_host"`
	DatabaseTable    *string   `json:"database_table"`
	DatabaseUsername *string   `json:"database_username"`
	DatabasePassword *string   `json:"database_password"`
	JwtSecret        *string   `json:"jwt_secret"`
	CacheDeadline    time.Time `json:"-"`
	Version          *string   `json:"version"`
	VersionTag       *string   `json:"versionTag"`
}

var config = new(LocalConfig)

func GetConfig() *LocalConfig {
	if config.CacheDeadline.Before(time.Now()) {
		if err := LoadConfig("cnf/conf.json", config); err != nil {
			Fatal("loading file failed")
			return nil
		}
		config.CacheDeadline = time.Now().Add(time.Second * 60)
	}
	return config
}

func (lc *LocalConfig) GetDatabaseUsername() string {
	return *lc.DatabaseUsername
}

func (lc *LocalConfig) GetDatabasePassword() string {
	return *lc.DatabasePassword
}

func (lc *LocalConfig) GetDatabaseHost() string {
	return *lc.DatabaseHost
}

func (lc *LocalConfig) GetDatabaseTable() string {
	return *lc.DatabaseTable
}

func (lc *LocalConfig) GetJwtSecret() []byte {
	return []byte(*lc.JwtSecret)
}
