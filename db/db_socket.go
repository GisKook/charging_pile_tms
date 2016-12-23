package db

import (
	"database/sql"
	"fmt"
	"github.com/giskook/charging_pile_tms/charging_pile"
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/lib/pq"
	"log"
	"strconv"
	"strings"
	"time"
)

type DbSocket struct {
	Db             *sql.DB
	ChargePile     map[uint64]*charging_pile.ChargingPile
	ChargingPrices map[uint64][]*charging_pile.ChargingPrice

	Listener *pq.Listener
}

func NewDbSocket(db_config *conf.DBConfigure) (*DbSocket, error) {
	conn_string := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", db_config.User, db_config.Passwd, db_config.Host, db_config.Port, db_config.DbName)

	log.Println(conn_string)
	db, err := sql.Open(db_config.User, conn_string)

	if err != nil {
		return nil, err
	}
	log.Println("db open success")

	reportProblem := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			log.Println(err.Error())
		}
	}

	return &DbSocket{
		Db:             db,
		ChargePile:     make(map[uint64]*charging_pile.ChargingPile),
		Listener:       pq.NewListener(conn_string, 60*time.Second, time.Minute, reportProblem),
		ChargingPrices: make(map[uint64][]*charging_pile.ChargingPrice),
	}, nil
}

func (db_socket *DbSocket) LoadAll() error {
	st, err := db_socket.Db.Prepare("select cpid, station_id, terminal_type_id, rated_power, electric_current_type, voltage_input, voltage_output, electric_current_output, gun_number, ammeter_number,interface_type, baud_rate ,id from t_charge_pile")

	if err != nil {
		return err
	}

	r, e := st.Query()
	defer st.Close()

	if e != nil {
		return e
	}

	var sql_charging_pile_id sql.NullString
	var sql_station_id sql.NullInt64
	var sql_terminal_type_id sql.NullInt64
	var sql_rated_power sql.NullFloat64
	var sql_electric_current_type sql.NullInt64
	var sql_voltage_input sql.NullInt64
	var sql_voltage_output sql.NullInt64
	var sql_electric_current_output sql.NullInt64
	var sql_gun_number sql.NullInt64
	var sql_ammeter_number sql.NullFloat64
	var sql_interface_type sql.NullInt64
	var sql_baud_rate sql.NullInt64
	var sql_id sql.NullInt64

	for r.Next() {
		err = r.Scan(
			&sql_charging_pile_id,
			&sql_station_id,
			&sql_terminal_type_id,
			&sql_rated_power,
			&sql_electric_current_type,
			&sql_voltage_input,
			&sql_voltage_output,
			&sql_electric_current_output,
			&sql_gun_number,
			&sql_ammeter_number,
			&sql_interface_type,
			&sql_baud_rate,
			&sql_id,
		)
		charging_pile_id := GetStringValue(sql_charging_pile_id, "")
		station_id := uint64(GetInt64Value(sql_station_id, 0))
		terminal_type_id := uint8(GetInt64Value(sql_terminal_type_id, 0))
		rated_power := float32(GetFloat64Value(sql_rated_power, 0.0))
		electric_current_type := uint8(GetInt64Value(sql_electric_current_type, 0))
		voltage_input := uint32(GetInt64Value(sql_voltage_input, 0))
		voltage_output := uint32(GetInt64Value(sql_voltage_output, 0))
		electric_current_output := uint32(GetInt64Value(sql_electric_current_output, 0))
		gun_number := uint32(GetInt64Value(sql_gun_number, 0))
		ammeter_number := float32(GetFloat64Value(sql_ammeter_number, 0.0))
		interface_type := uint8(GetInt64Value(sql_interface_type, 0))
		baud_rate := uint8(GetInt64Value(sql_baud_rate, 0))
		id := uint32(GetInt64Value(sql_id, 0))

		if err != nil {
			log.Println(err.Error())
			continue
			//		return err
		}
		cpid, _ := strconv.ParseUint(charging_pile_id, 10, 64)
		db_socket.ChargePile[cpid] = &charging_pile.ChargingPile{
			ID:                    id,
			StationID:             station_id,
			TypeID:                terminal_type_id,
			RatedPower:            rated_power,
			ElectricCurrentType:   electric_current_type,
			VoltageInput:          voltage_input,
			VoltageOutput:         voltage_output,
			ElectricCurrentOutput: electric_current_output,
			GunNum:                gun_number,
			AmmeterNum:            ammeter_number,
			InterfaceType:         interface_type,
			BaudRate:              baud_rate,
		}

	}

	log.Println(db_socket.ChargePile)

	defer r.Close()

	return nil
}

func (db_socket *DbSocket) Listen(table string) error {
	return db_socket.Listener.Listen(table)
}

func (db_socket *DbSocket) ProcessNotify(_notify string) {
	log.Println(_notify)
	if strings.Contains(_notify, conf.GetConf().DB.NotifyTable) {
		notify := strings.TrimPrefix(_notify, conf.GetConf().DB.NotifyTable+"^")
		log.Println(notify)
		switch notify[0] {
		case 'U':
			db_socket.update(notify)
		case 'I':
			db_socket.insert(notify)
		case 'D':
			db_socket.del(notify)
		}
	} else if strings.Contains(_notify, conf.GetConf().DB.ListenPriceTable) {
		notify := strings.TrimPrefix(_notify, conf.GetConf().DB.ListenPriceTable+"^")
		log.Println(notify)
		db_socket.parse_payload_price(notify)
	}

}

func (db_socket *DbSocket) WaitForNotification() {
	for {
		select {
		case notify := <-db_socket.Listener.Notify:
			db_socket.ProcessNotify(notify.Extra)
			break
		case <-time.After(90 * time.Second):
			go func() {
				db_socket.Listener.Ping()
			}()
			// Check if there's more work available, just in case it takes
			// a while for the Listener to notice connection loss and
			// reconnect.
			log.Println("received no work for 90 seconds, checking for new work")
			break
		}
	}
}

func (db_socket *DbSocket) Close() {
	db_socket.Listener.Close()
	db_socket.Db.Close()
}

func (db_socket *DbSocket) CheckChargingPileID(cpid uint64) bool {
	_, ok := db_socket.ChargePile[cpid]
	return ok
}

func (db_socket *DbSocket) parse_payload(payload string) (uint64, *charging_pile.ChargingPile) {
	values := strings.Split(payload, "^")
	id, _ := strconv.ParseUint(values[1], 10, 32)
	station_id, _ := strconv.ParseUint(values[2], 10, 64)
	charging_pile_id, _ := strconv.ParseUint(values[11], 10, 64)
	terminal_type_id, _ := strconv.ParseUint(values[3], 10, 8)
	rated_power, _ := strconv.ParseFloat(values[4], 32)
	electric_current_type, _ := strconv.ParseUint(values[5], 10, 8)
	voltage_input, _ := strconv.ParseUint(values[6], 10, 32)
	voltage_output, _ := strconv.ParseUint(values[7], 10, 32)
	electric_current_output, _ := strconv.ParseUint(values[8], 10, 32)
	gun_number, _ := strconv.ParseUint(values[9], 10, 32)
	ammeter_number, _ := strconv.ParseFloat(values[10], 32)
	interface_type, _ := strconv.ParseUint(values[13], 10, 8)
	baud_rate, _ := strconv.ParseUint(values[14], 10, 8)

	log.Println(charging_pile_id)
	return charging_pile_id, &charging_pile.ChargingPile{
		ID:                    uint32(id),
		StationID:             station_id,
		TypeID:                uint8(terminal_type_id),
		RatedPower:            float32(rated_power),
		ElectricCurrentType:   uint8(electric_current_type),
		VoltageInput:          uint32(voltage_input),
		VoltageOutput:         uint32(voltage_output),
		ElectricCurrentOutput: uint32(electric_current_output),
		GunNum:                uint32(gun_number),
		AmmeterNum:            float32(ammeter_number),
		InterfaceType:         uint8(interface_type),
		BaudRate:              uint8(baud_rate),
	}
}

func (db_socket *DbSocket) insert(payload string) {
	tid, charging_pile := db_socket.parse_payload(payload)
	db_socket.ChargePile[tid] = charging_pile
	log.Println(db_socket.ChargePile)
}

func (db_socket *DbSocket) del(payload string) {
	tid, _ := db_socket.parse_payload(payload)
	delete(db_socket.ChargePile, tid)
	log.Println(db_socket.ChargePile)
}

func (db_socket *DbSocket) update(payload string) {
	tid, charging_pile := db_socket.parse_payload(payload)
	db_socket.ChargePile[tid] = charging_pile
	log.Println(db_socket.ChargePile)
}
