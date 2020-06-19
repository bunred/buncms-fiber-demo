module buncms

go 1.14

replace github.com/bunred/buncms => ../../buncms.com/buncms
replace github.com/99designs/gqlgen => github.com/arsmn/gqlgen v0.12.0

require (
	github.com/99designs/gqlgen v0.11.3
	github.com/bunred/buncms v1.0.1
	github.com/gofiber/fiber v1.11.1
	github.com/klauspost/compress v1.10.9 // indirect
	github.com/vektah/gqlparser/v2 v2.0.1
)
