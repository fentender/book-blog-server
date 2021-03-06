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
	"regexp"
	"strconv"
)

//GetReview 根据请求读取review，并返回
func GetReview(w http.ResponseWriter, r *http.Request) {
	reg := regexp.MustCompile(`[0-9]+`)
	str := reg.FindAllString(r.URL.Path, -1)
	bookID, _ := strconv.Atoi(str[0])
	reviewID, _ := strconv.Atoi(str[1])

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	review := DbGetReview(db, int32(bookID), int32(reviewID))
	buf, _ := json.Marshal(review)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}

//GetReviews 根据请求读取review[]，并返回
func GetReviews(w http.ResponseWriter, r *http.Request) {
	var i, j, pageNumber int = 0, 0, 1
	var review []Review
	reg := regexp.MustCompile(`[0-9]+`)
	BookID, _ := strconv.Atoi(reg.FindString(r.URL.Path))
	var BucketName string = "Reviews/" + string(BookID)
	if r.URL.Query()["pageNumber"] != nil {
		pageNumber, _ = strconv.Atoi(r.URL.Query()["pageNumber"][0])
	}

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		var objReview Review
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
			json.Unmarshal(v, &objReview)
			review = append(review, objReview)
			j++
			i++
		}
		return nil
	})

	var buf []byte
	if pageNumber != -1 {
		buf, _ = json.Marshal(Reviews{int32(i), review})
	} else {
		buf, _ = json.Marshal(Reviews{int32(i), review})
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
