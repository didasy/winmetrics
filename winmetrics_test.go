package winmetrics

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getSensorsQueryNamespaceFunction(query string, dst interface{}, namespace string) (err error) {
	dv := dst.(*[]Sensor)
	*dv = append(*dv, Sensor{})

	return
}

func getSensorsQueryNamespaceFunctionError(query string, dst interface{}, namespace string) (err error) {
	err = errors.New("error")

	return
}

func getHardwaresQueryNamespaceFunction(query string, dst interface{}, namespace string) (err error) {
	dv := dst.(*[]Hardware)
	*dv = append(*dv, Hardware{})

	return
}

func getHardwaresQueryNamespaceFunctionError(query string, dst interface{}, namespace string) (err error) {
	err = errors.New("error")

	return
}

func TestGetSensors(t *testing.T) {
	sensors, err := GetSensors(getSensorsQueryNamespaceFunction)
	assert.NoError(t, err)
	assert.NotEmpty(t, sensors)
}

func TestGetSensorsNil(t *testing.T) {
	sensors, err := GetSensors(nil)
	assert.NoError(t, err)
	assert.Empty(t, sensors)
}

func TestGetSensorsError(t *testing.T) {
	sensors, err := GetSensors(getSensorsQueryNamespaceFunctionError)
	assert.Error(t, err)
	assert.Empty(t, sensors)
}

func TestGetHardwares(t *testing.T) {
	hws, err := GetHardwares(getHardwaresQueryNamespaceFunction)
	assert.NoError(t, err)
	assert.NotEmpty(t, hws)
}

func TestGetHardwaresNil(t *testing.T) {
	hws, err := GetHardwares(nil)
	assert.NoError(t, err)
	assert.Empty(t, hws)
}

func TestGetHardwaresError(t *testing.T) {
	hws, err := GetHardwares(getHardwaresQueryNamespaceFunctionError)
	assert.Error(t, err)
	assert.Empty(t, hws)
}
