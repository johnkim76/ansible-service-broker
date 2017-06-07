package broker

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/openshift/ansible-service-broker/pkg/apb"
	"github.com/pborman/uuid"
)

func RunCommand(bin string, args ...string) {
	cmd := exec.Command(bin, args...) //.CombinedOutput()
	cmd.Stdin = os.Stdin
	cmd.Stdin = os.Stdout
	cmd.Stdin = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}

	return
}

func ProjectRoot() string {
	gopath := os.Getenv("GOPATH")
	rootPath := path.Join(gopath, "src", "github.com", "openshift",
		"ansible-service-broker")
	return rootPath
}

// TODO: This is going to have to be expanded much more to support things like
// parameters (will need to get passed through as metadata
func SpecToService(spec *apb.Spec) Service {
	parameterDescriptors := make(map[string]interface{})
	parameterDescriptors["parameters"] = spec.Parameters
	for k, v := range spec.Metadata {
		parameterDescriptors[k] = v
	}

	retSvc := Service{
		ID:          uuid.Parse(spec.Id),
		Name:        spec.Name,
		Description: spec.Description,
		Tags:        make([]string, len(spec.Tags)),
		Bindable:    spec.Bindable,
		// TODO: CREATE REAL PLAN here since this is where we want to add the
		// schema stuff
		Plans:    plans,
		Metadata: parameterDescriptors,
	}

	copy(retSvc.Tags, spec.Tags)
	return retSvc
}

func StateToLastOperation(state apb.State) LastOperationState {
	switch state {
	case apb.StateInProgress:
		return LastOperationStateInProgress
	case apb.StateSucceeded:
		return LastOperationStateSucceeded
	case apb.StateFailed:
		return LastOperationStateFailed
	default:
		return LastOperationStateFailed
	}
}
