package module

import (
	"encoding/json"
	"fmt"
	"main/config/topic"
	"main/model"
	"main/redis"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

	redigoRedis "github.com/gomodule/redigo/redis"
)

// @Summary 获取所有节点基本信息
func GetNodes(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var nodes []model.Sysnodeinfo
	db.Model(model.Sysnodeinfo{}).Find(&nodes)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": nodes,
	})
}

//获取所有节点详细信息
func GetNodesInfo(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var nodes []model.Sysnodeinfo
	db.Model(model.Sysnodeinfo{}).Find(&nodes)
}

// 获取节点详细信息
func GetNodeInfo(c *gin.Context) {
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	key := fmt.Sprintf("SYSNODEINFO:%v", c.Query("id"))
	fmt.Println(key)
	data := make(map[string]string)
	r, err := redigoRedis.Strings(redisConn.Do("HMGET", key, "id", "hostname", "description", "ip", "status",
	 "feprunstatus", "fepstatuschangetime", "starttime", "hearttime", "diskavailable", 
	 "cpuusage", "memusage", "heartinterval", "deadperiod"))
	if err != nil {
		fmt.Printf("redis hmget failed: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": nil,
		})
		return
	}
	fmt.Println(r)
	data["id"] = r[0]
	data["hostname"] = r[1]
	data["description"] = r[2]
	data["ip"] = r[3]
	data["status"] = r[4]
	data["feprunstatus"] = r[5]
	data["fepstatuschangetime"] = r[6]
	data["starttime"] = r[7]
	data["hearttime"] = r[8]
	data["diskavailable"] = r[9]
	data["cpuusage"] = r[10]
	data["memusage"] = r[11]
	data["heartinterval"] = r[12]
	data["deadperiod"] = r[13]
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

// 控制节点
func ControlNode(c *gin.Context) {
	//根据参数生成命令
	cmd := model.CommandSysMgr{}
	cmd.RequestId = uuid.NewV1().String()
	cmd.Request = topic.Request_command_sysmgr
	cmd.Hostname = c.Query("hostname")
	cmd.Process = "sysmgr"
	cmd.Operation = c.Query("action")
	cmd.Response = "yes"
	data, errJson := json.Marshal(cmd)
	//命令失败 给前端返回错误信息
	if errJson != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": "json error",
		})
		return
	}
	//获取redis连接
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	//协程定制参数以便超时退出
	done := make(chan bool)
	defer close(done)
	go func() {
		redisConn := redis.RedisPool.Get()
		defer redisConn.Close()
		psc := redigoRedis.PubSubConn{Conn: redisConn}
		psc.Subscribe(cmd.RequestId)
		for {
			switch v := psc.ReceiveWithTimeout(2 * time.Second).(type) {
			case redigoRedis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
				cmdres := model.CommandResponse{}
				json.Unmarshal(v.Data, &cmdres)
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  cmdres.Message,
				})
				psc.Unsubscribe(cmd.RequestId)
				done <- true
				return
			case redigoRedis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Printf("error: %v\n", v)
				psc.Unsubscribe(cmd.RequestId)
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "未收到UBox系统回应",
				})
				done <- true
				return
			}
		}
	}()
	_, err := redisConn.Do("PUBLISH", cmd.Command, data)
	if err != nil {
		fmt.Printf("redis set failed: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err,
		})
		return
	}
	<-done
}

func GetProcesses(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var processes []model.Processinfo
	db.Model(model.Processinfo{}).Where("hostid = ?", c.Query("nodeId")).Find(&processes)
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	for i := range processes {
		processe := &processes[i]
		key := fmt.Sprintf("PROCESSINFO:%v", processe.Id)
		r, err := redigoRedis.Strings(redisConn.Do("HMGET", key, "status", "starttime", "hearttime", "cpuusage", "memusage"))
		if err != nil {
			fmt.Printf("redis hmget failed: %v", err)
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusOK,
				"data": nil,
			})
			return
		}
		fmt.Println(r)
		status, _ := strconv.Atoi(r[0])
		processe.Status = int8(status)
		processe.Starttime = r[1]
		processe.Hearttime = r[2]
		cpuusage, _ := strconv.ParseFloat(r[3], 32)
		processe.Cpuusage = float32(cpuusage)
		memusage, _ := strconv.ParseFloat(r[4], 32)
		processe.Memusage = float32(memusage)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": processes,
	})
}

// 控制节点
func ControlProcess(c *gin.Context) {
	//根据参数生成命令
	cmd := model.CommandSysMgr{}
	cmd.RequestId = uuid.NewV1().String()
	cmd.Request = topic.Request_command_sysmgr
	cmd.Hostname = c.Query("hostname")
	cmd.Process = c.Query("processName")
	cmd.Operation = c.Query("action")
	cmd.Response = "yes"
	data, errJson := json.Marshal(cmd)
	//命令失败 给前端返回错误信息
	if errJson != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 500,
			"data": "json error",
		})
		return
	}
	//获取redis连接
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	//协程定制参数以便超时退出
	done := make(chan bool)
	defer close(done)
	go func() {
		redisConn := redis.RedisPool.Get()
		defer redisConn.Close()
		psc := redigoRedis.PubSubConn{Conn: redisConn}
		psc.Subscribe(cmd.RequestId)
		for {
			switch v := psc.ReceiveWithTimeout(2 * time.Second).(type) {
			case redigoRedis.Message:
				fmt.Printf("%s: message: %s\n", v.Channel, v.Data)
				cmdres := model.CommandResponse{}
				json.Unmarshal(v.Data, &cmdres)
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  cmdres.Message,
				})
				psc.Unsubscribe(cmd.RequestId)
				done <- true
				return
			case redigoRedis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				fmt.Printf("error: %v\n", v)
				psc.Unsubscribe(cmd.RequestId)
				c.JSON(http.StatusOK, gin.H{
					"code": http.StatusOK,
					"msg":  "未收到UBox系统回应",
				})
				done <- true
				return
			}
		}
	}()
	_, err := redisConn.Do("PUBLISH", cmd.Command, data)
	if err != nil {
		fmt.Printf("redis set failed: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  err,
		})
		return
	}
	<-done
}

func Test(c *gin.Context) {

	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	_, err := redisConn.Do("PUBLISH", "Request_command_sysmgr", c.ClientIP())
	if err != nil {
		fmt.Printf("redis set failed: %v", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}



func WatchPoint(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 防止跨站点的请求伪造
		},
	}
	// 升级为websocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		ws.Close()
	}(ws)
}

// ws测试
func WsHandler(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 防止跨站点的请求伪造
		},
	}
	fmt.Println("ws")
	// 升级为websocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer func(ws *websocket.Conn) {
		ws.Close()
		fmt.Println("ws close")
	}(ws)
	MsgHandler(c, ws)
}

func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	// 读取ws中的数据,发过来的是点号字符串：1，2，3，4，5
	_, message, err := ws.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return
	}
	//获取redis连接
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	cmdFep := model.CommandFep{}
	cmdFep.RequestId = uuid.NewV1().String()
	cmdFep.Request = topic.Request_command_fep
	cmdFep.Operation = "watchs"
	cmdFep.Response = "yes"
	cmdFep.Param = "{\"ids\":\"" + string(message) + "\"}"
	dataFep, _ := json.Marshal(cmdFep)
	//协程定制参数以便超时退出
	done := make(chan bool)
	defer close(done)
	go func() {
		redisConn := redis.RedisPool.Get()
		defer redisConn.Close()
		psc := redigoRedis.PubSubConn{Conn: redisConn}
		err := psc.Subscribe(cmdFep.RequestId)
		if err != nil {
			respone := model.CommandResponse{}
			respone.RequestId = cmdFep.RequestId
			respone.Command = cmdFep.Command
			respone.Status = "fail"
			respone.Message = "未收到Fep系统回应"
			data, _ := json.Marshal(respone)
			ws.WriteMessage(websocket.TextMessage, data)
			done <- false
			return
		}
		for {
			switch v := psc.ReceiveWithTimeout(2 * time.Second).(type) {
			case redigoRedis.Message:
				ws.WriteMessage(websocket.TextMessage, v.Data)
				psc.Unsubscribe(cmdFep.RequestId)
				done <- true
				return
			case redigoRedis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				psc.Unsubscribe(cmdFep.RequestId)
				respone := model.CommandResponse{}
				respone.RequestId = cmdFep.RequestId
				respone.Command = cmdFep.Command
				respone.Status = "fail"
				respone.Message = "未收到Fep系统回应"
				data, _ := json.Marshal(respone)
				ws.WriteMessage(websocket.TextMessage, data)
				done <- false
				return
			}
		}
	}()
	_, err = redisConn.Do("PUBLISH", cmdFep.Command, dataFep)
	if err != nil {
		fmt.Printf("redis set failed: %v", err)
		return
	}
	<-done
	cmdDp := model.CommandDp{}
	cmdDp.RequestId = uuid.NewV1().String()
	cmdDp.Request = topic.Request_command_dp
	cmdDp.Operation = "watchs"
	cmdDp.Response = "yes"
	cmdDp.Param = "{\"ids\":\"" + string(message) + "\"}"
	dataDp, _ := json.Marshal(cmdDp)
	go func() {
		redisConn := redis.RedisPool.Get()
		defer redisConn.Close()
		psc := redigoRedis.PubSubConn{Conn: redisConn}
		err := psc.Subscribe(cmdDp.RequestId)
		if err != nil {
			respone := model.CommandResponse{}
			respone.RequestId = cmdDp.RequestId
			respone.Command = cmdDp.Command
			respone.Status = "fail"
			respone.Message = "未收到Dp系统回应"
			data, _ := json.Marshal(respone)
			ws.WriteMessage(websocket.TextMessage, data)
			done <- false
			return
		}
		for {
			switch v := psc.ReceiveWithTimeout(2 * time.Second).(type) {
			case redigoRedis.Message:
				ws.WriteMessage(websocket.TextMessage, v.Data)
				psc.Unsubscribe(cmdDp.RequestId)
				done <- true
				return
			case redigoRedis.Subscription:
				fmt.Printf("%s: %s %d\n", v.Channel, v.Kind, v.Count)
			case error:
				psc.Unsubscribe(cmdDp.RequestId)
				respone := model.CommandResponse{}
				respone.RequestId = cmdDp.RequestId
				respone.Command = cmdDp.Command
				respone.Status = "fail"
				respone.Message = "未收到Dp系统回应"
				data, _ := json.Marshal(respone)
				ws.WriteMessage(websocket.TextMessage, data)
				done <- false
				return
			}
		}
	}()
	_, err = redisConn.Do("PUBLISH", cmdDp.Command, dataDp)
	if err != nil {
		fmt.Printf("redis set failed: %v", err)
		return
	}
	<-done
	psc := redigoRedis.PubSubConn{Conn: redisConn}
	psc.Subscribe(topic.Response_watch_point)
	defer psc.Unsubscribe()
	for {
		switch r := psc.Receive().(type) {
		case error:
			fmt.Println("redis sub receive error:", r)
			return
		case redigoRedis.Message:
			err = ws.WriteMessage(websocket.TextMessage, r.Data)
			if err != nil {
				fmt.Println("ws write error:", err)
				return
			}
		}
	}
}
