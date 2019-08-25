package models

import (
	"time"
)

// User ORM Model
// Written By Indexyz @ Cloudhammer/Seeds
type User struct {
	Id                 int       `gorm:"column:id;NOT NULL;PRIMARY KEY;" json:"-"`
	UserName           string    `gorm:"column:user_name;NOT NULL;"`
	Email              string    `gorm:"column:email;NOT NULL;"`
	Pass               string    `gorm:"column:pass; NOT NULL;" json:"-"`
	Passwd             string    `gorm:"column:passwd;NOT NULL;"`
	T                  int       `gorm:"column:t;NOT NULL;"`
	U                  int64     `gorm:"column:u;NOT NULL;"`
	D                  int64     `gorm:"column:d;NOT NULL;"`
	Plan               string    `gorm:"column:plan;NOT NULL;"`
	TransferEnable     int64     `gorm:"column:transfer_enable;NOT NULL;"`
	Port               int       `gorm:"column:port;NOT NULL;"`
	Switch             int       `gorm:"column:switch;NOT NULL;"`
	Enable             int       `gorm:"column:enable;NOT NULL;"`
	Type               int       `gorm:"column:type;NOT NULL;"`
	LastGetGiftTime    int       `gorm:"column:last_get_gift_time;NOT NULL;"`
	LastCheckInTime    int       `gorm:"column:last_check_in_time;NOT NULL;"`
	LastRestPassTime   int       `gorm:"column:last_rest_pass_time;NOT NULL;"`
	RegDate            time.Time `gorm:"column:reg_date;NOT NULL;"`
	InviteNum          int       `gorm:"column:invite_num;NOT NULL;"`
	Money              float64   `gorm:"column:money;NOT NULL;"`
	RefBy              int       `gorm:"column:ref_by;NOT NULL;"`
	ExpireTime         int       `gorm:"column:expire_time;NOT NULL;"`
	Method             string    `gorm:"column:method;NOT NULL;"`
	IsEmailVerify      int       `gorm:"column:is_email_verify;NOT NULL;"`
	RegIp              string    `gorm:"column:reg_ip;NOT NULL;"`
	NodeSpeedlimit     float64   `gorm:"column:node_speedlimit;NOT NULL;"`
	NodeConnector      int       `gorm:"column:node_connector;NOT NULL;"`
	IsAdmin            int       `gorm:"column:is_admin;NOT NULL;"`
	ImType             int       `gorm:"column:im_type;"`
	ImValue            string    `gorm:"column:im_value;type:text;"`
	LastDayT           int64     `gorm:"column:last_day_t;NOT NULL;"`
	SendDailyMail      int       `gorm:"column:sendDailyMail;NOT NULL;"`
	Class              int       `gorm:"column:class;NOT NULL;"`
	ClassExpire        time.Time `gorm:"column:class_expire;NOT NULL;"`
	ExpireIn           time.Time `gorm:"column:expire_in;NOT NULL;"`
	Theme              string    `gorm:"column:theme;NOT NULL;type:text;"`
	GaToken            string    `gorm:"column:ga_token;NOT NULL;type:text;"`
	GaEnable           int       `gorm:"column:ga_enable;NOT NULL;"`
	Pac                string    `gorm:"column:pac;type:longtext;"`
	Remark             string    `gorm:"column:remark;type:text;" json:"-"`
	NodeGroup          int       `gorm:"column:node_group;NOT NULL;"`
	AutoResetDay       int       `gorm:"column:auto_reset_day;NOT NULL;"`
	AutoResetBandwidth float64   `gorm:"column:auto_reset_bandwidth;NOT NULL;"`
	Protocol           string    `gorm:"column:protocol;"`
	ProtocolParam      string    `gorm:"column:protocol_param;"`
	Obfs               string    `gorm:"column:obfs;"`
	ObfsParam          string    `gorm:"column:obfs_param;"`
	ForbiddenIp        string    `gorm:"column:forbidden_ip;type:longtext;"`
	ForbiddenPort      string    `gorm:"column:forbidden_port;type:longtext;"`
	DisconnectIp       string    `gorm:"column:disconnect_ip;type:longtext;"`
	IsHide             int       `gorm:"column:is_hide;NOT NULL;"`
	IsMultiUser        int       `gorm:"column:is_multi_user;NOT NULL;"`
	TelegramId         int64     `gorm:"column:telegram_id;"`
}


func (User) TableName() string {
	return "user"
}
