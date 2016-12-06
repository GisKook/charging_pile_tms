package server

import (
	"github.com/giskook/charging_pile_tms/pb"
	"github.com/golang/protobuf/proto"
	"log"
)

func ProcessNsq(message []byte) {
	command := &Report.Command{}
	err := proto.Unmarshal(message, command)
	if err != nil {
		log.Println("unmarshal error")
	} else {
		log.Printf("<IN NSQ> %s %x \n", command.Uuid, command.Tid)
		switch command.Type {
		case Report.Command_CMT_REQ_LOGIN:
			event_handler_rep_login(command.Uuid, command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REQ_SETTING:
			event_handler_rep_setting(command.Uuid, command.Tid, command.SerialNumber, command.Paras)
			break
		case Report.Command_CMT_REQ_PRICE:
			event_handler_rep_price(command.Uuid, command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_MODE:
			event_handler_rep_mode(command.Uuid, command.Tid, command.SerialNumber, command.Paras)
		case Report.Command_CMT_REQ_MAX_CURRENT:
			event_handler_rep_max_current(command.Uuid, command.Tid, command.SerialNumber, command.Paras)
		}
	}
}
