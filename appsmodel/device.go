package appsmodel

import "github.com/agungdhewe/dwtpl"

type Device struct {
	Type  dwtpl.DeviceType
	Model string
	OS    string
}
