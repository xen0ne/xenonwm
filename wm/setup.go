package wm

import (
	"log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xwindow"
)

var (
	x *xgbutil.XUtil
	// Root is an xwindow windo corresponding to the root window
	Root *xwindow.Window
)

// Setup sets up event listening on the root window and assigns callbacks
func Setup(xu *xgbutil.XUtil) {
	x = xu
	Root = xwindow.New(x, x.RootWin())
	// setup events on root window
	mask := xproto.EventMaskPropertyChange |
		xproto.EventMaskFocusChange |
		xproto.EventMaskButtonPress |
		xproto.EventMaskButtonRelease |
		xproto.EventMaskStructureNotify |
		xproto.EventMaskSubstructureNotify |
		xproto.EventMaskSubstructureRedirect

	err := Root.Listen(mask)
	if err != nil {
		log.Fatalf("Could not listen to Root window events: %s\n", err)
	}

}