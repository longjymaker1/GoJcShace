package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// Users 用户表
type Users struct {
	Id            int64           `pk:"auto"`
	Name          string          `orm:"size(32)"`
	Email         string          `orm:"size(64)"`
	Status        int64           `orm:"size(10);default(1)"`
	Integral      int64           `orm:"size(10);default(0)"`
	CreateTime    time.Time       `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time       `orm:"auto_now;type(datetime)"`
	UserLoginInfo *UsersLoginInfo `orm:"rel(one)"`
}

// UsersLoginInfo 用户注册信息表
type UsersLoginInfo struct {
	Id         int64     `pk:"auto;column(id)"`
	UserID     int64     `orm:"column(user_id)" description:"用户id"`
	LoginName  string    `orm:"size(16);column(login_name)" description:"用户登录账户"`
	Pwd        string    `orm:"size(32);column(password)" description:"用户登录密码"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

// Attention 用户关注表
type Attention struct {
	Id         int64     `pk:"auto"`
	UserId     *Users    `orm:"rel(fk)"`
	AttUserId  *Users    `orm:"rel(fk)"`
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateTime time.Time `orm:"auto_now;type(datetime)"`
}

// Country 国家表--我们的目标是P遍全球
type Country struct {
	Id         int64     `pk:"auto"`
	Country    string    `orm:"size(32)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
}

// Province 省份表 -- 省、直辖市
type Province struct {
	Id         int64     `pk:"auto"`
	Province   string    `orm:"size(32)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
	CountryId  *Country  `orm:"rel(one)"`
}

// City 城市表 -- 市, 直辖市区
type City struct {
	Id         int64     `pk:"auto"`
	City       string    `orm:"size(32)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
	ProvinceId *Province `orm:"rel(one)"`
}

// District 城市区域表 -- 区
type District struct {
	Id         int64     `pk:"auto"`
	District   string    `orm:"size(32)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
	CityId     *City     `orm:"rel(one)"`
}

// ArticleType 文章类型
type ArticleType struct {
	Id              int64     `pk:"auto"`
	ArticleTypeName string    `orm:"size(32)"` // LF, ZJ, XY, ZL, blackList
	CreateTime      time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime      time.Time `orm:"type(datetime);auto_now"`
}

// Article 文章主表
type Article struct {
	Id            int64        `pk:"auto"`
	Title         string       `orm:"size(64)"`
	ViewNum       int32        `orm:"size(10);default(0)"`
	CommentNum    int32        `orm:"size(10);default(0)"`
	LikeNum       int32        `orm:"size(10);default(0)"`
	User          *Users       `orm:"rel(one)"`
	ArticleTypeId *ArticleType `orm:"rel(one)"`
	CountryId     *Country     `orm:"rel(one)"`
	ProvinceId    *Province    `orm:"rel(one)"`
	CityId        *City        `orm:"rel(one)"`
	DistrictId    *District    `orm:"rel(one)"`
	MainPhotoPath string       `orm:"size(128);default('')'"`
	CreateTime    time.Time    `orm:"type(datetime);auto_now_add"`
	UpdateTime    time.Time    `orm:"type(datetime);auto_now"`
}

// ArticlePhotoPath 文章图片路径
type ArticlePhotoPath struct {
	id         int64     `pk:"auto"`
	PhotoPaths string    `orm:"size(256)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
}

// UserLikes 用户收藏(like),浏览表
type UserLikes struct {
	Id         int64     `pk:"auto;column(id)"`
	UserId     *Users    `orm:"rel(fk)"`
	ArticleId  *Article  `orm:"rel(fk)"`
	ActType    int32     `orm:"size(10);default(1)"` // 1浏览; 2收藏
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
}

// ArticleMessage 文章表, 文章内容
type ArticleMessage struct {
	id             int64             `pk:"auto;column(id)"`
	ArticleId      *Article          `orm:"rel(one)"`
	HookerNum      string            `orm:"size(10);default('1')"`  // 数量
	HookerAge      string            `orm:"size(10);default('20')"` // 年龄
	HookerQuality  string            `orm:"size(10);default('')"`   // 质量
	HookerShape    string            `orm:"size(10);default('')"`   // 外形
	ServiceContent string            `orm:"size(10);default('')"`   // 项目
	BusinessHours  string            `orm:"size(10);default('')"`   //时间
	Equipment      string            `orm:"size(10);default('')"`   // 环境
	Security       string            `orm:"size(10);default('')"`   // 安全
	Address        string            `orm:"size(10);default('')"`   // 地址
	ContactWay     string            `orm:"size(10);default('')"`   // 联系方式
	Evaluation     string            `orm:"size(10);default('1')"`  // 总和评价
	PhotoList      *ArticlePhotoPath `orm:"rel(one)"`               // 图片列表
	PriceP         float64           `orm:"digits(12);decimals(4)"` // 一次价格
	PricePp        float64           `orm:"digits(12);decimals(4)"` // 两次价格
	PriceY         float64           `orm:"digits(12);decimals(4)"` // 夜价格
	Content        string            `orm:"size(512);default('')"`  // 正文
	CreateTime     time.Time         `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime     time.Time         `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
}

type Comment struct {
	id          int64  `pk:"auto;column(id)"`
	ComContent  string `orm:"size(512);default('')"` // 评论正文
	CommentType int32  `orm:"size(10);default(1)"`   // 1评论文章; 2回复评论
	CommentId   int64  `orm:"size(32);default(10)"`  // type为1写文章id; 2为comment id
}

func init() {
	orm.RegisterModel(new(Users), new(UsersLoginInfo), new(Attention), new(Country),
		new(Province), new(City), new(District), new(ArticleType), new(Article),
		new(ArticlePhotoPath), new(UserLikes), new(ArticleMessage), new(Comment))
}
