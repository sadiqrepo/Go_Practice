package main

import(
	"fmt"
)

type dtx struct{
	Lat, Long float64
}

var r map[string]dtx

var n = map[string]dtx{
	"Atari" : dtx{
		21.1, 32.3,
	},
	"Nintendo" : dtx{
		30.5, 34.5,
	},
}

var t = map[string]dtx{
	"Google" : {21.7, 32.7},
	"Oracle" : {10.7, 34.7},
}

func main(){

	r = make(map[string]dtx)
	r["Bell Labs"] = dtx{
		22.0, 11.0,
	}
	r["Xerox"] = dtx{
		34.0, 65.0,
	}

	fmt.Println(r)
	fmt.Println(n)
	fmt.Println(t)
}