package session

import (
	"errors"
	"net/http"
	"simple-user-inventory/server/quick"
	"simple-user-inventory/server/utils"

	"github.com/gorilla/sessions"
	echoS "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const sessionName = "__suis__session"
const uuidKey = "__u_u_id"

func NewSessionStroe(secret string) sessions.Store {
	return sessions.NewCookieStore([]byte(secret))
}

func NewSessionOptions() *sessions.Options {
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
		c.Logger().Error(err)
		return "", quick.ServiceError()
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
