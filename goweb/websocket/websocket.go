package websocket
import (
	"fmt"
	"net/http"
	"github.com/gorilla/websocket"
)

// 启动程序
func StartWebSocket() {

	http.HandleFunc("/acc", wsPage)
	http.ListenAndServe(":9999", nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())

	// currentTime := uint64(time.Now().Unix())
	// client := NewClient(conn.RemoteAddr().String(), conn, currentTime)

	// go client.read()
	// go client.write()
}