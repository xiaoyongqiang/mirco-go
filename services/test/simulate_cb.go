package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"noruserverV4/banks/xmpay/params"
	"noruserverV4/utils"
	"time"
)

var (
	outTradeNo  = flag.String("otn", "", "out_trade_no")
	totalFee    = flag.String("tf", "", "total_fee")
	orderStatus = flag.String("os", "0000", "orderstatus")
	memo        = flag.String("memo", "", "memo")
	notifyURL   = flag.String("url", "https://api.ctvrtv.com/gateway/trade/notify/xmpay", "notify url")
)

func main() {
	flag.Parse()

	if *outTradeNo == "" ||
		*totalFee == "" ||
		*orderStatus == "" ||
		*memo == "" ||
		*notifyURL == "" {
		flag.Usage()
		return
	}

	nr := &params.NotifyReq{
		Success: true,
		Data: params.NotifyDataRep{
			TotalFee:    *totalFee,
			OutTradeNO:  *outTradeNo,
			OrderStatus: *orderStatus,
			Memo:        *memo,
			TimeStamp:   fmt.Sprintf("%d", time.Now().Unix()),
			NonceStr:    utils.RandStr(8),
		},
	}

	//nr := &params.NotifyReq{
	//	Success: true,
	//	Data: params.NotifyDataRep{
	//		TotalFee:    "1.00",
	//		OutTradeNO:  "Yf20170905095843210054914U",
	//		OrderStatus: "0000",
	//		Memo:        "GC201709050959298458:xm-alipay-01",
	//		TimeStamp:   "1504577190",
	//		NonceStr:    "392075515",
	//	},
	//}

	sign, _ := utils.Md5(nr.Data, "rrkx68Uc_23iinndoU", nil)
	//log.Printf("sign:%s\n", sign)
	//return
	nr.Data.Sign = sign
	b, _ := json.Marshal(nr)
	req, err := http.NewRequest("POST", *notifyURL, bytes.NewBuffer(b))
	rep, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Do err:%v\n", err)
	}
	defer rep.Body.Close()

	data, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		log.Printf("ReadAll, Err:%v\n", err)
	}

	log.Printf("body:%s\n", data)

}
