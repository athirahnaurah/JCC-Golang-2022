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
  layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.NilaiMahasiswa, error) {
  var mahasiswa []models.NilaiMahasiswa
  db, err := config.MySQL()

  if err != nil {
    log.Fatal("Cant connect to MySQL", err)
  }

  queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
  rowQuery, err := db.QueryContext(ctx, queryText)

  if err != nil {
    log.Fatal(err)
  }

  for rowQuery.Next() {
    var mhs models.NilaiMahasiswa
    if err = rowQuery.Scan(&mhs.ID,
      &mhs.Nama,
      &mhs.MataKuliah,
      &mhs.Nilai,
      &mhs.IndeksNilai,
     ); err != nil {
      return nil, err
    }

    //  Change format string to datetime for created_at and updated_at
    // mhs.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

    if err != nil {
      log.Fatal(err)
    }

    // movie.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
    // if err != nil {
    //   log.Fatal(err)
    // }

    mahasiswa = append(mahasiswa, mhs)
  }
  return mahasiswa, nil
}

// Insert Mahasiswa
func Insert(ctx context.Context, mhs models.NilaiMahasiswa) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  
  queryText := fmt.Sprintf("INSERT INTO %v (Nama, MataKuliah, Nilai) values('%v','%v',%v,)", table,
    mhs.Nama,
    mhs.MataKuliah,
    mhs.Nilai)
  _, err = db.ExecContext(ctx, queryText)

  if err != nil {
    return err
  }
  return nil
}

// Update Movie
func Update(ctx context.Context, mhs models.NilaiMahasiswa, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }

  queryText := fmt.Sprintf("UPDATE %v set Nama ='%s', Nilai = %d, MataKuliah = '%s', where id = %s",
    table,
    mhs.Nama,
    mhs.Nilai,
    mhs.MataKuliah,
    id,
  )

  _, err = db.ExecContext(ctx, queryText)
  if err != nil {
    return err
  }

  return nil
}

// Delete Movie
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