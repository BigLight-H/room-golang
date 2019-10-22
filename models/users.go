/*
@Time : 2019-10-13 17:40
@Author : biglight
@File : users
@Software: GoLand
*/
package models

import "time"

type Users struct {
	Id      int
	Name    string
	Pwd     string
	Mobile  string
	Email   string
	Logo   string
	Created time.Time
	Updated time.Time
}

func (m *Users) TableName() string {
	return TableName("users")
}
