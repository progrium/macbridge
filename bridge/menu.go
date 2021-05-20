package bridge

import (
	"encoding/base64"

	"github.com/progrium/macbridge/handle"
	"github.com/progrium/macbridge/shell"
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type Menu struct {
	shell.Menu
}

func (m *Menu) Resource() (*handle.Handle, interface{}) {
	return &m.Handle, &m.Menu
}

func (m *Menu) Apply(target objc.Object) (objc.Object, error) {
	if target == nil {
		menu := cocoa.NSMenu_New()
		menu.SetAutoenablesItems(true)
		for _, i := range m.Items {
			menu.AddItem(NSMenuItem(i))
		}
		target = menu.Object
	}
	return target, nil
}

func NSMenuItem(i shell.MenuItem) cocoa.NSMenuItem {
	if i.Separator {
		return cocoa.NSMenuItem_Separator()
	}
	obj := cocoa.NSMenuItem_New()
	obj.SetTitle(i.Title)
	obj.SetEnabled(i.Enabled)
	obj.SetToolTip(i.Tooltip)
	if i.Checked {
		obj.SetState(cocoa.NSControlStateValueOn)
	}
	if i.Icon != "" {
		b, err := base64.StdEncoding.DecodeString(i.Icon)
		if err == nil {
			data := core.NSData_WithBytes(b, uint64(len(b)))
			img := cocoa.NSImage_InitWithData(data)
			img.SetSize(core.Size(16, 16))
			obj.SetImage(img)
		}
	}
	if i.Title == "Quit" {
		obj.SetTarget(cocoa.NSApp())
		obj.SetAction(objc.Sel("terminate:"))
	}
	// if i.OnClick != nil && i.OnClick.Caller != nil {
	// 	t, sel := i.OnClick.Callback()
	// 	obj.SetTarget(t)
	// 	obj.SetAction(sel)
	// }
	return obj
}
