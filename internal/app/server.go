package app

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/kimanimichael/mk-device-manager/internal/adapters/database/sqlc/gensql"
	"github.com/kimanimichael/mk-device-manager/internal/devices"
	"github.com/kimanimichael/mk-device-manager/internal/devices/api"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func NewServer() *http.Server {
	loadConfigs()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in this environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database: ", dbURL)

	db := sqlcdatabase.New(conn)
	deviceRepoSQL := devices.NewDeviceRepositorySQL(db)

	deviceService := devices.NewDeviceService(deviceRepoSQL)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	}))

	deviceHandler := devicesapi.NewDeviceHandler(deviceService)
	deviceHandler.RegisterRoutes(router)

	actualRouter := chi.NewRouter()
	actualRouter.Mount("/mk", router)

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("couldn't find a port in this environment")
	}

	srv := &http.Server{
		Handler: actualRouter,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %s", portString)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	return srv
}

func loadConfigs() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
}
