package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mycaserver/pkg/ca"
	"net/http"
)

var server *http.Server
var running bool

func Run() {
	if running {
		return
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/csr-template", getCsrTemplate)
	mux.HandleFunc("/csr", signCsrHandler)
	server = &http.Server{
		Addr:    ":8111",
		Handler: mux,
	}
	running = true
	fmt.Printf("server listening at %v, http", server.Addr)
	if server.ListenAndServe() != nil {
		running = false
		log.Print("can't start http server @ 8111")
	}
	running = false
}

func signCsrHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	csr := &ca.CertificateSigningRequest{}

	err = json.Unmarshal(reqBody, csr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var sync chan int = make(chan int, 1)
	go signCsrRoutine(w, csr, sync)
	// 如果一个channel已经被关闭，那么在该channel上进行读取操作将不会阻塞，并且会立即返回一个零值和一个布尔值，表示该channel是否已经关闭。
	<-sync
}

func signCsrRoutine(w http.ResponseWriter, csr *ca.CertificateSigningRequest, sync chan<- int) {
	defer close(sync)
	theCert, err := ca.CA.SignX509(csr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error happen:%v", err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("Context-Type", "application/json")
	jsonByte, _ := json.Marshal(theCert)
	w.Write(jsonByte)

}

func getCsrTemplate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// 创建一个csr模版实例
	csr := ca.CertificateSigningRequest{
		SubjectCountry:            []string{"China"},
		SubjectOrganization:       []string{"Qinghua"},
		SubjectOrganizationalUnit: []string{"ComputerScience"},
		SubjectProvince:           []string{"Shanxi"},
		SubjectLocality:           []string{"Xian"},

		SubjectCommonName: "harry.test.com",
		EmailAddresses:    []string{"ex.example@qq.com"},
		DNSNames:          []string{"localhost"},
	}
	// 序列化 & 返回
	csrBytes, err := json.Marshal(csr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(csrBytes)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("err when reading body"))
		return
	}
	w.Write(body)
}
