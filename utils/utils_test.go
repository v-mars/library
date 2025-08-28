package utils

import (
	"fmt"
	"testing"
)

func TestObj2Json(t *testing.T) {
	aa := struct {
		Aa string                 `json:"aa"`
		Bb string                 `json:"bb"`
		Cc map[string]interface{} `json:"cc"`
		Dd []string               `json:"dd"`
	}{Aa: "a1", Bb: "b1", Cc: map[string]interface{}{"cm": "msg", "cdd": 1}, Dd: []string{"d1", "d2"}}
	fmt.Println(Any2Json(aa))
	fmt.Println(Any2Yaml(aa))
}
