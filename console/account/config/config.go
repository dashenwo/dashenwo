package config

import "time"

var (
	ConfPath = "./conf/"
	AppId    = "com.dashenwo.srv.account"
)

// 数据库配置信息
type Database struct {
	LogMode     bool
	AutoMigrate bool
	Engine      string
	Host        string
	Port        string
	User        string
	Password    string
	Name        string

	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

// es配置信息
type Elasticsearch struct {
	Hosts     []string  //host地址
	Sniff     bool      //是否使用监听机制，新加入节点或者有节点死掉
	BasicAuth basicAuth //如果设置了用户名密码认证
}

type basicAuth struct {
	UserName string //用户名
	PassWord string //密码
}
