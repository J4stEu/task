package response

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Response(w http.ResponseWriter, statusCode int, responseData any) {
	var out []byte
	var err error
	switch data := responseData.(type) {
	case []byte:
		out = data
	default:
		out, err = json.Marshal(responseData)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	w.WriteHeader(statusCode)
	if _, err = w.Write(out); err != nil {
		fmt.Println(err)
	}
}
