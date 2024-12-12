package module

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	redigoRedis "github.com/gomodule/redigo/redis"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"main/config/topic"
	"main/model"
	"main/redis"
	"net/http"
	"strconv"
	"time"
)

// // @brief：耗时统计函数
// func timeCost() func() {
// 	start := time.Now()
// 	return func() {
// 		tc := time.Since(start)
// 		fmt.Printf("time cost = %v\n", tc)
// 	}
// }

// 获取设备信息
func GetDevices(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var devices []model.Rtu
	db.Model(model.Rtu{}).Where("type = ?", c.Query("type")).Find(&devices)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": devices,
	})
}

// 获取采集设备信息
func GetCollectDevices(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var devices []model.Rtu
	db.Model(model.Rtu{}).Where("type = 0").Find(&devices)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": devices,
	})
}

// 获取转发设备信息
func GetForwardDevices(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var devices []model.Rtu
	db.Model(model.Rtu{}).Where("type = 3").Find(&devices)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": devices,
	})
}

// 获取采集设备和转发设备数量
func GetDeviceCount(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var collectCount int64
	db.Model(model.Rtu{}).Where("type = 0").Count(&collectCount)
	var forwardCount int64
	db.Model(model.Rtu{}).Where("type = 3").Count(&forwardCount)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"collectCount": collectCount,
		"forwardCount": forwardCount,
	})
}

// 获取设备通道信息
func GetDeviceChannels(c *gin.Context) {
	// 获取数据库连接
	db := model.Db
	var channels []model.Channel
	db.Model(model.Channel{}).Where("rtuid = ? and forwardflag = 0", c.Query("deviceId")).Find(&channels)
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	for i := range channels {
		channel := &channels[i]
		key := fmt.Sprintf("CHANNEL:%v", channel.Id)
		r, err := redigoRedis.Strings(redisConn.Do("HMGET", key, "status", "starttime", "hbtime", "issrc"))
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
		channel.Status = int8(status)
		channel.Starttime = r[1]
		channel.Hbtime = r[2]
		issrc, _ := strconv.Atoi(r[3])
		channel.Issrc = int8(issrc)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": channels,
	})
}

// 获取模拟量点信息
func GetAnalogPoints(c *gin.Context) {
	// 获取数据库连接
	id := c.Query("deviceId")
	// 获取数据库连接
	db := model.Db
	var rtuAnalog []model.Analog
	db.Model(model.Analog{}).Where("rtuid = ?", id).Find(&rtuAnalog)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rtuAnalog,
	})
}

// 获取数字量点信息
func GetDigitalPoints(c *gin.Context) {
	// 获取数据库连接
	id := c.Query("id")
	// 获取数据库连接
	db := model.Db
	var rtuStatus []model.Status
	db.Model(model.Status{}).Where("rtuid = ?", id).Find(&rtuStatus)
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": rtuStatus,
	})
}

// 获取协议参数
func GetProtocolParam(c *gin.Context) {
	//根据参数生成命令
	cmd := model.CommandChannel{}
	cmd.RequestId = uuid.NewV1().String()
	cmd.Request = topic.Request_command_channel
	channelid, _ := strconv.Atoi(c.Query("channelid"))
	cmd.ChannelId = int32(channelid)
	cmd.Operation = "protol_param"
	cmd.AddParam("param", c.Query("param"))
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
	fmt.Println(string(data))
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
	_, err := redisConn.Do("PUBLISH", cmd.Request, string(data))
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

// websocket获取测点实时数据
func WsGetRealTimeData(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 防止跨站点的请求伪造
		},
	}
	// 升级为websocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer ws.Close()

	// 读取ws中的数据,发过来的是redis中的Key json格式
	_, message, err := ws.ReadMessage()
	if err != nil {
		log.Println(err)
		return
	}
	type keys struct {
		Keys []string `json:"keys"`
	}
	var key keys
	err = json.Unmarshal(message, &key)
	if err != nil {
		log.Println(err)
		return
	}
	// 获取redis连接
	redisConn := redis.RedisPool.Get()
	defer redisConn.Close()
	dataMap := make(map[string]map[string]string)
	for {

		for _, k := range key.Keys {
			redisConn.Send("hmget", k, "value", "qualitycode", "timestamp")
		}
		redisConn.Flush()
		for _, k := range key.Keys {
			r, err := redisConn.Receive()
			if err != nil {
				log.Println(err)
				continue
			}
			if r == nil {
				continue
			}
			if r.([]interface{})[0] == nil {
				continue
			}
			if r.([]interface{})[1] == nil {
				continue
			}
			if r.([]interface{})[2] == nil {
				continue
			}
			value := string(r.([]interface{})[0].([]byte))
			qualitycode := string(r.([]interface{})[1].([]byte))
			timestamp := string(r.([]interface{})[2].([]byte))
			data := map[string]string{
				"value":       value,
				"qualitycode": qualitycode,
				"timestamp":   timestamp,
			}
			dataMap[k] = data
		}
		data, err := json.Marshal(dataMap)
		if err != nil {
			log.Println(err)
			continue
		}
		err = ws.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
