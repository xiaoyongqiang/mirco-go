package main

import (
	"fmt"
	"log"
	"net/http"
	//"github.com/astaxie/session"
	//_ "github.com/astaxie/session/providers/memory"
	"swingbaby-go/apis/ctrls"
	"swingbaby-go/config"

	"github.com/julienschmidt/httprouter"
	"github.com/koding/multiconfig"
	"github.com/pilu/xrequestid"
	"github.com/urfave/negroni"
)

// var globalSessions *session.Manager

// //然后在init函数中初始化
// func init() {
// 	 globalSessions, _ = session.NewManager("memory", "gosessionid",
// }
// m.MustLoad(config.Config)

// err = config.InitializeConn()
// if err != nil {3600)
// 	 go globalSessions.GC()
// }

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
	router.POST("/wechat/auth", ctrls.NewWechat)
	router.POST("/class/create", ctrls.NewClass)
	router.POST("/class/editmobil", ctrls.ClassEditmobil)
	router.POST("/class/editname", ctrls.ClassEditname)
	router.POST("/class/tremove", ctrls.ClassTremove)
	router.POST("/class/join", ctrls.ClassJoin)
	router.POST("/class/apply", ctrls.ClassApply)
	router.POST("/class/premove", ctrls.ClassPremove)
	router.GET("/class/mlist", ctrls.ClassMlist)
	router.GET("/class/plist", ctrls.ClassPlist)
	router.GET("/class/tlist", ctrls.ClassTlist)
	router.POST("/photo/create", ctrls.NewAlbum)
	router.POST("/photo/add", ctrls.NewPhoto)
	router.POST("/photo/addimg", ctrls.PhotoAddImg)
	n := negroni.New(negroni.NewRecovery())
	n.Use(xrequestid.New(16))
	n.Use(negroni.HandlerFunc(ctrls.ControlBases))
	n.UseHandler(router)
	n.Run(fmt.Sprintf("%s:%d", config.Config.ApiConf.Host, config.Config.ApiConf.Port))
}
