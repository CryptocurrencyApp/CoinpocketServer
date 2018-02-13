package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Article_20180212_145304 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Article_20180212_145304{}
	m.Created = "20180212_145304"

	migration.Register("Article_20180212_145304", m)
}

// Run the migrations
func (m *Article_20180212_145304) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	sql := "CREATE TABLE article(" +
		"id int AUTO_INCREMENT PRIMARY KEY," +
		"user_id nvarchar(80) NOT NULL, " +
		"user_name nvarchar(128) NOT NULL, " +
		"url nvarchar(512) NOT NULL, " +
		"image nvarchar(512), " +
		"comment nvarchar(256) NOT NULL, " +
		"good int DEFAULT 0, " +
		"bad int DEFAULT 0, " +
		"created_at datetime, " +
		"updated_at datetime);"
	m.SQL(sql)
}

// Reverse the migrations
func (m *Article_20180212_145304) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE article;")
}
