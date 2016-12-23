package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_login(uuid string, tid uint64, serial uint32, param []*Report.Param) {

	var result uint8 = 0
	var id uint32 = 0
	var station_id uint64 = 0
	if !GetServer().DB.CheckChargingPileID(tid) {
		result = 1
	} else {
		id = GetServer().DB.ChargePile[tid].ID
		station_id = GetServer().DB.ChargePile[tid].StationID

		log.Println(result)
		log.Println(id)
		log.Println(station_id)
	}
	command := &Report.Command{
		Type: Report.Command_CMT_REP_LOGIN,
		Uuid: conf.GetConf().Uuid,
		Tid:  tid,
		Paras: []*Report.Param{
			&Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(result),
			},
			&Report.Param{
				Type:  Report.Param_UINT32,
				Npara: uint64(id),
			},
			&Report.Param{
				Type:  Report.Param_UINT64,
				Npara: station_id,
			},
			&Report.Param{ // status
				Type:  Report.Param_UINT8,
				Npara: param[2].Npara,
			},
			&Report.Param{ // time stamp
				Type:  Report.Param_UINT64,
				Npara: param[3].Npara,
			},
		},
	}

	data, _ := proto.Marshal(command)

	GetServer().MQ.Send(uuid, data)
}
