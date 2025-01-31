package controllers

import (
    "github.com/Encedeus/panel/config"
    "github.com/Encedeus/panel/ent"
    encMiddleware "github.com/Encedeus/panel/middleware"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

type Controller interface {
    registerRoutes(*Server)
}

func registerControllerRoutes(srv *Server, cs ...Controller) {
    for _, c := range cs {
        c.registerRoutes(srv)
    }
}

type Server struct {
    *echo.Echo
    DB *ent.Client
}

func NewEmptyServer(db *ent.Client) *Server {
    srv := &Server{
        Echo: echo.New(),
        DB:   db,
    }

    return srv
}

func WrapServerWithDefaults(srv *Server, _ *ent.Client) {
    srv.Use(encMiddleware.JSONSyntaxMiddleware)
    srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "PATCH", "HEAD"},
        AllowHeaders:     []string{"Accept", "Content-Type", "Authorization"},
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowCredentials: true,
    }))

    InitRouter(srv)
}

func NewDefaultServer(db *ent.Client) *Server {
    srv := NewEmptyServer(db)
    WrapServerWithDefaults(srv, db)

    return srv
}

func InitRouter(srv *Server) {
    registerControllerRoutes(srv,
        AuthController{},
        RoleController{},
        UserController{},
        APIKeyController{},
    )
}

func StartServer(srv *Server) {
    srv.Logger.Fatal(srv.Start(config.Config.Server.URI()))
}

func StartDefaultServer(db *ent.Client) {
    StartServer(NewDefaultServer(db))
}
