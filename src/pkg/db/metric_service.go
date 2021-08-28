package db

import (
	"pkg/configuration"
)

type MetricService struct {
	config 	configuration.Configuration
	db		*TestDataCache
}

func NewMetricService(config configuration.Configuration, db *TestDataCache) *MetricService {
	return &MetricService{config, db}
}

