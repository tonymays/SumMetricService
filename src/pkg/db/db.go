package db

import (
	"errors"
	"strconv"
)

// MetricDataModel
type MetricDataModel struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// TestDataCache
type TestDataCache struct {
	MetricData []*MetricDataModel
}

// GetMetricData - returns all jsonData in the testDataCache
func (tc *TestDataCache) GetMetricData() []*MetricDataModel {
	return tc.MetricData
}

// GetMetric - returns a specific jsonData entry from the testDataCache
func (tc *TestDataCache) GetMetric(id string) (*MetricDataModel, error) {
	for _, td := range tc.MetricData {
		if td.Id == id {
			return td, nil
		}
	}
	return &MetricDataModel{}, errors.New(id + " Not found!")
}

// DelMetric - returns a specific jsonData entry from the testDataCache
func (tc *TestDataCache) DelMetric(id string) error {
	var upArray []*MetricDataModel
	var found bool
	for _, td := range tc.MetricData {
		if td.Id == id {
			found = true
		} else {
			upArray = append(upArray, td)
		}
	}
	if !found {
		return errors.New(id + " Not found!")
	}
	tc.MetricData = upArray
	return nil
}

// PostMetric - adds a new jsonData entry to the testDataCache
func (tc *TestDataCache) PostMetric(td *MetricDataModel) {
	td.Id = strconv.Itoa(len(tc.MetricData))
	tc.MetricData = append(tc.MetricData, td)
}

// NewTestDataCache returns a pointer to a new testDataCache struct
func NewTestDataCache(td []*MetricDataModel) *TestDataCache {
	return &TestDataCache{
		td,
	}
}