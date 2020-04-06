// controllers
// Handles requests made to the server
package controllers

import (
	"encoding/json"
	"github.com/miles1618/udemy/rest/mvc/services"
	"github.com/miles1618/udemy/rest/mvc/utils"
	"net/http"
	"strconv"
)

// GetUser
// Queries datastore for a user given a user_id
func GetUser(res http.ResponseWriter, req *http.Request){
	// Parse any params
	paramUserId := req.URL.Query().Get("user_id")
	// Convert params to expected types
	userId, convErr := strconv.ParseInt(paramUserId, 10, 64)
	// No user_id in params
	if convErr != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request" ,
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}
	// No user found in data source
	user, getUserErr := services.GetUser(userId)
	if getUserErr != nil {
		apiErr := &utils.ApplicationError{
			Message: "user not found",
			StatusCode: http.StatusNotFound,
			Code: "not_found",
		}
		jsonValue, _ := json.Marshal(apiErr)
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	// User found - marshall for response
	jsonValue, jsonErr := json.Marshal(user)
	if jsonErr != nil{
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(jsonErr.Error()))
		return
	}
	res.Header().Set("Content-Type", "application/json")
	_, writeErr := res.Write(jsonValue)
	if writeErr != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(writeErr.Error()))
		return
	}
}
