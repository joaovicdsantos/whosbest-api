package app

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/joaovicdsantos/whosbest-api/app/handler"
)

type App struct {
	mux *http.ServeMux
}

func (a *App) Initialize(db *sql.DB) {
	a.mux = http.NewServeMux()
	a.SetRoutes(db)
}

func (a *App) SetRoutes(db *sql.DB) {
	// User routes
	userRoutes := new(handler.UserRoutes)
	userRoutes.DB = db
	a.mux.Handle("/user", userRoutes)
	a.mux.Handle("/user/", userRoutes)
	a.mux.HandleFunc("/login", userRoutes.Login)

	// Leaderboard routes
	leaderboardRoutes := new(handler.LeaderboardRoutes)
	leaderboardRoutes.DB = db
	a.mux.Handle("/leaderboard", leaderboardRoutes)
	a.mux.Handle("/leaderboard/", leaderboardRoutes)

	// Competitor routes
	competitorRoutes := new(handler.CompetitorRoutes)
	competitorRoutes.DB = db
	a.mux.Handle("/competitor", competitorRoutes)
	a.mux.Handle("/competitor/", competitorRoutes)
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(":3001", a.mux))
}
