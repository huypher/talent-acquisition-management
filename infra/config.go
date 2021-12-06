package infra

func NewConfig() (*Config, error) {
	config := &Config{}
	return config, nil
}

type Config struct {
	Postgres PostgresConfig `json:"postgres"`
}

type PostgresConfig struct {
	Address     string `json:"addr"`
	Database    string `json:"db"`
	Username    string `json:"talent"`
	Password    string `json:"password"`
	ReadTimeout int    `json:"readtimeout"`
}
