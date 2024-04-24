package concorrencia

import (
	"reflect"
	"testing"
	"time"
)

func slowStubVerificadorWebSite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func Benchmark_verificaWebSites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma Url qualquer"
	}

	for i := 0; i < b.N; i++ {
		VerificaWebsites(slowStubVerificadorWebSite, urls)
	}
}

/// aqui é do cara lá

func mockVerificadorWebsite(url string) bool {
	if url == "waat://furhurterwe.geds" {
		return false
	}
	return true
}

func TestVerificaWebsites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	esperado := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	resultado := VerificaWebsites(mockVerificadorWebsite, websites)

	if !reflect.DeepEqual(esperado, resultado) {
		t.Fatalf("esperado %v, resultado %v", esperado, resultado)
	}
}
