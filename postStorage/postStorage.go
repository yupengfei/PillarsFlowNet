package postStorage

import (
	"PillarsFlowNet/mongoUtility"
	"PillarsFlowNet/utility"
	"labix.org/v2/mgo/bson"
)
func StoreToPost(post * utility.Post) (* utility.Post, error){
	
	err := mongoUtility.PostCollection.Insert(post)
	if err != nil {
		return post, err
	}
	mongoUtility.PostCollection.Find(bson.M{"postcode":post.Id}).One(post)
	return post, err
}

func GetAllPostByTargetCode(targetCode * string) ([] utility.Post, error) {
	var postSlice [] utility.Post
	iter := mongoUtility.PostCollection.Find(bson.M{"code":*targetCode, "posttype":2}).Iter()
	err := iter.All(&postSlice)
	if err != nil {
		return postSlice, err
	}
	return postSlice, err
}