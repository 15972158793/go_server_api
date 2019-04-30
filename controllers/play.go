package controllers

import (
	"yukee/webapi/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

//玩法相关接口
type PlayController struct {
	beego.Controller
}

// @Title addplay
// @Description POST添加新的玩法信息
// @Param	body body models.Play true  "玩法的json信息"
// @Success 200 {string} 返回用户的唯一标识符uid
// @Failure 403 body is empty
// @router /addplay [post]
func (pc *PlayController) AddPlay() {
	var play models.Play
	json.Unmarshal(pc.Ctx.Input.RequestBody, &play)
	uid := models.AddPlay(play)
	pc.Data["json"] = uid
    pc.ServeJSON()
}

// @Title getplayinfo
// @Description 获取玩法信息
// @Success 200 {Object} models.Play
// @Failure 403 nil
// @Param	id	path string true "玩法ID"
// @router /getplayinfo/:id [get]
func (pc *PlayController) GetPlayInfo(){
	id := pc.Ctx.Input.Param(":id")
	value := models.GetPlayInfo(id)
	pc.Data["json"] = value
	pc.ServeJSON()
}