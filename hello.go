package hello

import (
	"golang.org/x/text/language"
	"rsc.io/sampler"
	"rsc.io/quote/v3"
)

func Hello() string {
	return sampler.Hello(language.English)
}

func Proverb() string {
	return quote.Concurrency()
}
