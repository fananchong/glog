package main

import (
	"flag"

	"github.com/fananchong/glog"
)

func main() {
	flag.Parse()
	glog.GetLogger().SetLogDir("./logs")
	glog.GetLogger().SetLogLevel(0)
	glog.GetLogger().Infoln("aaaaaaaaaaaaaaaaaaaaaaaaaaaa")
}
