package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver    string `mapstructure:"DB_DRIVER"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPass      string `mapstructure:"DB_PASS"`
	DBName      string `mapstructure:"DB_NAME"`
	WSPort      string `mapstructure:"WS_PORT"`
	JWTSecret   string `mapstructure:"JWT_SECRET"`
	JWTExpireIn int    `mapstructure:"JWT_EXPIRE_IN"`
	TokenAuth   *jwtauth.JWTAuth
}

// Executada antes da função main
func init() {

}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
