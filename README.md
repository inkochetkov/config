## config

    config loader options

### git

    go get "github.com/inkochetkov/config/pkg/config"

### loading the universal structure

loads the config as it is in the file, does not require a structure description

example

    conf, err := config.NewConfigAny("config.yaml")
    if err!= nil{

    }

    log.Println(conf)

### custom structure loading 

loads the config according to the structure, the structure is in the form of an example, how you can organize

example

    conf, err := config.NewConfig("config.yaml")
    if err!= nil{

    }

    log.Println(conf)