package db

import (
	"errors"
	"github.com/gofrs/uuid"
)

// ---- MetricDataModel ----
type MetricDataModel struct {
	Id			string	`json:"id,omitempty"`
	Key 		string 	`json:"key,omitempty"`
	Value 		int 	`json:"value,omitempty"`
	EntryTime	string	`json:"entry_time,omitempty"`
}

// ---- NewTestDataCache ----
func NewTestDataCache(td []*MetricDataModel) *TestDataCache {
	// returns a pointer to a new testDataCache struct
	return &TestDataCache{
		td,
	}
}

// --- TestDataCache ----
type TestDataCache struct {
	MetricData []*MetricDataModel
}

// ---- genUuid ----
func genId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

// ---- GetMetricData ----
func (tc *TestDataCache) GetMetricData() []*MetricDataModel {
	// returns all jsonData in the testDataCache
	return tc.MetricData
}

// ---- GetMetric ----
func (tc *TestDataCache) GetMetric(id string) (*MetricDataModel, error) {
	// returns a specific jsonData entry from the testDataCache
	for _, td := range tc.MetricData {
		if td.Id == id {
			return td, nil
		}
	}
	return &MetricDataModel{}, errors.New(id + " Not found!")
}

// ---- DelMetric ----
func (tc *TestDataCache) DelMetric(id string) error {
	// returns a specific jsonData entry from the testDataCache
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

// ---- PostMetric -----
func (tc *TestDataCache) PostMetric(td *MetricDataModel) {
	// adds a new jsonData entry to the testDataCache
	td.Id = genId()
	tc.MetricData = append(tc.MetricData, td)
}