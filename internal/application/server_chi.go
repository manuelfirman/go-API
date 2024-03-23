package application

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	handler "github.com/manuelfirman/go-API/internal/handler/chi"
	"github.com/manuelfirman/go-API/internal/repository"
	"github.com/manuelfirman/go-API/internal/service"
)

// ConfigServer is the configuration for the server
type ConfigServer struct {
	// Addr is the address to listen on
	Addr string
	// MySQLDSN is the DSN for the MySQL database
	MySQLDSN string
}

// New creates a new instance of the server
func New(cfg ConfigServer) *ServerChi {
	// default config
	defaultCfg := ConfigServer{
		Addr:     ":8080",
		MySQLDSN: "",
	}
	if cfg.Addr != "" {
		defaultCfg.Addr = cfg.Addr
	}
	if cfg.MySQLDSN != "" {
		defaultCfg.MySQLDSN = cfg.MySQLDSN
	}

	return &ServerChi{
		addr:     defaultCfg.Addr,
		mysqlDSN: defaultCfg.MySQLDSN,
	}
}

// ServerChi is the default implementation of the server
type ServerChi struct {
	// addr is the address to listen on
	addr string
	// mysqlDSN is the DSN for the MySQL database
	mysqlDSN string
}

// Run runs the server
func (s *ServerChi) Run() (err error) {
	// dependencies
	// - database: connection
	db, err := sql.Open("mysql", s.mysqlDSN)
	if err != nil {
		return
	}
	defer db.Close()
	// - database: ping
	err = db.Ping()
	if err != nil {
		return
	}

	// - router
	router := chi.NewRouter()
	//   middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// endpoints

	// - products
	buildProductsRouter(router, db)
	// - sellers
	buildSellersRouter(router, db)

	// run
	err = http.ListenAndServe(s.addr, router)
	return
}

// *buildProductsRouter builds the router for the products endpoints
func buildProductsRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewProductMySQL(db)
	sv := service.NewProductDefault(rp)
	hd := handler.NewProductDefault(sv)

	// define the routes of the products
	router.Route("/api/v1/products", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Create())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetByID())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})
}

// *buildSellersRouter builds the router for the sellers endpoints
func buildSellersRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewSellerMySQL(db)
	sv := service.NewSellerDefault(rp)
	hd := handler.NewSellerDefault(sv)

	// define the routes of the sellers
	router.Route("/api/v1/sellers", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Create())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.GetByID())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})
}
