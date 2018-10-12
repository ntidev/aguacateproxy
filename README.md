AguacateProxy
=============

![Image of AguacateProxy](https://github.com/ntidev/logo.png)

This is a very simple proxy used to be able to forward an address/port from one point to another.

It is not intented to be used (for now) in production as it doesn't care (at all) about any type of security and/or dynamic configuration.

This proxy is entirely based on the library [tcpproxy](https://godoc.org/github.com/google/tcpproxy) from GoLang.


## Installation

The simplest way to install this is with Docker.

    ```    
    $ docker run -p "8080:8080" -v "${PWD}/config.yml:/go/src/app/config.yml" -it ntidev/aguacateproxy
    ```

Unless you make the container run on the host's network, you will have to map all the ports that you will use on the configuration file.

## Configuration Sample

The configuration is based on a simple file called `config.yml`

    ```
    endpoints: 
    - { name: "aguacateproxy_test", from: ":8080", to: "1.2.3.4:80" }
    - { name: "aguacateproxy_test_sni", from: ":8080", to: "1.2.3.4:80", type: "sni", domain: "aguacateproxy.xyz" }  

    ```

The configuration consists of an array of `endpoints` which has the following structure:

* *name*: A descriptive name for the route *(Required)*
* *from*: The `From` piece of the redirection  *(Required)*
* *to*: The `To` piece of the redirection  *(Required)*
* *type*: As of right now the only available type is `sni`, this parameter is optional
* *domain*: In case the type `sni` is used, this parameter contains the domain piece for the SNI route


## Contributors

* Benjamín Visón [GitHub](http://github.com/bvisonl)