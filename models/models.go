package models

import (
   "fmt"
  	"belajargolang/db"
    "errors"
)

type Rental struct {
	Id          string `json:"id"`
	Brand       string `json:"brand"`
	Year        int    `json:"year"`
	OwnerId     string `json:"owner_id"`
	RentPrice   int    `json:"rent_per_hour"`
	IsAvailable int    `json:"availability"`
}

func CekData(id string) (Rental, error){
  var rental = Rental{}
  conn, err := db.Connect()
	if err != nil {
		return rental, err
	}
	defer conn.Close()

  sql := "SELECT * FROM rental WHERE id="+id

  rows, err := conn.Query(sql)
  if err != nil {
    return rental, err
  }
  defer rows.Close()


  for rows.Next() {
    var errGet = rows.Scan(&rental.Id, &rental.Brand, &rental.Year, &rental.OwnerId, &rental.RentPrice, &rental.IsAvailable)
    if errGet != nil {
      return rental, errGet
    }
  }

  // rentalPtr := &rental

  if rental.Id == "" {
    return rental, errors.New("Data tidak ada")
  }


  return rental, nil
}

func TampilData() ([]Rental, error) {
  var data []Rental
	conn, err := db.Connect()
	if err != nil {
		return data, err
	}
	defer conn.Close()

	sql := "SELECT * FROM rental "

	rows, err := conn.Query(sql)
	if err != nil {
		return data, err
	}

	defer rows.Close()

	for rows.Next() {
		var each = Rental{}
		var err = rows.Scan(&each.Id, &each.Brand, &each.Year, &each.OwnerId, &each.RentPrice, &each.IsAvailable)
		if err != nil {
			return data, err
		}
		data = append(data, each)
	}
	if err := rows.Err(); err != nil {
		return data, err
	}
  return data, nil
}

func CariData(key string, value string) ([]Rental, error) {
  var data []Rental
	conn, err := db.Connect()
	if err != nil {
		return data, err
	}
	defer conn.Close()

	sql := fmt.Sprintf("SELECT * FROM rental WHERE %s='%s'", key, value)
	rows, err := conn.Query(sql)
	if err != nil {
		return data, err
	}
	defer rows.Close()
  fmt.Println(sql)

	for rows.Next() {
		var each = Rental{}
		var err = rows.Scan(&each.Id, &each.Brand, &each.Year, &each.OwnerId, &each.RentPrice, &each.IsAvailable)
		if err != nil {
			return data, err
		}
		data = append(data, each)
	}
	if err := rows.Err(); err != nil {
		return data, err
	}
  return data, nil
}

func UpdateData (r Rental, value string) (string, error) {

  if r.Brand == value {
    return "Error", errors.New("Perbarui dibatalkan, input sama")
  }

  conn, err := db.Connect()
	if err != nil {
		return "DB Error", err
	}
	defer conn.Close()

  sqlUpd := "UPDATE rental SET brand = ? WHERE id = ?"
  // DB UPDATE
  _, err = conn.Exec(sqlUpd, value, r.Id)
  if err != nil {
    return "Update Error", err
  }

  response := fmt.Sprintf("Perbarui Brand '%v' Berhasil dengan Nama '%v' dari nomor id: %v", r.Brand, value, r.Id)
  return response , nil
}

func DeleteData (r Rental) (string, error) {
  conn, err := db.Connect()
	if err != nil {
		return "DB Error", err
	}
	defer conn.Close()

  sqlUpd := "DELETE FROM rental WHERE id = ?"
  // DB UPDATE
  _, err = conn.Exec(sqlUpd, r.Id)
  if err != nil {
    return "Delete Error", err
  }

  response := fmt.Sprintf("Hapus Data Mobil Rental %v dengan ID pemilik %v Berhasil ", r.Brand, r.OwnerId)
  return response , nil
}
