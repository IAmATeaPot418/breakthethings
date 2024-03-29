package customtypes

import (
	v1 "k8s.io/api/core/v1"
)

// PodSpec is a description of a pod.
// This is a wrapper around v1.PodSpec with custom utility methods
type PodSpec struct {
	// This overrides the Containers field in the v1.PodSpec, therefore making sure that callers
	// only access the Containers array through utility functions in this package.
	Containers struct{}
	v1.PodSpec
}

// AllContainers returns a list of all containers in the Pod, both Init and Regular
func (p *PodSpec) AllContainers() []v1.Container {
	allContainers := make([]v1.Container, 0, len(p.PodSpec.InitContainers)+len(p.PodSpec.Containers))
	allContainers = append(allContainers, p.PodSpec.InitContainers...)
	allContainers = append(allContainers, p.PodSpec.Containers...)
	return allContainers
}

// NonInitContainers returns a list of all regular (non-init) containers in the Pod
func (p *PodSpec) NonInitContainers() []v1.Container {
	return p.PodSpec.Containers
}

// InitContainers returns a list of all init containers in the Pod
func (p *PodSpec) InitContainers() []v1.Container {
	return p.PodSpec.InitContainers
}
