package midware

import (
	"context"
	"net/http"

	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/houseme/mobiledetect"
)

func MobileDetect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device := appsmodel.Device{}

		// detek device yang dipakai user saat ini
		currentDevice := mobiledetect.New(r, nil)
		if currentDevice.IsTablet() {
			device.Type = dwtpl.DeviceTablet
		} else if currentDevice.IsMobile() {
			device.Type = dwtpl.DeviceMobile
		} else {
			device.Type = dwtpl.DeviceDesktop
		}

		ctx := context.WithValue(r.Context(), appsmodel.DeviceKeyName, device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
