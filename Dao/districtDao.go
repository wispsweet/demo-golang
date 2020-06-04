package Dao

import "database/sql"

func DistrictIdByName(title string) int {
	var row *sql.Row

	link := Connect()
	defer link.Close()

	row = link.QueryRow("select id from vip_district where name = " + title + " limit 1")

	var id int
	err := row.Scan(&id)
	if err != nil && err != sql.ErrNoRows {
		return 0
	} else {
		row.Scan(&id)
		return id
	}
}
