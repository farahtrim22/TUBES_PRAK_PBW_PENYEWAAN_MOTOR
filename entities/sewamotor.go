package entities

type Customer struct {
	Id 				int64
	NamaLengkap 	string `validate:"required" label:"Nama Lengkap"`
	NIK 			string `validate:"required"`
	JenisKelamin 	string `validate:"required" label:"Jenis Kelamin"`
	Alamat 			string `validate:"required"`
	NomorTelepon 	string `validate:"required" label:"Nomor Telepon"`
}