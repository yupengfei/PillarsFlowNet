package utility

func SliceResultToOutMessage(command * string, opResult interface{}, errorCode int, userCode * string) [] byte {
	sysError := &Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * OutMessage

	//TO DO 
	//the return result mayno right, be careful!
	//theorically, is opResult is nil, it should return "null"
	out = &(OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result:*(ObjectToJsonString(opResult)),
					})

	
	var result = ObjectToJsonByte(out)
	return result
}

func BoolResultToOutMessage(command * string, opResult interface{}, errorCode int, userCode * string) [] byte {
	sysError := &Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * OutMessage
	
	tempout := OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result: *ObjectToJsonString(opResult),
					}
		out = & tempout

	var result = ObjectToJsonByte(out)
	return result
}

func StringResultToOutMessage(command * string, opResult * string, errorCode int, userCode * string) [] byte {
	sysError := &Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * OutMessage
	
	tempout := OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result: *opResult,
					}
		out = & tempout

	var result = ObjectToJsonByte(out)
	return result
}