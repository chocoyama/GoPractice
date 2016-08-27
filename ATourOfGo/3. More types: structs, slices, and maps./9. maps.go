package main

import (
    "fmt"
)

type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

// リテラル
var m2 = map[string]Vertex{
    "Bell Labs": Vertex{
        4069433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}

// リテラル(型省略)
var m3 = map[string]Vertex{
    "Bell Labs": {4069433, -74.39967,},
    "Google": {37.42202, -122.08408,},
}

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        4069433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
    fmt.Println(m2)
    fmt.Println(m3)
}
