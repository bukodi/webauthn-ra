module github.com/bukodi/webauthn-ra

go 1.17

require (
	github.com/go-chi/chi/v5 v5.0.7
	github.com/swaggest/rest v0.2.18
	github.com/swaggest/swgui v1.4.3
	github.com/swaggest/usecase v1.1.0
)

require (
	github.com/fxamacker/webauthn v0.6.1
	github.com/rs/cors v1.8.2
	gorm.io/driver/sqlite v1.2.6
	gorm.io/gorm v1.22.5
)

replace github.com/fxamacker/webauthn => github.com/bukodi/webauthn v0.6.2-0.20220120090524-e5db2cd7d66d
