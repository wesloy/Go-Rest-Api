package config

import "github.com/spf13/viper"

// criando ponteiro para função privada
var cfg *config

// init é a primeira função a ser executada pelo módulo
func init() {

	// criando algumas configurações básicas caso o arquivo de configuração não seja localizado
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")

}

// carregando as configurações usando viper e lendo de um arquivo de configurações
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	// validando se houve erro, diferente de arquivo não encontrado e repassando
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// criando um ponteiro
	cfg = new(config)

	// Carregando as infos recuperadas do arquivo de configuração com o Viper
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.database"),
	}

	// qdo não há erro retorna nulo, já que acima estamos adicionando informações em ponteiros
	return nil
}

// externalizando as configurações para consumo, via ponteiro
func GetDB() DBConfig {
	return cfg.DB
}
func GetServerPort() string {
	return cfg.API.Port
}
