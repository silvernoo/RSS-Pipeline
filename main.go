package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/silvernoo/RSS-Pipeline/pipeline"
	"github.com/silvernoo/RSS-Pipeline/rss"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
)

var node []string

func main() {
	if len(os.Args) < 3 {
		println("demo : ./rss 3600 http://127.0.0.1:6800/jsonrpc")
	}
	sec, _ := strconv.Atoi(os.Args[1])
	ticker := time.NewTicker(time.Second * time.Duration(sec))
	file, _ := os.OpenFile("data", os.O_CREATE|os.O_APPEND|os.O_RDONLY, 0600)
	open, _ := os.Open("rss.json")
	ReadFile(file)
	jsonRaw, _ := ioutil.ReadAll(open)
	_ = file.Close()
	for range ticker.C {
		var arr []string
		_ = json.Unmarshal(jsonRaw, &arr)
		for _, item := range arr {
			pipeline.Pineline(seen(rss.GetRssItem(item)))
		}
	}
}

func ReadFile(file *os.File) {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if len(line) > 1 {
			node = append(node, line[:len(line)-1])
		}
		if err != nil {
			break
		}
	}
}

func seen(item []string) []string {
	return CheckChangeItem(item)
}

func CheckChangeItem(item []string) []string {
	var result []string
	for _, i := range item {
		hash := GetMD5Hash(i)
		if !SliceExists(node, hash) {
			node = append(node, hash)
			file, _ := os.OpenFile("data", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
			_, _ = file.WriteString(hash)
			_, _ = file.WriteString("\n")
			_ = file.Close()
			result = append(result, i)
		}
	}
	return result
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}
	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}
	return false
}
