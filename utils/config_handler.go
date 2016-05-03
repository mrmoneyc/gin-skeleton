package utils

import (
    "os"
    "io/ioutil"
    "github.com/naoina/toml"
)

type tomlConfig struct {
    App struct {
        Name string
        Version string
    }
    
    DB struct {
        SQLite struct {
            DBFile string
        }
        MySQL struct {
            Host string
            DBName string
            UserName string
            Password string
        }
        Driver string
    }
}

func ReadConfig() (tomlConfig, error) {
    f, err := os.Open("config.toml")
    defer f.Close()
    
    buf, err := ioutil.ReadAll(f)
    
    var config tomlConfig
    err = toml.Unmarshal(buf, &config)
    
    return config, err
}