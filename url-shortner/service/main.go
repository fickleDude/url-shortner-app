package services

import "database/sql"

var db *sql.DB

type Models struct {
	Airports     Airports
	JsonResponse JsonResponse
}

func New(dbPool *sql.DB) Models {
	db = dbPool
	return Models{}
}
