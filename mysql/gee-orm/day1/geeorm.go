package geeorm

import (
	"database/sql"
	"tryDemo/mysql/gee_orm/day1/log"
	"tryDemo/mysql/gee_orm/day1/session"
)

// 用户交互
type Engine struct {
	db *sql.DB
}

func NewEngine(drive, source string) (e *Engine, err error) {
	db, err := sql.Open(drive, source)
	if err != nil {
		log.Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}
	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine *Engine) NewSuccession() *session.Session {
	return session.New(engine.db)
}
