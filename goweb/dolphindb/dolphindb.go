package dolphindb
import(
    "context"
    "fmt"
    "github.com/dolphindb/api-go/api"
	"main/config"
)

func ConnectDb() (db api.DolphinDB, err error) {
	// 读取配置文件
	configBase, err := config.InitConfig()
	if err != nil {
		fmt.Printf("读取配置信息失败：%v", err)
	}
	host := fmt.Sprintf("%s:%d", configBase.Dolphindb.Host, configBase.Dolphindb.Port)
	db,err = api.NewSimpleDolphinDBClient(context.TODO(), host, configBase.Dolphindb.Username, configBase.Dolphindb.Password)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	return db, err
}
