package websocket

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

func WebSocketHandler(w http.ResponseWriter, req *http.Request){
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, req, nil)
	if err != nil{
		http.NotFound(w, req)
		return
	}
	conn.SetCloseHandler(func(code int, text string) error {
		return nil
	})

}

func GetMuxRouter() *mux.Router{
	rtr := mux.NewRouter()
	rtr.HandleFunc("/col/id:[a-zA-Z0-9\\-\\_]*/p:[a-zA-Z0-9\\-\\_]*", WebSocketHandler)
	return rtr
}
