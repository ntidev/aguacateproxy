package main

import "github.com/google/tcpproxy"
import "io/ioutil"
import "fmt"
import "gopkg.in/yaml.v2"

type Endpoint struct {
	Name   string
	From   string
	To     string
	Type   string
	Domain string
}

type Config struct {
	Endpoints []Endpoint
}

func main() {

	configData, yamlErr := ioutil.ReadFile("/etc/aguacateproxy/config.yml")
	if yamlErr != nil {
		panic(yamlErr)
	}

	config := Config{}
	configErr := yaml.Unmarshal([]byte(configData), &config)
	if configErr != nil {
		panic(configErr)
	}

	var p tcpproxy.Proxy
	for _, endpoint := range config.Endpoints {
		fmt.Printf("Adding endpoint %s, From %s to %s\n", endpoint.Name, endpoint.From, endpoint.To)
		if endpoint.Type == "sni" {
			fmt.Printf("Domain for prebious endpoint: %s\n", endpoint.Domain)
			p.AddSNIRoute(endpoint.From, endpoint.Domain, tcpproxy.To(endpoint.To))
		} else {
			p.AddRoute(endpoint.From, tcpproxy.To(endpoint.To))
		}
	}

	fmt.Printf("Ready to accept connections\n")
	p.Run()
}
