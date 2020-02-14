package common

import "encoding/json"

// 定时任务
type Job struct {
	Name     string `json:"name"`     //任务名
	Command  string `json:"command"`  //shell 命令
	CronExpr string `json:"cronExpr"` // cron
}

// http 接口应答
type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

// 应答

func BuildResponse(errno int, msg string, data interface{}) (resp []byte, err error) {
	var (
		response Response
	)

	response.Data = data
	response.Errno = errno
	response.Msg = msg

	resp, err = json.Marshal(response)

	return
}
