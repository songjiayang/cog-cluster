package handler

import "github.com/songjiayang/cog-cluster/pkg/cog"

type PredictionInput struct {
	Version string `json:"version"`

	cog.Input `json:",inline"`
}
