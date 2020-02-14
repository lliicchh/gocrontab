package master

import (
	"encoding/json"
	"fmt"
	"gocrontab/common"
	"net"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	//单例对象
	G_apiServer *ApiServer
)

// 保存任务
// POST job={"name":"job1", "command":"echo xxx", "cronExpr":"* * * * *"}
func handleJobSave(resp http.ResponseWriter, req *http.Request) {
	var (
		err     error
		bytes   []byte
		postJob string
		job     common.Job
		oldJob  *common.Job
	)

	// 1. 解析表单
	fmt.Println("begin parse")
	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	// 2. 获取job字段
	postJob = req.PostForm.Get("job")

	fmt.Println(postJob)
	// 3. 反序列化
	if err = json.Unmarshal([]byte(postJob), &job); err != nil {
		goto ERR
	}


	// 4.保存到etcd
	if oldJob, err = G_jobMgr.SaveJob(&job); err != nil {
		return
	}

	// 5. 返回正常应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}

ERR:
	if bytes, err = common.BuildResponse(-1, err.Error(), nil); err == nil {
		resp.Write(bytes)
	}

}
func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}
	httpServer = &http.Server{
		ReadTimeout:  time.Duration(G_config.ApiReadTimeout) * time.Millisecond,
		WriteTimeout: time.Duration(G_config.ApiWriteTimeout) * time.Millisecond,
		Handler:      mux,
	}
	G_apiServer = &ApiServer{httpServer: httpServer}

	go httpServer.Serve(listener)
	return
}
