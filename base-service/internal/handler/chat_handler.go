package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mackelaophu/ChatApp/base-service/internal/model"
	"github.com/mackelaophu/ChatApp/base-service/internal/service"
	"github.com/mackelaophu/ChatApp/base-service/internal/websocket"
	"github.com/olahol/melody"
)

type ChatHandler struct {
	Service *service.ChatService
}

func NewChatHandler(Service *service.ChatService) *ChatHandler {
	return &ChatHandler{Service: Service}
}

func (Handler *ChatHandler) Init() {
	ws := websocket.NewWebSocket()
	m := ws.M
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		err := m.HandleRequest(w, r)
		if err == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	http.HandleFunc("/register", Handler.processRegister)
	m.HandleMessage(Handler.processDispatch)
	const port = "8080"
	fmt.Println("listtening on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
func (Handler *ChatHandler) processDispatch(s *melody.Session, msg []byte) {
	var chatObj model.ChatObjc
	if err := json.Unmarshal(msg, &chatObj); err != nil {
		fmt.Println("parse json failed")
	}
	switch chatObj.ChatType {
	case model.GetAllUser:
		Handler.GetAllUser(s)
	case model.Disconnect:
		Handler.processDisconnect(s)
	case model.Message:
		Handler.processMessage(&chatObj)
	}
}
func (Handler *ChatHandler) processRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "only support POST method", http.StatusMethodNotAllowed)
		return
	}
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "parse json failed", http.StatusBadRequest)
		return
	}
	Handler.Service.SaveUser(&user)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
func (Handler *ChatHandler) GetAllUser(s *melody.Session) {
	users := Handler.Service.GetAllUser()
	resp := model.ChatObjResponse{
		ChatType: model.GetAllUser,
		Data:     users,
	}
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("invalid json")
		return
	}
	s.Write(data)
}
func (Handler *ChatHandler) processDisconnect(s *melody.Session) {

}
func (Handler *ChatHandler) processMessage(chatObj *model.ChatObjc) {

}
