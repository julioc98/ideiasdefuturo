package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/julioc98/ideiasdefuturo/internal/app"
	"github.com/julioc98/ideiasdefuturo/internal/domain"
	"github.com/julioc98/ideiasdefuturo/internal/infra/gateway"
	"github.com/julioc98/ideiasdefuturo/internal/infra/handler"
	"github.com/julioc98/ideiasdefuturo/internal/infra/repository"
	"github.com/julioc98/ideiasdefuturo/pkg/krypto"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/urfave/negroni"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Println("Start Ideias de Futuro API")

	dbURL := os.Getenv("DATABASE_URL")

	db, err := provideDB(dbURL, true)
	if err != nil {
		log.Panicf("database step %s", err.Error())
	}

	userRepo := repository.NeWUserGorm(db)

	userUseCase := app.NewUserUseCase(userRepo, &krypto.Hash{}, &gateway.Auth{}, validator.New(), &gateway.Console{})

	// go-guardian
	keeper := jwt.StaticSecret{
		ID:        "secret-id",
		Secret:    []byte("secret"),
		Algorithm: jwt.HS256,
	}

	guard := gateway.NewGuardian(keeper, userUseCase, &krypto.Hash{})

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
		negroni.HandlerFunc(guard.AuthMiddleware),
	)

	userHandler := handler.NewUserRestHandler(userUseCase, guard)
	userHandler.SetUserRoutes(r.PathPrefix("/users").Subrouter(), *n)

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	err = http.ListenAndServe(":"+os.Getenv("PORT"), r)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo a Ideias de Futuro API!!"
	log.Println(msg)
	_, _ = w.Write([]byte(msg))
}

func provideDB(dbURL string, migrate bool) (*gorm.DB, error) {
	dialector := postgres.Open(dbURL)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database err: %w", err)
	}

	if !migrate {
		return db, nil
	}

	if dbc := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`); dbc.Error != nil {
		return nil, fmt.Errorf("failed to migrate database err: %w", err)
	}

	err = db.AutoMigrate(domain.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database err: %w", err)
	}

	return db, nil
}
