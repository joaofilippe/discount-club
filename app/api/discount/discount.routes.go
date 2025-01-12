package discountwebserver

func (ws *WebServer) BuildRoutes() {
	ws.server.POST("/discounts", ws.CreateDiscount)
}
