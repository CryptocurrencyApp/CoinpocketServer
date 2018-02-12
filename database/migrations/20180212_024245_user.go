package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20180212_024245 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20180212_024245{}
	m.Created = "20180212_024245"

	migration.Register("User_20180212_024245", m)
}

// Run the migrations
func (m *User_20180212_024245) Up() {
	sql := "ALTER TABLE user MODIFY COLUMN name nvarchar(100) NOT NULL," +
		"MODIFY COLUMN sex nvarchar(10) NOT NULL," +
		"MODIFY COLUMN birthday datetime," +
		"MODIFY COLUMN mail nvarchar(100) NOT NULL," +
		"MODIFY COLUMN password nvarchar(300) NOT NULL," +
		"MODIFY COLUMN salt nvarchar(300) NOT NULL;"

		m.SQL(sql)
}

// Reverse the migrations
func (m *User_20180212_024245) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
