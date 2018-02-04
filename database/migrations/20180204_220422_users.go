package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Users_20180204_220422 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Users_20180204_220422{}
	m.Created = "20180204_220422"

	migration.Register("Users_20180204_220422", m)
}

// Run the migrations
func (m *Users_20180204_220422) Up() {
	sql := "CREATE TABLE users(" +
		"id nvarchar(80) PRIMARY KEY," +
		"name nvarchar(100), " +
		"sex nvarchar(10), " +
		"birthday datetime, " +
		"mail nvarchar(100), " +
		"password nvarchar(300), " +
		"salt nvarchar(300), " +
		"created_at datetime, " +
		"update_at datetime);"

	m.SQL(sql)

}

// Reverse the migrations
func (m *Users_20180204_220422) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE users;")
}
