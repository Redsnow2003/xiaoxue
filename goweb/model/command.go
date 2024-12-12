package model

import (
	"encoding/json"
)

type Command struct {
	RequestId string `json:"requestId"`
	Request string `json:"request"`
	Response string `json:"response"`
	Param string `json:"param"`
}

type CommandSysMgr struct {
	Command
	Operation string `json:"operation"`
	Hostname string `json:"hostname"`
	Process string `json:"process"`
	Description string `json:"description"`
}

type CommandResponse struct {
	Command
	Status string `json:"status"`
	Message string `json:"message"`
}

type CommandChannel struct {
	Command
	ChannelId int32 `json:"channelid"`
	Operation string `json:"operation"`
}

type CommandProcess struct {
	Command
	Operation string `json:"operation"`
}

type CommandFep struct {
	CommandProcess
}

type CommandDp struct {
	CommandProcess
}

type CommandDolphindb struct {
	CommandProcess
}

func (c *Command) AddParam(key string, value interface{}) {
	mapParam := make(map[string]interface{})
	json.Unmarshal([]byte(c.Param), &mapParam)
	mapParam[key] = value
	data, _ := json.Marshal(mapParam)
	c.Param = string(data)
}

func (c *Command) GetParamString(key string) string {
	mapParam := make(map[string]interface{})
	json.Unmarshal([]byte(c.Param), &mapParam)
	return mapParam[key].(string)
}


func (c *Command) GetParamInt64(key string) int64 {
	mapParam := make(map[string]interface{})
	json.Unmarshal([]byte(c.Param), &mapParam)
	return mapParam[key].(int64)
}

func (c *Command) GetParamFloat64(key string) float64 {
	mapParam := make(map[string]interface{})
	json.Unmarshal([]byte(c.Param), &mapParam)
	return mapParam[key].(float64)
}

