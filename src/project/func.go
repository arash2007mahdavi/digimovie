package project

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Message1 struct {
	Success bool   `json:"success"`
	Text    string `json:"text"`
}

type Message2 struct {
	Success bool   `json:"success"`
	ID      int    `json:"id"`
	Text    string `json:"text"`
}

type Movie struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Country  string `json:"country"`
	Year     int    `json:"year"`
	IMDB     string `json:"imdb"`
}

var DataBase = make(map[int]Movie)

func Check_Api_Key(w http.ResponseWriter, r *http.Request) error {
	if r.Header.Get("api-key") != "09945092968" {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "wrong api-key"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return fmt.Errorf("wrong api-key")
	}
	return nil
}

var id_counter int

func Set_Movie(w http.ResponseWriter, r *http.Request) {
	check_err := Check_Api_Key(w, r)
	if check_err != nil {
		return
	}
	if len(r.URL.Query().Get("fullname")) == 0 || len(r.URL.Query().Get("country")) == 0 || len(r.URL.Query().Get("year")) == 0 || len(r.URL.Query().Get("imdb")) == 0 {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "not enoght args"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return
	}
	if r.Method != http.MethodPost {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "incorrect method"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return
	}
	id := id_counter + 1
	id_counter++
	fullname := r.URL.Query().Get("fullname")
	country := r.URL.Query().Get("country")
	year := r.URL.Query().Get("year")
	year_int, _ := strconv.Atoi(year)
	imdb := r.URL.Query().Get("imdb")
	movie := Movie{ID: id, Fullname: fullname, Country: country, Year: year_int, IMDB: imdb}
	DataBase[id] = movie
	w.WriteHeader(200)
	message := Message2{Success: true, ID: id, Text: "movie set"}
	json_message, _ := json.Marshal(message)
	fmt.Fprintf(w, "%s", string(json_message))
}

func Get_Movie(w http.ResponseWriter, r *http.Request) {
	check_err := Check_Api_Key(w, r)
	if check_err != nil {
		return
	}
	if len(r.URL.Query().Get("fullname")) == 0 {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "not enoght args"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "incorrect method"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return
	}
	fullname := r.URL.Query().Get("fullname")
	fullname = strings.ToLower(fullname)
	for _, value := range DataBase {
		if strings.ToLower(value.Fullname) == fullname {
			w.WriteHeader(200)
			json_value, _ := json.Marshal(value)
			fmt.Fprintf(w, "%s", string(json_value))
			return
		}
	}
	w.WriteHeader(403)
	message := Message1{Success: false, Text: "there isn't any movie with this name"}
	json_message, _ := json.Marshal(message)
	fmt.Fprintf(w, "%s", string(json_message))
}

func Movies_List(w http.ResponseWriter, r *http.Request) {
	check_err := Check_Api_Key(w, r)
	if check_err != nil {
		return
	}
	if r.Method != http.MethodGet {
		w.WriteHeader(403)
		message := Message1{Success: false, Text: "incorrect method"}
		json_message, _ := json.Marshal(message)
		fmt.Fprintf(w, "%s", string(json_message))
		return
	}
	w.WriteHeader(200)
	json_list, _:= json.Marshal(DataBase)
	fmt.Fprintf(w, "%s", string(json_list))
}

func Show_Version(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	message := Message1{Success: true, Text: "version 1.0"}
	json_message, _:= json.Marshal(message)
	fmt.Fprintf(w, "%s", string(json_message))
}