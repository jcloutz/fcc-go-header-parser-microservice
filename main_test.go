package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(indexHandler))
	defer server.Close()

	// http.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req, _ := http.NewRequest("Get", server.URL, nil)
	req.Header.Set("Accept-Language", "en-US,en;q=0.8")
	res, err := http.DefaultClient.Do(req)
	http.Client
	// res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("Should be able to get response.")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal("Status should be 200")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Body should not be empty")
	}

	var hr HeaderResponse
	err = json.Unmarshal(b, &hr)
	if err != nil {
		t.Fatal("Unable to unmarshal json into HeaderResponse struct")
	}

	if !strings.Contains(*hr.IP, "127.0.0.1") {
		t.Fatal("IP should contain 127.0.0.1")
	}

	if !strings.Contains(*hr.UserAgent, "Go-http-client") {
		t.Fatal("UserAgent should contain Go-http-client")
	}

	if !strings.Contains(*hr.Language, "en-US") {
		t.Fatal("Accept Language (Language) should contain en-US")
	}
}

func TestEmptyLanguage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(indexHandler))
	defer server.Close()

	// http.Header.Set("Accept-Language", "en-US,en;q=0.8")
	req, _ := http.NewRequest("Get", server.URL, nil)
	req.Header.Set("Accept-Language", "")
	res, err := http.DefaultClient.Do(req)
	// res, err := http.Get(server.URL)
	if err != nil {
		t.Fatal("Should be able to get response.")
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatal("Status should be 200")
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal("Body should not be empty")
	}

	var hr HeaderResponse
	err = json.Unmarshal(b, &hr)
	if err != nil {
		t.Fatal("Unable to unmarshal json into HeaderResponse struct")
	}

	if !strings.Contains(*hr.IP, "127.0.0.1") {
		t.Fatal("IP should contain 127.0.0.1")
	}

	if !strings.Contains(*hr.UserAgent, "Go-http-client") {
		fmt.Println(*hr.UserAgent)
		t.Fatal("UserAgent should contain Go-http-client")

	}

	if *hr.Language != "" {
		t.Fatal("Accept-Language should be empty")
	}
}
