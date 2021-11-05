//go:build !test

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/z7zmey/php-parser/node"
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
	dumpConfig  = flag.Bool("dump-conf", false, "dump default scanners config to stdout")
	configFile  = flag.String("conf", "", "load scanners config from file")
	scanMasks   = flag.String("mask", "*.php*", "scan masks, use ';' as separator")
	minScore    = flag.Float64("score", 0, "minimal score to threat file as suspect")
)

func makeHandler(cfg *Config, callback func(string, scanners.Scores)) func(f *pipe.File) {
	return func(f *pipe.File) {
		parser := php7.NewParser(f.Body.Bytes(), f.Path)
		parser.Parse()

		result := scanners.Scores{}

		root := parser.GetRootNode()
		if root == nil {
			return
		}

		var wg sync.WaitGroup

		scns := cfg.MakeScanners()

		details := make(chan *scanners.Score, len(scns))

		for _, s := range scns {
			wg.Add(1)

			go func(s scanners.Scanner, n node.Node) {
				n.Walk(s)

				if sc := s.Score(); sc > 0 {
					details <- &scanners.Score{
						Scanner: s.Name(),
						Score:   sc,
					}
				}

				wg.Done()
			}(s, root)
		}

		wg.Wait()

		close(details)

		for s := range details {
			result = append(result, s)
		}

		if total := result.Sum(); total > *minScore {
			callback(f.Path, result)
		}
	}
}

func loadConfig(path string) (c *Config, err error) {
	var (
		cfg Config
		fd  *os.File
	)

	if fd, err = os.Open(path); err != nil {
		err = fmt.Errorf("open: %w", err)

		return
	}

	defer fd.Close()

	if err = cfg.Decode(fd); err != nil {
		err = fmt.Errorf("decode: %w", err)

		return
	}

	return &cfg, nil
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	flag.Parse()

	if *showVersion {
		fmt.Println(appName, "git:", gitVersion, gitHash, "build at", buildDate)

		return
	}

	cfg := defaultConfig

	if *dumpConfig {
		_ = cfg.Encode(os.Stdout)

		return
	}

	if *configFile != "" {
		ucf, err := loadConfig(*configFile)
		if err != nil {
			log.Fatal("config:", err)
		}

		cfg.Merge(ucf)
	}

	handler := makeHandler(cfg, func(path string, score scanners.Scores) {
		var sb strings.Builder

		sb.WriteString(path)

		if *showReport {
			var report []string

			for _, d := range score {
				report = append(report, fmt.Sprintf("(%s:%.1f)", d.Scanner, d.Score))
			}

			sb.WriteString(fmt.Sprintf(" [%s=%.1f]", strings.Join(report, "+"), score.Sum()))
		}

		fmt.Println(sb.String())
	})

	masks := strings.Split(*scanMasks, ";")
	p := pipe.New(*numWorkers, masks, handler)

	if err := p.Walk(".", os.DirFS(".")); err != nil {
		log.Fatal(err)
	}
}
