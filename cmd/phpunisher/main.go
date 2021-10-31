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

func makeHandler(cfg *Config, callback func(path string, s scores)) func(f *pipe.File) {
	return func(f *pipe.File) {
		parser := php7.NewParser(f.Body.Bytes(), f.Path)
		parser.Parse()

		details := scores{}

		if root := parser.GetRootNode(); root != nil {
			for _, s := range cfg.MakeScanners() {
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

	handler := makeHandler(cfg, func(path string, s scores) {
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
	})

	masks := strings.Split(*scanMasks, ";")
	p := pipe.New(*numWorkers, masks, handler)

	if err := p.Walk(".", os.DirFS(".")); err != nil {
		log.Fatal(err)
	}
}
