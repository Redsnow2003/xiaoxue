package model

// 用户表
type User struct {
	Id         uint64 `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 用户id
	UserName   string `gorm:"column:username;type:varchar(255);not null" json:"username"` // 用户名
	Password   string `gorm:"column:password;type:varchar(255);not null" json:"password"` // 用户密码
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`         // 用户姓名
	Role       int8   `gorm:"column:role;type:int;not null" json:"role"`                  // 用户角色
	Img        string `gorm:"column:img;type:varchar(255)" json:"img"`                    // 用户头像
	LoginTime  string `gorm:"column:login_time;type:datetime" json:"login_time"`          // 登录时间
	LogoutTime string `gorm:"column:logout_time;type:datetime" json:"logout_time"`        // 退出时间
}

// 表名
func (User) TableName() string {
	return "user"
}

// 节点信息表
type Sysnodeinfo struct {
	Id 	   			uint64 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 节点id
	Hostname  		string 	`gorm:"column:hostname;type:varchar(255);not null" json:"hostname"` // 节点名
	Description  	string 	`gorm:"column:description;type:varchar(255);not null" json:"description"` // 节点描述
	Ip			 	string 	`gorm:"column:ip;type:varchar(255);not null" json:"ip"` // 节点ip
	Heartinterval	int8	`gorm:"column:heartinterval;type:int;not null" json:"heartinterval"` // 心跳间隔
	Deadperiod		int8	`gorm:"column:deadperiod;type:int;not null" json:"deadperiod"` // 超时时间
}

// 表名
func (Sysnodeinfo) TableName() string {
	return "sysnodeinfo"
}

// 子进程信息表
type Processinfo struct {
	Id 	   			uint64 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 进程id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // 进程名
	Description  	string 	`gorm:"column:description;type:varchar(255);not null" json:"description"` // 进程描述
	Hostid		 	uint64 	`gorm:"column:hostid;type:int;not null" json:"hostid"` // 节点id
	Hostname 		string 	`gorm:"column:hostname;type:varchar(255);not null" json:"hostname"` // 节点名
	Type 			int8 	`gorm:"column:type;type:int;not null" json:"type"` // 进程类型
	Enable			int8	`gorm:"column:enable;type:int;not null" json:"enable"` // 是否启用
	Delaytime 		int8 	`gorm:"column:delaytime;type:int;not null" json:"delaytime"` // 延迟时间
	Show 			int8 	`gorm:"column:show;type:int;not null" json:"show"` // 是否显示
	Status 			int8 	`gorm:"column:status;type:int;not null" json:"status"` // 状态
	Starttime 		string 	`gorm:"column:starttime;type:datetime" json:"starttime"` // 启动时间
	Hearttime 		string 	`gorm:"column:hearttime;type:datetime" json:"hearttime"` // 心跳时间
	Cpuusage 		float32 `gorm:"column:cpuusage;type:float;not null" json:"cpuusage"` // cpu使用率
	Memusage 		float32 `gorm:"column:memusage;type:float;not null" json:"memusage"` // 内存使用率
	Heartinterval 	int8 	`gorm:"column:heartinterval;type:int;not null" json:"heartinterval"` // 心跳间隔
	Deadperiod 		int8 	`gorm:"column:deadperiod;type:int;not null" json:"deadperiod"` // 超时时间
}

// 表名
func (Processinfo) TableName() string {
	return "processinfo"
}

// 子进程表
type Process struct {
	Id 	   			uint64 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 进程id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // 进程名
	Description  	string 	`gorm:"column:description;type:varchar(255);not null" json:"description"` // 进程描述
}

// 表名
func (Process) TableName() string {
	return "process"
}

// 通道表
type Channel struct {
	Id 	   			uint32 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 通道id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // 通道名
	Aliasname 		string 	`gorm:"column:aliasname;type:varchar(255)" json:"aliasname"` // 通道别名
	Description  	string 	`gorm:"column:description;type:varchar(255)" json:"description"` // 通道描述
	Hostid		 	uint32 	`gorm:"column:hostid;type:bigint;not null" json:"hostid"` // 节点id
	Appprotocolid 	uint32 	`gorm:"column:appprotocolid;type:bigint;not null" json:"appprotocolid"` // 应用层协议
	Rtuid 			uint32 	`gorm:"column:rtuid;type:bigint;not null" json:"rtuid"` // rtuid
	Pollflag 		int8 	`gorm:"column:pollflag;type:int;not null" json:"pollflag"` // 轮询标志
	Forwardflag 	int8 	`gorm:"column:forwardflag;type:int;not null" json:"forwardflag"` // 转发标志
	Mediatype 		int8 	`gorm:"column:mediatype;type:int;not null" json:"mediatype"` // 媒体类型
	Address			string 	`gorm:"column:address;type:varchar(255)" json:"address"` // 地址
	Linkpara 		string 	`gorm:"column:linkpara;type:varchar(255)" json:"linkpara"` // 链接参数
	Availableflag 	int8 	`gorm:"column:availableflag;type:int;not null" json:"availableflag"` // 是否可用
	Errorratelimit  float32 `gorm:"column:errorratelimit;type:float;not null" json:"errorratelimit"` // 错误率限制
	Invalidtimeout  int32   `gorm:"column:invalidtimeout;type:bigint;not null" json:"invalidtimeout"` // 无效超时
	Heartinterval	int8 	`gorm:"column:heartinterval;type:int;not null" json:"heartinterval"` // 心跳间隔
	Deadperiod 		int8 	`gorm:"column:deadperiod;type:int;not null" json:"deadperiod"` // 超时时间
	Starttime 		string 	`gorm:"column:starttime;type:datetime" json:"starttime"` // 启动时间
	Status 			int8 	`gorm:"column:status;type:int;not null" json:"status"` // 状态
	Statechangetime string 	`gorm:"column:statechangetime;type:datetime" json:"statechangetime"` // 状态改变时间
	Hbtime 			string 	`gorm:"column:hbtime;type:datetime" json:"hbtime"` // 心跳时间
	Issrc 			int8 	`gorm:"column:issrc;type:int;not null" json:"issrc"` // 是否源
	Isopen 			int8 	`gorm:"column:isopen;type:int;not null" json:"isopen"` // 是否开启
	Samplecorrect 	int64 	`gorm:"column:samplecorrect;type:int;not null" json:"samplecorrect"` // 采样校正
	Sampleerror 	int64 	`gorm:"column:sampleerror;type:int;not null" json:"sampleerror"` // 采样错误
	Sampletotal 	int64 	`gorm:"column:sampletotal;type:int;not null" json:"sampletotal"` // 采样总数
	Sampleerrorrate float32 `gorm:"column:sampleerrorrate;type:float;not null" json:"sampleerrorrate"` // 采样错误率
	Totalcorrect 	int64 	`gorm:"column:totalcorrect;type:int;not null" json:"totalcorrect"` // 总校正
	Totalerror 		int64 	`gorm:"column:totalerror;type:int;not null" json:"totalerror"` // 总错误
	Totalframe 		int64 	`gorm:"column:totalframe;type:int;not null" json:"totalframe"` // 总帧数
	Totalerrorrate 	float32 `gorm:"column:totalerrorrate;type:float;not null" json:"totalerrorrate"` // 总错误率
	Totalvalidtime 	int64 	`gorm:"column:totalvalidtime;type:int;not null" json:"totalvalidtime"` // 总有效时间
	Totalinvalidtime int64 	`gorm:"column:totalinvalidtime;type:int;not null" json:"totalinvalidtime"` // 总无效时间
	Checktime 		string 	`gorm:"column:checktime;type:datetime" json:"checktime"` // 检查时间
	Ssl				int8 	`gorm:"column:ssl;type:int" json:"ssl"` // 是否ssl
	SslType 		int8 	`gorm:"column:ssl_type;type:int" json:"ssl_type"` // ssl类型
	ServerCertPath 	string 	`gorm:"column:server_cert_path;type:varchar(255)" json:"server_cert_path"` // 服务端证书路径
	ClientCertPath 	string 	`gorm:"column:client_cert_path;type:varchar(255)" json:"client_cert_path"` // 客户端证书路径
	ClientKeyPath 	string 	`gorm:"column:client_key_path;type:varchar(255)" json:"client_key_path"` // 客户端密钥路径
	ClientCertPass 	string 	`gorm:"column:client_cert_pass;type:varchar(255)" json:"client_cert_pass"` // 客户端证书密码
}

// 表名
func (Channel) TableName() string {
	return "channel"
}

//RTU表
type Rtu struct {
	Id 	   			uint32 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // rtu id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // rtu名
	Aliasname 		string 	`gorm:"column:aliasname;type:varchar(255)" json:"aliasname"` // rtu别名
	Description  	string 	`gorm:"column:description;type:varchar(255)" json:"description"` // rtu描述
	Type 			int8 	`gorm:"column:type;type:int;not null" json:"type"` // rtu类型
	Isavailable 	int8 	`gorm:"column:isavailable;type:int;not null" json:"isavailable"` // 是否可用
	Issrc 			int8 	`gorm:"column:issrc;type:int;not null" json:"issrc"` // 是否源
	Scanallinterval int32 	`gorm:"column:scanallinterval;type:int;not null" json:"scanallinterval"` // 扫描间隔
	Scanaccinterval int32 	`gorm:"column:scanaccinterval;type:int;not null" json:"scanaccinterval"` // 扫描间隔
	Ycmax 			int32 	`gorm:"column:ycmax;type:int;not null" json:"ycmax"` // 遥测最大值
	Yxmax 			int32 	`gorm:"column:yxmax;type:int;not null" json:"yxmax"` // 遥信最大值
	Ymmax 			int32 	`gorm:"column:ymmax;type:int;not null" json:"ymmax"` // 遥脉最大值
	Ykmax 			int32 	`gorm:"column:ykmax;type:int;not null" json:"ykmax"` // 遥控最大值
	Ytmax 			int32 	`gorm:"column:ytmax;type:int;not null" json:"ytmax"` // 遥调最大值
	Timesyncinterval int32 	`gorm:"column:timesyncinterval;type:int;not null" json:"timesyncinterval"` // 时间同步间隔
	Heartinterval 	int32 	`gorm:"column:heartinterval;type:int;not null" json:"heartinterval"` // 心跳间隔
	Deadperiod 		int32 	`gorm:"column:deadperiod;type:int;not null" json:"deadperiod"` // 超时时间
	Starttime 		string 	`gorm:"column:starttime;type:datetime" json:"starttime"` // 启动时间
	Status 			int8 	`gorm:"column:status;type:int;not null" json:"status"` // 状态
	Statechangetime string 	`gorm:"column:statechangetime;type:datetime" json:"statechangetime"` // 状态改变时间
	Hbtime 			string 	`gorm:"column:hbtime;type:datetime" json:"hbtime"` // 心跳时间
	Isopen 			int8 	`gorm:"column:isopen;type:int;not null" json:"isopen"` // 是否开启
	Totalvalidtime 	int32 	`gorm:"column:totalvalidtime;type:int;not null" json:"totalvalidtime"` // 总有效时间
	Totalinvalidtime int32 	`gorm:"column:totalinvalidtime;type:int;not null" json:"totalinvalidtime"` // 总无效时间
	Checktime 		string 	`gorm:"column:checktime;type:datetime" json:"checktime"` // 检查时间
}

// 表名
func (Rtu) TableName() string {
	return "rtu"
}

// 遥测表
type Analog struct {
	Id 	   			uint32 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 遥测id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // 遥测名
	Aliasname 		string 	`gorm:"column:aliasname;type:varchar(255)" json:"aliasname"` // 遥测别名
	Description  	string 	`gorm:"column:description;type:varchar(255)" json:"description"` // 遥测描述
	Rtuid 			uint32 	`gorm:"column:rtuid;type:bigint;not null" json:"rtuid"` // rtuid
	Pointnum 		uint32 	`gorm:"column:pointnum;type:bigint;not null" json:"pointnum"` // 点号
	Type 			int8 	`gorm:"column:type;type:int;not null" json:"type"` // 类型
	Param 			string  `gorm:"column:param1;type:varchar(255)" json:"param1"` // 参数4
	Scalefactor 	float32 `gorm:"column:scalefactor;type:float;not null" json:"scalefactor"` // 比例因子
}

// 表名
func (Analog) TableName() string {
	return "analog"
}

// 遥信表
type Status struct {
	Id 	   			uint32 	`gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`      // 遥信id
	Name 			string 	`gorm:"column:name;type:varchar(255);not null" json:"name"` // 遥信名
	Aliasname 		string 	`gorm:"column:aliasname;type:varchar(255)" json:"aliasname"` // 遥信别名
	Description  	string 	`gorm:"column:description;type:varchar(255)" json:"description"` // 遥信描述
	Rtuid 			uint32 	`gorm:"column:rtuid;type:bigint;not null" json:"rtuid"` // rtuid
	Pointnum 		uint32 	`gorm:"column:pointnum;type:bigint;not null" json:"pointnum"` // 点号
	Type 			int8 	`gorm:"column:type;type:int;not null" json:"type"` // 类型
	Param 			string  `gorm:"column:param1;type:varchar(255)" json:"param1"` // 参数4
}

// 表名
func (Status) TableName() string {
	return "status"
}