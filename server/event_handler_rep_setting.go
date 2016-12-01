package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func event_handler_rep_setting(uuid string, tid uint64, serial uint32, param []*Report.Param) {
	log.Println("rep")
	if len(param) == 1 {
		mode := uint8(param[0].Npara)
		command := &Report.Command{
			Type: Report.Command_CMT_REP_SETTING,
			Uuid: conf.GetConf().Uuid,
			Tid:  tid,
			Paras: []*Report.Param{
				&Report.Param{
					Type:  Report.Param_UINT8,
					Npara: uint64(mode),
				},
			},
		}

		if mode == 1 { // 1 means interfacetype
			command.Paras = append(command.Paras, &Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(GetServer().DB.GetInterfaceType(tid)),
			})
		} else if mode == 3 {
			command.Paras = append(command.Paras, &Report.Param{
				Type:  Report.Param_UINT8,
				Npara: uint64(GetServer().DB.GetBaudRate(tid)),
			})
		} else if mode == 2 {
			command.Paras = append(command.Paras, &Report.Param{
				Type:  Report.Param_BYTES,
				Bpara: GetServer().DB.GetWifiAndPasswd(tid),
			})

		}

		data, _ := proto.Marshal(command)

		GetServer().MQ.Send(uuid, data)
	}
}
