package wechat

import (
	"testing"

	"github.com/iGoogle-ink/gopay/pkg/xlog"
)

func TestPaySignOfJSAPIp(t *testing.T) {
	jsapi, err := client.PaySignOfJSAPI("prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("jsapi:%#v", jsapi)
}

func TestPaySignOfApp(t *testing.T) {
	app, err := client.PaySignOfApp("prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("app:%#v", app)
}

func TestPaySignOfApplet(t *testing.T) {
	applet, err := client.PaySignOfApplet("prepayid")
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debugf("applet:%#v", applet)
}
