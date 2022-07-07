package server

import (
	"errors"
	"log"
	"net"

	"github.com/google/uuid"

	"github.com/uladzislaubarysau/faraway_testing/internal/helpers"
	"github.com/uladzislaubarysau/faraway_testing/storage"
)

type Server struct {
	storage *storage.WordsOfWisdom
}

func New() *Server {
	return &Server{
		storage: storage.NewWordsOfWisdom(),
	}
}

func (s *Server) Start(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Panic(err)
	}

	log.Println("server started on:", address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("listener accept:", err)
			continue
		}

		go func() {
			if err := s.handle(conn); err != nil {
				log.Println("handle err:", err)
			}
		}()
	}
}

func (s *Server) handle(conn net.Conn) (err error) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println("can't close connection:", err)
		}
	}()

	_, err = helpers.ReadMessage(conn)
	if err != nil {
		return
	}

	challenge := []byte(uuid.New().String())
	if err = helpers.WriteMessage(conn, challenge); err != nil {
		return
	}

	pow, err := helpers.ReadMessage(conn)
	if err != nil {
		return
	}

	if !helpers.Verify(challenge, pow) {
		err = errors.New("proof of work not valid")
		return
	}

	resp := s.storage.GetWordOfWisdom()

	return helpers.WriteMessage(conn, resp)
}
