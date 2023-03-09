module github.com/sufficit/sufficit-quepasa/whatsmeow

require github.com/sufficit/sufficit-quepasa/whatsapp v0.0.0-00010101000000-000000000000

require (
	github.com/gosimple/slug v1.13.1 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.11
	github.com/sirupsen/logrus v1.8.1
	go.mau.fi/libsignal v0.1.0 // indirect
	go.mau.fi/whatsmeow v0.0.0-20230204181151-b1f00ea99464
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/sufficit/sufficit-quepasa/whatsmeow => ./

replace github.com/sufficit/sufficit-quepasa/whatsapp => ../whatsapp

go 1.18
