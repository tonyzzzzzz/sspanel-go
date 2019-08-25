package models

// SsNodeOnlineLog ORM Model
// Written By Indexyz @ Cloudhammer/Seeds
type SsNodeOnlineLog struct {
	ID         int `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
	NodeID     int `gorm:"column:node_id;NOT NULL;"`
	OnlineUser int `gorm:"column:online_user;NOT NULL;"`
	LogTime    int `gorm:"column:log_time;NOT NULL;"`
}

func (SsNodeOnlineLog) TableName() string {
	return "ss_node_online_log"
}
