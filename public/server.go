package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

type coinsData struct {
	Id   string `json:"id"`
	Nama string `json:"nama"`
	Pel  string `json:"layanan"`
	Tgl  string `json:"tgl"`
}

func requestData(data string) []byte {
	res, err := http.Get(data)
	if err != nil {
		fmt.Print(err.Error())
		fmt.Print("error linknya")
	}
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
		fmt.Print("gak kebaca jsonnya")
	}
	fmt.Println(string(resData))
	return resData
}

func pageRead(w http.ResponseWriter, r *http.Request) {
	readTable(w, r, "index.html")
}

func updateDataPage(w http.ResponseWriter, r *http.Request) {
	flpath2 := path.Join("view", "updatePage.html")
	tmpl, err := template.ParseFiles(flpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jdata coinsData
	jdata.Id = r.FormValue("ids")
	jdata.Nama = r.FormValue("namas")
	jdata.Pel = r.FormValue("pels")
	jdata.Tgl = r.FormValue("tgls")

	err = tmpl.Execute(w, jdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(r.FormValue("result"))
}

func updateData(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		var jdata coinsData
		jdata.Id = r.FormValue("id")
		jdata.Nama = r.FormValue("nama")
		jdata.Pel = r.FormValue("layanan")
		jdata.Tgl = r.FormValue("tgl")
		datas := url.Values{
			"id":      {jdata.Id},
			"nama":    {jdata.Nama},
			"layanan": {jdata.Pel},
			"tgl":     {jdata.Tgl},
		}

		buf := bytes.Buffer{}
		enc := gob.NewEncoder(&buf)
		errs := enc.Encode(jdata)
		if errs != nil {
			fmt.Print("error encode struct")
		}
		fmt.Println(buf.Bytes())
		requesto, errto := http.PostForm("http://localhost:1333/update", datas)
		if errto != nil {
			fmt.Println("Post Error")
		}
		fmt.Println("data diupdate : ", jdata)
		var rest map[string]interface{}
		json.NewDecoder(requesto.Body).Decode(&rest)
	}
	r.Method = "GET"
	http.Redirect(w, r, "http://localhost:8080/table", http.StatusFound)
	//readTable(w, r)
}

func delData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var client = &http.Client{}
		num := r.FormValue("num")
		fmt.Println("nilai NUM :", num)
		request, err := http.NewRequest("DELETE", "http://localhost:1333/del/"+string(num), nil)
		if err != nil {
			fmt.Println("error request")
		}
		//client.Do(request)
		response, err := client.Do(request)
		if err != nil {
			fmt.Println("error client do")
		}
		defer response.Body.Close()
	}
	r.Method = "GET"
	http.Redirect(w, r, "http://localhost:8080/table", http.StatusFound)
	//readTable(w, r)
}

func addData(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nama := r.FormValue("nama")
		pel := r.FormValue("layanan")
		tgl := r.FormValue("tgl")
		var jdata coinsData
		jdata.Id = id
		jdata.Nama = nama
		jdata.Pel = pel
		jdata.Tgl = tgl
		datas := url.Values{
			"id":      {jdata.Id},
			"nama":    {jdata.Nama},
			"layanan": {pel},
			"tgl":     {tgl},
		}

		buf := bytes.Buffer{}
		enc := gob.NewEncoder(&buf)
		err := enc.Encode(jdata)
		if err != nil {
			fmt.Print("error encode struct")
		}
		fmt.Println("id :", id, "nama :", nama, "pel :", pel, "tgl :", tgl)
		fmt.Println(buf.Bytes())
		requesto, errto := http.PostForm("http://localhost:1333/add", datas)
		if errto != nil {
			fmt.Println("Post Error")
		}
		var rest map[string]interface{}
		json.NewDecoder(requesto.Body).Decode(&rest)

		//request, err := http.NewRequest("POST", "http://localhost:1333/add", bytes.NewBuffer(buf.Bytes()))
		//request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	r.Method = "GET"
	http.Redirect(w, r, "http://localhost:8080/table", http.StatusFound)
	//readTable(w, r)
}

func readTable(w http.ResponseWriter, r *http.Request, page string) {
	//num := r.FormValue("numid")
	dat := requestData("http://localhost:1333/read")
	//tess := string(dat)
	//tes2, err := json.Marshal(tess)
	//fmt.Println("tipe dari DAT :  ", reflect.TypeOf(tess))
	//fmt.Println("tipe dari tes :  ", reflect.TypeOf(tes2))
	//fmt.Println("isi tes2 : ", tes2)

	flpath2 := path.Join("view", page)
	tmpl, err := template.ParseFiles(flpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//tes := `{"id":"sadfasd", "nama":"sdafsa", "pelayanan":"sdafds", "tgl":"sdassda"}`
	var result []coinsData
	errs := json.Unmarshal(dat, &result)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Printf("%+v\n", result)
	err = tmpl.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(r)
	fmt.Print(r.Method)
}

func logins(w http.ResponseWriter, r *http.Request) {
	requestData("http://localhost:1333/user")

	flpath2 := path.Join("view", "index.html")
	tmpl, err := template.ParseFiles(flpath2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]interface{}{
		"title": "page",
		"name":  "superman",
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(r)
}

func main() {
	http.HandleFunc("/login", logins)
	http.HandleFunc("/table", pageRead)
	http.HandleFunc("/tes", addData)
	http.HandleFunc("/delete", delData)
	http.HandleFunc("/updatepage", updateDataPage)
	http.HandleFunc("/update", updateData)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		flpath := path.Join("view", "home.html")
		tmpl, err := template.ParseFiles(flpath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := map[string]interface{}{
			"title": "learn golang",
			"name":  "superman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(r)
	})

	//http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
	//	temp := template.Must(template.ParseFiles("view/home.html"))
	//	if err := temp.Execute(w, nil); err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//	}
	//})
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("src"))))

	fmt.Println("server running in 8080")
	http.ListenAndServe(":8080", nil)
}
