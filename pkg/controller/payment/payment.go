package payment

import "github.com/gin-gonic/gin"

type Payment struct{
	payAccount int
}

func newPayment() *Payment{

	return &Payment{payAccount:1}
}

func getPayment(){
	return
}

func (p *Payment)GetPaymentInfo(c *gin.Context){
	return
}