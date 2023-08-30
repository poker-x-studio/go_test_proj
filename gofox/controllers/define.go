package controllers

// 返回统一的格式
type RespJson struct {
	Code    int
	Message string
	Data    interface{}
}
