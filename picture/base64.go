package picture
import (
	"encoding/base64"
)

const (  
    base64Table = "143QRSTUabcdVqXYZHijKLAWDCABDstEFGuvwxyzGHIJklmnopqr234560178912"  
)

coder := base64.NewEncoding(base64Table)

func Base64Encode(src * []byte) * []byte {  
	var encoded * []byte
    return []byte(coder.EncodeToString(src))  
}  
  
func Base64Decode(src []byte) ([]byte, error) {  
    return coder.DecodeString(string(src))  
}   