package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// conf é uma struct que armazena as configurações do nosso projeto.
type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	jwtSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

// cfg é uma variável global que armazenará as configurações carregadas.
var cfg *conf

// LoadConfig carrega as configurações a partir de um arquivo .env e retorna uma instância de conf.
func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config") // Define o nome do arquivo de configuração.
	viper.SetConfigType("env")        // Define o tipo do arquivo de configuração.
	viper.AddConfigPath(path)         // Adiciona o caminho onde o arquivo de configuração está localizado.
	viper.SetConfigFile(".env")       // Define o arquivo de configuração específico a ser carregado.
	viper.AutomaticEnv()              // Permite que as variáveis de ambiente sobrescrevam as configurações.

	err := viper.ReadInConfig() // Lê o arquivo de configuração.
	if err != nil {
		panic(err) // Encerra o programa se houver um erro ao ler o arquivo de configuração.
	}

	err = viper.Unmarshal(&cfg) // Converte as configurações para a struct conf.
	if err != nil {
		panic(err) // Encerra o programa se houver um erro ao converter as configurações.
	}

	// Cria uma nova instância de JWTAuth usando a chave secreta do JWT.
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.jwtSecret), nil)

	return cfg, err
}

/*
formatos de arquivos que podemos trabalhar:
yaml
toml
env
json
*/
