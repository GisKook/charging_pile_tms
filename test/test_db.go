package main

import (
	"fmt"
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/db"
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
	// create a db socket
	db_socket, e := db.NewDbSocket(configuration.DB)
	checkError(e)
	e = db_socket.LoadAll()
	checkError(e)

	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
