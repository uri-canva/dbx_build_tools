package procfs

func GetProcessDescendents(parentPid int) ([]int, error) {
	// TODO: gopsutil doesn't implement this, have to implement it ourselves
	// https://github.com/shirou/gopsutil/issues/111
	return ChildPids(parentPid);
}

// Return all immediate child process for a pid
func ChildPids(pid int) ([]int, error) {
	p, err := newProcess(pid)
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
