package mahasiswa

import (
	"Tugas14/config"
	"Tugas14/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
  table          = "mahasiswa"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
  var mahasiswa []models.Mahasiswa
  db, err := config.MySQL()

  if err != nil {
    log.Fatal("Cant connect to MySQL", err)
  }

  queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
  rowQuery, err := db.QueryContext(ctx, queryText)

  if err != nil {
    log.Fatal(err)
  }

  for rowQuery.Next() {
    var mhs models.Mahasiswa
    if err = rowQuery.Scan(&mhs.Id,
      &mhs.Nama,
    ); err != nil {
      return nil, err
    }

    if err != nil {
      log.Fatal(err)
    }

    mahasiswa = append(mahasiswa, mhs)
  }
  return mahasiswa, nil
}

// Insert Mahasiswa
func Insert(ctx context.Context, mhs models.Mahasiswa) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  
  queryText := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", table, mhs.Nama)
  _, err = db.ExecContext(ctx, queryText)

  if err != nil {
    return err
  }
  return nil
}

// Update Mahasiswa
func Update(ctx context.Context, mhs models.Mahasiswa, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }

  queryText := fmt.Sprintf("UPDATE %v set nama ='%s' where id = %s", table, mhs.Nama, id)

  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }

  return nil
}

// Delete Mahasiswa
func Delete(ctx context.Context, id string) error {
  db, err := config.MySQL()
  if err != nil {
      log.Fatal("Can't connect to MySQL", err)
  }

  queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

  s, err := db.ExecContext(ctx, queryText)

  if err != nil && err != sql.ErrNoRows {
      return err
  }

  check, err := s.RowsAffected()
  fmt.Println(check)
  if check == 0 {
      return errors.New("id tidak ada")
  }

  if err != nil {
      fmt.Println(err.Error())
  }

  return nil
}