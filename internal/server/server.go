package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/TimNikolaev/Pocketoid/internal/repository"
	"github.com/zhashkevych/go-pocket-sdk"
)

type AuthorizationServer struct {
	server       *http.Server
	pocketClient *pocket.Client
	repository   repository.Repository
	redirectURL  string
}

func NewAuthorizationServer(pck *pocket.Client, rps repository.Repository, rdUrl string) *AuthorizationServer {
	return &AuthorizationServer{
		pocketClient: pck,
		repository:   rps,
		redirectURL:  rdUrl,
	}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":8888",
		Handler: s,
	}

	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := s.repository.Get(chatID, repository.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResp, err := s.pocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.repository.Save(chatID, authResp.AccessToken, repository.AccessTokens); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("chat_id: %d\nrequest_token: %s\naccess_token: %s\n", chatID, requestToken, authResp.AccessToken)

	// w.Header().Add("Location", s.redirectURL)
	// w.WriteHeader(http.StatusMovedPermanently)
	http.Redirect(w, r, s.redirectURL, http.StatusMovedPermanently)
}
