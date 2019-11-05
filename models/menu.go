package models

import "time"

type Menu struct {
	Id        string    `xorm:"varchar(32) NOT NULL pk unique COMMENT('用户表')" json:"id"`
	RecordId  string    `xorm:"varchar(32) DEFAULT NULL COMMENT('记录 ID')" json:"record_id"`
	Name      string    `xorm:"varchar(32) DEFAULT NULL COMMENT('菜单权限名称')" json:"name"`
	Component string    `xorm:"varchar(32) DEFAULT NULL COMMENT('前端组件地址')" json:"component"`
	Path      string    `xorm:"varchar(32) DEFAULT NULL COMMENT('页面路径 前端组件URL')" json:"path"`
	Title     string    `xorm:"varchar(32) DEFAULT NULL COMMENT('菜单标题')" json:"title"`
	Icon      string    `xorm:"varchar(32) DEFAULT NULL COMMENT('菜单图标')" json:"icon"`
	SortOrder int16     `xorm:"smallint(4) DEFAULT '50' COMMENT('排序值')" json:"sort_order"`
	ParentId  string    `xorm:"varchar(32) DEFAULT NULL COMMENT('父ID 上级菜单ID')" json:"parent_id"`
	Level     string    `xorm:"smallint(3)  DEFAULT '0' COMMENT('层级')" json:"level"`
	Status    int8      `xorm:"tinyint(1) DEFAULT '0' COMMENT('正常0 冻结1')" json:"status"`
	Url       string    `xorm:"varchar(32) DEFAULT NULL COMMENT('菜单为网页连接 是URL')" json:"url"`
	CreatedAt time.Time `xorm:"datetime DEFAULT NULL COMMENT('创建时间')" json:"created_at"`
	UpdatedAt time.Time `xorm:"datetime DEFAULT NULL COMMENT('更新时间')" json:"updated_at"`
}
