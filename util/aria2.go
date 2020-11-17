package util

import (
	"context"
	"github.com/saltbo/zpan/config"
	"github.com/zyxar/argo/rpc"
	"log"
	"strconv"
	"time"
)

var Aria2Client rpc.Protocol

func StartAria2Rpc(conf *config.Config)  {
	var err error
	Aria2Client, err = rpc.New(context.Background(), "http://localhost:" + strconv.Itoa(conf.Aria2.Port), conf.Aria2.RPCSecret, 7*time.Second, nil)
	if err != nil {
		log.Printf("Start aria2 rpc failed [%s].", err)
	}
	log.Print("Connected to aria2.")
}