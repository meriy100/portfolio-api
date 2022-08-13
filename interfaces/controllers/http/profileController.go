package http

import (
	"github.com/meriy100/portfolio-api/usecase/ports"
	"net/http"
)

type profileControllerInputPortFactory func(ports.ProfileOutputPort, ports.PostRepository, ports.ProfileRepository) ports.ProfileInputPort
type profileControllerOutputPortFactory func(w http.ResponseWriter) ports.ProfileOutputPort

type ProfileController struct {
	ProfileRepository ports.ProfileRepository
	PostRepository    ports.PostRepository
	InputPortFactory  profileControllerInputPortFactory
	OutputPortFactory profileControllerOutputPortFactory
}

func NewProfileController(
	postRepository ports.PostRepository,
	profileRepository ports.ProfileRepository,
	inputFactory profileControllerInputPortFactory,
	outputFactory profileControllerOutputPortFactory,
) *ProfileController {
	return &ProfileController{
		profileRepository,
		postRepository,
		inputFactory,
		outputFactory,
	}
}

func (pc *ProfileController) ShowProfile(w http.ResponseWriter, r *http.Request) {
	if err := pc.newInputPort(w).ShowProfile(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProfileController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := pc.newInputPort(w).UpdateProfile(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (pc *ProfileController) newInputPort(w http.ResponseWriter) ports.ProfileInputPort {
	return pc.InputPortFactory(pc.OutputPortFactory(w), pc.PostRepository, pc.ProfileRepository)
}
