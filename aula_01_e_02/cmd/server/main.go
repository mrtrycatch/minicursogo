package main

import "github.com/mrtrycatch/minicursogo/configs"

func main() {
	conf, _ := configs.LoadConf(".")
	println(conf.DatabaseName)
}
