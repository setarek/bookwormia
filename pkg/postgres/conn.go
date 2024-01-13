package postgres

import (
	"bookwormia/pkg/config"
	"bookwormia/pkg/logger"
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Postgres struct {
	host     string
	port     string
	db       string
	user     string
	password string

	conn *sql.DB
}

func newPostgre(host, port, db, user, password string) *Postgres {
	return &Postgres{
		host:     host,
		port:     port,
		db:       db,
		user:     user,
		password: password,
	}
}

func (p *Postgres) ping() {
	ticker := time.NewTicker(2 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				pingContext, can := context.WithTimeout(context.Background(), 5*time.Second)
				defer can()
				err := p.conn.PingContext(pingContext)
				if err != nil {
					time.Sleep(5 * time.Second)
					go p.connect()
					<-quit
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func (p *Postgres) connect() {
	var err error
POSTGRESQLTRY:

	p.conn, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		p.host, p.port, p.user, p.password, p.db))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while opening database connection")
		goto POSTGRESQLTRY
	}

	err = p.conn.Ping()
	if err != nil {
		logger.Logger.Error().Err(err).Msg("error while pinging database connection")
		goto POSTGRESQLTRY
	}

	p.ping()
}

func InitDB(config *config.Config) (*sql.DB, error) {
	client := newPostgre(
		config.GetString("db_host"),
		config.GetString("db_port"),
		config.GetString("db_database"),
		config.GetString("db_user"),
		config.GetString("db_password"),
	)

	// todo: add timeout context parameter in config
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client.ping()
	client.connect()

	return client.conn, nil
}
