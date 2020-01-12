package middleware

import (
	"net/http"

	"github.com/jimweng/thirdparty/redis_related/redigo/session_advance/store"
	"github.com/jimweng/thirdparty/redis_related/redigo/session_advance/utils"
	"github.com/labstack/echo"
)

func Middleware(s store.Store) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			cookie, err := ectx.Cookie("sessionID")
			if err != nil {
				sessionID := utils.ReturnUUidNewV4()
				ectx.SetCookie(&http.Cookie{
					Name:  "sessionID",
					Value: sessionID,
				})
				ectx.Set("sessionID", sessionID)
				s.Set(sessionID, store.Session{})
				return hf(ectx)
			}

			sessionID := cookie.Value
			ectx.Set("sessionID", sessionID)
			return hf(ectx)
		}
	}
}
