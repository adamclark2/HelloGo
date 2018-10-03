package main

import(
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port string
	Name string
}

var cfg *Config;

func GetConfig() *Config{
	if(cfg != nil){
		return cfg;
	}

	dat, err := ioutil.ReadFile("config.json")
	if(err != nil){
		panic(err)
	}
	if(dat == nil){
		fmt.Printf("data from config.json is nil");
		panic(dat);
	}

	json.Unmarshal(dat, &cfg);
	return cfg
}

/// Test by printf
func Test_File() {
	c := GetConfig();
	if(c == nil){
		fmt.Printf("Config is null\n");
		panic(c)
	}
	fmt.Printf("%#v\n", c);

	fmt.Printf("\n\nFile Test Success!\n");
}