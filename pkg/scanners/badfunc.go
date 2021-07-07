package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/walker"

	"github.com/s0rg/phpunisher/pkg/set"
)

const bfName = "bad-func"

var (
	badFuncs = []string{
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
		"eval",
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
	}
)

type BadFunc struct {
	stub
	scorer
	list set.Strings
}

func NewBadFunc(score float64) *BadFunc {
	bf := &BadFunc{
		scorer: scorer{Step: score, name: bfName},
		list:   make(set.Strings),
	}

	bf.list.FromList(badFuncs)

	return bf
}

// EnterNode is invoked at every node in hierarchy.
func (bf *BadFunc) EnterNode(w walker.Walkable) bool {
	n, ok := w.(node.Node)
	if !ok {
		return false
	}

	switch n.(type) {
	case *expr.FunctionCall:
		fc, ok := w.(*expr.FunctionCall)
		if !ok {
			return false
		}

		nm, ok := fc.Function.(*name.Name)
		if !ok {
			return false
		}

		for _, p := range nm.Parts {
			np, ok := p.(*name.NamePart)
			if !ok {
				continue
			}

			if bf.list.Has(np.Value) {
				bf.scorer.Up()
			}
		}
	}

	return true
}
