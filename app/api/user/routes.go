package userwebserver

func (ws *WebServer) ConfigureRoutes() {
	ws.group.POST("/", ws.CreateUser)
}
