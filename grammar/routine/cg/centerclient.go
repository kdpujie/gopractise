package cg

import (
	"encoding/json"
	"errors"
	"routine/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

//增加玩家
func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(*player)
	if err != nil {
		return err
	}
	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

//玩家退出
func (client *CenterClient) RemovePlayer(name string) error {
	resp, _ := client.Call("removeplayer", name)
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}

//玩家列表
func (client *CenterClient) ListPlayer(params string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", params)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

//广播
func (client *CenterClient) Broadcast(message string) error {
	m := &Message{Content: message} //构造消息
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
