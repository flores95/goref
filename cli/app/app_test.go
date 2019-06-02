package app

import (
	"reflect"
	"testing"

	"github.com/flores95/golang-curriculum-c-5/cli/frameworks"
	"github.com/flores95/golang-curriculum-c-5/cli/frameworks/logging"
	"github.com/flores95/golang-curriculum-c-5/cli/processes"
)

func createMockProcesses() []processes.Processor {
	p := new(processes.MockProcessor)
	p.DoFunc = func() {
		// nothing to do here :)
	}
	p.NameFunc = func() string {
		return "Mock Processor"
	}
	return []processes.Processor{p}
}

func createMockWorkers() []frameworks.Worker {
	w := new(frameworks.MockWorker)
	w.NameFunc = func() string {
		return "Mock Worker"
	}
	w.TypeFunc = func() string {
		return "MOCK"
	}
	w.SetLoggerFunc = func(logging.Logger) {
		// nothing to do here :)
	}
	w.GetLoggerFunc = func() logging.Logger {
		return new(logging.MockLogger)
	}
	return []frameworks.Worker{w}
}

func TestNewApp(t *testing.T) {
	procs := createMockProcesses()
	fws := createMockWorkers()

	tests := []struct {
		name  string
		wantA App
	}{
		{
			name: "NewApp should take processes and workers and create a new App",
			wantA: App{
				procs: procs,
				fws:   fws,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := NewApp(procs, fws); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("NewApp() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func TestApp_GetProcessorByName(t *testing.T) {
	procs := createMockProcesses()
	fws := createMockWorkers()

	tests := []struct {
		name     string
		a        App
		n        string
		wantProc processes.Processor
	}{
		{
			name:     "Should be able to retrieve processor by name",
			a:        NewApp(procs, fws),
			n:        "Mock Processor",
			wantProc: procs[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotProc := tt.a.GetProcessorByName(tt.n); !reflect.DeepEqual(gotProc, tt.wantProc) {
				t.Errorf("App.GetProcessorByName() = %v, want %v", gotProc, tt.wantProc)
			}
		})
	}
}
