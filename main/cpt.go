package main

import (
	"fmt"
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/db"
	"github.com/giskook/charging_pile_tms/server"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// read configuration
	configuration, err := conf.ReadConfig("./conf.json")

	checkError(err)
	// create a mq socket
	mq_socket := server.NewNsqSocket(configuration.Nsq)
	// create a db socket
	db_socket, e := db.NewDbSocket(configuration.DB)
	checkError(e)
	db_socket.LoadAll()
	db_socket.LoadAllPrices()
	// create server
	server := server.NewServer(mq_socket, db_socket)
	server.Start()

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	server.Stop()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
