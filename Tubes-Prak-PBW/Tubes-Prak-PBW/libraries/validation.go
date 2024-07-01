package libraries

import (
	"Tubes_PBW/config"
	"database/sql"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type Validation struct {
	conn *sql.DB
}

func NewValidation() *Validation {
	conn, err := config.DBconnection()

	if err != nil {
		panic(err)
	}

	return &Validation{
		conn: conn,
	}
}

func (v *Validation) Init() (*validator.Validate, ut.Translator) {
	// memanggil package translator
	translator := en.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("en")

	validate := validator.New()

	// register default translation (en)
	en_translations.RegisterDefaultTranslations(validate, trans)

	// mengubah label default
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		labelName := field.Tag.Get("label")
		return labelName
	})

	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} tidak boleh kosong", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterValidation("isunique", func(fl validator.FieldLevel) bool {
		params := fl.Param()
		split_params := strings.Split(params, "-")

		tableName := split_params[0]
		fieldName := split_params[1]
		fieldValue := fl.Field().String()

		isEdit := fl.Parent().FieldByName("Id").Int() > 0 

		if isEdit {
			return true
		}

		return v.checkIsUnique(tableName, fieldName, fieldValue)
	})

	validate.RegisterTranslation("isunique", trans, func(ut ut.Translator) error {
		return ut.Add("isunique", "{0} sudah digunakan", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("isunique", fe.Field())
		return t
	})

	validate.RegisterTranslation("eqfield", trans, func(ut ut.Translator) error {
		return ut.Add("eqfield", "{0} Harus Sama Dengan Password", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("eqfield", fe.Field())
		return t
	})

	validate.RegisterTranslation("len", trans, func(ut ut.Translator) error {
		return ut.Add("len", "{0} harus memiliki 16 karakter", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("len", fe.Field())
		return t
	})

	validate.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "{0} harus memiliki setidaknya 4 karakter", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.Field())
		return t
	})

	validate.RegisterValidation("dategte", func(fl validator.FieldLevel) bool {
		tanggalSewa := fl.Top().FieldByName("TanggalSewa").String()
		tanggalKembali := fl.Field().String()

		layout := "2006-01-02"

		tSewa, err := time.Parse(layout, tanggalSewa)
		if err != nil {
			return false
		}

		tKembali, err := time.Parse(layout, tanggalKembali)
		if err != nil {
			return false
		}

		return tKembali.After(tSewa)
	})

	validate.RegisterTranslation("dategte", trans, func(ut ut.Translator) error {
		return ut.Add("dategte", "{0} Tidak bisa di pilih", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("dategte", fe.Field())
		return t
	})

	validate.RegisterValidation("dategteToday", func(fl validator.FieldLevel) bool {
		tanggalSewa := fl.Field().String()

		layout := "2006-01-02" // Sesuaikan format tanggal sesuai dengan input Anda

		tSewa, err := time.Parse(layout, tanggalSewa)
		if err != nil {
			return false
		}

		today := time.Now().Format(layout)
		tToday, err := time.Parse(layout, today)
		if err != nil {
			return false
		}

		return !tSewa.Before(tToday)
	})

	validate.RegisterTranslation("dategteToday", trans, func(ut ut.Translator) error {
		return ut.Add("dategteToday", "{0} Tidak bisa dipilih", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("dategteToday", fe.Field())
		return t
	})




	return validate, trans
}

func (v *Validation) Struct(s interface{}) interface{} {

	validate, trans := v.Init()

	vErrors := make(map[string]interface{})

	err := validate.Struct(s)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			vErrors[e.StructField()] = e.Translate(trans)
		}
	}

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil

}

func (v *Validation) checkIsUnique(tableName, fieldName, fieldValue string) bool {

	row, _ := v.conn.Query("select "+fieldName+" from "+tableName+" where "+fieldName+" = ?", fieldValue)

	defer row.Close()

	var result string
	for row.Next() {
		row.Scan(&result)
	}

	return result != fieldValue
}