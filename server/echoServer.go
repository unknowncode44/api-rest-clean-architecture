package server

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/unknowncode44/api-rest-clean-architecture/config"
	"github.com/unknowncode44/api-rest-clean-architecture/database"
)

// la estructura echoServer la utilizaremos para el servidor, la db y la configuracion
// que importaremos desde la libreria echo, de nuestro archivo de database y nuestro archivo de configuracion
type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

// la funcion NewEchoServer recibe los parametros de configuracion y database
// y sera la encargada de levantar nuestro servidor en el puerto especificado en el archivo de configuracion
// y utilizar nuestra configuracion de base de datos
func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()

	// seteamos el built-in logger de echo para que nos lance mensages de debug en la consola
	echoApp.Logger.SetLevel(log.DEBUG)

	// devolvemos un objeto con nuestra estructura echoServer
	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

// creamos la funcion para iniciar el servidor
func (s *echoServer) Start() {
	// usamos middleware para:
	s.app.Use(middleware.Recover()) // <-- evita el panic cuando haya crashes en el server y lo mantiene on
	s.app.Use(middleware.Logger())  // <-- el logger nos manda mensages a la consola con la actividad del server

	// agregamos un check para comprobar correcto funcionamiento
	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	// definimos la direccion y puerto del servidor
	serveUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serveUrl))

}
