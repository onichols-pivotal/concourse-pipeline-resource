package out

import (
	"fmt"

	"github.com/robdimsdale/concourse-pipeline-resource/concourse"
	"github.com/robdimsdale/concourse-pipeline-resource/logger"
)

const (
	apiPrefix = "/api/v1"
)

type OutCommand struct {
	logger        logger.Logger
	binaryVersion string
}

func NewOutCommand(
	binaryVersion string,
	logger logger.Logger,
) *OutCommand {
	return &OutCommand{
		logger:        logger,
		binaryVersion: binaryVersion,
	}
}

func (c *OutCommand) Run(input concourse.OutRequest) (concourse.OutResponse, error) {
	if input.Source.Target == "" {
		return concourse.OutResponse{}, fmt.Errorf("%s must be provided", "target")
	}

	if input.Source.Username == "" {
		return concourse.OutResponse{}, fmt.Errorf("%s must be provided", "username")
	}

	if input.Source.Password == "" {
		return concourse.OutResponse{}, fmt.Errorf("%s must be provided", "password")
	}

	c.logger.Debugf("Received input: %+v\n", input)

	return concourse.OutResponse{}, fmt.Errorf("out is not implemented yet")
}
