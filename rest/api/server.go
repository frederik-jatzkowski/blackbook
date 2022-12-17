package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/frederik-jatzkowski/blackbook/user"
)

type Server struct {
	frontend string
	user     *user.Service
	mux      *http.ServeMux
	server   *http.Server
}

func NewServer(auth *user.Service) (*Server, error) {
	mux := http.NewServeMux()
	server := &Server{
		frontend: os.Getenv("FRONTEND_ADDRESS"),
		user:     auth,
		mux:      mux,
		server:   &http.Server{Handler: mux, Addr: ":8080"},
	}

	if server.frontend == "" {
		return server, fmt.Errorf("missing env variable FRONTEND_ADDRESS")
	}

	return server, nil
}

func (server Server) cors(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST")
		w.Header().Add("Access-Control-Allow-Origin", server.frontend)
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, *")
		handler(w, r)
	}
}

func (server *Server) Start() error {
	var (
		err error
	)

	// user service
	server.mux.HandleFunc("/user/create", server.cors(server.user.HandleCreate))
	server.mux.HandleFunc("/user/activate", server.cors(server.user.HandleActivate))
	server.mux.HandleFunc("/user/login", server.cors(server.user.HandleLogin))
	server.mux.HandleFunc("/user/logout", server.cors(server.user.HandleLogout))
	server.mux.HandleFunc("/user/sessionCheck", server.cors(server.user.HandleSessionCheck))
	server.mux.HandleFunc("/user/update", server.cors(server.user.HandleUpdate))
	server.mux.HandleFunc("/user/delete", server.cors(server.user.HandleDelete))

	err = server.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
