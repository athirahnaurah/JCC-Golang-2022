package models

type Nilai struct {
	Id           uint64 `json:"id"`
	Indeks       string `json:"indeks"`
	Skor         uint64 `json:"skor"`
	MahasiswaId  uint64 `json:"mahasiswa_id"`
	MataKuliahId uint64 `json:"mk_id"`
}