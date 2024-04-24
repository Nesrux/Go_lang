package concorrencia

import (
	"testing"
	"time"
)

func slowStubVerificadorWebSite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func Benchmark_verificaWebSites(b *testing.B) {
	urls := make([]string, 100)
	for i:=0; i < len(urls); i++{
		urls[i] = "uma Url qualquer"
	}

	for i := 0; i <b.N; i++ {
		VerificaWebsites(slowStubVerificadorWebSite, urls)
	}
}
