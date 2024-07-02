package entities

type Customer struct {
	Id 				int64
	NamaLengkap 	string `validate:"required" label:"Nama Lengkap"`
	NIK 			string `validate:"required"`
	JenisKelamin 	string `validate:"required" label:"Jenis Kelamin"`
	Alamat 			string `validate:"required"`
	NomorTelepon 	string `validate:"required" label:"Nomor Telepon"`
}

type User struct {
	Id				int64
	NamaLengkap 	string
	Email 			string
	Username 		string
	Password		string
}	

type Motor struct
{
	Id 				int64
	Merek 			string `validate:"required" label:"Merek"`
	Tipe 			string `validate:"required" label:"Tipe Motor"`
	JenisMotor		string `validate:"required" label:"Jenis Motor"`
	TahunProduksi	string `validate:"required" label:"Tahun Produksi"`
	Warna			string `validate:"required" label:"Warna"`
	Stok			string `validate:"required" label:"Stok"`
	Harga			string `validate:"required" label:"Harga"`
}