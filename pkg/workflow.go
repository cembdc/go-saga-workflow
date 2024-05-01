package workflow

import "log"

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
	for stepName, stepFunc := range w.Steps {
		log.Printf("[Executing step] %s", stepName)

		if err := stepFunc.Transaction(); err != nil {
			log.Printf("[Error executing step] %s: %s", stepName, err)

			for n, f := range w.Steps {
				log.Printf("[Compensating step] %s", n)
				f.Compensate()

				if stepName == n {
					break
				}
			}
		}
	}

	return nil
}
