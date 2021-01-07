package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

//DbCreateBook 创建一个Book键值对
func DbCreateBook(db *bolt.DB, BookName string, Autor string, Info string) {
	var book Book
	book.BookName = BookName
	book.Autor = Autor
	book.Info = Info

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Book"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		id, _ := b.NextSequence()
		book.BookId = int32(id)

		buf, err := json.Marshal(book)

		if err != nil {
			return err
		}
		fmt.Println("创建Book: ", book)
		return b.Put(itob(book.BookId), buf)
	})

	if err != nil {
		fmt.Println("Book添加失败")
	}
}

//DbGetBook 得到一个Book对象
func DbGetBook(db *bolt.DB, BookID int32) Book {
	var book Book = Book{}

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))
		v := b.Get([]byte(itob(BookID)))
		if v == nil {
			fmt.Println("Book不存在")
		}
		err := json.Unmarshal(v, &book)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
	return book
}

//DbDelBook 删除一个Book
func DbDelBook(db *bolt.DB, BookID int32) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Book"))
		fmt.Println(string(b.Get([]byte(itob(BookID)))))
		err := b.Delete([]byte(itob(BookID)))
		return err
	})
}
