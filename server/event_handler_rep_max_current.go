package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_max_current(uuid string, tid uint64, serial uint32, param []*Report.Param) {
	log.Println("rep max_current")
	cp := GetServer().DB.ChargePile[tid]
	if cp != nil {
		max_current := cp.ElectricCurrentOutput

		command := &Report.Command{
			Type: Report.Command_CMT_REP_MAX_CURRENT,
			Uuid: conf.GetConf().Uuid,
			Tid:  tid,
			Paras: []*Report.Param{
				&Report.Param{
					Type:  Report.Param_UINT8,
					Npara: uint64(max_current),
				},
			},
		}
		data, _ := proto.Marshal(command)

		GetServer().MQ.Send(uuid, data)
	}
}
