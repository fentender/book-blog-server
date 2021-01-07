package swagger

import (
	"fmt"
	"github.com/boltdb/bolt"
)

//DbCreateToken 创建一个Token键值对
func DbCreateToken(db *bolt.DB, Username string, Token string) {

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Token"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		fmt.Println("创建Token:", Username, Token)
		return b.Put([]byte(Username), []byte(Token))
	})

	if err != nil {
		fmt.Println("Token添加失败")
	}
}

//DbKeyofToken 得到Token的用户名
func DbKeyofToken(db *bolt.DB, Token string) []byte {
	var k, v []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Token"))
		c := b.Cursor()

		for k, v = c.First(); k != nil; k, v = c.Next() {
			if string(v) == Token {
				break
			}
		}
		return nil
	})
	return k
}

//DbGetToken 得到目标键值对
func DbGetToken(db *bolt.DB, Username string) []byte {
	var token []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Token"))

		token = b.Get([]byte(Username))
		return nil
	})
	return token
}

//DbDelToken 删除Token键值对
func DbDelToken(db *bolt.DB, Username string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Token"))
		err := b.Delete([]byte(Username))
		if err != nil {
			fmt.Println("Token of", string(Username), "was not existed.")
		}
		fmt.Println("Token of", string(Username), "was deleted.")
		return err
	})
}
