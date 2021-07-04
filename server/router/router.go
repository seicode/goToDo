package router

import (
	s "goToDo/server/service"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router is exported and used in main.go
func Router() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323", "http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: logFormat(),
		Output: os.Stdout,
	}))
	e.GET("/", s.GetAllTask)
	e.GET("/api/task", s.GetAllTask)
	e.POST("/api/task", s.CreateTask)
	e.PUT("/api/task/:id", s.TaskComplete)
	e.PUT("/api/undoTask/:id", s.UndoTask)
	e.DELETE("/api/deleteTask/:id", s.DeleteTask)
	e.DELETE("/api/deleteAllTask", s.DeleteAllTask)
	e.Logger.Fatal(e.Start(":1323"))
	return e
}

func logFormat() string {
	// Refer to https://github.com/tkuchiki/alp
	var format string
	format += "time:${time_rfc3339}\t"
	format += "host:${host}\t"
	format += "status:${status}\t"
	format += "method:${method}\t"
	format += "uri:${uri}\t"
	format += "reqtime_ns:${latency}\n"
	return format
}
