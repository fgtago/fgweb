package midware

import (
	"context"
	"net/http"

	"github.com/fgtago/fgweb/appsmodel"
	"github.com/houseme/mobiledetect"
)

func MobileDetect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device := appsmodel.Device{}

		// detek device yang dipakai user saat ini
		currentDevice := mobiledetect.New(r, nil)
		if currentDevice.IsTablet() {
			device.Type = "tablet"
		} else if currentDevice.IsMobile() {
			device.Type = "mobile"
		} else {
			device.Type = "desktop"
		}

		ctx := context.WithValue(r.Context(), appsmodel.DeviceKeyName, device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
