module github.com/enaldo1709/budget-manager/app

go 1.21.1

replace github.com/enaldo1709/budget-manager/domain/model => ../domain/model

replace github.com/enaldo1709/budget-manager/domain/usecase => ../domain/usecase

replace github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api => ../infrastructure/entry-points/web-api

replace github.com/enaldo1709/budget-manager/infrastructure/helpers/configutil => ../infrastructure/helpers/configutil

replace github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter => ../infrastructure/adapters/postgresql-adapter

require github.com/enaldo1709/budget-manager/domain/model v0.0.0-20230322031830-6842ea4aeb05

require (
	github.com/bytedance/sonic v1.10.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/chenzhuoyu/iasm v0.9.0 // indirect
	github.com/enaldo1709/budget-manager/domain/usecase v0.0.0-20230322031830-6842ea4aeb05 // indirect
	github.com/enaldo1709/budget-manager/infrastructure/adapters/postgresql-adapter v0.0.0-00010101000000-000000000000 // indirect
	github.com/enaldo1709/budget-manager/infrastructure/entry-points/web-api v0.0.0-00010101000000-000000000000 // indirect
	github.com/enaldo1709/budget-manager/infrastructure/helpers/configutil v0.0.0-00010101000000-000000000000 // indirect
	github.com/gabriel-vasile/mimetype v1.4.2 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.15.4 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/leodido/go-urn v1.2.4 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.1.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.11 // indirect
	golang.org/x/arch v0.5.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
