package storage
import (

	"PillarsFlowNet/utility"
	// "fmt"
)
func StoreToPost(post * utility.Post) (bool, error){
	
	err := PostCollection.Insert(post)
	if err != nil {
		return false, err
	}
	return true, err
}