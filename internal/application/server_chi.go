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
	// - middlewares
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	// - ping endpoint
	buildPing(router)

	// endpoints
	// - products
	buildProductsRouter(router, db)
	// - sellers
	buildSellersRouter(router, db)
	// - buyers
	buildBuyersRouter(router, db)
	// - warehouses
	buildWarehousesRouter(router, db)
	// - employees
	buildEmployeesRouter(router, db)
	// - sections
	buildSectionsRouter(router, db)

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

// *buildBuyersRouter builds the router for the buyers endpoints
func buildBuyersRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewBuyerMySQL(db)
	sv := service.NewBuyerDefault(rp)
	hd := handler.NewBuyerDefault(sv)

	// define the routes of the buyers
	router.Route("/api/v1/buyers", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Save())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.Get())
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

// *buildWarehousesRouter builds the router for the warehouses endpoints
func buildWarehousesRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewWarehouseMySQL(db)
	sv := service.NewWarehouseDefault(rp)
	hd := handler.NewWarehouseDefault(sv)

	// define the routes of the warehouses
	router.Route("/api/v1/warehouses", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Save())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.Get())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})
}

// *buildEmployeesRouter builds the router for the employees endpoints
func buildEmployeesRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewEmployeeMySQL(db)
	sv := service.NewEmployeeDefault(rp)
	hd := handler.NewEmployeeDefault(sv)

	// define the routes of the employees
	router.Route("/api/v1/employees", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Save())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.Get())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})
}

// *buildSectionsRouter builds the router for the sections endpoints
func buildSectionsRouter(router *chi.Mux, db *sql.DB) {
	// instance dependences
	rp := repository.NewSectionMySQL(db)
	sv := service.NewSectionDefault(rp)
	hd := handler.NewSectionDefault(sv)

	// define the routes of the sections
	router.Route("/api/v1/sections", func(r chi.Router) {
		// endpoints
		r.Post("/", hd.Save())
		r.Get("/", hd.GetAll())
		r.Get("/{id}", hd.Get())
		r.Patch("/{id}", hd.Update())
		r.Delete("/{id}", hd.Delete())
	})
}

func buildPing(router *chi.Mux) {
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
}
