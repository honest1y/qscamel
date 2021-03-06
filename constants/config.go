package constants

import "runtime"

// DefaultConfigContent is the default config config.
const DefaultConfigContent = `concurrency: 0
log_file: ~/.qscamel/qscamel.log
log_level: info
pid_file: ~/.qscamel/qscamel.pid
database_file: ~/.qscamel/db
`

var (
	// DefaultConcurrency is default num of objects being migrated concurrently.
	DefaultConcurrency = runtime.NumCPU() * 10
)

// Path store all path related constants.
const (
	Path         = "~/.qscamel"
	ConfigPath   = Path + "/qscamel.yaml"
	DatabasePath = Path + "/db"
	LogPath      = Path + "/qscamel.log"
	PIDPath      = Path + "/qscamel.pid"
)

// DefaultMultipartBoundarySize is the default multipart boundary size.
// 2 * 1024 * 1024 * 1024 = 2147483648 B = 2 GB
const DefaultMultipartBoundarySize = 2147483648
