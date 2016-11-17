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
	Db         *sql.DB
	ChargePile map[uint64]*charging_pile.ChargingPile

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
		Db:         db,
		ChargePile: make(map[uint64]*charging_pile.ChargingPile),
		Listener:   pq.NewListener(conn_string, 60*time.Second, time.Minute, reportProblem),
	}, nil
}

func (db_socket *DbSocket) LoadAll() error {
	st, err := db_socket.Db.Prepare("select code, terminal_type_id, rated_power, electric_current_type, voltage_input, voltage_output, electric_current_output, gun_number, ammeter_number from t_charge_pile")

	if err != nil {
		return err
	}

	r, e := st.Query()

	if e != nil {
		return e
	}

	defer st.Close()
	var charging_pile_id uint64
	var terminal_type_id uint8
	var rated_power float32
	var electric_current_type uint8
	var voltage_input uint32
	var voltage_output uint32
	var electric_current_output uint32
	var gun_number uint32
	var ammeter_number float32

	for r.Next() {
		err = r.Scan(&charging_pile_id,
			&terminal_type_id,
			&rated_power,
			&electric_current_type,
			&voltage_input,
			&voltage_output,
			&electric_current_output,
			&gun_number,
			&ammeter_number)
		if err != nil {
			return err
		}
		db_socket.ChargePile[charging_pile_id] = &charging_pile.ChargingPile{
			TypeID:                terminal_type_id,
			RatedPower:            rated_power,
			ElectricCurrentType:   electric_current_type,
			VoltageInput:          voltage_input,
			VoltageOutput:         voltage_output,
			ElectricCurrentOutput: electric_current_output,
			GunNum:                gun_number,
			AmmeterNum:            ammeter_number,
		}

	}

	defer r.Close()

	return nil
}

func (db_socket *DbSocket) Listen(table string) error {
	return db_socket.Listener.Listen(table)
}

func (db_socket *DbSocket) WaitForNotification() {
	for {
		select {
		case notify := <-db_socket.Listener.Notify:
			log.Println(notify.Extra)
			switch notify.Extra[0] {
			case 'U':
				db_socket.update(notify.Extra)
			case 'I':
				db_socket.insert(notify.Extra)
			case 'D':
				db_socket.del(notify.Extra)
			}
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
	charging_pile_id, _ := strconv.ParseUint(values[0], 10, 64)
	terminal_type_id, _ := strconv.ParseUint(values[1], 10, 8)
	rated_power, _ := strconv.ParseFloat(values[2], 32)
	electric_current_type, _ := strconv.ParseUint(values[3], 10, 8)
	voltage_input, _ := strconv.ParseUint(values[4], 10, 32)
	voltage_output, _ := strconv.ParseUint(values[5], 10, 32)
	electric_current_output, _ := strconv.ParseUint(values[6], 10, 32)
	gun_number, _ := strconv.ParseUint(values[7], 10, 32)
	ammeter_number, _ := strconv.ParseFloat(values[8], 32)

	return charging_pile_id, &charging_pile.ChargingPile{
		TypeID:                uint8(terminal_type_id),
		RatedPower:            float32(rated_power),
		ElectricCurrentType:   uint8(electric_current_type),
		VoltageInput:          uint32(voltage_input),
		VoltageOutput:         uint32(voltage_output),
		ElectricCurrentOutput: uint32(electric_current_output),
		GunNum:                uint32(gun_number),
		AmmeterNum:            float32(ammeter_number),
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
