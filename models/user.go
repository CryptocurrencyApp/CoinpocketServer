package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math/rand"
	"strconv"
	"time"
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
	CreatedAt time.Time `json:"created_at" form:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `json:"updated_at" orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(u *User) (i string, err error) {
	o := orm.NewOrm()
	rand.Seed(time.Now().UnixNano())
	u.Id = strconv.Itoa(rand.Intn(15))
	u.Salt =  strconv.Itoa(rand.Intn(20))
	u.Password = hash.ToHash(u.Password)

	id, err := o.Insert(&u)
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
	v := User{Id: m.Id}

	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
		return nil
	}
	return err
}

//func Login(username, password string) bool {
//	for _, u := range UserList {
//		if u.Username == username && u.Password == password {
//			return true
//		}
//	}
//	return false
//}
