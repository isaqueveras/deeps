package core

import (
	"log"
	"time"
)

// Logger is used for logging formatted messages.
type Logger interface {
	// Printf must have the same semantics as log.Printf.
	Printf(format string, args ...any)
}

// NewDeeps ...
func NewDeeps(name string, scripts ...func()) {
	now := time.Now()
	log.Println("Iniciando servidor:", name)

	for _, fn := range scripts {
		if fn == nil {
			continue
		}
		go fn()
	}

	// read config file
	// defaultLogger := Logger(log.New(os.Stderr, "", log.LstdFlags))

	// defaultLogger.Printf("Olá mundo")

	// go func() {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/metrics/profile", pprof.Profile)
	// log.Fatal(http.ListenAndServe(":7777", mux))
	// }()

	wait()
	log.Printf("Finalizando servidor %s em %s", name, time.Since(now).String())
}
