# webauthn-ra
Webauthn based Registration Authority (a'la X.509 RA)

## Build
```
cd _ui
npm run build
cd ..
go build
./webauthn-ra 
```
Then open http://localhost:3000/app


## Dev mode
```
go run main.go
cd _ui
npm run serve
```
Then open http://localhost:8080/dev