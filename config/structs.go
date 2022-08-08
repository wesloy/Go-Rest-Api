package config

// função privada
type config struct {
	API APIConfig
	DB  DBConfig
}

// estruturas públicas
type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}
