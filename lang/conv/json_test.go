package conv

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func Jso() {
	var dd = map[string]string{"ab": "ab1", "ac": "ac1", "ad": "ad1"}
	var aa = map[string]json.RawMessage{}
	marshal, err := json.Marshal(dd)
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}
	err = json.Unmarshal(marshal, &aa)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Println("marshal:", string(marshal))
	fmt.Println("aa:", aa)
	for k, v := range aa {
		fmt.Println("k:", k)
		fmt.Println("v:", string(v))
	}
}

func Jsov2() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	//marshal, err := json.Marshal(&ee)
	//if err != nil {
	//	fmt.Println("Marshal err:", err)
	//	return
	//}
	err := json.Unmarshal([]byte(ee), &ff)
	//err = json.Unmarshal(marshal, &ff)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}

func Jsov4() {
	var ee = "kkabcccc"
	var ff = json.RawMessage{}
	jsonIterator := jsoniter.ConfigCompatibleWithStandardLibrary
	// 序列化
	structJson, err := jsonIterator.Marshal(ee)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	// 打印输出结果
	fmt.Println("输出序列化结果: ", structJson, string(structJson))

	//fmt.Println("marshal:", marshal, string(marshal), ee)
	fmt.Println("ff:", ff)
	fmt.Println("ff:", string(ff))
}
