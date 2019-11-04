package models

import "time"

type User struct {
	Id         int       `xorm:"NOT NULL pk autoincr int(11) comment('主键1111')" json:"id"`
	TenantCode string    `xorm:" varchar(12) default '000000' comment('租户编号')" json:"tenant_code"`
	Account    string    `xorm:"varchar(45) default null comment('账号')" json:"account"`
	Password   string    `xorm:"varchar(45) default null comment('密码')" json:"password"`
	Name       string    `xorm:"varchar(20) default null comment('昵称')" json:"name"`
	RealName   string    `xorm:" varchar(10) default null comment('真名')" json:"real_name"`
	Email      string    `xorm:"varchar(45) default null comment('邮箱')" json:"email"`
	Phone      string    `xorm:"varchar(45) default null comment('手机')" json:"phone"`
	Birthday   time.Time `xorm:"datetime default null comment('生日')" json:"birthday"`
	Sex        int       `xorm:"smallint(6) default null comment('性别')" json:"sex"`
	RoleId     string    `xorm:"varchar(255) default null comment('角色id')" json:"role_id"`
	DeptId     string    `xorm:"varchar(255) default null comment('部门id')" json:"dept_id"`
	CreateUser int       `xorm:"int(11) default null comment('创建人')" json:"create_user"`
	CreateTime time.Time `xorm:"datetime default null comment('创建时间')" json:"create_time"`
	UpdateUser int       `xorm:"int(11) default null comment('修改人')" json:"update_user"`
	UpdateTime time.Time `xorm:"datetime default null comment('修改时间')" json:"update_time"`
	Status     int       `xorm:"int(2) default null comment('状态')" json:"status"`
	IsDeleted  int       `xorm:"int(2) default '0' comment('是否已删除')" json:"is_deleted"`
}
