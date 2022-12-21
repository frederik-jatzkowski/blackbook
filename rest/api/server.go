package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/frederik-jatzkowski/blackbook/group"
	"github.com/frederik-jatzkowski/blackbook/user"
)

type Server struct {
	frontend string
	user     *user.Service
	group    *group.Service
	mux      *http.ServeMux
	server   *http.Server
}

func NewServer(user *user.Service, group *group.Service) (*Server, error) {
	mux := http.NewServeMux()
	server := &Server{
		frontend: os.Getenv("FRONTEND_ADDRESS"),
		user:     user,
		group:    group,
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

	// group service
	server.mux.HandleFunc("/group/create", server.cors(server.group.HandleCreate))
	server.mux.HandleFunc("/group/getAll", server.cors(server.group.HandleGetAll))
	server.mux.HandleFunc("/group/invite", server.cors(server.group.HandleInvite))
	server.mux.HandleFunc("/group/accept", server.cors(server.group.HandleAccept))
	server.mux.HandleFunc("/group/decline", server.cors(server.group.HandleDecline))
	server.mux.HandleFunc("/group/update", server.cors(server.group.HandleUpdate))
	server.mux.HandleFunc("/group/leave", server.cors(server.group.HandleLeave))

	err = server.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
