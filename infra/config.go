package infra

func NewConfig() (*Config, error) {
	config := &Config{}
	return config, nil
}

type Config struct {
	Postgres  PostgresConfig `json:"postgres"`
	Migration Migration      `json:"migration"`
}

type PostgresConfig struct {
	Address     string `json:"addr"`
	Database    string `json:"db"`
	Username    string `json:"user"`
	Password    string `json:"password"`
	ReadTimeout int    `json:"readtimeout"`
}

type Migration struct {
	Version uint `json:"version"`
}
