# SimpleSQLInjector
Simple go program to exploit blind boolean based SQL injections over Post method. 

Building: 

Build on Linux:

```bash
go build -o SimpleSQLI -ldflags "-s -w" ./main.go ./tests.go && upx brute ./SimpleSQLI && ./SimpleSQLI
```
Build from Linux for Windows:

```bash
env GOOS=windows GOARCH=amd64 go build -o SimpleSQLI -ldflags "-s -w" ./main3.go ./tests.go &&
```

Build on Windows:

```powershell
go build -o ./SimpleSQLI -ldflags "-s -w" ./main3.go ./tests.go
```

---

Once we have these done, the app is ready to be used

```bash
❯ ./SimpleSQLI -h

   _____ ____    __       ____        _           __  _           
  / ___// __ \  / /      /  _/___    (_)__  _____/ /_(_)___  ____ 
  \__ \/ / / / / /       / // __ \  / / _ \/ ___/ __/ / __ \/ __ \
 ___/ / /_/ / / /___   _/ // / / / / /  __/ /__/ /_/ / /_/ / / / /
/____/\___\_\/_____/  /___/_/ /_/_/ /\___/\___/\__/_/\____/_/ /_/ 
                               /___/                              
Bruno Ríos Castelló                                         


Usage of ./SimpleSQLI:
  -d string
    	Data: Data que es tramita per POST al servidor. Ex: usuari=SQLI&contrasenya=test
  -db string
    	Database: Base de dades on es farà la consulta.
  -es string
    	Error string: Distintiu de la resposta errònia del servidor.
  -q string
    	Query: Codi SQL per interactuar amb la base de dades. (default "select schema_name from information_schema.schemata")
  -sc int
    	Status Code: Codi d'estat que torna el servidor si es fa una petició satisfactòria.
  -u string
    	Url: Url víctima. Ex: http://ip:port
```
