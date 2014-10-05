package post

import (
	"time"
	"PillarsFlowNet/storage"
	"PillarsFlowNet/utility"
	"PillarsFlowNet/authentication"
	// "fmt"
	"strings"
)
// type Post struct {
//     PostCode string
//     TargetCode string
//     Message string
//     ReplyTo string
//     UserCode string
//     PostTime string
//     Deleted int
//     DeletedTime string
// }
func Post(userCodeAndParameter * string) ([] byte, *string) {
	inputParameters := strings.SplitN(*userCodeAndParameter, "@", 2)
	auth := authentication.GetAuthInformation(&(inputParameters[0]))
	var errorCode int
	if (auth == false) {
		errorCode = 3
	}
	var post * utility.Post
	//var toUserCode * string
	if (errorCode == 0) {
		post, _ := utility.ParsePostMessage(&(inputParameters[1]))
		//toUserCode = &(chart.To)

		post.PostCode = *(utility.GenerateCode(&inputParameters[0]))
		post.UserCode = inputParameters[0]
		
		post.Deleted = 0
		post.DeletedTime = time.Now().Format("2006-01-02 15:04:05")
		storage.StoreToPost(post)
	}
	return utility.ObjectToJsonByte(post), nil
}