package vessel

import (
	"context"

	dockerClient "github.com/docker/docker/client"
	"github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projecteru2/barrel/store"
	"github.com/projecteru2/barrel/types"
)

// ContainerVessel .
type ContainerVessel interface {
	ListContainers() ([]types.ContainerInfo, error)
	// GetContainerByID(containerID string) (types.ContainerInfo, error)
	UpdateContainer(context.Context, types.ContainerInfo) error
	DeleteContainer(context.Context, types.ContainerInfo) error
}

// Vessel .
type Vessel interface {
	Hostname() string
	ContainerVessel() ContainerVessel
	CalicoIPAllocator() CalicoIPAllocator
	FixedIPAllocator() FixedIPAllocator
}

type vessel struct {
	hostname         string
	containerVessel  ContainerVessel
	fixedIPAllocator FixedIPAllocator
}

// NewVessel .
func NewVessel(hostname string, cliv3 clientv3.Interface, dockerCli *dockerClient.Client, driverName string, stor store.Store) Vessel {
	return vessel{
		hostname:         hostname,
		fixedIPAllocator: NewFixedIPAllocator(NewIPPoolManager(cliv3, dockerCli, driverName, hostname), stor),
	}
}

func (v vessel) Hostname() string {
	return v.hostname
}

func (v vessel) ContainerVessel() ContainerVessel {
	return v.containerVessel
}

func (v vessel) CalicoIPAllocator() CalicoIPAllocator {
	return v.fixedIPAllocator
}

func (v vessel) FixedIPAllocator() FixedIPAllocator {
	return v.fixedIPAllocator
}
