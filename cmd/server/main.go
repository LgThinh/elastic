package main

import (
	"es-tranform/conf"
	"es-tranform/pkg/handler"
	"es-tranform/pkg/repo"
	"es-tranform/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//db, err := service.ConnectDB(conf.AppConfig{
	//	DBHost: conf.LoadEnv().DBHost,
	//	DBPort: conf.LoadEnv().DBPort,
	//	DBUser: conf.LoadEnv().DBUser,
	//	DBPass: conf.LoadEnv().DBPass,
	//	DBName: conf.LoadEnv().DBName,
	//})
	//// error handling
	//if err != nil {
	//	fmt.Println("Đã có lỗi xảy ra: ", err)
	//	return
	//}

	es := service.StartES(conf.GetConfig().ESAddress)

	EsRepo := repo.NewRepo(es)

	EsHandler := handler.NewEsHandlers(EsRepo)

	router := gin.Default()
	router.POST("/v1/input", EsHandler.CreateQuery)
	router.POST("/v2/input", EsHandler.CreateQueryFLex)

	router.Run(":8080")
}
