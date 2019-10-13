/*
@Time : 2019-10-13 17:51
@Author : biglight
@File : permissions
@Software: GoLand
*/
package models

import "time"

type Permissions struct {
	Id      int
	RolesId int
	UsersId int
	Created time.Time
	Updated time.Time
}

func (m *Permissions) TableName() string {
	return TableName("permissions")
}