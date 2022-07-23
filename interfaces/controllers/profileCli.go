package controllers

import (
	"fmt"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type InputPortFactory func(ports.ProfileOutputPort, ports.PostRepository, ports.ProfileRepository) ports.ProfileInputPort
type OutputPortFactory func() ports.ProfileOutputPort

type ProfileCli struct {
	ProfileRepository ports.ProfileRepository
	PostRepository    ports.PostRepository
	InputPortFactory  InputPortFactory
	OutputPortFactory OutputPortFactory
}

func NewProfileCli(
	postRepository ports.PostRepository,
	profileRepository ports.ProfileRepository,
	inputFactory InputPortFactory,
	outputFactory OutputPortFactory,
) *ProfileCli {
	return &ProfileCli{
		profileRepository,
		postRepository,
		inputFactory,
		outputFactory,
	}
}

func (pc *ProfileCli) UpdateProfile() {
	if err := pc.newInputPort().UpdateProfile(); err != nil {
		fmt.Println(err)
	}
}

func (pc *ProfileCli) newInputPort() ports.ProfileInputPort {
	return pc.InputPortFactory(pc.OutputPortFactory(), pc.PostRepository, pc.ProfileRepository)
}
