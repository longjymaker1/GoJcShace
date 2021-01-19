package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Users 用户表
type Users struct {
	Id            int64           `pk:"auto;column('Id')"`
	Name          string          `orm:"size(32)"`
	Email         string          `orm:"size(64)"`
	Status        int64           `orm:"size(10);default(1)"`
	Integral      int64           `orm:"size(10);default(0)"`
	Jcb           float64         `orm:"digits(12);decimals(4)"` // 会员当前金币数量
	AllJcb        float64         `orm:"digits(12);decimals(4)"` // 累计兑换金币
	CreateTime    time.Time       `orm:"auto_now_add;type(datetime)"`
	UpdateTime    time.Time       `orm:"auto_now;type(datetime)"`
	UserLoginInfo *UsersLoginInfo `orm:"rel(one)"`
	VipLevel      *VipLevel       `orm:"rel(fk)"`        // 当前会员你等级id
	VipEndTime    time.Time       `orm:"type(datetime)"` // VIP等级截止时间
}

// UsersLoginInfo 用户注册信息表
type UsersLoginInfo struct {
	Id         int64     `pk:"auto;column(Id)"`
	UserId     int64     `orm:"column(user_Id)" description:"用户Id"`
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
	Id         int64     `pk:"auto"`
	PhotoPaths string    `orm:"size(256)"`
	CreateTime time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time `orm:"type(datetime);auto_now"`
}

// UserLikes 用户收藏(like),浏览表
type UserLikes struct {
	Id         int64     `pk:"auto;column(Id)"`
	UserId     *Users    `orm:"rel(fk)"`
	ArticleId  *Article  `orm:"rel(fk)"`
	ActType    int32     `orm:"size(10);default(1)"` // 1浏览; 2收藏
	CreateTime time.Time `orm:"auto_now_add;type(datetime);column(create_time);comment('创建时间')"`
	UpdateTime time.Time `orm:"auto_now;type(datetime);column(update_time);comment('更新时间')"`
}

// ArticleMessage 文章表, 文章内容
type ArticleMessage struct {
	Id             int64             `pk:"auto;column(Id)"`
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

// Comment 评论表
type Comment struct {
	Id          int64  `pk:"auto;column(Id)"`
	ComContent  string `orm:"size(512);default('')"` // 评论正文
	CommentType int32  `orm:"size(10);default(1)"`   // 1评论文章; 2回复评论
	CommentId   int64  `orm:"size(32);default(10)"`  // type为1写文章Id; 2为comment Id
}

// CreateVirtualJcCurrency JC币表，记录创建时间、金额、数量、状态、兑换数量
type VirtualJcb struct {
	Id          int64     `pk:"auto;column(Id)"`
	Number      int64     `orm:"size(10);default(0)"`    // 创建数量
	FaceValue   int64     `orm:"size(10);default(0)"`    // 面值
	Status      int64     `orm:"size(10);default(1)"`    // 状态1正常，0不可用，2全部兑换
	Money       float64   `orm:"digits(12);decimals(4)"` // 对应RMB金额
	ExchangeNum int64     `orm:"size(10);default(0)"`    // 已兑换数量
	CreateTime  time.Time `orm:"type(datetime);auto_now_add"`
	UpdateTime  time.Time `orm:"type(datetime);auto_now"`
}

// VipLevel 会员类型等级表，记录会员类型及标准
type VipLevel struct {
	Id        int64  `pk:"auto;column('Id')"`
	LevelName string `orm:"size(10);default('')"` // vip等级名称, 会员, VIP1--周, VIP2--月, VIP3--季, VIP4--年
	JcbNum    int64  `orm:"size(10);default(0)"`  // 每个等级需要的JCB数量
	ValidDay  int64  `orm:"size(10)ldefault(10)"` // 不同VIP等级有效时间
}

// VirtualJcbExchangeLog JcB兑换记录
type VirtualJcbExchangeLog struct {
	Id         int64       `pk:"auto"`
	UserId     *Users      `orm:"rel(fk)"`
	JcbId      *VirtualJcb `orm:"rel(fk)"`
	CreateTime time.Time   `orm:"type(datetime);auto_now_add"`
	UpdateTime time.Time   `orm:"type(datetime);auto_now"`
}

func init() {
	orm.RegisterModel(new(Users), new(UsersLoginInfo), new(Attention), new(Country),
		new(Province), new(City), new(District), new(ArticleType), new(Article),
		new(ArticlePhotoPath), new(UserLikes), new(ArticleMessage), new(Comment),
		new(VirtualJcb), new(VirtualJcbExchangeLog), new(VipLevel))
}
