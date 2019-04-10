package ui

import (
	"fmt"
	"github.com/askovpen/gossiped/pkg/msgapi"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func (a *App) ViewMsg(areaId int, msgNum uint32) (string, tview.Primitive, bool, bool) {
	msg, err := msgapi.Areas[areaId].GetMsg(msgNum)
	if err != nil {
		modal := tview.NewModal().
			SetText(err.Error()).
			AddButtons([]string{"Quit"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				a.Pages.SwitchToPage("AreaList")
			})
		return fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum), modal, true, true
	}
	msgapi.Areas[areaId].SetLast(msgNum)
	header := tview.NewTextView().
		SetWrap(false)
	header.SetBorder(true).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorBlue).
		SetTitle(" " + msgapi.Areas[areaId].GetName() + " ").
		SetTitleAlign(tview.AlignLeft).
		SetTitleColor(tcell.ColorYellow)
	htxt := fmt.Sprintf(" Msg  : %-34s %-36s\n",
		fmt.Sprintf("%d of %d", msgNum, msgapi.Areas[areaId].GetCount()), "Pvt")
	htxt += fmt.Sprintf(" From : %-34s %-15s %-18s\n",
		msg.From,
		msg.FromAddr.String(),
		msg.DateWritten.Format("02 Jan 06 15:04:05"))
	htxt += fmt.Sprintf(" To   : %-34s %-15s %-18s\n",
		msg.To,
		msg.ToAddr.String(),
		msg.DateArrived.Format("02 Jan 06 15:04:05"))
	htxt += fmt.Sprintf(" Subj : %-50s",
		msg.Subject)
	header.SetText(htxt)
	body := tview.NewTextView().SetWrap(true).SetWordWrap(true)
	body.SetDynamicColors(true)
	body.SetText(msg.ToView(a.showKludges))
	body.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			a.Pages.SwitchToPage("AreaList")
			a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
		}
	})
	body.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRight {
			if msgNum == msgapi.Areas[areaId].GetCount() {
				a.Pages.SwitchToPage("AreaList")
				a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
			} else {
				if a.Pages.HasPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum+1)) {
					a.Pages.SwitchToPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum+1))
					a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
				} else {
					a.Pages.AddPage(a.ViewMsg(areaId, msgNum+1))
					a.Pages.SwitchToPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum+1))
					a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
				}
			}
		} else if event.Key() == tcell.KeyLeft {
			if msgNum <= 1 {
				a.Pages.SwitchToPage("AreaList")
				a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
			} else {
				if a.Pages.HasPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum-1)) {
					a.Pages.SwitchToPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum-1))
					a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
				} else {
					a.Pages.AddPage(a.ViewMsg(areaId, msgNum-1))
					a.Pages.SwitchToPage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum-1))
					a.Pages.RemovePage(fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum))
				}
			}
		} else if event.Key() == tcell.KeyCtrlH {
			a.showKludges = !a.showKludges
			body.SetText(msg.ToView(a.showKludges))
		} else if event.Key() == tcell.KeyCtrlQ {
		} else if event.Key() == tcell.KeyInsert {
			a.Pages.AddPage(a.InsertMsg(areaId))
			a.Pages.SwitchToPage(fmt.Sprintf("InsertMsg-%s", msgapi.Areas[areaId].GetName()))
		}

		return event
	})

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 6, 1, false).
		AddItem(body, 0, 1, true)
	return fmt.Sprintf("ViewMsg-%s-%d", msgapi.Areas[areaId].GetName(), msgNum), layout, true, true
}
