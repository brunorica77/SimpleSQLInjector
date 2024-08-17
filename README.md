# SimpleSQLInjector
 Simple go program to exploit blind boolean based SQL injections over Post method. Simple go program to exploit blind boolean based SQL injections over Post method. Simple go program to exploit blind boolean based SQL injections over Post method.

Instalation: 

Build in linux:

```bash
go build -o SimpleSQLI ./main.go ./tests.go -ldflags "-s -w" && upx brute ./SimpleSQLI && ./SimpleSQLI
```
Build from Linux for Windows:

```bash
env GOOS=windows GOARCH=amd64 go build -o ./SimpleSQLI -ldflags "-s -w" ./main3.go ./tests.go &&
```

Build in Windows:

```powershell
go build -o ./SimpleSQLI -ldflags "-s -w" ./main3.go ./tests.go
```
