package db

import (
	"database/sql"
	"github.com/giskook/charging_pile_tms/charging_pile"
	"github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	"time"
)

func (db_socket *DbSocket) LoadAllPrices() error {
	st, err := db_socket.Db.Prepare("SELECT id, charge_station_id, start_time, end_time, electricity_price, service_price FROM t_electricity_price")

	if err != nil {
		return err
	}

	r, e := st.Query()
	defer st.Close()

	if e != nil {
		return e
	}

	var sql_id sql.NullInt64
	var sql_station_id sql.NullInt64
	var sql_start_time pq.NullTime
	var sql_end_time pq.NullTime
	var sql_electricity_price sql.NullFloat64
	var sql_service_price sql.NullFloat64

	for r.Next() {
		err = r.Scan(
			&sql_id,
			&sql_station_id,
			&sql_start_time,
			&sql_end_time,
			&sql_electricity_price,
			&sql_service_price,
		)
		id := uint64(GetInt64Value(sql_id, 0))
		station_id := uint64(GetInt64Value(sql_station_id, 0))
		start_time, _ := GetTimeValue(sql_start_time)
		end_time, _ := GetTimeValue(sql_end_time)
		electricity_price := float32(GetFloat64Value(sql_electricity_price, 0.0))
		service_price := float32(GetFloat64Value(sql_service_price, 0))

		if err != nil {
			log.Println(err.Error())
			return err
		}
		db_socket.ChargingPrices[station_id] =
			append(db_socket.ChargingPrices[station_id], &charging_pile.ChargingPrice{
				ID:              id,
				Start_hour:      uint8(start_time.Hour()),
				Start_min:       uint8(start_time.Minute()),
				End_hour:        uint8(end_time.Hour()),
				End_min:         uint8(end_time.Minute()),
				Elec_unit_price: uint16(electricity_price * 100),
				Service_price:   uint16(service_price * 100),
			})
	}

	log.Println(db_socket.ChargePile)

	defer r.Close()

	return nil
}

func (db_socket *DbSocket) parse_payload_price(notify string) {
	switch notify[0] {
	case 'U':
		db_socket.update_price(notify)
	case 'I':
		db_socket.insert_price(notify)
	case 'D':
		db_socket.del_price(notify)
	}

}

func (db_socket *DbSocket) parse_payload_price_common(payload string) (uint64, *charging_pile.ChargingPrice) {
	values := strings.Split(payload, "^")
	id, _ := strconv.ParseUint(values[1], 10, 64)
	station_id, _ := strconv.ParseUint(values[2], 10, 64)
	start_time_string := values[4]
	end_time_string := values[5]
	elec_unit_price, _ := strconv.ParseFloat(values[6], 32)
	service_price, _ := strconv.ParseFloat(values[7], 32)
	start_time, _ := time.Parse(time.Stamp, start_time_string)
	end_time, _ := time.Parse(time.Stamp, end_time_string)
	log.Println(start_time)

	return station_id, &charging_pile.ChargingPrice{
		ID:              id,
		Start_hour:      uint8(start_time.Hour()),
		Start_min:       uint8(start_time.Minute()),
		End_hour:        uint8(end_time.Hour()),
		End_min:         uint8(end_time.Minute()),
		Elec_unit_price: uint16(elec_unit_price * 100),
		Service_price:   uint16(service_price * 100),
	}
}

func (db_socket *DbSocket) insert_price(payload string) {
	tid, charging_price := db_socket.parse_payload_price_common(payload)
	log.Println(db_socket.ChargingPrices[tid])
	db_socket.ChargingPrices[tid] = append(db_socket.ChargingPrices[tid], charging_price)
	log.Println(db_socket.ChargingPrices[tid])
}

func (db_socket *DbSocket) del_price(payload string) {
}

func (db_socket *DbSocket) update_price(payload string) {
}
