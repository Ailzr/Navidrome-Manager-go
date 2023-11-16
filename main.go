package main

import (
	_ "MusicManager/conf"
	"MusicManager/router"
)

func main() {
	router.SetRouter()
}
