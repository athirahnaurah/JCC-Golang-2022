package nilai

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
  table          = "nilai"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Nilai, error) {
  var nilai []models.Nilai
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
    var poin models.Nilai
    if err = rowQuery.Scan(&poin.Id,
      &poin.Skor,
	  &poin.Indeks,
    ); err != nil {
      return nil, err
    }

    if err != nil {
      log.Fatal(err)
    }

    nilai = append(nilai, poin)
  }
  return nilai, nil
}

// Insert Nilai
func Insert(ctx context.Context, nilai models.Nilai) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
  queryText := fmt.Sprintf("INSERT INTO %v (indeks, skor, mahasiswa_id, mata_kuliah_id) values('%v',%v,%v,%v)", table, nilai.Indeks, nilai.Skor,nilai.MahasiswaId,nilai.MataKuliahId)
  _, err = db.ExecContext(ctx, queryText)

  if err != nil {
    return err
  }
  return nil
}


// Update Nilai
func Update(ctx context.Context, nilai models.Nilai, id string) error {
  db, err := config.MySQL()
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }

  queryText := fmt.Sprintf("UPDATE %v set skor = %d where id = %s", table, nilai.Skor, id)

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