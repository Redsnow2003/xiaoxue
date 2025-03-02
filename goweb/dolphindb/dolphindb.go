package dolphindb
import(
    "context"
    "fmt"
    "github.com/dolphindb/api-go/api"
	"main/config"
)

func ConnectDb() (db api.DolphinDB, err error) {
	// 读取配置文件
	configBase := config.GetConfig()
	host := fmt.Sprintf("%s:%d", configBase.Dolphindb.Host, configBase.Dolphindb.Port)
	db,err = api.NewSimpleDolphinDBClient(context.TODO(), host, configBase.Dolphindb.Username, configBase.Dolphindb.Password)
	if err != nil {
		// Handle exceptions
		panic(err)
	}
	return db, err
}
