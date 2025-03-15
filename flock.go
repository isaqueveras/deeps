package core

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/venustre"
)

type flock struct {
	w *venustre.Watch
}

// Flock returns a structure of the zap log
// implementation and methods of the lib's main interface.
func Flock(id string, name string) *flock {
	return &flock{w: venustre.Watcher(id, name, venustre.Impl(&attach{}))}
}

// Attach ...
func (f *flock) Attach(id, name string, script func(ctx *venustre.Context) error, interval ...time.Duration) {
	var opts = []venustre.Option{
		venustre.WithID(venustre.ID(id)),
		venustre.WithName(name),
		venustre.WithScript(script),
	}

	if len(interval) != 0 {
		opts = append(opts, venustre.WithInterval(interval[0]))
	} else {
		opts = append(opts, venustre.WithNotUseLoop())
	}

	go f.w.Go(opts...)
}

func (f *flock) Wait() {
	f.w.Wait()
}

// wait responsible for keeping routines running
func wait() {
	venustre.Wait()
	// zap.L().Info("Shutting down script server")
}

type attach struct{}

// Go calls the given function in a new goroutine.
// It blocks until the new goroutine can be added without the number of
// active goroutines in the group exceeding the configured limit.
func (a *attach) Go(fn func() error) {
	if err := fn(); err != nil {
		panic(err)
	}
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (a *attach) Wait() error { return nil }

// Before contém a lógica para ser executada antes da execução de uma rotina.
func (a *attach) Before(ctx *venustre.Context) error {
	ctx.Id = venustre.ID(strings.Split(uuid.New().String(), "-")[4])
	// ctx.Info("Inicializando rotina: "+ctx.Name,
	// 	Any("cid", ctx.Id),
	// 	Any("rid", ctx.RoutineID),
	// 	Any("wid", ctx.Watcher.Id),
	// )
	return nil
}

// After contains the logic to be executed after executing a routine.
func (a *attach) After(ctx *venustre.Context) error {
	ctx.NewIndicator("latency").Add(ctx.GetLatency())
	// ctx.Info("Finalizando rotina: "+ctx.Name,
	// 	Any("cid", ctx.Id),
	// 	Any("rid", ctx.RoutineID),
	// 	Any("wid", ctx.Watcher.Id),
	// )
	return nil
}

// Init contains the logic for saving information from a routine to the database.
func (a *attach) Init(ctx *venustre.Context) error {
	// TODO: (@isaque-brisa) save routine information in the database
	// ctx.Info("Inicializando watcher: "+ctx.Watcher.Name,
	// 	Any("rid", ctx.RoutineID),
	// 	Any("wid", ctx.Watcher.Id),
	// )
	return nil
}

// Event receives events from routines (scripts).
// Processes each type of event data.
func (a *attach) Event(ctx *venustre.Context, event venustre.Event) {
	if metric, ok := event.(venustre.EventMetric); ok {
		for _, indicator := range metric.Indicators {
			_ = indicator
			// ctx.Info("Indicador da rotina: "+ctx.Name,
			// 	Any("type", "indicator"),
			// 	Any("key", indicator.GetKey()),
			// 	Any("value", indicator.GetValue()),
			// 	Any("cid", ctx.Id),
			// 	Any("rid", ctx.RoutineID),
			// 	Any("wid", ctx.Watcher.Id),
			// )
		}

		for _, item := range metric.Histograms {
			_ = item

			// values, times := item.GetValues()
			// ctx.Info("Histograma da rotina: "+ctx.Name,
			// 	Any("type", "histogram"),
			// 	Any("key", item.GetKey()),
			// 	Any("values", values),
			// 	Any("times", times),
			// 	Any("cid", ctx.Id),
			// 	Any("rid", ctx.RoutineID),
			// 	Any("wid", ctx.Watcher.Id),
			// )
		}

		if metric.Metadata != nil && len(metric.Metadata) != 0 {
			// ctx.Info("Metadados da rotina: "+ctx.Name,
			// 	Any("type", "metadata"),
			// 	zap.Object("metadata", utils.JSONB(metric.Metadata)),
			// 	Any("cid", ctx.Id),
			// 	Any("rid", ctx.RoutineID),
			// 	Any("wid", ctx.Watcher.Id),
			// )
		}
	}
}
