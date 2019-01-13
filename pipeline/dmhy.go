package pipeline

import (
	bytes2 "bytes"
	"net/http"
	"os"
	"strings"
)

func Pineline(item []string) {
	client := &http.Client{}
	for _, str := range item {
		rawJson := "{\"jsonrpc\":\"2.0\",\"method\":\"aria2.addUri\",\"id\":\"1\",\"params\":[[\"{{magnet}}\"],{}]}"
		postData := strings.Replace(rawJson, "{{magnet}}", str, -1)
		println(postData)
		req, _ := http.NewRequest("POST", os.Args[2], bytes2.NewBuffer([]byte(postData)))
		_, _ = client.Do(req)
	}
}
