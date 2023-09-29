package handler

import (
	"blogspot/flow"
	"blogspot/parser"
	"blogspot/util"
	"net/http"
	"strconv"
	"time"
)

type userHandler struct {
	parser    parser.IUserParser
	presenter util.IPresenterJSON
	flow      flow.IUserFlow
}

func NewUserHandler(parser parser.IUserParser, presenter util.IPresenterJSON, flow flow.IUserFlow) *userHandler {
	return &userHandler{
		parser:    parser,
		presenter: presenter,
		flow:      flow,
	}
}

func (h *userHandler) Register(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Register(user)
	if err != nil {
		h.presenter.SendError(w, "Error Insert")
		return
	}

	h.presenter.SendSuccess(w)
}

func (h *userHandler) Login(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	user, err = h.flow.Login(user)
	if err != nil {
		h.presenter.SendError(w, "Error Login")
		return
	}

	userTypeCookie := http.Cookie{
		Name:     "userType",
		Value:    strconv.Itoa(user.UserType),
		HttpOnly: true,
		Path:     "/",
	}

	userIDCookie := http.Cookie{
		Name:     "userID",
		Value:    strconv.Itoa(user.ID),
		HttpOnly: true,
		Path:     "/",
	}

	http.SetCookie(w, &userTypeCookie)
	http.SetCookie(w, &userIDCookie)

	h.presenter.SendSuccess(w)
}

func (h *userHandler) Logout(w http.ResponseWriter, r *http.Request) {
	userTypeCookie := http.Cookie{
		Name:     "userType",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), // Set an expiration time in the past
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &userTypeCookie)

	// Clear user ID cookie
	userIDCookie := http.Cookie{
		Name:     "userID",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, &userIDCookie)

	h.presenter.SendSuccess(w)
}

func (h *userHandler) Update(w http.ResponseWriter, r *http.Request) {
	user, err := h.parser.ParseUserEntity(r)
	if err != nil {
		h.presenter.SendError(w, "Error Parsing")
		return
	}

	err = h.flow.Update(user)
	if err != nil {
		h.presenter.SendError(w, "Error Update")
		return
	}

	h.presenter.SendSuccess(w)
}
