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
	"strings"
)

//GetUser 根据请求从数据库中读取User信息，并返回
func GetUser(w http.ResponseWriter, r *http.Request) {
	username := strings.Replace(r.URL.Path, "/users/", "", -1)

	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user := DbGetUser(db, username)
	buf, _ := json.Marshal(user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
	w.WriteHeader(http.StatusOK)
	w.Write(buf)
}
