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
	sql := "CREATE TABLE Users(" +
		"Id nvarchar(80) PRIMARY KEY," +
		"Name nvarchar(100), " +
		"Sex nvarchar(10), " +
		"Birthday datetime, " +
		"Mail nvarchar(100), " +
		"Password nvarchar(300), " +
		"Salt nvarchar(300), " +
		"crated_at datetime, " +
		"update_at datetime);"

	m.SQL(sql)

}

// Reverse the migrations
func (m *Users_20180204_220422) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
