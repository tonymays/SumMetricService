package db

import (
	//"fmt"
	"pkg"
	"pkg/configuration"
	"time"
)

// ---- MetricService ----
type MetricService struct {
	config 	configuration.Configuration
	db		*TestDataCache
}

// ---- NewMetricService -----
func NewMetricService(config configuration.Configuration, db *TestDataCache) *MetricService {
	return &MetricService{config, db}
}

// ---- MetricService.AddMetric ----
func (rcvr *MetricService) AddMetric(m root.Metric) (root.Metric, error) {
	currentTime := time.Now().UTC()
	mdm := MetricDataModel{}
	mdm.Key = m.Key
	mdm.Value = m.Value
	mdm.EntryTime = currentTime.String()
	m.EntryTime = currentTime.String()
	rcvr.db.PostMetric(&mdm)
	return m, nil
}

// ---- MetricService.GetMetrics ----
func (rcvr *MetricService) GetMetrics() ([]root.Metric, error) {
	data := rcvr.db.GetMetricData()
	var ms []root.Metric
	for _, d := range data {
		var m root.Metric
		m.Id = d.Id
		m.Key = d.Key
		m.Value = d.Value
		m.EntryTime = d.EntryTime
		ms = append(ms, m)
	}
	if len(ms) == 0 {
		return []root.Metric{}, nil
	}
	return ms, nil
}

// ---- MetricService.ShowActiveMetrics ----
func (rcvr *MetricService) ShowActiveMetrics() ([]root.Metric, error) {
	data := rcvr.db.GetMetricData()
	var ms []root.Metric
	for _, d := range data {
		var m root.Metric
		m.Id = d.Id
		m.Key = d.Key
		m.Value = d.Value
		m.EntryTime = d.EntryTime
		if rcvr.CanCount(m) {
			ms = append(ms, m)
		}
	}
	if len(ms) == 0 {
		return []root.Metric{}, nil
	}
	return ms, nil
}

// ---- MetricService.SumMetrics ----
func (rcvr *MetricService) SumMetrics(key string) (root.MetricSum, error) {
	if rcvr.config.ClearOnSum == "on" {
		rcvr.ClearOutdatedMetrics(key)
	}
	data := rcvr.db.GetMetricData()
	var ms root.MetricSum
	ms.Key = key
	ms.Sum = 0
	var m root.Metric
	for _, d := range data {
		m.EntryTime = d.EntryTime
		if rcvr.CanCount(m) {
			ms.Sum += d.Value
		}
	}
	return ms, nil
}

// ---- MetricService.ClearOutDatedMetrics ----
func (rcvr *MetricService) ClearOutdatedMetrics(key string) error {
	data := rcvr.db.GetMetricData()
	for _, d := range data {
		if d.Key == key {
			var m root.Metric
			m.EntryTime = d.EntryTime
			if !rcvr.CanCount(m) {
				rcvr.db.DelMetric(d.Id)
			}
		}
	}
	return nil
}

func (rcvr *MetricService) CanCount(m root.Metric) bool {
	countStrategy := rcvr.config.CountStrategy
	now := time.Now()
	then := now.Add(time.Duration(-countStrategy) * time.Minute)
	return (m.EntryTime >= then.String())
}
