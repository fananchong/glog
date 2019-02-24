package main

import (
	"flag"
	"time"

	"github.com/fananchong/glog"
)

func main() {
	flag.Parse()
	glog.GetLogger().SetLogDir(".")
	glog.GetLogger().SetLogLevel(0)
	glog.GetLogger().SetFlushInterval(500 * time.Millisecond)
	glog.GetLogger().Infoln("aaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	glog.GetLogger().Infoln("bbbbbbbbbbbbbbbbbb")
}
