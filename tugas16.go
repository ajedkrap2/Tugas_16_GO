package main

import (
  "belajargolang/models"
	"encoding/json"
	"fmt"
	"net/http"
  "strings"
	// "strconv"
)

var host = ":8080"

type Rental = models.Rental

type response struct {
  Success bool `json:"success"`
  Message string `json:"message"`
  Data []Rental `json:"data"`
}

var mobilPath = "/mobil"

func main() {
	http.HandleFunc(mobilPath, tampilMobil)
	http.HandleFunc(mobilPath+"/", updateMobil)

	http.HandleFunc("/", func(web http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(web, "Welcome to your paradise")
	})

	fmt.Println("Server Running at port localhost" + host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		fmt.Printf("could not run the server %v \n", err)
		return
	}
}

func tampilMobil(web http.ResponseWriter, req *http.Request) {
	web.Header().Set("Content-Type", "application/json")
  fmt.Println(req)
	switch req.Method {
	case "GET":
    getData, errGet := models.TampilData()
    if errGet != nil {
      http.Error(web, errGet.Error(), http.StatusInternalServerError)
      return
    }
		var result, err = json.Marshal(response{true, "List Data" ,getData})
		if err != nil {
			http.Error(web, err.Error(), http.StatusInternalServerError)
			return
		}
		web.Write(result)
		return
	case "POST":
    var key, value string
    if errPost := req.ParseForm(); errPost != nil {
        http.Error(web, errPost.Error(), http.StatusInternalServerError)
        return
    }
    form := len(req.Form)
    if form > 1 {
      http.Error(web, "Form Should Not More Than One", http.StatusBadRequest)
      return
    } else if form < 1 {
      http.Error(web, "Form Empty", http.StatusBadRequest)
      return
    } else {
      for k,v := range req.Form {
        key = k ; value = v[0]
      }
      getData, errGet := models.CariData(key, value)
      if errGet != nil {
        http.Error(web, errGet.Error(), http.StatusInternalServerError)
        return
      }
      fmt.Println(getData)
      var result, err = json.Marshal(response{true, "Cari Data Sukses" ,getData})
      if err != nil {
        http.Error(web, err.Error(), http.StatusInternalServerError)
        return
      }
      web.Write(result)
      return
    }
	default:
		http.Error(web, "", http.StatusBadRequest)
	}
}

func updateMobil(web http.ResponseWriter, req *http.Request){
  web.Header().Set("Content-Type", "application/json")
  fmt.Println(req)
	switch req.Method {
  case "PATCH":
    var value string
    if errPost := req.ParseForm(); errPost != nil {
        http.Error(web, errPost.Error(), http.StatusInternalServerError)
    }
    form := len(req.Form)
    if form > 1 {
      http.Error(web, "Form Should Not More Than One", http.StatusBadRequest)
      return
    } else if form < 1 {
      http.Error(web, "Form Empty", http.StatusBadRequest)
      return
    } else {
      id := strings.TrimPrefix(req.URL.Path, mobilPath+"/")
      for _,v := range req.Form {
         value = v[0]
      }
      cekData, errCek := models.CekData(id)
      if errCek != nil {
        http.Error(web, errCek.Error(), http.StatusBadRequest)
      }
      update, errU := models.UpdateData(cekData, value)
      if errU != nil {
        http.Error(web, errU.Error(), http.StatusInternalServerError)
      }
      cekData.Brand = value
      var data = [] Rental {cekData}
      var result, errJs = json.Marshal(response{true, update ,data})
      if errJs != nil {
        http.Error(web, errJs.Error(), http.StatusInternalServerError)
        return
      }
      web.Write(result)
      return
    }
  case "DELETE":
    id := strings.TrimPrefix(req.URL.Path, mobilPath+"/")
    cekData, errCek := models.CekData(id)
    if errCek != nil {
      http.Error(web, errCek.Error(), http.StatusBadRequest)
    }
    delete, errD := models.DeleteData(cekData)
    if errD != nil {
      http.Error(web, errD.Error(), http.StatusInternalServerError)
    }
    var data = [] Rental{cekData}
    var result, errJs = json.Marshal(response{true, delete ,data})
    if errJs != nil {
      http.Error(web, errJs.Error(), http.StatusInternalServerError)
      return
    }
    web.Write(result)
    return
  default:
    http.Error(web, "", http.StatusBadRequest)
  }
}
