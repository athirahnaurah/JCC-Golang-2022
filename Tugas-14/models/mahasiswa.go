package models

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matkul"`
	IndeksNilai string `json:"indeks"`
	Nilai       uint64 `json:"nilai"`
	ID          uint64 `json:"id"`
}