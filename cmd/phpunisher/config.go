//go:build !test

package main

import (
	"fmt"
	"io"

	"gopkg.in/yaml.v3"

	"github.com/s0rg/phpunisher/pkg/scanners"
)

type Config struct {
	Funcs         *FuncsBlacklistSettings  `yaml:"funcs,omitempty"`
	LongStr       *LongStringsSettings     `yaml:"long-str,omitempty"`
	ArrayOps      *ArrayOperationsSettings `yaml:"array-ops,omitempty"`
	ArrayCall     *ScoreSettings           `yaml:"array-call,omitempty"`
	Escapes       *ScoreSettings           `yaml:"escapes,omitempty"`
	Evals         *ScoreSettings           `yaml:"evals,omitempty"`
	SingleInclude *ScoreSettings           `yaml:"include,omitempty"`
}

type ScoreSettings struct {
	Score float64 `yaml:"score"`
}

type FuncsBlacklistSettings struct {
	ScoreSettings `yaml:",inline"`
	Names         []string `yaml:"names"`
}

type LongStringsSettings struct {
	ScoreSettings `yaml:",inline"`
	MinLength     int `yaml:"min_length"`
}

type ArrayOperationsSettings struct {
	ScoreSettings `yaml:",inline"`
	MaxRate       float64 `yaml:"max_rate"`
}

var badNames = []string{
	"exec",
	"expect_popen",
	"passthru",
	"system",
	"shell_exec",
	"popen",
	"proc_open",
	"pcntl_exec",
	"pcntl_alarm",
	"pcntl_exec",
	"pcntl_fork",
	"pcntl_get_last_error",
	"pcntl_getpriority",
	"pcntl_setpriority",
	"pcntl_signal",
	"pcntl_signal_dispatch",
	"pcntl_sigprocmask",
	"pcntl_sigtimedwait",
	"pcntl_sigwaitinfo",
	"pcntl_strerror",
	"pcntl_wait",
	"pcntl_waitpid",
	"pcntl_wexitstatus",
	"pcntl_wifcontinued",
	"pcntl_wifexited",
	"pcntl_wifsignaled",
	"pcntl_wifstopped",
	"pcntl_wstopsig",
	"pcntl_wtermsig",
	"assert",
	"str_rot13",
	"base64_decode",
	"gzinflate",
	"gzuncompress",
	"preg_replace",
	"chr",
	"hexdec",
	"decbin",
	"bindec",
	"ord",
	"str_replace",
	"substr",
	"goto",
	"unserialize",
	"explode",
	"strchr",
	"strstr",
	"chunk_split",
	"strtok",
	"addcslashes",
	"runkit_function_rename",
	"rename_function",
	"call_user_func_array",
	"call_user_func",
	"register_tick_function",
	"register_shutdown_function",
	"fsockopen",
	"extract",
}

var defaultConfig = &Config{
	Funcs: &FuncsBlacklistSettings{
		ScoreSettings: ScoreSettings{Score: 0.1},
		Names:         badNames,
	},
	LongStr: &LongStringsSettings{
		ScoreSettings: ScoreSettings{Score: 0.2},
		MinLength:     64,
	},
	ArrayOps: &ArrayOperationsSettings{
		ScoreSettings: ScoreSettings{Score: 0.2},
		MaxRate:       0.2,
	},
	ArrayCall:     &ScoreSettings{Score: 0.1},
	Escapes:       &ScoreSettings{Score: 0.1},
	Evals:         &ScoreSettings{Score: 0.2},
	SingleInclude: &ScoreSettings{Score: 0.2},
}

func (c *Config) Encode(w io.Writer) (err error) {
	if e := yaml.NewEncoder(w).Encode(c); e != nil {
		return fmt.Errorf("yaml: %w", e)
	}

	return nil
}

func (c *Config) Decode(r io.Reader) (err error) {
	if e := yaml.NewDecoder(r).Decode(c); e != nil {
		return fmt.Errorf("yaml: %w", e)
	}

	return nil
}

func (c *Config) MakeScanners() (rv []scanners.Scanner) {
	rv = []scanners.Scanner{}

	if sc := c.Funcs; sc.Score > 0 {
		rv = append(rv, scanners.NewFuncsBlacklist(sc.Score, sc.Names))
	}

	if sc := c.LongStr; sc.Score > 0 {
		rv = append(rv, scanners.NewLongStrings(sc.Score, sc.MinLength))
	}

	if sc := c.ArrayOps; sc.Score > 0 {
		rv = append(rv, scanners.NewArrayOperations(sc.Score, sc.MaxRate))
	}

	if sc := c.ArrayCall; sc.Score > 0 {
		rv = append(rv, scanners.NewArrayCall(sc.Score))
	}

	if sc := c.Escapes; sc.Score > 0 {
		rv = append(rv, scanners.NewEscapes(sc.Score))
	}

	if sc := c.Evals; sc.Score > 0 {
		rv = append(rv, scanners.NewEvals(sc.Score))
	}

	if sc := c.SingleInclude; sc.Score > 0 {
		rv = append(rv, scanners.NewSingleInclude(sc.Score))
	}

	return rv
}

func (c *Config) Merge(o *Config) {
	if sc := o.Funcs; sc != nil {
		c.Funcs = sc
	}

	if sc := o.LongStr; sc != nil {
		c.LongStr = sc
	}

	if sc := o.ArrayOps; sc != nil {
		c.ArrayOps = sc
	}

	if sc := o.ArrayCall; sc != nil {
		c.ArrayCall = sc
	}

	if sc := o.Escapes; sc != nil {
		c.Escapes = sc
	}

	if sc := o.Evals; sc != nil {
		c.Evals = sc
	}

	if sc := o.SingleInclude; sc != nil {
		c.SingleInclude = sc
	}
}
