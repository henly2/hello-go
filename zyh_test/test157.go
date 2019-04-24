package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var cfj string

func main() {

	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333", "cell": "415.555.5555"}
	data, err := json.MarshalIndent(c, "", "      ") //这里返回的data值，类型是[]byte
	if err != nil {
		log.Println("ERROR:", err)
	}
	fmt.Println(string(data))
	WriteConfig("/host.json", data)
}
func WriteConfig(f string, data interface{}) {

	cfj = " ./host.json"
	if cfj == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	_, err := f.WriteBytes(cfj, jsonByte)
	if err != nil {
		log.Fatalln("write config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	log.Println("write config file:", cfj, "successfully")

}
