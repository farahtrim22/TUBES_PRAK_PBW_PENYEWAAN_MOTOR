package entities

type Customer struct {
	Id           int64
	NamaLengkap  string `validate:"required" label:"Nama Lengkap"`
	NIK          string `validate:"required,len=16,isunique=customer-nik" isedit:"true"`
	JenisKelamin string `validate:"required" label:"Jenis Kelamin"`
	Alamat       string `validate:"required"`
	NomorTelepon string `validate:"required" label:"Nomor Telepon"`
}

type User struct {
	Id          int64
	NamaLengkap string `validate:"required" label:"Nama Lengkap"`
	Email       string `validate:"required,email,isunique=user-email"`
	Username    string `validate:"required,gte=4,isunique=user-username"`
	Password    string `validate:"required"`
	Cpassword   string `validate:"required,eqfield=Password" label:"Konfirmasi Password"`
}

type Motor struct {
	Id            int64
	Merek         string `validate:"required" label:"Merek"`
	Tipe          string `validate:"required" label:"Tipe Motor"`
	JenisMotor    string `validate:"required" label:"Jenis Motor"`
	TahunProduksi string `validate:"required" label:"Tahun Produksi"`
	Warna         string `validate:"required" label:"Warna"`
	Stok          string `validate:"required" label:"Stok"`
	Harga         string `validate:"required" label:"Harga"`
}

type Sewa struct {
	Id             int64
	CustomerId     int64 `validate:"required" label:"Customer"`
	MotorId        int64 `validate:"required" label:"Motor"`
	CustomerName   string
	TipeMotor      string
	TanggalSewa    string `validate:"required,dategteToday" label:"Tanggal Sewa"`
	TanggalKembali string `validate:"required,dategte" label:"Tanggal Pengembalian"`
}
