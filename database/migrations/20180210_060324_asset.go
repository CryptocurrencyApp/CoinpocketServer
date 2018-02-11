package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Asset_20180210_060324 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Asset_20180210_060324{}
	m.Created = "20180210_060324"

	migration.Register("Asset_20180210_060324", m)
}

// Run the migrations
func (m *Asset_20180210_060324) Up() {
	sql := "CREATE TABLE asset(" +
		"id int AUTO_INCREMENT PRIMARY KEY," +
		"user_id nvarchar(100)," +
		"coin_id nvarchar(100)," +
		"amount nvarchar(100)" +
		");"

	m.SQL(sql)
}

// Reverse the migrations
func (m *Asset_20180210_060324) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE asset;")
}
