package swagger

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

//DbCreateReview 创建一个Review键值对
func DbCreateReview(db *bolt.DB, BookID int32, Content string, Autor string) {
	var Review Review
	var BucketName string = "Reviews/" + string(BookID)
	Review.Content = Content
	Review.Autor = Autor

	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		id, _ := b.NextSequence()
		Review.ID = int32(id)

		buf, err := json.Marshal(Review)
		if err != nil {
			return err
		}
		fmt.Println("创建Review: ", Review)
		return b.Put(itob(Review.ID), buf)
	})

	if err != nil {
		fmt.Println("Review添加失败")
	}
}

//DbGetReview 得到一个Review对象
func DbGetReview(db *bolt.DB, BookID int32, ReviewID int32) Review {
	var review Review = Review{}
	var BucketName string = "Reviews/" + string(BookID)

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		v := b.Get([]byte(itob(ReviewID)))
		if v == nil {
			fmt.Println("Review不存在")
		}
		err := json.Unmarshal(v, &review)
		if err != nil {
			fmt.Println(err)
		}

		return nil
	})
	return review
}

//DbDelReview 删除一个Review
func DbDelReview(db *bolt.DB, BookID int32, ReviewID int32) {
	var BucketName string = "Reviews/" + string(BookID)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		fmt.Println(string(b.Get([]byte(itob(ReviewID)))))
		err := b.Delete([]byte(itob(ReviewID)))
		return err
	})
}
