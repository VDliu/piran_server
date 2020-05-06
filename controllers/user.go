package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"pirain_server/bean/req"
	"pirain_server/common"
	"pirain_server/models"
	"fmt"
)


type UserController struct {
	beego.Controller
}

type JSONDemo struct {
	Code int
	Msg  string
}



func (c *UserController) Register() {

	res := common.Res{}

	req := req.RegisterReq{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	if err != nil {
		res.SetRes(1,"error parms",false)
	} else {
		err = models.AddUser(req.Nick_name,req.Invite_code,req.Password)
		if err != nil {
			res.SetRes(0,"Register ok",true)
		}else {
			res.SetRes(0,err.Error(),false)
		}
	}

	c.Data["json"] = res
	c.ServeJSON()
}

func (c *UserController) UserCharge() {
	res := common.Res{}
	req := models.Charge{Amount:-10.0,User:1,Charge_man:"huzi"}
	//err := json.Unmarshal(c.Ctx.Input.RequestBody, &req)
	//if err != nil {
	//	res.SetRes(1,"error parms",false)
	//
	//}else {
	//	res.SetRes(0,"Charge ok",true)
	//}

	err := models.ChargeC(&req)
	if err != nil {
		res.SetMsg(err.Error())
	}else{
		res.SetMsg("ok")
	}

	c.Data["json"] = res
	c.ServeJSON()

}



func (c *UserController) Get() {
	users := models.Users()
	if users != nil {
		c.Data["Users"] = users
		c.TplName = "index.tpl"
	}else {
		c.Ctx.WriteString("error")
	}

}

func (c *UserController) Trans() {
	tran := models.Transition{FromId:1,ToId:3,Amount:1,FromInvite:"huzi",ToInvite:"222"}
	s := make([]models.Transition,0)
	s2 := append(s,tran)

	fmt.Printf("controoler---- %d len\\n",len(s2))
	err := models.Transtion(s2)

	if err != nil {
		c.Ctx.WriteString(err.Error())
	}else {
		c.Ctx.WriteString("tans ok")
	}

}

