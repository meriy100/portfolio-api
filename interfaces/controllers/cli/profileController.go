package cli

import (
	"fmt"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type profileControllerInputPortFactory func(ports.ProfileOutputPort, ports.PostRepository, ports.ProfileRepository, ports.ContentDeliveryRepository) ports.ProfileInputPort
type profileControllerOutputPortFactory func() ports.ProfileOutputPort

type ProfileCli struct {
	ProfileRepository         ports.ProfileRepository
	PostRepository            ports.PostRepository
	contentDeliveryRepository ports.ContentDeliveryRepository
	InputPortFactory          profileControllerInputPortFactory
	OutputPortFactory         profileControllerOutputPortFactory
}

func NewProfileCli(
	postRepository ports.PostRepository,
	profileRepository ports.ProfileRepository,
	contentDeliveryRepository ports.ContentDeliveryRepository,
	inputFactory profileControllerInputPortFactory,
	outputFactory profileControllerOutputPortFactory,
) *ProfileCli {
	return &ProfileCli{
		profileRepository,
		postRepository,
		contentDeliveryRepository,
		inputFactory,
		outputFactory,
	}
}

func (pc *ProfileCli) ShowProfile() {
	if err := pc.newInputPort().ShowProfile(); err != nil {
		fmt.Println(err)
	}
}

func (pc *ProfileCli) UpdateProfile() {
	if err := pc.newInputPort().UpdateProfile(); err != nil {
		fmt.Println(err)
	}
}

func (pc *ProfileCli) newInputPort() ports.ProfileInputPort {
	return pc.InputPortFactory(pc.OutputPortFactory(), pc.PostRepository, pc.ProfileRepository, pc.contentDeliveryRepository)
}
