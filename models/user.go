package models

import "time"

type User struct {
	Id           string    `xorm:"varchar(32) NOT NULL pk unique COMMENT('用户表')" json:"id"`
	RecordId     string    `xorm:"varchar(32) DEFAULT NULL COMMENT('记录 ID')" json:"record_id"`
	Username     string    `xorm:"varchar(20) DEFAULT NULL COMMENT('用户名')" json:"username"`
	Password     string    `xorm:"varchar(32) DEFAULT NULL COMMENT('密码')" json:"password"`
	RealName     string    `xorm:"varchar(20) DEFAULT NULL COMMENT('真实姓名')" json:"real_name"`
	Email        string    `xorm:"varchar(20) DEFAULT NULL COMMENT('邮箱')" json:"email"`
	Phone        string    `xorm:"varchar(11) DEFAULT NULL COMMENT('手机号')" json:"phone"`
	Status       int8       `xorm:"tinyint(1)  DEFAULT '0' COMMENT('正常0 冻结1')" json:"status"`
	Avatar       string    `xorm:"varchar(128) DEFAULT NULL COMMENT('头像')" json:"avatar"`
	DepartmentId string    `xorm:"varchar(32) DEFAULT NULL COMMENT('所属于公司 所属公司ID')" json:"department_id"`
	CreatedAt    time.Time `xorm:"datetime DEFAULT NULL COMMENT('创建时间')" json:"created_at"`
	UpdatedAt    time.Time `xorm:"datetime DEFAULT NULL COMMENT('更新时间')" json:"updated_at"`
}
