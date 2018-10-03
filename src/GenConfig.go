package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

func main() {
	con := Config{};
	fmt.Printf("test:\n----------\n%#v\n\n", con);

	b, err := json.MarshalIndent(con, "", "   ");
	if(err != nil){
		panic(err)
	}

	ioutil.WriteFile("config.json", b, 0700)

	fmt.Printf("Testing File:----\n")
	Test_File();
}
