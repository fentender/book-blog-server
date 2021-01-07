package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

//dbCreateBookshelf 创建一个Bookshelf键值对
func dbCreateBookshelf(db *bolt.DB, Username string, BookshelfName string) {
	var BucketName string = "Users/" + Username + "/Bookshelfs"

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

		fmt.Println("创建Bookshelf: ", BucketName, BookshelfName)

		return b.Put([]byte(BookshelfName), []byte(BookshelfName))
	})

	if err != nil {
		fmt.Println("Bookshelf添加失败")
	}
}

//dbCreateBookshelfBook 创建一个Book of Bookshelf键值对
func dbCreateBookshelfBook(db *bolt.DB, Username string, BookshelfName string, Book BookshelfBookshelf) {
	var BucketName = "Users/" + Username + "/Bookshelfs/" + BookshelfName

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}

		buf, err := json.Marshal(Book)

		if err != nil {
			return err
		}

		fmt.Println("创建Book of Bookshelf: ", Book)

		return b.Put(itob(Book.BookId), buf)
	})

	if err != nil {
		fmt.Println("Book of Bookshelf添加失败")
	}
}

//dbDelBookshelf 删除一个Bookshelf
func dbDelBookshelf(db *bolt.DB, Username string, BookshelfName string) {
	var BucketName string = "Users/" + Username + "/Bookshelfs"

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		fmt.Println(string(b.Get([]byte(BookshelfName))))

		BucketName = "Users/" + Username + "/Bookshelfs/" + BookshelfName
		tx.DeleteBucket([]byte(BucketName))
		err := b.Delete([]byte(BookshelfName))
		return err
	})
}

//dbDelBookshelfBook 删除一个Bookshelf中的Book
func dbDelBookshelfBook(db *bolt.DB, Username string, BookshelfName string, BookID int32) {
	var BucketName = "Users/" + Username + "/Bookshelfs/" + BookshelfName

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		fmt.Println(string(b.Get(itob(BookID))))

		err := b.Delete(itob(BookID))
		return err
	})
}
