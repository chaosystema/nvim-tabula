package database

import (
	"errors"

	"github.com/javiorfo/nvim-tabula/go/database/engine"
	"github.com/javiorfo/nvim-tabula/go/database/engine/model"
)

type Executor interface {
    // TODO return error
	Run() 
    GetTables()
    GetTableInfo()
}

func Context(eng string, option model.Option, data model.Data) error {
	switch eng {
	case engine.POSTGRES:
		return run(engine.Postgres{Data: data}, option)
	case engine.MONGO:
		return run(engine.Mongo{Data: data}, option)
	case engine.MYSQL:
		return run(engine.MySql{Data: data}, option)
	default:
		return errors.New("engine does not exist")
	}
}

func run(executor Executor, option model.Option) error {
    switch option {
    case model.RUN:
        executor.Run()
        return nil
    case model.TABLES:
        executor.GetTables()
        return nil
    case model.TABLE_INFO:
        executor.GetTableInfo()
        return nil
    default:
	    return errors.New("Option does not exist")
    }
}
