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
	var phase_mode uint8 = 0
	var auth_mode uint8 = 0
	var lock_mode uint8 = 0
	if !GetServer().DB.CheckChargingPileID(tid) {
		result = 1
	} else {
		id = GetServer().DB.ChargePile[tid].ID
		station_id = GetServer().DB.ChargePile[tid].StationID
		phase_mode = uint8(GetServer().DB.ChargePile[tid].VoltageInput)
		auth_mode = GetServer().DB.ChargePile[tid].AuthMode
		lock_mode = GetServer().DB.ChargePile[tid].LockMode

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
			&Report.Param{
				Type:  Report.Param_UINT64,
				Npara: uint64(phase_mode),
			},
			&Report.Param{
				Type:  Report.Param_UINT64,
				Npara: uint64(auth_mode),
			},
			&Report.Param{
				Type:  Report.Param_UINT64,
				Npara: uint64(lock_mode),
			},
		},
	}

	data, _ := proto.Marshal(command)

	GetServer().MQ.Send(uuid, data)
}
