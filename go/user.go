package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

//DbCreateUser 创建一个User键值对
func DbCreateUser(db *bolt.DB, Username string, Password string) {
	var User User
	var BucketName string = "User"

	User.Username = Username
	User.Password = Password

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

		buf, err := json.Marshal(User)
		if err != nil {
			return err
		}
		fmt.Println("创建User: ", User)

		return b.Put([]byte(User.Username), buf)
	})

	if err != nil {
		fmt.Println("User添加失败")
	}
}

//DbGetUser 得到一个User对象
func DbGetUser(db *bolt.DB, Username string) User {
	var User User = User{}
	var BucketName string = "User"

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		v := b.Get([]byte(Username))
		if v == nil {
			fmt.Println("User不存在")
		}
		err := json.Unmarshal(v, &User)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
	return User
}

//DbDelUser 删除一个User
func DbDelUser(db *bolt.DB, Username string) {
	var BucketName string = "User"

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		fmt.Println(string(b.Get([]byte(Username))))
		err := b.Delete([]byte(Username))
		return err
	})
}
