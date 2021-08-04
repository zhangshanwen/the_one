package code

/*
成功code 100000 开头
:说明 前端以定义的code 去定义对应的msg，切换语言
*/
const (
	BaseSuccessCode = 100000              // 基数成功code
	Success         = BaseSuccessCode + 1 // 普通成功
	CreateSuccess   = BaseSuccessCode + 2 // 创建成功
)
