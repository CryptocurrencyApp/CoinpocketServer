package models

import (
	"strconv"
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
	"math/rand"
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

func AddUser(u User) (i string, err error) {
	o := orm.NewOrm()
	rand.Seed(time.Now().UnixNano())

	id, err := o.Insert(&u)
	fmt.Printf("ID: %d, ERR: %v\n", id, err)
	if err != nil {
		return "", err
	}
	u.Id = strconv.Itoa(rand.Intn(15))
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

//
//func GetAllUsers() map[string]*User {
//	return UserList
//}
//
//func UpdateUser(uid string, uu *User) (a *User, err error) {
//	if u, ok := UserList[uid]; ok {
//		if uu.Username != "" {
//			u.Username = uu.Username
//		}
//		if uu.Password != "" {
//			u.Password = uu.Password
//		}
//		if uu.Profile.Age != 0 {
//			u.Profile.Age = uu.Profile.Age
//		}
//		if uu.Profile.Address != "" {
//			u.Profile.Address = uu.Profile.Address
//		}
//		if uu.Profile.Gender != "" {
//			u.Profile.Gender = uu.Profile.Gender
//		}
//		if uu.Profile.Email != "" {
//			u.Profile.Email = uu.Profile.Email
//		}
//		return u, nil
//	}
//	return nil, errors.New("User Not Exist")
//}
//
//func Login(username, password string) bool {
//	for _, u := range UserList {
//		if u.Username == username && u.Password == password {
//			return true
//		}
//	}
//	return false
//}
//
//func DeleteUser(uid string) {
//	delete(UserList, uid)
//}
