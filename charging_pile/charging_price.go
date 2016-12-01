package charging_pile

import ()

type ChargingPrice struct {
	ID              uint64
	Start_hour      uint8
	Start_min       uint8
	End_hour        uint8
	End_min         uint8
	Elec_unit_price uint16
	Service_price   uint16
}
