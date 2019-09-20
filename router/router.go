package router

import "github.com/gin-gonic/gin"
//import "github.com/SaasMarket/pkg/controller/payment"
import "../pkg/controller/payment"

type Router struct{
	Threads int
	Router	*gin.Engine

}

type router struct{
	Threads int
	Router	*gin.Engine

}


//func (r *Router)SetupRouter(){
//
//}

func SetupRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/test", )

	paymentGroup := r.Group("/", )

	paymentController := payment.Payment{1}

	{
		paymentGroup.POST("/payment", paymentController.GetPaymentInfo)

	}

	return r
}
