package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/z7zmey/php-parser/php7"

	"github.com/s0rg/phpunisher/pkg/pipe"
	"github.com/s0rg/phpunisher/pkg/scanners"
)

var (
	GitHash    string
	BuildDate  string
	minScore   = flag.Float64("score", 0, "minimal score to threat file as suspect (default: 0)")
	logVerbose = flag.Bool("verbose", false, "show scan details for found suspects")
	scanMasks  = flag.String("mask", "*.php*", "scan masks, use ';' as separator")
	numWorkers = flag.Int("workers", 2, "workers count (scan parallelism)")
)

type score struct {
	Scanner string
	Score   float64
}

type scores []*score

func (s scores) Len() int           { return len(s) }
func (s scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool { return s[i].Score < s[j].Score }

func (s scores) Sum() (rv float64) {
	for i := 0; i < len(s); i++ {
		rv += s[i].Score
	}

	return
}

func buildScanners() []scanners.Scanner {
	return []scanners.Scanner{
		scanners.NewEvalExpr(0.2),
		scanners.NewSingleInclude(0.2),
		scanners.NewArrayCall(0.1),
		scanners.NewBadString(0.1),
		scanners.NewBadFunc(0.1),
		scanners.NewArrayOperations(0.2),
	}
}

func makeHandler(callback func(path string, s scores)) func(f *pipe.File) {
	return func(f *pipe.File) {
		parser := php7.NewParser(f.Body.Bytes(), f.Path)
		parser.Parse()

		for _, e := range parser.GetErrors() {
			log.Printf("scanner: parse error on %s: %v", f.Path, e)
		}

		root := parser.GetRootNode()
		if root == nil {
			log.Printf("scanner: no root node for %s", f.Path)

			return
		}

		details := scores{}

		for _, s := range buildScanners() {
			root.Walk(s)

			if sc := s.Score(); sc > 0 {
				details = append(details, &score{
					Scanner: s.Name(),
					Score:   sc,
				})
			}
		}

		if total := details.Sum(); total > *minScore {
			callback(f.Path, details)
		}
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("git", GitHash, "build at", BuildDate)
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	verbose := *logVerbose
	report := log.New(os.Stdout, "", 0)

	reportSuspect := func(path string, s scores) {
		var sb strings.Builder

		sb.WriteString(path)

		if verbose {
			sb.WriteString(fmt.Sprintf(" [%.1f]\n", s.Sum()))
			sort.Sort(s)

			for _, d := range s {
				sb.WriteString(fmt.Sprintf("\t%s %.1f\n", d.Scanner, d.Score))
			}
		}

		report.Println(sb.String())
	}

	p := pipe.New(
		*numWorkers,
		strings.Split(*scanMasks, ";"),
		makeHandler(reportSuspect),
	)

	root := args[0]

	if err := p.Walk(".", os.DirFS(root)); err != nil {
		log.Fatal(err)
	}
}
