package main

import (
	"embed"
	"fmt"

	"github.com/Yiling-J/piper"
	"github.com/Yiling-J/piper/example/config"
)

//go:embed config/*
var configFS embed.FS

func main() {
	err := piper.Load("config/dev.toml")
	if err != nil {
		panic(err)
	}
	// should be spf14
	fmt.Println(piper.GetString(config.Params.Githubuser))
	// should be true
	fmt.Println(piper.GetBool(config.Params.Debug))
	// should be 5000
	fmt.Println(piper.GetInt(config.Caches.Assets.Maxage))
	// should be fallback
	fmt.Println(piper.GetString(config.Build.Useresourcecachewhen))
	// should be []string{"foo1", "foo2"}
	fmt.Println(piper.GetStringSlice(config.Params.Listoffoo))

	piper.Reset()
	err = piper.Load("config/prod.toml")
	if err != nil {
		panic(err)
	}
	// should be spf15
	fmt.Println(piper.GetString(config.Params.Githubuser))
	// should be false
	fmt.Println(piper.GetBool(config.Params.Debug))
	// should be 3000
	fmt.Println(piper.GetInt(config.Caches.Assets.Maxage))
	// should be fallback
	fmt.Println(piper.GetString(config.Build.Useresourcecachewhen))
	// should be []string{"foo1", "foo2"}
	fmt.Println(piper.GetStringSlice(config.Params.Listoffoo))
}
