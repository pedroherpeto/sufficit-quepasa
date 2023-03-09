module github.com/sufficit/sufficit-quepasa/models

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/sufficit/sufficit-quepasa/whatsapp v0.0.0-00010101000000-000000000000
	github.com/sufficit/sufficit-quepasa/whatsmeow v0.0.0-00010101000000-000000000000
)

require (
	github.com/Rhymen/go-whatsapp v0.1.1 // indirect
	github.com/golang/protobuf v1.5.0 // indirect
	github.com/joncalhoun/migrate v0.0.2 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/sufficit/sufficit-quepasa/library v0.0.0-00010101000000-000000000000 // indirect
)

require (
	filippo.io/edwards25519 v1.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-chi/jwtauth v4.0.4+incompatible
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/jinzhu/copier v0.3.5
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.5.2
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/skip2/go-qrcode v0.0.0-20191027152451-9434209cb086
	github.com/stretchr/testify v1.4.0 // indirect
	go.mau.fi/libsignal v0.1.0 // indirect
	go.mau.fi/whatsmeow v0.0.0-20230204181151-b1f00ea99464 // indirect
	golang.org/x/crypto v0.5.0
	golang.org/x/sys v0.4.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace github.com/sufficit/sufficit-quepasa/library => ../library

replace github.com/sufficit/sufficit-quepasa/whatsmeow => ../whatsmeow

replace github.com/sufficit/sufficit-quepasa/whatsapp => ../whatsapp

replace github.com/sufficit/sufficit-quepasa/models => ./

go 1.18
