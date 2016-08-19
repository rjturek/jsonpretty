package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println(jsonPrettyPrint(string(bytes)))
}
