package main

import "github.com/google/tcpproxy"
import "io/ioutil"
import "fmt"
import "gopkg.in/yaml.v2"

type Endpoint struct {
	Name 	string
	From 	string
	To	 	string
	Type 	string
	Domain 	string
}

type Config struct {
	Endpoints		[]Endpoint
}

func main() {
	
	configData, yamlErr := ioutil.ReadFile("./config.yml")
	if(yamlErr != nil) {
		panic(yamlErr)
	}

	config := Config{}
	configErr := yaml.Unmarshal([]byte(configData), &config)
	if configErr != nil {
		panic(configErr)
	}

	var p tcpproxy.Proxy	
	for i := 0; i < len(config.Endpoints); i++ {
		fmt.Printf("Adding endpoint " + config.Endpoints[i].Name + ", From " + config.Endpoints[i].From + " to " + config.Endpoints[i].To + "\n")
		if(config.Endpoints[i].Type == "sni") {
			fmt.Printf("Domain for prebious endpoint: " + config.Endpoints[i].Domain + "\n")
			p.AddSNIRoute(config.Endpoints[i].From, config.Endpoints[i].Domain, tcpproxy.To(config.Endpoints[i].To))
		} else {
			p.AddRoute(config.Endpoints[i].From, tcpproxy.To(config.Endpoints[i].To))
		}
		
	}

	fmt.Printf("Ready to accept connections" + "\n");
	p.Run()
}
