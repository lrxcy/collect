package main

import (
	"fmt"
	"net/http"

	"github.com/jimweng/thirdparty/redis_related/redigo/session_advance/middleware"
	"github.com/jimweng/thirdparty/redis_related/redigo/session_advance/store/memorystore"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func FromContext(ectx echo.Context) *logrus.Entry {
	sessionID := ectx.Get("sessionID").(string)
	return logrus.WithField("sessionID", sessionID)
}

func main() {
	sessionsStore := memorystore.NewMemoryStore()
	// sessionsStore := NewRedisStore()

	e := echo.New()
	e.Use(middleware.Middleware(sessionsStore))

	e.GET("/", func(ectx echo.Context) error {
		log := FromContext(ectx)
		log.Info("Hello world")

		sessionID := ectx.Get("sessionID").(string)
		s, err := sessionsStore.Get(sessionID)
		if err != nil {
			log.Errorf("err: %v", err)
		}

		log.Infof("Visits: %d", s.VisitCount)
		response := fmt.Sprintf("Hello World #%d\n", s.VisitCount)

		s.VisitCount = s.VisitCount + 1
		err = sessionsStore.Set(sessionID, s)
		if err != nil {
			log.Errorf("err: %v", err)
		}

		return ectx.String(http.StatusOK, response)
	})

	e.Start(":5000")
}
