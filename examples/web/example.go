package main

import (
	"database/sql"
	"fmt"

	"github.com/tombuente/ee"
)

type db struct{}
type service struct {
	db db
}
type httpHandler struct {
	service service
}

func (db db) getItem(id int) (string, error) {
	if id >= 100 {
		return "", ee.NewSQLError(ee.SQLInternal, sql.ErrTxDone)
	}
	if id >= 10 {
		return "", ee.NewSQLError(ee.SQLNotFound, sql.ErrNoRows)
	}

	return fmt.Sprintf("item%v", id), nil
}

func (s service) getItem(id int) (string, error) {
	item, err := s.db.getItem(id)
	if err != nil {
		switch ee.UnpackErrKind(err) {
		case ee.SQLNotFound:
			return "", ee.NewError(ee.NotFound, err)
		default:
			return "", ee.NewError(ee.Internal, err)
		}
	}

	return item, nil
}

func (h httpHandler) getItem(id int) string {
	item, err := h.service.getItem(id)
	if err != nil {
		switch ee.UnpackErrKind(err) {
		case ee.NotFound:
			return "item was not found"
		default:
			return "internal server error"
		}
	}

	return item
}

func main() {
	db := db{}
	service := service{db: db}
	httpHandler := httpHandler{service: service}

	fmt.Println(httpHandler.getItem(1))
	fmt.Println(httpHandler.getItem(10))
	fmt.Println(httpHandler.getItem(100))
}
