package swagger

import (
	"github.com/boltdb/bolt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"log"
	"net/http"
)

func authorizedValid(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, errors := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		})
	db, err := bolt.Open("./database/bookblog.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	if errors == nil && token.Valid {
		tokenString, _ := token.SignedString([]byte(SecretKey))
		if DbKeyofToken(db, tokenString) != nil {
			db.Close()
			log.Println("Token验证成功")
			next(rw, r)
		} else {
			db.Close()
			log.Println("Token验证失败")
			rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
			rw.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
			rw.WriteHeader(http.StatusUnauthorized)
			rw.Write([]byte("Token is not valid"))
		}
	} else {
		db.Close()
		log.Println("Token验证失败")
		rw.Header().Set("Content-Type", "application/json; charset=UTF-8")
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type")
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Unauthorized access to this resource"))
	}
}
