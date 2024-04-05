package config

import (
	"strings"
	"sync"

	"github.com/spf13/viper"
)

// creamos los tipos que usaremos en la configuracion
type (
	Config struct {
		// la configuracion general sera del servidor y la db
		Server *Server
		Db     *Db
	}

	// del servidor solo necesitamos el puerto
	Server struct {
		Port int
	}

	// de la db necesitamos varios datos
	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		TimeZone string
	}
)

var (
	once           sync.Once // usamos sync.Once para sincronizar solo una vez la configuracion
	configInstance *Config   // la instancia configInstance contendra la configuracion
)

// el metodo GetConfig retornara una unica instancia de la configuracion que actuara como un Singleton
// no podremos llamar a GetConfig dos veces.
func GetConfig() *Config {
	// usamos viper para setear la config y usarla through nuestra app
	once.Do(func() {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err := viper.ReadInConfig()

		if err != nil {
			panic(err)
		}

		if err := viper.Unmarshal(&configInstance); err != nil {
			panic(err)
		}
	})

	// devolvemos nuestra instancia de configuracion
	return configInstance

}
