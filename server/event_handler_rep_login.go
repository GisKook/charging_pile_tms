package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
)

func event_handler_rep_login(uuid string, tid uint64, serial uint32, param []*Report.Param) {

	var result uint8 = 0
	if !GetServer().DB.CheckChargingPileID(tid) {
		result = 1
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
		},
	}

	data, _ := proto.Marshal(command)

	GetServer().MQ.Send(uuid, data)

}
