/*
client
*/
package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/lizrice/secure-connections/utils"
)

type CandyRequest struct {
	Money      int64  `json:"money"`
	CandyType  string `json:"candyType"`
	CandyCount int64  `json:"candyCount"`
}

func main() {
	k := flag.String("k", "", "candy type")
	c := flag.Int64("c", 0, "count of candies")
	m := flag.Int64("m", 0, "money")
	flag.Parse()

	candy := CandyRequest{CandyType: *k, Money: *m, CandyCount: *c}

	if !requestValid() {
		log.Fatalln("Wrong arguments")
	}

	var data bytes.Buffer
	err := json.NewEncoder(&data).Encode(candy)
	if err != nil {
		log.Fatal(err)
	}

	client := getClietn()
	resp, err := client.Post("https://127.0.0.1:3333/buy_candy", "application/json", bytes.NewBuffer(data.Bytes()))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func getClietn() *http.Client {
	data, err := ioutil.ReadFile("../ca/minica.pem")
	if err != nil {
		log.Println(err)
	}
	cp, err := x509.SystemCertPool()
	if err != nil {
		log.Println(err)
	}
	cp.AppendCertsFromPEM(data)

	config := &tls.Config{
		InsecureSkipVerify:    true,
		ClientAuth:            tls.RequireAndVerifyClientCert,
		RootCAs:               cp,
		GetCertificate:        utils.CertReqFunc("../ca/client/cert.pem", "../ca/client/key.pem"),
		VerifyPeerCertificate: utils.CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func requestValid() bool {
	flags := map[string]bool{"k": false, "c": false, "m": false}
	flag.Visit(func(f *flag.Flag) {
		flags[f.Name] = true
	})
	for _, b := range flags {
		if !b {
			return false
		}
	}
	return true
}
