## config

    config loader options

    uses yaml 

### git

    go get "github.com/inkochetkov/config/pkg/config"

### loading the universal structure

loads the config as it is in the file, does not require a structure description

example

    conf, err := config.NewConfigAny("config.yaml")
    if err!= nil{
        ...
    }

    log.Println(conf)

### custom structure loading 

loads the config according to the structure, the structure is in the form of an example, how you can organize

example

    type Server struct {
	    Host string `yaml:"host"`
	    Port string `yaml:"port" default:"8080"`
    }


    conf, err := config.NewConfig("config.yaml", &Server{})
    if err!= nil{
        ...
    }

    log.Println(conf)

### License

Released under the MIT License