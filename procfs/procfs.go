package procfs

import (
	"os"
	"syscall"

	"github.com/shirou/gopsutil/v3/process"
)

func GetVszRssBytes(pid int) (vszBytes uint64, rssBytes uint64, err error) {
	p, err := newProcess(castPid(pid))
	if err != nil {
		return 0, 0, err
	}
	info, err := p.MemoryInfo()
	if err != nil {
		return 0, 0, err
	}
	return info.VMS, info.RSS, nil
}

func newProcess(pid int32) (*process.Process, error) {
	p, err := process.NewProcess(pid)
	return p, handleESRCH(err)
}

func children(p *process.Process) ([]*process.Process, error) {
	cs, err := p.Children()
	// gopsutil uses pgrep to find children, so we need to handle
	// errors that occur when the process has no children.
	if err != nil && err.Error() == "exit status 1" {
		return []*process.Process{}, nil
	}
	return cs, handleESRCH(err)
}

func handleESRCH(err error) error {
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err = syscall.ESRCH
	}
	// Under certain circumstances, the error returned for the
	// underlying read is ESRCH. It seems mostly likely to be a
	// race inside the procfs.
	if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ESRCH {
		err = syscall.ESRCH
	}
	return err
}

func castPid(pid int) int32 {
	pid32 := int32(pid)
	if int(pid32) != pid {
		panic("only 32 bit pids are supported")
	}
	return pid32
}
