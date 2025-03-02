package config

import (
	"fmt"           // 格式化
	"os"            // 操作系统
	"path"          // 路径
	"path/filepath" // 文件路径
	"runtime"       // 运行时

	"gopkg.in/yaml.v3" // yaml解析
)

var configFile []byte // 配置文件内容
var configFilePath string
// 配置文件数据结构
type Config struct {
	// 日志配置
	Log struct {
		EnableConsole     bool   `yaml:"enableConsole"`     // 是否启用控制台
		ConsoleJSONFormat bool   `yaml:"consoleJSONFormat"` // 控制台是否输出json格式
		ConsoleLevel      string `yaml:"consoleLevel"`      // 控制台日志级别
		EnableFile        bool   `yaml:"enableFile"`        // 是否启用文件
		FileJSONFormat    bool   `yaml:"fileJSONFormat"`    // 文件是否输出json格式
		FileLevel         string `yaml:"fileLevel"`         // 文件日志级别
		FileLocation      string `yaml:"fileLocation"`      // 文件日志位置
		MaxAge            int    `yaml:"maxAge"`            // 日志文件最大保存时间（天）
		MaxSize           int    `yaml:"maxSize"`           // 日志文件最大大小（MB）
		Compress          bool   `yaml:"compress"`          // 是否压缩日志文件
	}

	// 服务配置
	Webapi struct {
		Uri string `yaml:"uri"` // 服务地址
	}

	// 数据库配置
	MySqlnd struct {
		Username string `yaml:"username"` // 用户名
		Password string `yaml:"password"` // 密码
		Host     string `yaml:"host"`     // 主机
		Port     string `yaml:"port"`     // 端口
		Database string `yaml:"database"` // 数据库
		History  string `yaml:"history"`  // 历史数据库
	}

	// redis配置
	Redis struct {
		Host     string `yaml:"host"`     // 主机
		Port     int    `yaml:"port"`     // 端口
		Password string `yaml:"password"` // 密码
	}

	// dolphindb配置
	Dolphindb struct {
		Host     string `yaml:"host"`     // 主机
		Port     int    `yaml:"port"`     // 端口
		Username string `yaml:"username"` // 用户名
		Password string `yaml:"password"` // 密码
		Db       string `yaml:"db"`       // 数据库分区
	}

	//system 配置
	System struct {
		Enable_prevent_losses     bool   `yaml:"enable_prevent_losses" json:"enable_prevent_losses"`         // 是否开启防亏损选项
		Enable_limit_same_channel bool   `yaml:"enable_limit_same_channel" json:"enable_limit_same_channel"` // 是否开启限制同一通道同时操作
		Enable_allow_same_number  bool   `yaml:"enable_allow_same_number" json:"enable_allow_same_number"`   // 是否开启允许同号码同时充值
		To_up_info_type           int    `yaml:"to_up_info_type" json:"to_up_info_type"`                     // 提交给上游信息 0:按号段识别 1:按产品信息配置的产品运营商
		Supplier_order_no_up      int    `yaml:"supplier_order_no_up" json:"supplier_order_no_up"`           // 供货单状态与上游不一致时处理策略，0:自动下架供货商+短信通知 1:短信通知
		Order_save_days           int    `yaml:"order_save_days" json:"order_save_days"`                     // 订单保存天数，最低30天，不设置则永久保留
		System_maintenance_time   string `yaml:"system_maintenance_time" json:"system_maintenance_time"`     // 系统维护时间段，空表示不启动维护功能
		Notify_number             string `yaml:"notify_number" json:"notify_number"`                         // 系统通知号码，多个用逗号隔开
		Card_onsume_order         int    `yaml:"card_onsume_order" json:"card_onsume_order"`                 // 卡密消费顺序 10:先库存再通道 1:先通道再库存
		Query_exclude_failure     bool   `yaml:"query_exclude_failure" json:"query_exclude_failure"`         // 查询排除失败订单
		Card_arvchine_days        int    `yaml:"card_arvchine_days" json:"card_arvchine_days"`               // 已售卡券归档天数，最低15天，不设置不归档
		Card_check_repeat         bool   `yaml:"card_check_repeat" json:"card_check_repeat"`                 // 卡券入库是否检验重复
	}
}

// 定义Config全局对象实例
var config *Config

// 保存配置文件
func SaveConfig() error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFilePath, data, 0644)
}

// 初始化配置文件
func init() {
	var err error
	configFilePath = filepath.Join(getCurrentAbPathByCaller(), "config.yaml")
	fmt.Println(configFilePath)
	configFile, err = os.ReadFile(configFilePath)
	if err != nil {
		fmt.Printf("Read config yaml file err %v", err)
	}
	config, err = initConfig()
	if err != nil {
		fmt.Printf("读取配置信息失败：%v", err)
	}
	config.System.Notify_number = "123456"
	// 保存配置到文件中
	err = SaveConfig()
	if err != nil {
		fmt.Printf("保存配置文件失败：%v", err)
	}
	print(config.System.Notify_number)
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 初始化配置文件
func initConfig() (config *Config, err error) {

	// 解析配置文件
	err = yaml.Unmarshal(configFile, &config)
	return config, err
}

// 获取配置文件
func GetConfig() *Config {
	return config
}
