package controller

import (
	"github.com/redhat-developer/openshift-jenkins-operator/pkg/controller/jenkinsimage"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, jenkinsimage.Add)
}
