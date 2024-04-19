package main

import (
	"datacontrol/modHttp"
	"datacontrol/modUtility"
	"fmt"
)

func main() {
	err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = modHttp.Http_Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = modHttp.Http_Start()
	if err != nil {
		fmt.Println(err)
		return
	}

}
