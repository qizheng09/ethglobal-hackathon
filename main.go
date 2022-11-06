package main

import (
	"fmt"
	"hackathon/controller"
	"log"
	"os"

	"hackathon/config"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	// init config
	if err := config.InitConfig(); err != nil {
		log.Println(err)
		return
	}

	// access log
	ac := accesslog.File("./log/access.log")
	ac.SetFormatter(&accesslog.JSON{
		HumanTime: true,
		Indent:    "    ",
	})
	defer ac.Close()

	file, err := os.OpenFile("./log/iris.log", os.O_CREATE|os.O_APPEND, 0600)
	defer file.Close()

	// bootstrap server
	server := iris.New()
	server.Validator = validator.New()
	server.Logger().SetOutput(file)
	server.Use(recover.New())
	server.UseRouter(ac.Handler)

	userAPI := server.Party("/hackathon")
	{
		userAPI.Get("/hello", controller.Hello)
		userAPI.Post("/trade_check", controller.TradeCheck)
	}

	err = server.Listen(fmt.Sprintf(":%d", config.EnvConfig.Server.Port))
	if err != nil {
		log.Println(err)
	}
}
