package util

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Config struct {
	PgUrl             string `mapstructure:"POSTGRES_URL"`
	PgPool            int    `mapstructure:"POSTGRES_POOL"`
	DbDriver          string `mapstructure:"DB_DRIVER"`
	ServersAddress    string `mapstructure:"ADDRESS"`
	AuthorizationPort int    `mapstructure:"AUTHORIZATION_PORT"`
}

func AuthPort() int {
	config, err := LoadConfig(".")
	if err != nil {
		logrus.Fatal("cannot load config: ", err)
	}
	return config.AuthorizationPort
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func RandomInt(min, max int64) int64 {
	return min + int64(rand.Int63n(max-min+1))
}

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case EUR, USD, CAD:
		return true
	}
	return false
}
