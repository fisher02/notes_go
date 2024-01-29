package models

type R struct {
	Success bool        `json:"success"` // 表示请求是否成功
	Message string      `json:"message"` // 返回的消息
	Data    interface{} `json:"data"`    // 返回的数据
}
