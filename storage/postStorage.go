package storage
import (

	"PillarsFlowNet/utility"
	"labix.org/v2/mgo/bson"
)
func StoreToPost(post * utility.Post) (* utility.Post, error){
	
	err := PostCollection.Insert(post)
	if err != nil {
		return post, err
	}
	PostCollection.Find(bson.M{"postcode":post.PostCode}).One(post)
	return post, err
}

func GetAllPostByTargetCode(targetCode * string) ([] utility.Post, error) {
	var postSlice [] utility.Post
	iter := PostCollection.Find(bson.M{"targetcode":*targetCode}).Iter()
	err := iter.All(&postSlice)
	if err != nil {
		return postSlice, err
	}
	return postSlice, err
}