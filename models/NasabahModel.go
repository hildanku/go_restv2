package models

type Nasabah struct {
	Id            uint   `json:"id"`
	NamaNasabah   string `json:"nama_nasabah"`
	AlamatNasabah string `json:"alamat_nasabah"`
	SaldoNasabah  int    `json:"saldo_nasabah"`
	StatusNasabah string `json:"status_nasabah"`
}

func (b *Nasabah) TableName() string {
	return "Nasabah"
}
