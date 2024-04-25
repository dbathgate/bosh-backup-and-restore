package deployment

func NewParallelExecutor(maxInFlight int) ParallelExecutor {
	return ParallelExecutor{
		maxInFlight: maxInFlight,
	}
}

type ParallelExecutor struct {
	maxInFlight int
}

func (s ParallelExecutor) Run(executables []Executable) []DeploymentError {
	var errors []DeploymentError

	guard := make(chan bool, s.maxInFlight)
	errs := make(chan DeploymentError, len(executables))

	for _, executable := range executables {
		guard <- true
		go func(executable Executable) {
			errs <- executable.Execute()
			<-guard
		}(executable)
	}

	for range executables {
		err := <-errs
		if err.Errs != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
