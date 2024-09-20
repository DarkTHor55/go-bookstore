package utils
import (
	"encoding/json"
	"net/http"
	"io"
)
func ParseBody(r *http.Request ,x interface{}){
	if body ,err:=io.ReadAll(r.body); err == nil{
		if err := json.Inmarshal([]byte(body),x); err != nil{
			return 
		}
	}
}