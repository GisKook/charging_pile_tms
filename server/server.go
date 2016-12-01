package server

import (
	"github.com/giskook/charging_pile_tms/conf"
	"github.com/giskook/charging_pile_tms/db"
)

type Server struct {
	MQ *NsqSocket
	DB *db.DbSocket
}

var g_server *Server

func NewServer(nsq_socket *NsqSocket, db_socket *db.DbSocket) *Server {
	g_server = &Server{
		MQ: nsq_socket,
		DB: db_socket,
	}

	return g_server
}

func (server *Server) Start() {
	server.MQ.Start()
	server.DB.Listen(conf.GetConf().DB.ListenPriceTable)
	server.DB.Listen(conf.GetConf().DB.NotifyTable)
	server.DB.WaitForNotification()
}

func (server *Server) Stop() {
	server.MQ.Stop()
	server.DB.Listener.UnlistenAll()
	server.DB.Close()
}

func GetServer() *Server {
	return g_server
}
