package scanners

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/node/name"
	"github.com/z7zmey/php-parser/walker"
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
}

func NewBadFunc(score float64) *BadFunc {
	return &BadFunc{
		scorer: scorer{Step: score, name: bfName},
	}
}

func isBadFunc(name string) (yes bool) {
	for i := 0; i < len(badFuncs); i++ {
		if yes = name == badFuncs[i]; yes {
			return
		}
	}

	return
}

// EnterNode is invoked at every node in hierarchy
func (bf *BadFunc) EnterNode(w walker.Walkable) bool {
	switch w.(node.Node).(type) {
	case *expr.FunctionCall:
		fc := w.(*expr.FunctionCall)
		if n, ok := fc.Function.(*name.Name); ok {
			for _, p := range n.Parts {
				if np, ok := p.(*name.NamePart); ok && isBadFunc(np.Value) {
					bf.scorer.Up()
				}
			}
		}
	}

	return true
}
