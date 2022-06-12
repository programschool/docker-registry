module github.com/programschool/docker-registry/middleware

go 1.18

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/labstack/echo/v4 v4.7.2
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/net v0.0.0-20220607020251-c690dde0001d // indirect
	golang.org/x/sys v0.0.0-20220610221304-9f5ed59c137d // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858 // indirect
)

require (
	github.com/programschool/docker-registry/config v0.0.0-incompatible
	github.com/programschool/docker-registry/library v0.0.0-incompatible
)

replace (
	github.com/programschool/docker-registry/config => ../config
	github.com/programschool/docker-registry/library => ../library
)
