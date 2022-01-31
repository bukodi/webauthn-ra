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
- Built-in webserver
- Webauthn
- OpenAPI V3 (swaggest)
- GORM

## Data domains:
### Audit log
Append only

### Config
Transaction log, with signer info
Cumulative state hash
Time machine function
No anonymization support
Rollback on concurrent modification

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

## Useful docs:
REST API recommendation:
https://fidoalliance.org/specs/fido-v2.0-rd-20180702/fido-server-v2.0-rd-20180702.html
Icons:
https://jossef.github.io/material-design-icons-iconfont/


