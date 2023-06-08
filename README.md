# Workshop FTGO: REST API & TDD

## Init
### install packages
Pada tahap ini kita perlu melakukan instalasi pada package yang akan digunakan dalam project, yaitu Gorm beserta dengan driver postgres, Gin sebagai web framework untuk mempercepat pembuatan web service, dan testify untuk membantu proses assertion dalam test
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

go get -u github.com/gin-gonic/gin
go get -u github.com/stretchr/testify
```

### Init Database Configuration
buatlah sebuah folder `config`, dan file `database.go` di dalamnya, pada file ini kita akan menginisiasi koneksi ke database menggunakan gorm
```go
package config

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database!")
	}
	return db
}
```
Nanti function InitDatabase ini akan kita panggil saat menjalankan aplikasi dengan database untuk aplikasi kita, dan saat menjalankan testing dengan database yang khusus kita peruntukan untuk melakukan testing

### Init app configuration
pada tahap ini kita akan menginisasi app engine kita, sehingga nantinya bisa fleksibel digunakan baik untuk menjalankan app nya, atau bisa juga digunakan saat menjalankan file testing nya. 

buatlah folder `routes`, dan file `main.go`
```go
package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	app *gin.Engine
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	app = gin.Default()
	// your app endpoints here
	return app
}
```

dan file `main.go` sebagai pintu gerbang menjalankan aplikasi kita,
```go
package main

import (
	"workshoptdd/config"
	"workshoptdd/routes"
)

func main() {
	db := config.InitDatabase("your_db_dsn_here")
	app := routes.InitRoutes(db)
	app.Run(":8000")
}
```

Bisa terlihat ketika kita memanggil `InitDatabase` kita bisa memasukan `DSN` database kita yang digunakan untuk proses development / menjalankan aplikasi. Begitu pula nanti ketika kita menginisasi testing, kita bisa memasukan `DSN` khusus untuk database yang kita peruntukan untuk keperluan testing saja, sehingga tidak menganggu database aplikasi utama kita

## Testing
### Init test main
Untuk mempermudah proses test, kita bisa membuat sebuah folder `test` yang nanti akan berisi file-file test kita. Pada project kali ini kita akan membuat sebuah file test, dengan nama `app_test.go` 
Dan kita perlu menginisiasi database, dan app yang akan kita gunakan dalam testing. Kita bisa memanfaakan function TestMain untuk proses ini. 
Referensi TestMain https://medium.com/@aldinofrizal/golang-testing-lifecycle-hack-httptest-example-3a96dc8181ed
```go
package test

func TestMain(m *testing.M) {
	db = config.InitDatabase("your_db_test")
	app = routes.InitRoutes(db)

	m.Run()
}
```

### First TestCase
Selanjutnya kita sudah bisa membuat test cases untuk TDD kita, berikut adalah contoh sederhana
```go
func TestHealthCheck(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)
	
	app.ServeHTTP(w, req)
	
	response := w.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, _ := io.ReadAll(response.Body)

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "OK", responseBody["message"])
}
```

Untuk menjalankan test bisa menggunakan
```bash
go test -v ./...
```

Selanjutnya setelah kita berhasil menjalankan test nya, yang tentu nya masih akan error karena kita belum menyiapkan endpoint `/healthcheck`. Kita perlu membuat endpoint /healthcheck dan menyesuaikan request dan ataupun response nya agar memenuhi test case yang sudah kita buat.

Selanjutnya untuk membuat testing nya passed atau berhasil, kita bisa menambahkan sebuah endpoint baru di `routes/main.go`
```go
app.GET("/healthcheck", func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
})
```

Dan kita bisa kembali menjalankan test nya, jika masih terdapat error, maka kita kembali menyesuaikan code fitur kita, atau pada kasus ini adalah endpoint `/healthcheck` dan jalankan kembali testnya sampai berhasil.

Jika sudah berhasil, maka kita bisa melanjutkan untuk membuat fitur atau test case berikutnya, dilanjutkan dengan membuat code fitur nya.
