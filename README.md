# webauthn-ra
Webauthn based Registration Authority (a'la X.509 RA)

### Build
```
cd _ui
npm run build
cd ..
go build
./webauthn-ra 
```
Then open http://localhost:3000/app


### Dev mode
```
go run main.go
cd _ui
npm run serve
```
Then open http://localhost:8080/dev

## Used technologies
- UI written in Vuetify with Typescript
- Built-in webserver with automatic certificate management
- Webauthn
- OpenAPI V3 (swaggest)
- GORM
- Code generation (Go->TypeScript)

## Coding style
### File names within a package
- `init.go` : application level Init() function or go level init() function
- `api.go` : Exported Go level API definitions
- `openapi.go` : Exported OpenAPI implementation, based on  [swaggest](http://github.com/swaggest) 
- `foo_test.go` : Black box tests. The test should use the `foo_test` package

### Line level conventions
- In tests use `t.Fatalf("%+v", err )` instead of `t.Fatal(err)` to get stack trace


## Data domains:
### Audit log
Append only

### Config
Transaction log, with signer info
Cumulative state hash
Time machine function
No anonymization support
Concurrent modification isn't permitted
Can't reference to other domain

### Generic data
Multiple version of a record
Anonymization support
Explicit previous record

## TODO
- List of last logins (list of issued JTWs)
- Rate limiting (distributed way), provide for  
    Provide backlog length for health monitoring and scale up/down service 
- Add request/response debug logger
- Add global error handling (`window.onerror`) forward to server 
- Introduce modules: Init()

## Useful docs:
REST API recommendation:
https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html
Icons:
https://jossef.github.io/material-design-icons-iconfont/

## Review of dependencies (TODO)
### GORM
### Chi router
### Ulid
### Koanf


