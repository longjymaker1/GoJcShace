package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Users 用户表
type Users struct {
	Id            int64           `pk:"auto;column(id)"`
	Name          string          `orm:"size(32);column(name);comment('昵称')"`
	Email         string          `orm:"size(64);column(email);comment('邮箱')"`
	Status        int64           `orm:"size(10);column(status);default(1);comment('状态,正常1,异常0')"`
	Integral      int64           `orm:"size(10);column(integral);comment('积分,RMB:积分,1:1');default(0)"`
	CreateTime    time.Time       `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime    time.Time       `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
	UserLoginInfo *UsersLoginInfo `orm:"rel(one)"`
	Collect       *Collect        `orm:"rel(one)`
}

// UsersLoginInfo 用户注册信息表
type UsersLoginInfo struct {
	Id        int64  `pk:"auto;column(id)"`
	UserID    int64  `orm:"column(user_id);comment('用户id')"`
	LoginName string `orm:"size(16);column(login_name);comment('用户登录账户')"`
	Pwd       string `orm:"size(32);column(password);comment('用户登录密码')"`
	User      *Users `orm:"reverse(one)"`
}

// Collect 用户收藏表， 文章与用户是多对多的关系
type Collect struct {
	Id         int64     `pk:"auto;column(id)"`
	UserID     int64     `orm:"column(user_id);comment('用户id')`
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
	User       *Users    `orm:"reverse(one)"`
}

func init() {
	orm.RegisterModel(new(Users), new(UsersLoginInfo))
}
