package restapi

import (
	controller "capsulego/controller/ordercontroller"
	"capsulego/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	Result     map[string]interface{} `json:"result"`
	VarResult  string                 `json:"var_result"`
	VarMessage string                 `json:"var_message"`
}

var order = new(models.Order)

func SaveOrder(c *gin.Context) {
	var response response
	var order models.Order
	c.Bind(&order)
	err := controller.SaveOrder(order)
	if err != nil {
		response.VarResult = "0"
		response.VarMessage = "error"
		c.JSON(http.StatusOK, response)
		return
	}
	response.VarResult = "1"
	//callback is x
	// Will output  :   x({\"foo\":\"bar\"})
	c.JSON(http.StatusOK, response)
}
