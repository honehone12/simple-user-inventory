package session

import (
	"errors"
	"net/http"
	"simple-user-inventory/server/context"
	"simple-user-inventory/server/quick"
	"time"

	"simple-user-inventory/server/utils"

	uuidG "github.com/google/uuid"

	"github.com/gorilla/sessions"
	echoS "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	sessionName  = "__suis__session"
	uuidKey      = "__ui"
	createdAtKey = "__ca"
)

var (
	ErrorSessionNotStored  = errors.New("session not stored")
	ErrorSessionExpired    = errors.New("session is expired")
	ErrorSessionParseError = errors.New("could not parsed stored value")
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
	// these options are only for browsers
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

	now := time.Now().Unix()
	sess.Options = NewSessionOptions()
	sess.Values[uuidKey] = uuid
	sess.Values[createdAtKey] = now
	sess.Save(c.Request(), c.Response())
	return nil
}

func Get(c echo.Context) (string, error) {
	sess, err := echoS.Get(sessionName, c)
	if err != nil {
		return "", err
	}

	createdAt, ok := sess.Values[createdAtKey]
	if !ok {
		return "", ErrorSessionNotStored
	}
	expiration, ok := createdAt.(int64)
	if !ok {
		return "", ErrorSessionParseError
	}
	expiration += 60 * 60
	if time.Now().Unix() > expiration {
		return "", ErrorSessionExpired
	}

	uuid, ok := sess.Values[uuidKey]
	if !ok {
		return "", ErrorSessionNotStored
	}

	uuidStr, ok := uuid.(string)
	if !ok {
		return "", ErrorSessionParseError
	}

	return uuidStr, nil
}

func GetAndVerify(c echo.Context) (SessionData, error) {
	uuid, err := Get(c)
	if err == ErrorSessionNotStored ||
		err == ErrorSessionExpired ||
		err == ErrorSessionParseError {

		return SessionData{
			Status: NotStored,
			Uuid:   "",
			Id:     0,
		}, err
	} else if err != nil {
		return SessionData{
			Status: Error,
			Uuid:   "",
			Id:     0,
		}, err
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

func RequireSession(c echo.Context) (SessionData, error) {
	sess, err := GetAndVerify(c)
	switch sess.Status {
	case Ok:
		return sess, nil
	case Error:
		c.Logger().Error(err)
		return sess, quick.ServiceError()
	case Rejected:
		fallthrough
	case NotStored:
		c.Logger().Warn(err)
		return sess, quick.NotAllowed()
	default:
		c.Logger().Error("not implemented")
		c.Logger().Error(err)
		return sess, quick.ServiceError()
	}
}
