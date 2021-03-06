package consul

import (
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"log"
)

func main() {
	viper.AddRemoteProvider("consul", "localhost:8500", "bastion/cobank/chains.wc")
	viper.SetConfigType("json") // Need to explicitly set this to json
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(viper.Get("name"))
	fmt.Println(viper.Get("network.url"))
}
