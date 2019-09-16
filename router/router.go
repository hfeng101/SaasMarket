package router

import "github.com/gin-gonic/gin"
import "../pkg/controller/payment"


type Router struct{
	Threads int
	Router	*gin.Engine

}


//func (r *Router)SetupRouter(){
//
//}

func setupRouter() *gin.Engine{
	r := gin.Default()

	r.GET("/test", )

	paymentGroup := r.Group("/", )

	paymentController := payment

	{
		paymentGroup.POST("/payment",)

	}

	return r
}
