package config

import (
    "log"
    "path/filepath"
    "github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
    var err error
    v := viper.New()
    v.SetConfigType("yaml")
    v.SetConfigName(env)
    if env == "test" {
        v.AddConfigPath("../config/")
    } else {
        v.AddConfigPath("config/")
    }
    err = v.ReadInConfig()

    if err != nil {
        log.Fatal("error on parsing configration file")
    }
    config = v
}

func relativePath(basedir string, path *string) {
    p := *path
    if p != "" && p[0] != "/" {
        *path = filepath.Join(basedir, p)
    }
}

func GetConfig() *viper.Viper {
    return config
}
