1. Buatlah database seperti di function di bawah
    db = &DbConfig{
			host:         PriorityString(fang.GetString("db.host"), os.Getenv("DB_HOST"), "localhost"),
			port:         PriorityString(fang.GetString("db.port"), os.Getenv("DB_PORT"), "5432"),
			name:         PriorityString(fang.GetString("db.name"), os.Getenv("DB_NAME"), "postgres"),
			user:         PriorityString(fang.GetString("db.user"), os.Getenv("DB_USER"), "postgres"),
			pass:         PriorityString(fang.GetString("db.pass"), os.Getenv("DB_PASS"), "1234"),
			sslmode:      givenSSLMode,
			tz:           PriorityString(fang.GetString("db.tz"), os.Getenv("DB_TZ"), "UTC"),
			maxOpenConns: mOpenConns,
			maxIdleConns: mIdleConns,
			connLifetime: ltConn,
		}
    nameDB          postgres
    userDB          postgres
    passwordDB      1234

    atau bisa sesuai keinginan dengan catatan mengganti config tersebut

2. Jalankan go mod tidy
    otomatis akan terbuat table karena sudah menggunakan feature migration DB 

3. Export postman Collection
    lalu lakukan register -> login -> masukan token di header postmant lalu hit endpoint yg di inginkan