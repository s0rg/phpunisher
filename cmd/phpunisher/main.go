//go:build !test

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/z7zmey/php-parser/php7"

	"github.com/s0rg/phpunisher/pkg/pipe"
	"github.com/s0rg/phpunisher/pkg/scanners"
)

const (
	appName = "PHPunisher"
)

var (
	gitHash     string
	gitVersion  string
	buildDate   string
	numWorkers  = flag.Int("workers", 2, "workers count (scan parallelism)")
	showReport  = flag.Bool("report", false, "show report for found suspects")
	showVersion = flag.Bool("version", false, "show version")
	scanMasks   = flag.String("mask", "*.php*", "scan masks, use ';' as separator")
	minScore    = flag.Float64("score", 0, "minimal score to threat file as suspect")
)

func buildScanners() []scanners.Scanner {
	return []scanners.Scanner{
		scanners.NewFuncsBlacklist(0.1),
		scanners.NewArrayCall(0.1),
		scanners.NewEscapes(0.1),
		scanners.NewEvals(0.2),
		scanners.NewLongStrings(0.2),
		scanners.NewSingleInclude(0.2),
		scanners.NewArrayOperations(0.2),
	}
}

func makeHandler(callback func(path string, s scores)) func(f *pipe.File) {
	return func(f *pipe.File) {
		parser := php7.NewParser(f.Body.Bytes(), f.Path)
		parser.Parse()

		details := scores{}

		if root := parser.GetRootNode(); root != nil {
			for _, s := range buildScanners() {
				root.Walk(s)

				if sc := s.Score(); sc > 0 {
					details = append(details, &score{
						Scanner: s.Name(),
						Score:   sc,
					})
				}
			}
		}

		if total := details.Sum(); total > *minScore {
			callback(f.Path, details)
		}
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()

	if *showVersion {
		fmt.Println(appName, "git:", gitVersion, gitHash, "build at", buildDate)

		return
	}

	if flag.NArg() != 1 {
		flag.Usage()

		return
	}

	reportSuspect := func(path string, s scores) {
		var sb strings.Builder

		sb.WriteString(path)

		if *showReport {
			var report []string

			for _, d := range s {
				report = append(report, fmt.Sprintf("(%s:%.1f)", d.Scanner, d.Score))
			}

			sb.WriteString(fmt.Sprintf(" [%s=%.1f]", strings.Join(report, "+"), s.Sum()))
		}

		fmt.Println(sb.String())
	}

	p := pipe.New(
		*numWorkers,
		strings.Split(*scanMasks, ";"),
		makeHandler(reportSuspect),
	)

	if err := p.Walk(".", os.DirFS(flag.Arg(0))); err != nil {
		log.Fatal(err)
	}
}
