package cgroups

import (
	"gdocker/cgroups/subsystem"
	"github.com/sirupsen/logrus"
)

type CGroupManager struct {
	Path string
}

func NewCGroupManager(path string) *CGroupManager {
	return &CGroupManager{path}
}

func (c *CGroupManager) Set(res *subsystem.ResourceConfig) {
	for _, subsystem := range subsystem.Subsystems {
		err := subsystem.Set(c.Path, res)
		if err != nil {
			logrus.Errorf("set %s err: %v", subsystem.Name(), err)
		}
	}
}
