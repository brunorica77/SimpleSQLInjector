package main

import (
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"
)



type Data_struct struct {
	value url.Values
	set   string
}

func (d *Data_struct) process_data(dataraw string) error {
	var err error
  if dataraw == ""{
    err = errors.New("Error: És necessari introduir la data que tramitada. Per a més informació utilitzar el paràmetre -h/--help")
    return err
  } else if !strings.Contains(dataraw, "SQLI") {
		err = errors.New("Error: Falta paraula clau SQLI")
    return err
  }

	d.value = url.Values{}
	parsedstr := strings.Split(dataraw, "&")
	for _, i := range parsedstr {

		if !strings.Contains(i, "=") {
			err = errors.New("Error: Data introduïda no valida. Per a més informació utilitzar el paràmetre -h/--help")
			return err
		}

		splited := strings.SplitN(i, "=", 2)
		d.value.Add(splited[0], splited[1])
		if splited[1] == "SQLI" {
			d.set = splited[0]
		}

	}
	return err
}

func (d *Data_struct) payload_load(payload string) {
	d.value.Set(d.set, payload)
}

func DataInit(dataraw string) (*Data_struct, error) {
	var dataf = &Data_struct{}
	var err error = dataf.process_data(dataraw)
	if err != nil {
		return nil, err
	}
	return dataf, err
}

//============================================================================================================================================

type QuerySQL struct {
	selectsql string
	fromsql   string
	extrasql  string
}

func QueryInit(queryraw string, db string) (*QuerySQL, error) {
	var queryf = &QuerySQL{}
	var err = queryf.process_query(queryraw, db)
	if err != nil {
		return nil, err
	}
	return queryf, err
}

// de moment treure els error
func (q *QuerySQL) process_query(queryraw string, db string) error {
	var known_querys = map[string][3]string{
		"show databases": {"schema_name", "from information_schema.schemata", ""},
		"show tables":    {"TABLE_NAME,0x3a,TABLE_SCHEMA", "from INFORMATION_SCHEMA.TABLES", ""},
		"desc":           {"COLUMN_NAME,0x3a,TABLE_NAME,0x3a,TABLE_SCHEMA", "from INFORMATION_SCHEMA.COLUMNS", "TABLE_NAME='%v'"},
	}
	for k, v := range known_querys {
		if strings.Contains(queryraw, k) {
			q.selectsql = v[0]
			q.fromsql = v[1]
			splited_query := strings.Split(queryraw, " ")
			for sqk, sqv := range splited_query {
				if strings.Contains(k, sqv) {
					splited_query[sqk] = ""
				}
			}
			if known_querys[k][2] != "" {
				q.extrasql += "where "
				q.extrasql += fmt.Sprintf(known_querys[k][2], strings.Join(splited_query, ""))
				if db != "" {
					q.extrasql += fmt.Sprintf(" and TABLE_SCHEMA='%v'", db)
				}
			} else if db != "" {
				q.extrasql += fmt.Sprintf("where TABLE_SCHEMA='%v'", db)
			}
      return nil
			//fmt.Printf("%v", splited_query)

		}
	}
  	if !strings.Contains(queryraw,"select") {
    		var err = errors.New("Consulta SQL incorrecta.")
    		return err
  	}
	splited_query := strings.Split(queryraw, " ")
	for i := slices.Index(splited_query, "select") + 1; i <= len(splited_query)-1; i++ {

    if splited_query[i-1] == "from" || i == slices.Index(splited_query, "from")+1 {
      if db != ""{
        q.fromsql+=" from "+db+"."+splited_query[i]
      } else {
        q.fromsql +=" from "+splited_query[i] + " "
      }

		} else if i > slices.Index(splited_query,"select") && i < slices.Index(splited_query,"from") || slices.Index(splited_query,"from") == -1 {
			q.selectsql += splited_query[i] + " "

		} else if splited_query[i] != "from"{
      //else if len(splited_query)-1 > slices.Index(splited_query, "from")+1 && i > slices.Index(splited_query, "from")+1
			q.extrasql += splited_query[i] + " "

		}
		q.selectsql = strings.ReplaceAll(q.selectsql,",",",0x3a,")
	}
	return nil
}

