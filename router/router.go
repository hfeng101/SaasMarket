package router

import "github.com/gin-gonic/gin"
//import "github.com/SaasMarket/pkg/controller/payment"
import "../pkg/controller/payment"
import "../pkg/controller/user"

type Router struct{
	Threads int
	Router	*gin.Engine

}

func SetupRouter() {
	server := ":8080"

	r := gin.Default()

	r.GET("/test", )

	userGroup := r.Group("/user", )
	userController := user

	paymentGroup := r.Group("/payment", )
	paymentController := payment.Payment{1}
	{
		paymentGroup.POST("/payment", paymentController.GetPaymentInfo)
	}

	//server := "127.0.0.1:8080"
	r.Run(server)
}
