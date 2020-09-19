package access_token_utils

import (
	"net/http"
	"strings"
)

// ExtractToken will extract the token from the Request
func ExtractToken(r *http.Request) *string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return &strArr[1]
	}
	return nil
}
