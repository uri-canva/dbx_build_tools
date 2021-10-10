package procfs

import (
	"github.com/shirou/gopsutil/v3/process"
)

func GetProcessDescendants(parentPid int) ([]int, error) {
	parentProcess, err := newProcess(castPid(parentPid))
	if err != nil {
		return nil, err
	}

	descendantsPids := []int{}
	stack := []*process.Process{parentProcess}
	var lastErr error
	for len(stack) > 0 {
		top := len(stack) - 1
		parent := stack[top]
		stack = stack[:top]
		children, err := children(parent)
		if err != nil {
			lastErr = err
			continue
		}
		for _, child := range children {
			descendantsPids = append(descendantsPids, int(child.Pid))
			stack = append(stack, child)
		}
	}
	return descendantsPids, lastErr
}

// Return all immediate child process for a pid
func ChildPids(pid int) ([]int, error) {
	p, err := newProcess(castPid(pid))
	if err != nil {
		return nil, err
	}

	childPids := []int{}
	children, err := p.Children()
	if err != nil {
		return nil, err
	}

	for _, child := range children {
		childPids = append(childPids, int(child.Pid))
	}
	return childPids, nil
}
