package usecase

import (
	"context"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type ProfileInteractor struct {
	outputPort                ports.ProfileOutputPort
	postRepository            ports.PostRepository
	profileRepository         ports.ProfileRepository
	contentDeliveryRepository ports.ContentDeliveryRepository
}

func NewProfileInteractor(outputPort ports.ProfileOutputPort, postRepository ports.PostRepository, profileRepository ports.ProfileRepository, contentDeliveryRepository ports.ContentDeliveryRepository) ports.ProfileInputPort {
	return &ProfileInteractor{
		outputPort,
		postRepository,
		profileRepository,
		contentDeliveryRepository,
	}
}
func (pi *ProfileInteractor) ShowProfile() error {
	profile, err := pi.profileRepository.Find()
	if err != nil {
		return pi.outputPort.OutputFindProfileError(err)
	}

	return pi.outputPort.OutputProfile(profile)
}

func (pi *ProfileInteractor) UpdateProfile() error {
	post, err := pi.postRepository.FetchPost(253)
	if err != nil {
		return pi.outputPort.OutputFetchPostError(err)
	}

	profile, err := post.ToProfile()
	if err != nil {
		return pi.outputPort.OutputToProfileError(err)
	}
	if err := pi.profileRepository.Save(profile); err != nil {
		return pi.outputPort.OutputProfileSaveError(err)
	}

	if err := pi.contentDeliveryRepository.Deploy(context.Background()); err != nil {
		return pi.outputPort.OutputDeployError(err)
	}

	return pi.outputPort.OutputSuccessUpdate()
}
