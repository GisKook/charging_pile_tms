package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_price(uuid string, tid uint64, serial uint32, param []*Report.Param) {
	log.Println("rep price")
	var paras []*Report.Param
	prices := GetServer().DB.ChargingPrices[tid]
	for _, price := range prices {
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(price.Start_hour),
		})
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(price.Start_min),
		})
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(price.End_hour),
		})
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT8,
			Npara: uint64(price.End_min),
		})
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT16,
			Npara: uint64(price.Elec_unit_price),
		})
		paras = append(paras, &Report.Param{
			Type:  Report.Param_UINT16,
			Npara: uint64(price.Service_price),
		})
	}

	command := &Report.Command{
		Type:  Report.Command_CMT_REP_PRICE,
		Uuid:  conf.GetConf().Uuid,
		Tid:   tid,
		Paras: paras,
	}

	data, _ := proto.Marshal(command)

	GetServer().MQ.Send(uuid, data)
}
