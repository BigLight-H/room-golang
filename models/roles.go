/*
@Time : 2019-10-13 17:44
@Author : biglight
@File : roles
@Software: GoLand
*/
package models

import "time"

type Roles struct {
	Id      int
	Role    string
	Created time.Time
	Updated time.Time
}

func (m *Roles) TableName() string {
	return TableName("roles")
}