package models

import (
	_"errors"
	"strconv"
	"time"
)
type Play struct {
	Id       string   `description:"玩法唯一标志"`
	GameId   string   `description:"玩法ID"`
	Level    string   `description:"玩法难度`
	Time     string   `description:"创建时间`
}

var (
	PlayList map[string]*Play
)

func init() {
	PlayList = make(map[string]*Play)
	p := Play{ "user_84423253","1","1", "2019-4-13 00:00:00"}
	PlayList["game_11111"] = &p
}

func AddPlay(u Play) string {
	//纳秒级别
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	PlayList[u.Id] = &u
	return u.Id
}

func GetPlayInfo(id string) *Play{
	for _, u := range PlayList {
		if u.Id == id {
			return u
		}
	}
	return nil
}

