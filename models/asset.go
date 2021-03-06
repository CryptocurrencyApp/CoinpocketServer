package models

import (
	"errors"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Asset struct {
	Id     int64  `orm:"auto;pk"`
	UserId string `json:"user_id";orm:"size(128)"`
	CoinId string `json:"coin_id";orm:"size(128)"`
	Amount string `json:"amount"`
}

func init() {
	orm.RegisterModel(new(Asset))
}

// AddAsset insert a new Asset into database and returns
// last inserted Id on success.
func AddAsset(m *Asset) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAssetById retrieves Asset by Id. Returns error if
// Id doesn't exist
func GetAssetById(userId string) (v *[]Asset, errMsg string) {
	o := orm.NewOrm()
	table := o.QueryTable("asset")

	var asset []Asset

	cnt, _ := table.Filter("UserId", userId).All(&asset, "CoinId", "Amount")
	if cnt == 0 {
		return nil, "No result found."
	} else {
		return &asset, ""
	}
}

// GetAllAsset retrieves all Asset matches certain condition. Returns empty list if
// no records exist
func GetAllAsset(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Asset))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Asset
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAsset updates Asset by Id and returns error if
// the record to be updated doesn't exist
func UpdateAssetById(m *Asset) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("asset").Filter("CoinId", m.CoinId).Filter("UserId", m.UserId).Update(orm.Params{
		"Amount": m.Amount,
	})
	return
}

// DeleteAsset deletes Asset by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAsset(m *Asset) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable("asset").Filter("CoinId", m.CoinId).Filter("UserId", m.UserId).Delete()
	return
}
