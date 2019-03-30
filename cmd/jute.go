package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"sync"
	"time"
)

func main() {
	f, err := os.Open("zookeeper.jute")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	wg := new(sync.WaitGroup)
	wg.Add(1)
	ch := make(chan string)

	go lex(ch, wg)

	s := bufio.NewScanner(f)
	commentRegexp := regexp.MustCompile(`//.*$`)
	angleBracketRegexp := regexp.MustCompile(`>([A-Za-z0-9_])`)
	var comment bool
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		switch {
		case l == "/**" && !comment:
			comment = true
			continue
		case l == "*/" && comment:
			comment = false
			continue
		case comment:
			continue
		}
		l = commentRegexp.ReplaceAllString(l, "")
		if l == "" {
			continue
		}
		l = angleBracketRegexp.ReplaceAllString(l, "> $1")
		l = strings.ReplaceAll(l, ";", " ;")

		sw := bufio.NewScanner(strings.NewReader(l))
		sw.Split(bufio.ScanWords)
		for sw.Scan() {
			ch <- sw.Text()
		}
		if err := sw.Err(); err != nil {
			panic(err)
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	close(ch)
	wg.Wait()
}

func lex(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	chT := make(chan interface{})
	chK := make(chan string)
	chI := make(chan string)

	wgp := new(sync.WaitGroup)
	wgp.Add(1)

	go parse(chT, chK, chI, wgp)

	for s := range ch {
		switch s {
		case "module":
			chK <- "module"
		case "class":
			chK <- "class"

		case "byte":
			chT <- byte(0)
		case "boolean":
			chT <- false
		case "int":
			chT <- uint32(0)
		case "long":
			chT <- uint64(0)
		case "float":
			chT <- float32(0)
		case "double":
			chT <- float64(0)
		case "ustring":
			chT <- ""
		case "buffer":
			chT <- []byte(nil)

		case "{":
			chK <- "{"
		case "}":
			chK <- "}"
		case ";":
			chK <- ";"

		default:
			switch {
			case len(s) >= 3 && s[0:3] == "map":
				panic("not yet implemented")
			case len(s) >= 6 && s[0:6] == "record":
				panic("not yet implemented")
			case len(s) >= 6 && s[0:6] == "vector":
				chT <- []interface{}(nil)
				chI <- s[7 : len(s)-1]

			default:
				//panic("unknown: " + s)
				chI <- s
			}
		}
	}

	wgp.Wait()
}

func parse(chT <-chan interface{}, chK, chI <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	p := &parser{chT: chT, chK: chK, chI: chI}

	for s := p.module; s != nil; s = s() {
	}
}

type state func() state

type generator struct {
	module  *string
	class   *string
	f       *os.File
	buf     *bytes.Buffer
	imports map[string]struct{}
	path    string
}

var upperRegexp = regexp.MustCompile(`[A-Z]`)

func (g *generator) generate(v interface{}) {
	switch v := v.(type) {
	case module:
		g.module = &v.Name
	case class:
		if g.module == nil {
			panic("error")
		}

		g.class = &v.Name

		dn := strings.ReplaceAll(*g.module, ".", "/")

		fn := upperRegexp.ReplaceAllString(*g.class, "_$0")
		fn = strings.ReplaceAll(fn, "_A_C_L", "_ACL")
		fn = strings.ReplaceAll(fn, "_S_A_S_L", "_SASL")
		fn = strings.ReplaceAll(fn, "_T_T_L", "_TTL")
		fn = strings.TrimLeft(fn, "_")
		fn = strings.ToLower(fn) + ".go"

		err := os.MkdirAll(dn, 0755)
		if err != nil {
			panic(err)
		}
		f, err := os.Create(filepath.Join(dn, fn))
		if err != nil {
			panic(err)
		}
		g.f = f
		g.buf = new(bytes.Buffer)
		g.imports = make(map[string]struct{})
		g.path = filepath.Join(dn, fn)

		fmt.Fprintf(f, "package %s\n\n", filepath.Base(dn))
		fmt.Fprintf(g.buf, "type %s struct {\n", *g.class)
	case field:
		if g.module == nil || g.class == nil {
			panic("error")
		}

		switch t := v.Type.(type) {
		case ptyp:
			fmt.Fprintf(g.buf, "    %s%s %s `jute:\"%s\"`\n", strings.ToUpper(v.Name[0:1]), v.Name[1:], t.Type, v.Name)
		case ctyp:
			n := t.Type
			ps := strings.Split(n, ".")
			p := strings.Join(ps[0:len(ps)-1], ".")
			n = ps[len(ps)-1]
			if len(p) > 0 {
				if p != *g.module {
					n = ps[len(ps)-2] + "." + n
					g.imports[p] = struct{}{}
				}
			} else {
				if n == "ustring" {
					n = "string"
				}
			}
			if t.Name == "vector" {
				n = "[]" + n
			}
			fmt.Fprintf(g.buf, "    %s%s %s `jute:\"%s\"`\n", strings.ToUpper(v.Name[0:1]), v.Name[1:], n, v.Name)
		}
	case nil:
		fmt.Fprintf(g.buf, "}\n")
		if len(g.imports) > 0 {
			fmt.Fprintf(g.f, "import (\n")
			for p := range g.imports {
				fmt.Fprintf(g.f, "    \"code.witches.io/go/zookeeper/cmd/%s\"\n", strings.ReplaceAll(p, ".", "/"))
			}
			fmt.Fprintf(g.f, ")\n\n")
		}

		_, err := io.Copy(g.f, g.buf)
		if err != nil {
			panic(err)
		}

		g.f.Close()
		g.f = nil

		err = exec.Command("go", "fmt", g.path).Run()
		if err != nil {
			panic(err)
		}
	}
}

type parser struct {
	chT <-chan interface{}
	chK <-chan string
	chI <-chan string

	stack []interface{}
	g     generator
}

func (p *parser) push(v interface{}) {
	//fmt.Printf("%#v\n", v)
	p.stack = append(p.stack, v)
	p.g.generate(v)
}

func (p *parser) pop() {
	if len(p.stack) == 0 {
		panic("stack empty")
	}

	p.stack = p.stack[0 : len(p.stack)-2]
	p.g.generate(nil)
}

func (p *parser) error(format string, a ...interface{}) state {
	return func() state {
		panic(fmt.Sprintf(format, a...))
		return p.error(format, a)
	}
}

func (p *parser) module() state {
	var k string
	select {
	case k = <-p.chK:
	case <-time.After(10 * time.Millisecond):
		return nil
	}
	if k != "module" {
		return p.error("expected module")
	}
	i := <-p.chI
	k = <-p.chK
	if k != "{" {
		return p.error("expected opening brace")
	}
	p.push(module{i})
	return p.class
}

func (p *parser) class() state {
	k := <-p.chK
	if k != "class" {
		if k == "}" {
			return p.module
		}
		return p.error("expected class")
	}
	i := <-p.chI
	k = <-p.chK
	if k != "{" {
		return p.error("expected opening brace")
	}
	p.push(class{i})
	return p.field
}

func (p *parser) field() state {
	var typ typ
	select {
	case t := <-p.chT:
		rv := reflect.ValueOf(t)
		if rv.Kind() == reflect.Slice && rv.Type().Elem().Kind() == reflect.Interface {
			tc := <-p.chI
			typ = ctyp{"vector", tc}
		} else {
			typ = ptyp{rv.Type()}
		}
	case i := <-p.chI:
		typ = ctyp{i, i}
	case k := <-p.chK:
		if k != "}" {
			return p.error("expected closing brace")
		}
		p.pop()
		return p.class()
	}
	i := <-p.chI
	k := <-p.chK
	if k != ";" {
		return p.error("expected semicolon")
	}
	p.push(field{typ, i})
	return p.field
}

type module struct {
	Name string
}

type class struct {
	Name string
}

type field struct {
	Type typ
	Name string
}

type typ interface{}

type ptyp struct {
	Type reflect.Type
}

type ctyp struct {
	Name string
	Type string
}
