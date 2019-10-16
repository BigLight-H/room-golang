/*
@Time : 2019-10-13 17:45
@Author : biglight
@File : types
@Software: GoLand
*/
package models

import (
	"time"
)

type Types struct {
	Id      int
	Type    string
	Money   float64
	Pid     int
	Path    string
	Created time.Time
	Updated time.Time
}

func (m *Types) TableName() string {
	return TableName("types")
}
