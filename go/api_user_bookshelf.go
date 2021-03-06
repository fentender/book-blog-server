/*
 * Book Blog API
 *
 * This is a blog about books.
 *
 * API version: 0.0.7
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//AddBookInBookshelf 根据请求在数据库中创建Book数据
func AddBookInBookshelf(w http.ResponseWriter, r *http.Request) {
	var book BookshelfBookshelf

	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	str = strings.Replace(str, "/bookshelfs", "", -1)
	strs := strings.Split(str, "/")
	username := strs[0]
	bookshelfName := strs[1]
	bookID, _ := strconv.Atoi(strs[2])

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bookInfo := DbGetBook(db, int32(bookID))
	book.BookId = bookInfo.BookId
	book.BookName = bookInfo.BookName

	dbCreateBookshelfBook(db, username, bookshelfName, book)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(201)
}

//CreateBookshelf 根据请求在数据库中创建Bookshelf数据
func CreateBookshelf(w http.ResponseWriter, r *http.Request) {
	var bookshelfInfo BookshelfInfo

	json.NewDecoder(r.Body).Decode(&bookshelfInfo)
	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	username := strings.Replace(str, "/bookshelfs", "", -1)

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbCreateBookshelf(db, username, bookshelfInfo.Name)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(201)
}

//DeleteBookInBookshelf 根据请求删除数据库中对应数据
func DeleteBookInBookshelf(w http.ResponseWriter, r *http.Request) {
	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	str = strings.Replace(str, "/bookshelfs", "", -1)
	strs := strings.Split(str, "/")
	username := strs[0]
	bookshelfName := strs[1]
	bookID, _ := strconv.Atoi(strs[2])

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbDelBookshelfBook(db, username, bookshelfName, int32(bookID))

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(204)
}

//DeleteBookshelf 根据请求删除数据库中对应数据
func DeleteBookshelf(w http.ResponseWriter, r *http.Request) {
	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	str = strings.Replace(str, "/bookshelfs", "", -1)
	strs := strings.Split(str, "/")
	username := strs[0]
	bookshelfName := strs[1]

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dbDelBookshelf(db, username, bookshelfName)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(204)
}

//GetBookshelf 根据请求在数据库中查询Bookshelf数据并返回
func GetBookshelf(w http.ResponseWriter, r *http.Request) {
	var i, j, pageNumber int = 0, 0, 1
	var bookshelf Bookshelf
	var book BookshelfBookshelf

	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	str = strings.Replace(str, "/bookshelfs", "", -1)
	strs := strings.Split(str, "/")
	username := strs[0]
	bookshelfName := strs[1]

	var BucketName = "Users/" + username + "/Bookshelfs/" + bookshelfName
	if r.URL.Query()["pageNumber"] != nil {
		pageNumber, _ = strconv.Atoi(r.URL.Query()["pageNumber"][0])
	}
	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if ((pageNumber-1)*8 > i || i >= pageNumber*8) && pageNumber != -1 {
				i++
				continue
			}
			json.Unmarshal(v, &book)
			bookshelf.Bookshelf = append(bookshelf.Bookshelf, book)
			j++
			i++
		}
		return nil
	})

	var buf []byte

	bookshelf.Num = int32(i)
	buf, _ = json.Marshal(bookshelf)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

//GetBookshelfs 根据请求在数据库中查询Bookshelf数据并返回
func GetBookshelfs(w http.ResponseWriter, r *http.Request) {
	var i, j, pageNumber int = 0, 0, 1
	var bookshelfs Bookshelfs

	str := strings.Replace(r.URL.Path, "/users/", "", -1)
	username := strings.Replace(str, "/bookshelfs", "", -1)

	var BucketName = "Users/" + username + "/Bookshelfs"
	if r.URL.Query()["pageNumber"] != nil {
		pageNumber, _ = strconv.Atoi(r.URL.Query()["pageNumber"][0])
	}

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			return nil
		}
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if ((pageNumber-1)*8 > i || i >= pageNumber*8) && pageNumber != -1 {
				i++
				continue
			}
			bookshelfs.Bookshelfs = append(bookshelfs.Bookshelfs, BookshelfsBookshelfs{string(v)})
			j++
			i++
		}
		return nil
	})

	var buf []byte

	bookshelfs.Num = int32(i)
	buf, _ = json.Marshal(bookshelfs)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
