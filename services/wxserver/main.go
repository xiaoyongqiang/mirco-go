package main

import (
	"fmt"
	"log"
	"net/http"
	"swingbaby-go/apis/ctrls"
	"swingbaby-go/config"

	"github.com/julienschmidt/httprouter"
	"github.com/koding/multiconfig"
	"github.com/pilu/xrequestid"
	"github.com/urfave/negroni"
)

func main() {
	var err error
	m := multiconfig.New()
	config.Config = new(config.CmdConfig)
	err = m.Load(config.Config)
	if err != nil {
		log.Fatalf("Load configuration failed. Error: %s\n", err.Error())
		log.Fatalf("config.InitialzeConn() failed. Error info: %s\n", err.Error())
	}
	defer func() {
		config.DBHandle.Close()
		config.RedisHandle.Close()
	}()

	router := httprouter.New()
	router.HandleMethodNotAllowed = false
	router.ServeFiles("/images/*filepath", http.Dir("./images"))
	router.POST("/class/create", ctrls.NewClass)
	router.GET("/class/mlist", ctrls.ClassMlist)
	router.GET("/wechat/oauthcb", ctrls.Callback)
	n := negroni.New(negroni.NewRecovery())
	n.Use(xrequestid.New(16))
	n.Use(negroni.HandlerFunc(ctrls.OauthBases))
	n.UseHandler(router)
	n.Run(fmt.Sprintf("%s:%d", config.Config.ApiConf.Host, config.Config.ApiConf.Port))
}
