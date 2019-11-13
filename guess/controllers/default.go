package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"guess/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var subject models.Subject
	err:=func() error{
		id,err:=c.GetInt("id")
		beego.Info(id)//保存info级别的日志，记录下每次请求的结果
		if err!=nil{
			id=1
		}
		subject,err=models.GetSubject(id)
		if err!=nil{
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err!=nil{
		c.Ctx.WriteString("wrong params")
	}//判断整个模块是否有错误，有则返回错误信息

	//处理查询到的题目的信息
	//选项（option）使用的是json格式存储，这里要把题目选项做jsondecode，转换成map格式
	var option map[string]string
	if err=json.Unmarshal([]byte(subject.Option),&option);err!=nil{
		c.Ctx.WriteString("wrong params,json decode")
	}
	c.Data["ID"]=subject.Id
	c.Data["Option"]=option
	c.Data["Img"]="/static"+subject.Img
	c.TplName="guess.tpl"
}

//答案提交部分
func(c *MainController)Post(){
	var subject models.Subject
	err:=func() error{
		id,err:=c.GetInt("id")
		beego.Info(id)//保存info级别的日志，记录下每次请求的结果
		if err!=nil{
			id=1
		}
		subject,err=models.GetSubject(id)
		if err!=nil{
			return errors.New("subject not exist")
		}
		return nil
	}()

	if err!=nil{
		c.Ctx.WriteString("wrong params")
	}//判断整个模块是否有错误，有则返回错误信息
	answer:=c.GetString("key")
	right:=models.Answer(subject.Id,answer)
	c.Data["Right"]=right
	c.Data["Next"]=subject.Id+1
	c.Data["ID"]=subject.Id
	c.TplName="guess.tpl"
}