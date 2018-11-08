package main

import (
	"fmt"

	"github.com/edubarbieri/ddm/data"
)

func main() {
	// serie := data.Serie{Name: "Teste", SearchKey: "teste", TvdbID: 123}
	// data.SaveSerie(&serie)
	// fmt.Println(serie)

	series := data.ListAllSeries()

	for _, s := range series {
		fmt.Println(s)
	}

	// serie, err := data.GetSerieBySearckKey("teste")
	// serie.Name = "Eduardo"
	// serie.SearchKey = "eduardo"
	// data.SaveSerie(&serie)
	// fmt.Println(serie.Name, err)

}
