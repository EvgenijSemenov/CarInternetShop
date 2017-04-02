package response

func BuildResponse(result map[string]interface{}, err error) map[string]interface{} {
	response := make(map[string]interface{})
	if err != nil {
		response = FailureResponse(err.Error())
	} else {
		response = SuccessResponse(result)
	}
	return response
}

func FailureResponse(failureMessage string) map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = "FAILURE"
	response["message"] = failureMessage
	return response
}

func SuccessResponse(result map[string]interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	response["status"] = "SUCCESS"
	response["result"] = result
	return response
}
