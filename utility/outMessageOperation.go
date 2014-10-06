package utility

func SliceResultToOutMessage(command * string, opResult interface{}, errorCode int, userCode * string) [] byte {
	sysError := &Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * OutMessage
	if opResult == nil {
		out = &(OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result:"{}",
					})
	} else {
		out = &(OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result:*(ObjectToJsonString(opResult)),
					})
	}
	
	var result = ObjectToJsonByte(out)
	return result
}

func BoolResultToOutMessage(command * string, opResult interface{}, errorCode int, userCode * string) [] byte {
	sysError := &Error {
						ErrorCode: errorCode,
						ErrorMessage: "",
					}
	var out * OutMessage
	if errorCode != 0 {
		tempout := OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result: "{}",
					}
		out = & tempout
	} else {
		tempout := OutMessage {
						Error: *sysError,
						Command: *command,
						UserCode: *userCode,
						Result: *ObjectToJsonString(opResult),
					}
		out = & tempout
	}

	var result = ObjectToJsonByte(out)
	return result
}