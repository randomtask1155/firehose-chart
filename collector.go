package main

import (
	"github.com/cloudfoundry/sonde-go/events"
)

// Metric container for all metrics
type Metric struct {
	Origin string                    `json:"origin"`
	Type   events.Envelope_EventType `json:"type"`
	Job    string                    `json:"job"`
	Index  string                    `json:"index"`
	Metric string                    `json:"metric"`
	Value  []float64
}

func (m *Metric) pushValue(v float64) {
	if len(m.Value) >= *maxSeries {
		m.Value = append(m.Value[1:*maxSeries], v)
	} else {
		m.Value = append(m.Value, v)
	}
}

func (m *Metric) pushDelta(d uint64) {
	if len(m.Value) >= *maxSeries {
		m.Value = append(m.Value[1:*maxSeries], float64(d))
	} else {
		m.Value = append(m.Value, float64(d))
	}
}

func parseEnvelope(e *events.Envelope) {

	for i := range mc {
		if e.GetEventType() == mc[i].Type && e.GetOrigin() == mc[i].Origin && e.GetJob() == mc[i].Job && e.GetIndex() == mc[i].Index {
			if e.GetEventType() == events.Envelope_ValueMetric && e.ValueMetric.GetName() == mc[i].Metric {
				mc[i].pushValue(e.ValueMetric.GetValue())
				if archiveEnabled {
					go archiveMetric(e, e.ValueMetric.GetName(), e.ValueMetric.GetValue(), e.ValueMetric.GetUnit())
				}
			} else if e.GetEventType() == events.Envelope_CounterEvent && e.CounterEvent.GetName() == mc[i].Metric {
				mc[i].pushDelta(e.CounterEvent.GetDelta())
				if archiveEnabled {
					go archiveMetric(e, e.CounterEvent.GetName(), e.CounterEvent.GetDelta(), e.CounterEvent.GetTotal())
				}
			}
		}
	}
}
