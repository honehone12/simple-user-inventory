package session

import (
	"errors"
	"net/http"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/utils"

	uuidG "github.com/google/uuid"

	"github.com/gorilla/sessions"
	echoS "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	sessionName = "__suis__session"
	uuidKey     = "__ui"
)

type SessionStatus uint8

const (
	Ok SessionStatus = 1 + iota
	NotStored
	Rejected
	Error
)

type SessionData struct {
	Status SessionStatus
	Uuid   string
	Id     uint
}

func NewSessionStroe(secret string) sessions.Store {
	return sessions.NewCookieStore([]byte(secret))
}

func NewSessionOptions() *sessions.Options {
	// these options are only valid for browsers
	return &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24,
		Secure:   !utils.IsDev(),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func Set(c echo.Context, uuid string) error {
	sess, err := echoS.Get(sessionName, c)
	if err != nil {
		c.Logger().Error(err)
		return quick.ServiceError()
	}

	sess.Options = NewSessionOptions()
	sess.Values[uuidKey] = uuid
	sess.Save(c.Request(), c.Response())
	return nil
}

func Get(c echo.Context) (string, error) {
	sess, err := echoS.Get(sessionName, c)
	if err != nil {
		return "", err
	}

	uuid, ok := sess.Values[uuidKey]
	if !ok {
		return "", nil
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		return "", errors.New("could not read stored uuid value as string")
	}

	return uuidStr, nil
}

func GetAndVerify(c echo.Context) (SessionData, error) {
	uuid, err := Get(c)
	if err != nil {
		return SessionData{
			Status: Error,
			Uuid:   "",
			Id:     0,
		}, err
	}
	if len(uuid) == 0 {
		return SessionData{
			Status: NotStored,
			Uuid:   "",
			Id:     0,
		}, nil
	}
	if _, err = uuidG.Parse(uuid); err != nil {
		return SessionData{
			Status: Rejected,
			Uuid:   "",
			Id:     0,
		}, err
	}

	ctrl := c.(*context.Context).User()
	id, err := ctrl.UuidToId(uuid)
	if err != nil {
		return SessionData{
			Status: Rejected,
			Uuid:   "",
			Id:     0,
		}, err
	}

	return SessionData{
		Status: Ok,
		Uuid:   uuid,
		Id:     id,
	}, nil
}
