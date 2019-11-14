package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"

	"github.com/atotto/clipboard"
	flag "github.com/spf13/pflag"
)

type options struct {
	alpha   bool
	numeric bool
	special bool
}

type generator struct {
	opts    options
	symbols []byte
	r       *rand.Rand
}

var alphas = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var nums = []byte("0123456789")
var specials = []byte("!@#$%^&*()_+=-`~.,")

func newGenerator(opts options) *generator {
	gen := &generator{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	var buf bytes.Buffer
	if opts.alpha {
		buf.Write(alphas)
	}
	if opts.numeric {
		buf.Write(nums)
	}
	if opts.special {
		buf.Write(specials)
	}
	gen.symbols = buf.Bytes()
	return gen
}
func (gen *generator) char() byte {
	i := gen.r.Intn(len(gen.symbols))
	return gen.symbols[i]
}

func (gen *generator) generate(size int) string {
	var buf bytes.Buffer
	for i := 0; i < size; i++ {
		buf.WriteByte(gen.char())
	}
	return buf.String()
}

func main() {
	var opts options
	flag.BoolVarP(&opts.alpha, "alpha", "a", true, "use alphabetical symbols")
	flag.BoolVarP(&opts.numeric, "numeric", "n", true, "use numeric symbols")
	flag.BoolVarP(&opts.special, "special", "s", false, "use special symbols")
	verbose := flag.BoolP("verbose", "v", false, "print out the generated password")
	length := flag.IntP("length", "l", 10, "password length to generate")
	flag.Parse()

	gen := newGenerator(opts)
	passwd := gen.generate(*length)

	if *verbose {
		fmt.Println(passwd)
	}
	if err := clipboard.WriteAll(passwd); err != nil {
		fmt.Println("cannot copy password to clipboard:", err)
	} else {
		fmt.Printf("password (%d chars) copied", *length)
	}
}
