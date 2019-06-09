package cli

import (
	"reflect"
	"testing"

	"github.com/flores95/goref/frameworks"
	"github.com/flores95/goref/frameworks/log"
	"github.com/flores95/goref/frameworks/process"
)

func createMockProcesses() []process.Processor {
	p := new(process.MockProcessor)
	p.DoFunc = func() {
		// nothing to do here :)
	}
	p.NameFunc = func() string {
		return "Mock Processor"
	}
	return []process.Processor{p}
}

func createMockWorkers() []frameworks.Worker {
	w := new(frameworks.MockWorker)
	w.NameFunc = func() string {
		return "Mock Worker"
	}
	w.TypeFunc = func() string {
		return "MOCK"
	}
	w.SetLoggerFunc = func(log.Logger) {
		// nothing to do here :)
	}
	w.GetLoggerFunc = func() log.Logger {
		return new(log.MockLogger)
	}
	return []frameworks.Worker{w}
}

func TestNewApp(t *testing.T) {
	logger := new(log.MockLogger)

	tests := []struct {
		name  string
		wantA App
	}{
		{
			name:  "NewApp should create app with injected logger",
			wantA: App{logger: logger},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := NewApp(logger); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("NewApp() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestApp_GetProcessorByName(t *testing.T) {
	procs := createMockProcesses()
	logger := new(log.MockLogger)

	tests := []struct {
		name     string
		a        App
		procs    []process.Processor
		n        string
		wantProc process.Processor
	}{
		{
			name:     "Should be able to retrieve processor by name",
			a:        NewApp(logger),
			procs:    procs,
			n:        "Mock Processor",
			wantProc: procs[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProc := tt.a.GetProcessorByName(tt.procs, tt.n); !reflect.DeepEqual(gotProc, tt.wantProc) {
				t.Errorf("App.GetProcessorByName() = %v, want %v", gotProc, tt.wantProc)
			}
		})
	}
}
