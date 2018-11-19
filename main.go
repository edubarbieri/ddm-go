package main

import (
	"fmt"

	"github.com/edubarbieri/ddm/tvdb"
)

func main() {
	client := tvdb.NewClient()
	err := client.Login("436CB4A29DEF63C1")
	if err == nil {
		fmt.Printf("Login success %+v", client)
		resp, err := client.SearchSeries("The big bang theory")
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("%+v", resp)
		}

		resp1, err1 := client.GetEpisode(80379, 10, 10)
		if err1 != nil {
			fmt.Println(err1)
		} else {
			fmt.Printf("%+v", resp1)
		}
	} else {
		fmt.Println(err)
	}

}
