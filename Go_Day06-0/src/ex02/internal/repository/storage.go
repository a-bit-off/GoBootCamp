package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"net"
	"time"
)

// подключение к базе данных по ссылке
func InitDbConnect(ctx context.Context, dbURL string) (dbpool *pgxpool.Pool, err error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return
	}
	// настройка конфигурации
	cfg.MaxConns = int32(5)                         // максимальное количество соединений
	cfg.MinConns = int32(1)                         // минимальное количество соединений
	cfg.HealthCheckPeriod = 1 * time.Minute         // период проверки состояния соединений
	cfg.MaxConnLifetime = 24 * time.Hour            // время жизни соединения
	cfg.MaxConnIdleTime = 30 * time.Minute          // время простоя соединения
	cfg.ConnConfig.ConnectTimeout = 1 * time.Second // таймаут соединени
	cfg.ConnConfig.DialFunc = (&net.Dialer{         // функция обработки набора соединений
		KeepAlive: cfg.HealthCheckPeriod,
		Timeout:   cfg.ConnConfig.ConnectTimeout}).DialContext

	dbpool, err = pgxpool.ConnectConfig(ctx, cfg) // создает пул подключений на основе конфигурации
	if err != nil {
		return
	}
	return
}
