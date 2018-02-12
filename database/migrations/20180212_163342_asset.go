package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Asset_20180212_163342 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Asset_20180212_163342{}
	m.Created = "20180212_163342"

	migration.Register("Asset_20180212_163342", m)
}

// Run the migrations
func (m *Asset_20180212_163342) Up() {
	sql := "ALTER TABLE asset " +
		"ADD UNIQUE INDEX user_coin(user_id, coin_id);"

	m.SQL(sql)
}

// Reverse the migrations
func (m *Asset_20180212_163342) Down() {
}
