module github.com/bukodi/webauthn-ra

go 1.20

require (
	github.com/bukodi/go-pkcs12 v0.0.0-20230406152834-f6d22e645fe1
	github.com/go-chi/chi/v5 v5.0.7
	github.com/swaggest/rest v0.2.18
	github.com/swaggest/swgui v1.4.3
	github.com/swaggest/usecase v1.1.0
)

require (
	github.com/elgopher/yala v0.20.0
	github.com/fxamacker/webauthn v0.6.1
	github.com/knadh/koanf v1.4.0
	github.com/oklog/ulid/v2 v2.0.2
	github.com/pulumi/pulumi-awsx/sdk v1.0.2
	github.com/pulumi/pulumi-eks/sdk v1.0.1
	github.com/pulumi/pulumi/sdk/v3 v3.58.0
	github.com/rs/cors v1.8.2
	gorm.io/driver/sqlite v1.2.6
	gorm.io/gorm v1.22.5
)

replace github.com/fxamacker/webauthn => github.com/bukodi/webauthn v0.6.2-0.20220120090524-e5db2cd7d66d
