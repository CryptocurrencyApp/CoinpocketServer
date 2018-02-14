package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id        int64     `json:"id";orm:"auto"`
	UserId    string    `json:"user_id";orm:"size(128)"`
	UserName  string    `json:"user_name";orm:"size(128)"`
	Comment   string    `json:"comment";orm:"size(256)"`
	Good      int       `json:"good"`
	Bad       int       `json:"bad"`
	Url       string    `json:"url";orm:"size(512)"`
	SiteTitle     string    `json:"site_title";orm:"size(512)"`
	Image     string    `json:"image";orm:"size(512)"`
	SiteName      string    `json:"site_name";orm:"size(512)"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Article))
}

// AddArticle insert a new Article into database and returns
// last inserted Id on success.
func AddArticle(m *Article) (i int64, err error) {
	o := orm.NewOrm()

	id, err := o.Insert(m)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		return 0, err
	}
	return m.Id, nil
}

// GetArticleById retrieves Article by Id. Returns error if
// Id doesn't exist
func GetArticleById(id int64) (v *Article, err error) {
	o := orm.NewOrm()
	v = &Article{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetArticlesByUserId(uid string) (ml *[]Article, err error) {
	o := orm.NewOrm()
	table := o.QueryTable("article")

	var article []Article

	_, err = table.Filter("UserId", uid).All(&article)
	if err != nil {
		return nil, err
	} else {
		return &article, nil
	}
}

func GetAllArticle() (ml *[]Article, err error) {
	o := orm.NewOrm()
	table := o.QueryTable("article")

	var article []Article

	_, err = table.All(&article)
	if err != nil {
		return nil, err
	} else {
		return &article, nil
	}
}

func ToggleGood2Article(id int64, isAdd bool) (err error) {
	o := orm.NewOrm()
	v := &Article{Id: id}
	evaluation := -1
	if isAdd {
		evaluation = 1
	}

	if err = o.Read(v); err == nil {
		var num int64
		v.Good = v.Good + evaluation
		if num, err = o.Update(v, "good"); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return nil
		}
	}
	return err
}

func ToggleBad2Article(id int64, isAdd bool) (err error) {
	o := orm.NewOrm()
	v := &Article{Id: id}
	evaluation := -1
	if isAdd {
		evaluation = 1
	}

	if err = o.Read(v); err == nil {
		var num int64
		v.Bad = v.Bad + evaluation
		if num, err = o.Update(v, "bad"); err == nil {
			fmt.Println("Number of records updated in database:", num)
			return nil
		}
	}
	return err
}

// DeleteArticle deletes Article by Id and returns error if
// the record to be deleted doesn't exist
func DeleteArticle(id int64) (err error) {
	o := orm.NewOrm()
	v := Article{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Article{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
