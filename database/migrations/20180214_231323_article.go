package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Article_20180214_231323 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Article_20180214_231323{}
	m.Created = "20180214_231323"

	migration.Register("Article_20180214_231323", m)
}

// Run the migrations
func (m *Article_20180214_231323) Up() {
	sql := "ALTER TABLE article ADD(" +
		"site_title nvarchar(512)," +
		"site_name nvarchar(512));"
	m.SQL(sql)
}

// Reverse the migrations
func (m *Article_20180214_231323) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
