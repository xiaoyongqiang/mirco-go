package merpaytype

import (
	"errors"
	"fmt"
	"math/rand"
	"noruserverV4/config"
	"noruserverV4/models"
	"time"

	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) GetConf(ctx context.Context, req *PayConfReq) (*PayConfResp, error) {
	mer := &models.Merchant{
		ID: req.MID,
	}

	assignPayType, err := mer.GetAssignPayType(req.PayType)
	if err != nil {
		return nil, err
	}

	resp := &PayConfResp{
		Data: &PayConf{},
	}

	pcModel := &models.PayChannel{}
	if assignPayType != nil {
		cID, ok := assignPayType["ccID"].(int64)
		if ok {
			pcModel.CCID = cID
		}

		cbName, ok := assignPayType["cbName"].(string)
		if ok {
			pcModel.CBName = cbName
		}

		cTypeName, ok := assignPayType["cTypeName"].(string)
		if ok {
			pcModel.CCTypeName = cTypeName
		}

		err := pcModel.GetFeeAndSettleTypeByID()
		if err != nil {
			return nil, err
		}
		resp.Data.CName = cbName
		resp.Data.CCID = cID
		resp.Data.CFee = pcModel.CCFee
		resp.Data.CTypeName = pcModel.CCTypeName
		resp.Data.SettleType = pcModel.CCSettleType
		return resp, nil
	}

	payInfos, err := pcModel.GetTypeInfoByPayType(req.PayType)
	if err != nil {
		return nil, err
	}

	if len(payInfos) == 0 {
		return nil, errors.New("Service Shutdown")
	}

	unuseListKey := fmt.Sprintf("unuse:%d:%d", req.MID, req.PayType)
	ids := config.RedisHandle.SMembers(unuseListKey).Val()

	if len(ids) > 0 {
		for _, id := range ids {
			if _, ok := payInfos[id]; ok == true {
				delete(payInfos, id)
			}
		}
	}

	var ranks int64
	var rankIds = make(map[int]int64)
	var rankSlice []int64
	idIndex := 0
	for _, v := range payInfos {
		rak, _ := v["rank"].(int64)
		ranks += rak
		id, _ := v["id"].(int64)
		rankIds[idIndex] = id
		rankSlice = append(rankSlice, rak)
		idIndex++
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	rd := rnd.Int63n(ranks)
	var rankScore int64 = 0
	cho := 0
	for k, v := range rankSlice {
		rankScore = v + rankScore
		if rd <= rankScore {
			cho = k
			break
		}
	}
	choStr := fmt.Sprintf("%d", rankIds[cho])
	name, _ := payInfos[choStr]["name"].(string)
	resp.Data.CName = name
	id, _ := payInfos[choStr]["id"].(int64)
	resp.Data.CCID = id
	tName, _ := payInfos[choStr]["tName"].(string)
	resp.Data.CTypeName = tName
	fee, _ := payInfos[choStr]["fee"].(int64)
	resp.Data.CFee = fee
	settleType, _ := payInfos[choStr]["settleType"].(int64)
	resp.Data.SettleType = settleType
	return resp, nil
}
