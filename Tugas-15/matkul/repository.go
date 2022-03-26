package matkul

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
  table          = "mata_kuliah"
)
 
// GetAll Matkul
func GetAll(ctx context.Context) ([]models.MataKuliah, error) {
 
  var mataKuliah []models.MataKuliah
 
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
    var mk models.MataKuliah
    if err = rowQuery.Scan(
		&mk.Id,
      	&mk.Nama,
     ); err != nil {
      return nil, err
    }
 
    mataKuliah = append(mataKuliah, mk)
  }
  return mataKuliah, nil
}
 
// Insert MataKuliah
func Insert(ctx context.Context, mk models.MataKuliah) error {
  db, err := config.MySQL()
 
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
 
  queryText := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", table,mk.Nama)
 
  _, err = db.ExecContext(ctx, queryText)
 
  if err != nil {
    return err
  }
  return nil
}
 
// Update Matkul
func Update(ctx context.Context, mk models.MataKuliah, id string) error {
 
  db, err := config.MySQL()
 
  if err != nil {
    log.Fatal("Can't connect to MySQL", err)
  }
 
  queryText := fmt.Sprintf("UPDATE %v set nama ='%s' where id = %s",table, mk.Nama, id)
  fmt.Println(queryText)
 
  _, err = db.ExecContext(ctx, queryText)
 
  if err != nil {
    return err
  }
 
  return nil
}
 
// Delete Matkul
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