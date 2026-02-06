package userwebserver

func (ws *WebServer) ConfigureRoutes() {
	ws.group.POST("/", ws.CreateUser)
	ws.group.POST("/login", ws.Login)
}
