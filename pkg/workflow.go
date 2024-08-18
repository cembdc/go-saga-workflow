package workflow

import (
	"log"
	"sort"
)

type LocalTransaction func() error
type CompensatingAction func() error

type Workflow struct {
	Steps map[string]SagaStep
}

type SagaStep struct {
	Transaction LocalTransaction
	Compensate  CompensatingAction
}

func NewWorkflow() *Workflow {
	return &Workflow{
		Steps: make(map[string]SagaStep),
	}
}

func (w *Workflow) AddStep(name string, step SagaStep) {
	w.Steps[name] = step
}

func (w *Workflow) Execute() error {
	w.sortSteps()

	var completedSteps []string

	for stepName, stepFunc := range w.Steps {
		log.Printf("[Executing step] %s", stepName)

		if err := stepFunc.Transaction(); err != nil {
			log.Printf("[Error executing step] %s: %s", stepName, err)

			// Perform compensating actions for all previously completed steps
			for i := len(completedSteps) - 1; i >= 0; i-- {
				n := completedSteps[i]
				log.Printf("[Compensating step] %s", n)
				if err := w.Steps[n].Compensate(); err != nil {
					log.Printf("[Error compensating step] %s: %s", n, err)
					// Handle compensation error if needed
				}
			}

			// Stop execution as soon as an error occurs
			return err
		}

		// Track successfully completed steps
		completedSteps = append(completedSteps, stepName)
	}

	return nil
}

func (w *Workflow) sortSteps() {
	keys := make([]string, 0, len(w.Steps))
	for k := range w.Steps {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	sortedSteps := make(map[string]SagaStep, len(w.Steps))
	for _, k := range keys {
		sortedSteps[k] = w.Steps[k]
	}

	w.Steps = sortedSteps
}
