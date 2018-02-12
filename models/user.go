package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"github.com/CryptocurrencyApp/CoinpocketServer/lib/random"
	"github.com/CryptocurrencyApp/CoinpocketServer/lib/hash"
)

type User struct {
	Id        string    `json:"id" orm:"pk"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Birthday  time.Time `json:"birthday"`
	Mail      string    `json:"mail"`
	Password  string    `json:"password"`
	Salt      string    `json:"salt"`
	CreatedAt time.Time `json:"created_at" orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u *User) (i string, err error) {
	o := orm.NewOrm()
	u.Id = random.RandString6(10)
	u.Salt = random.RandString6(20)
	u.Password = hash.ToHash(u.Password, u.Salt)

	id, err := o.Insert(u)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		return "", err
	}
	return u.Id, nil
}

func GetUserById(id string) (u *User, err error) {
	o := orm.NewOrm()
	u = &User{Id: id}
	if err = o.Read(u); err == nil {
		return u, nil
	}
	return nil, err
}

func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := &User{Id: m.Id}

	if err = o.Read(v); err == nil {
		// passwordは一旦更新できないものとする
		m.Password = v.Password
		m.Salt = v.Salt
		m.CreatedAt = v.CreatedAt

		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println(" Number of records updated in database:", num)
			return nil
		}
	}
	return err
}

func Login(request *User) (userId string, err error) {
	o := orm.NewOrm()
	user := &User{Mail: request.Mail}

	if err = o.Read(user, "mail"); err == nil {
		fmt.Println(user)
		return user.Id, nil
	}

	return
}
