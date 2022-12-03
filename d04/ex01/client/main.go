package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	order := initFlags()
	clientJson, err := json.Marshal(order)
	if err != nil {
		fmt.Printf("Bad arguments")
		os.Exit(1)
	}
	client := getClient()
	bodyPost := bytes.NewReader(clientJson)
	resp, err := client.Post("https://localhost:8080/buy_candy", "application/json; charset=UTF-8", bodyPost)
	must(err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	must(err)

	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
}

func initFlags() Order {
	var order Order
	flag.StringVar(&order.CandyType, "k", "", "CandyType use with two char")
	flag.IntVar(&order.CandyCount, "c", 0, "CandyCount use with int")
	flag.IntVar(&order.Money, "m", 0, "Money use with int")
	flag.Parse()
	return order
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../minica.pem")
	cp.AppendCertsFromPEM(data)

	// c, _ := tls.LoadX509KeyPair("signed-cert", "key")

	config := &tls.Config{
		// Certificates: []tls.Certificate{c},
		RootCAs:               cp,
		GetClientCertificate:  ClientCertReqFunc("cert.pem", "key.pem"),
		VerifyPeerCertificate: CertificateChains,
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}

func must(err error) {
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
}

func getCert(certfile, keyfile string) (c tls.Certificate, err error) {
	if certfile != "" && keyfile != "" {
		c, err = tls.LoadX509KeyPair(certfile, keyfile)
		if err != nil {
			fmt.Printf("Error loading key pair: %v\n", err)
		}
	} else {
		err = fmt.Errorf("I have no certificate")
	}
	return
}

// ClientCertReqFunc returns a function for tlsConfig.GetClientCertificate
func ClientCertReqFunc(certfile, keyfile string) func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
	c, err := getCert(certfile, keyfile)

	return func(certReq *tls.CertificateRequestInfo) (*tls.Certificate, error) {
		fmt.Println("Received certificate request: sending certificate")
		if err != nil || certfile == "" {
			fmt.Println("I have no certificate")
		} else {
			err := OutputPEMFile(certfile)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		Wait()
		return &c, nil
	}
}

func Wait() {
	fmt.Printf("[Press enter to proceed]")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println()
}

// OutputPEMFile reads info from a PEM file and displays it
func OutputPEMFile(filename string) error {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	for len(data) > 0 {
		var block *pem.Block
		block, data = pem.Decode(data)
		fmt.Printf("Type: %#v\n", block.Type)
		switch block.Type {
		case "CERTIFICATE":
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return err
			}
			fmt.Printf(CertificateInfo(cert))
		default:
			fmt.Println(block.Type)
		}
	}
	return nil
}

// CertificateInfo returns a string describing the certificate
func CertificateInfo(cert *x509.Certificate) string {
	if cert.Subject.CommonName == cert.Issuer.CommonName {
		return fmt.Sprintf("    Self-signed certificate %v\n", cert.Issuer.CommonName)
	}

	s := fmt.Sprintf("    Subject %v\n", cert.DNSNames)
	s += fmt.Sprintf("    Issued by %s\n", cert.Issuer.CommonName)
	return s
}

// CertificateChains prints information about verified certificate chains
func CertificateChains(rawCerts [][]byte, chains [][]*x509.Certificate) error {
	if len(chains) > 0 {
		fmt.Println("Verified certificate chain from peer:")

		for _, v := range chains {
			for i, cert := range v {
				fmt.Printf("  Cert %d:\n", i)
				fmt.Printf(CertificateInfo(cert))
			}
		}
		Wait()
	}

	return nil
}
