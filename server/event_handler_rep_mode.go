package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_mode(uuid string, tid uint64, serial uint32, param []*Report.Param) {
	log.Println("rep mode")
	cp := GetServer().DB.ChargePile[tid]
	if cp != nil {
		mode := cp.VoltageInput
		if mode == 380 {
			mode = 0
		} else if mode == 220 {
			mode = 2
		} else {
			mode = 1
		}

		command := &Report.Command{
			Type: Report.Command_CMT_REP_MODE,
			Uuid: conf.GetConf().Uuid,
			Tid:  tid,
			Paras: []*Report.Param{
				&Report.Param{
					Type:  Report.Param_UINT8,
					Npara: uint64(mode),
				},
			},
		}
		data, _ := proto.Marshal(command)

		GetServer().MQ.Send(uuid, data)
	}
}
