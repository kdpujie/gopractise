package cg

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	"learn.com/gopractise/grammar/routine/ipc"
)

var _ ipc.Server = &CenterServer{} //确认实现了Server接口
//消息
type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

//中央服务器
type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	//	rooms   []*Room
	mutex sync.RWMutex
}

//构造函数
func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}

//增加玩家
func (server *CenterServer) addPlayer(params string) error {
	player := NewPlayer()
	if err := json.Unmarshal([]byte(params), &player); err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock() //函数返回前执行unlock,有一定的性能开销
	server.players = append(server.players, player)
	return nil
}

//按名称移除
func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 { //只有一个玩家
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 { // 删除最后一个玩家
				server.players = server.players[:i]
			} else {
				server.players = append(server.players[:i], server.players[i+1:]...)
				fmt.Println("i=", i, "; 剩余元素:")
				for i, v := range server.players {
					fmt.Println("\t", i+1, ":", v)
				}
			}
			return nil
		}
	}
	return errors.New("Player not found.")
}

//列出玩家列表
func (server *CenterServer) listPlayer(params string) (players string, err error) {
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("No player online.")
	}
	return
}

//广播:向每个玩家发送消息
func (server *CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("No player online.")
	}
	return err
}

//处理命令
func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		if err := server.addPlayer(params); err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		if err := server.removePlayer(params); err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		if err := server.broadcast(params); err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &ipc.Response{Code: "200"}
}

//返回服务器名称
func (server *CenterServer) Name() string {
	return "CenterServer"
}
