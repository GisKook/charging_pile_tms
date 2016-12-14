package db

import (
	"database/sql"
	"log"
)

const SQL_WIFI_AND_PASSWD string = "SELECT w.ssid, w.password FROM t_wifi w INNER JOIN t_charge_pile p on w.charge_pile_id=p.id WHERE p.cpid=$1"

func (db_socket *DbSocket) GetWifiAndPasswd(cpid uint64) []byte {
	//cpid_string := strconv.FormatUint(cpid, 10)
	rows, err := db_socket.Db.Query(SQL_WIFI_AND_PASSWD, cpid)
	log.Println(cpid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var sql_wifi sql.NullString
	var sql_passwd sql.NullString
	var wifi string
	var passwd string
	for rows.Next() {
		if err := rows.Scan(&sql_wifi, &sql_passwd); err != nil {
			//log.Fatal(err)
		}
		wifi = GetStringValue(sql_wifi, "")
		passwd = GetStringValue(sql_passwd, "")
		log.Println(wifi)
		log.Println(passwd)
	}
	//	if err := rows.Err(); err != nil {
	//		log.Fatal(err)
	//	}
	wifi_len := byte(len(wifi))
	passwd_len := byte(len(passwd))

	result := []byte{wifi_len}
	result = append(result, []byte(wifi)...)
	result = append(result, passwd_len)
	result = append(result, []byte(passwd)...)

	return result
}

func (db_socket *DbSocket) GetInterfaceType(cpid uint64) uint8 {
	charging_pile, ok := db_socket.ChargePile[cpid]
	if ok {
		return charging_pile.InterfaceType
	}

	return 255
}

func (db_socket *DbSocket) GetBaudRate(cpid uint64) uint8 {
	charging_pile, ok := db_socket.ChargePile[cpid]
	if ok {
		return charging_pile.BaudRate
	}

	return 255
}

func (db_socket *DbSocket) GetStationID(cpid uint64) uint64 {
	charging_pile, ok := db_socket.ChargePile[cpid]
	if ok {
		return charging_pile.StationID
	}

	return 0
}
