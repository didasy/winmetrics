package winmetrics

import (
	"github.com/pkg/errors"
	"github.com/yusufpapurcu/wmi"
)

const (
	Namespace     = `ROOT\LibreHardwareMonitor`
	QuerySensor   = `SELECT * FROM Sensor`
	QueryHardware = `SELECT * FROM Hardware`
)

type Combined struct {
	Sensors   []Sensor   `json:"sensors,omitempty"`
	Hardwares []Hardware `json:"hardwares,omitempty"`
}

type Sensor struct {
	Name       string  `json:"name"`
	SensorType string  `json:"sensor_type"`
	Identifier string  `json:"identifier"`
	Value      float32 `json:"value"`
	Max        float32 `json:"max"`
	Min        float32 `json:"min"`
}

var DefaultWMIQueryNamespace = wmi.QueryNamespace

// GetSensors accept QueryNamespaceFn so it would be easier to test
func GetSensors(queryNamespace QueryNamespaceFn) (sensors []Sensor, err error) {
	if queryNamespace == nil {
		queryNamespace = DefaultWMIQueryNamespace
	}

	err = queryNamespace(QuerySensor, &sensors, Namespace)
	if err != nil {
		err = errors.Wrap(err, "failed to query for sensors")
		return
	}

	return
}

type Hardware struct {
	Name         string `json:"name"`
	HardwareType string `json:"hardware_type"`
	Identifier   string `json:"identifier"`
}

// GetHardwares accept QueryNamespaceFn so it would be easier to test
func GetHardwares(queryNamespace QueryNamespaceFn) (hardwares []Hardware, err error) {
	if queryNamespace == nil {
		queryNamespace = DefaultWMIQueryNamespace
	}

	err = queryNamespace(QueryHardware, &hardwares, Namespace)
	if err != nil {
		err = errors.Wrap(err, "failed to query for hardwares")
		return
	}

	return
}

type QueryNamespaceFn func(query string, dst interface{}, namespace string) (err error)
