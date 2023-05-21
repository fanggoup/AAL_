package main

import (
	"AAL_time/conf"
	"AAL_time/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_=r.Run(":"+conf.HttpPort)
}