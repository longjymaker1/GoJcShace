package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id            int64           `pk:"auto;column(id)"`
	Name          string          `orm:"size(32);column(name);comment('昵称')"`
	Email         string          `orm:"size(64);column(email);comment('邮箱')"`
	Status        int64           `orm:"size(10);column(status);default(1);comment('状态,正常1,异常0')"`
	Integral      int64           `orm:"size(10);column(integral);comment('积分,RMB:积分,1:1');default(0)"`
	CreateTime    time.Time       `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime    time.Time       `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
	UserLoginInfo *UsersLoginInfo `orm:"rel(one)"`
}

type UsersLoginInfo struct {
	Id        int64  `pk:"auto;column(id)"`
	UserID    int64  `orm:"column(user_id);comment('用户id')"`
	LoginName string `orm:"size(16);column(login_name);comment('用户登录账户')"`
	Pwd       string `orm:"size(32);column(password);comment('用户登录密码')"`
	User      *Users `orm:"reverse(one)"`
}

type Collect struct {
	Id int64 `pk:"auto;column(id)"`
}

func init() {
	orm.RegisterModel(new(Users), new(UsersLoginInfo))
}
