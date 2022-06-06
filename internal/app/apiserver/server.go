package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/model"
	"github.com/ArdentK/bmstu-6sem-ppo/internal/app/store"
	"github.com/google/uuid"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
)

const (
	sessionName        = "user"
	ctxKeyUser  ctxKey = iota
	ctxKeyRequestID
)

const (
	logFilePath = "../../../logrus.log"
)

var (
	errIncorrectEmailOrPassword = errors.New("incorrect email or password")
	errNotAuthenticated         = errors.New("not authenticated")
	errTemplateError            = errors.New("template error")
	errDB                       = errors.New("db error")
	errForm                     = errors.New("bad form")
)

type ctxKey int8

type server struct {
	router       *mux.Router
	logger       *logrus.Logger
	store        store.Store
	sessionStore sessions.Store
	tmpl         *template.Template
}

func newServer(store store.Store, sessionStore sessions.Store, templatesPath string) *server {
	templates := template.Must(template.ParseGlob(templatesPath))

	s := &server{
		router:       mux.NewRouter(),
		logger:       logrus.New(),
		store:        store,
		sessionStore: sessionStore,
		tmpl:         templates,
	}

	s.configureLogger()
	s.configureRouter()

	return s
}

func (s *server) configureLogger() {
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	s.logger.SetFormatter(&logrus.JSONFormatter{})
	s.logger.SetOutput(f)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(s.setRequestID)
	s.router.Use(s.logRequest)
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))

	s.router.HandleFunc("/", s.handleIndex()).Methods("GET")
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionCreate()).Methods("POST")

	s.router.HandleFunc("/competitions", s.handleCompetitionsIndex()).Methods("GET")
	s.router.HandleFunc("/competitions/new", s.handleCompetitionAddForm()).Methods("GET")
	s.router.HandleFunc("/competitions/new", s.handleCompetitionAdd()).Methods("POST")
	s.router.HandleFunc("/competitions/{id}", s.handleCompetitionUpdate()).Methods("POST")
	s.router.HandleFunc("/competitions/{id}", s.handleCompetitionEdit()).Methods("GET")
	s.router.HandleFunc("/competitions/{id}", s.handleCompetitionDelete()).Methods("DELETE")

	s.router.HandleFunc("/athlets", s.handleAthletsIndex()).Methods("GET")
	s.router.HandleFunc("/news", s.handleNewsIndex()).Methods("GET")

	private := s.router.PathPrefix("/private").Subrouter()
	private.Use(s.authenticateUser)
	private.HandleFunc("/whoami", s.handleWhoami()).Methods("GET")
}

func (s *server) handleCompetitionAddForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.tmpl.ExecuteTemplate(w, "create.html", nil)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errTemplateError)
			return
		}
	}
}

func (s *server) handleCompetitionAdd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		item := model.Competition{}
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		err := decoder.Decode(&item, r.PostForm)
		if err != nil {
			s.respond(w, r, http.StatusBadRequest, errForm)
			return
		}

		err = s.store.Competition().Create(&item)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errDB)
			return
		}

		http.Redirect(w, r, "/competitions", http.StatusFound)
	}
}

func (s *server) handleCompetitionUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, r, http.StatusBadGateway, "bad id")
			return
		}

		r.ParseForm()
		item := model.Competition{}
		decoder := schema.NewDecoder()
		decoder.IgnoreUnknownKeys(true)
		err = decoder.Decode(&item, r.PostForm)

		if err != nil {
			s.respond(w, r, http.StatusBadRequest, errForm)
			return
		}

		item.ID = id

		err = s.store.Competition().Update(&item)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errDB)
			return
		}

		http.Redirect(w, r, "/competitions", http.StatusFound)
	}
}

func (s *server) handleCompetitionEdit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, r, http.StatusBadGateway, "bad id")
			return
		}

		item, err := s.store.Competition().Find(id)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errDB)
			return
		}

		if item == nil {
			s.respond(w, r, http.StatusNotFound, "no record")
			return
		}

		err = s.tmpl.ExecuteTemplate(w, "edit.html", item)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errTemplateError)
			return
		}
	}
}

func (s *server) handleCompetitionDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.respond(w, r, http.StatusBadGateway, "bad id")
			return
		}

		err = s.store.Competition().Delete(id)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errDB)
			return
		}

		w.Header().Set("Content-type", "application/json")
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleNewsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elems, err := s.store.News().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		err = s.tmpl.ExecuteTemplate(w, "indexNews.html", struct {
			Items []*model.News
		}{
			Items: elems,
		})
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, errTemplateError)
			return
		}
	}
}

func (s *server) handleCompetitionsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		elems, err := s.store.Competition().GetAll()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		err = s.tmpl.ExecuteTemplate(w, "index.html", struct {
			Items []*model.Competition
		}{
			Items: elems,
		})
		if err != nil {
			http.Error(w, `Template errror`, http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleAthletsIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "athlets")
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := s.tmpl.ExecuteTemplate(w, "login.html", nil)
		if err != nil {
			http.Error(w, `Template errror`, http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (s *server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := s.logger.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		logger.Infof(
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (s *server) authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		id, ok := session.Values["user_id"]
		if !ok {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		u, err := s.store.User().Find(id.(int))
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
	})
}

func (s *server) handleWhoami() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		err := json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		err = s.store.User().Create(u)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitaze()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		// err := json.NewDecoder(r.Body).Decode(req)
		// if err != nil {
		// 	s.error(w, r, http.StatusBadRequest, err)
		// 	return
		// }

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, errIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(r, sessionName)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_id"] = u.ID
		err = s.sessionStore.Save(r, w, session)
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
