package midware

import (
	"context"
	"net/http"

	"github.com/agungdhewe/dwtpl"
	"github.com/fgtago/fgweb/appsmodel"
	"github.com/houseme/mobiledetect"
)

// MobileDetect is a middleware function that detects the device used by the user and sets the device type in the request context.
//
// It takes an http.Handler as a parameter and returns an http.Handler.
// The returned http.Handler wraps the next http.Handler with the mobile detection logic.
// The mobile detection is done using the mobiledetect package.
// The detected device type is stored in the request context with the key appsmodel.DeviceKeyName.
// The next http.Handler is then called with the modified request and response writer.
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
