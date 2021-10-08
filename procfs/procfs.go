package procfs

import (
	"github.com/shirou/gopsutil/v3/process"
	"os"
	"syscall"
)

func GetVszRssBytes(pid int) (vszBytes uint64, rssBytes uint64, err error) {
	p, err := newProcess(pid)
	if err != nil {
		return 0, 0, err
	}
	info, err := p.MemoryInfo()
	if err != nil {
		return 0, 0, err
	}
	return info.VMS, info.RSS, nil
}

func newProcess(pid int) (*process.Process, error){
	p, err := process.NewProcess(castPid(pid))
	if err != nil {
		if os.IsNotExist(err) {
			err = syscall.ESRCH
		}
		// Under certain circumstances, the error returned for the
		// underlying read is ESRCH. It seems mostly likely to be a
		// race inside the procfs.
		if pathErr, ok := err.(*os.PathError); ok && pathErr.Err == syscall.ESRCH {
			err = syscall.ESRCH
		}
		return nil, err
	}
	return p, err
}

func castPid(pid int) int32 {
	pid32 := int32(pid);
	if (int(pid32) != pid) {
		panic("only 32 bit pids are supported")
	}
	return pid32
}
