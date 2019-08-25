package models

import (
	"github.com/tonyzzzzzz/sspanel-go/utils"
)

// SsNode ORM Model
// Written By Indexyz @ Cloudhammer/Seeds
type SsNode struct {
	ID                     int     `gorm:"column:id;NOT NULL;PRIMARY KEY;"`
	Name                   string  `gorm:"column:name;NOT NULL;"`
	Type                   int     `gorm:"column:type;NOT NULL;"`
	Server                 string  `gorm:"column:server;NOT NULL;"`
	Method                 string  `gorm:"column:method;NOT NULL;"`
	Info                   string  `gorm:"column:info;NOT NULL;"`
	Status                 string  `gorm:"column:status;NOT NULL;"`
	Sort                   int     `gorm:"column:sort;NOT NULL;"`
	CustomMethod           int     `gorm:"column:custom_method;NOT NULL;"`
	TrafficRate            float64 `gorm:"column:traffic_rate;NOT NULL;"`
	NodeClass              int     `gorm:"column:node_class;NOT NULL;"`
	NodeSpeedlimit         float64 `gorm:"column:node_speedlimit;NOT NULL;"`
	NodeConnector          int     `gorm:"column:node_connector;NOT NULL;"`
	NodeBandwidth          int64   `gorm:"column:node_bandwidth;NOT NULL;"`
	NodeBandwidthLimit     int64   `gorm:"column:node_bandwidth_limit;NOT NULL;"`
	BandwidthlimitResetday int     `gorm:"column:bandwidthlimit_resetday;NOT NULL;"`
	NodeHeartbeat          int64   `gorm:"column:node_heartbeat;NOT NULL;"`
	NodeIP                 string  `gorm:"column:node_ip;type:text;"`
	NodeGroup              int     `gorm:"column:node_group;NOT NULL;"`
	CustomRss              int     `gorm:"column:custom_rss;NOT NULL;"`
	MuOnly                 int     `gorm:"column:mu_only;"`
	OnlineUser             int
}

// GetOnlineUserCount returns the online user of specified node
func (node *SsNode) GetOnlineUserCount() int {
	var onlineLog SsNodeOnlineLog
	db := utils.GetMySQLInstance().Database
	db.Where("node_id = ?", node.ID).Last(&onlineLog)
	return onlineLog.OnlineUser
}

func (SsNode) TableName() string {
	return "ss_node"
}
