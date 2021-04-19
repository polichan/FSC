package logic

import (
	"fmt"
	"fsc/global"
)

const (
	FSCPanelOptionLogin = iota
	FSCPanelOptionRunTarget
	FSCPanelOptionFreeRun
	FSCPanelOptionExit
	FSCPanelOptionWaitingUserInput
)

type PanelOptionItem struct {
	Option int
	Action func()
	Title  string
}

type IPanel interface {
	DrawTitle()
	DrawOptions()
	DrawOption()
	DrawLine()
	Draw()
	Ask()
	SetOption()
	SetCurrentOption()
	WatchingOption()
	SetOptions()
}

type FSCPanel struct {
	CurrentOption int
	options       []*PanelOptionItem
}

func (F FSCPanel) Draw() {
	F.DrawTitle()
	F.DrawLine()
	F.DrawOptions()
	F.DrawLine()
}

func (P *PanelOptionItem) SetPanelOptionItem(option int, action func(), title string) {
	P.Option = option
	P.Action = action
	P.Title = title
}

func NewFSCPanel() *FSCPanel {
	return &FSCPanel{options: make([]*PanelOptionItem, 0), CurrentOption: FSCPanelOptionWaitingUserInput}
}

func (F *FSCPanel) SetCurrentOption(option int) {
	F.CurrentOption = option
}

func (F *FSCPanel) SetOption(option PanelOptionItem) {
	F.options = append(F.options, &option)
}

func (F *FSCPanel) SetOptions(options []*PanelOptionItem) {
	F.options = options
}

func (F *FSCPanel) WatchingOption() {
	for F.CurrentOption != FSCPanelOptionExit {
		_, _ = fmt.Scanf("%d", &F.CurrentOption)
		hasOption := F.notifyOptionAction()
		if !hasOption {
			global.FSC_LOG.Info("请选择正确的选项！")
		}
		F.Draw()
		F.Ask()
	}
}

func (F *FSCPanel) notifyOptionAction() bool {
	hasOption := false
	// 根据用户的 option 调用相应 Action
	for _, option := range F.options {
		if option.Option == F.CurrentOption {
			hasOption = true
			option.Action()
			break
		}
	}
	return hasOption
}

func (F *FSCPanel) Ask() {
	fmt.Printf("\n请选择选项号: ")
}

func (F *FSCPanel) DrawLine() {
	fmt.Println("\n***********************************************")
}

func (F *FSCPanel) DrawOption(option PanelOptionItem) {
	fmt.Printf("\n*                %d. %s                  *\n", option.Option, option.Title)
}

func (F *FSCPanel) DrawTitle() {
	fmt.Print(`
	     _____ ____   ____ 
	    |  ___/ ___| / ___|
	    | |_  \___ \| |    
	    |  _|  ___) | |___ 
 	    |_|   |____/ \____|
                    `)
	fmt.Printf("\n                  FSC v%s\n", global.FSC_CONFIG.FSC.Version)
}

func (F *FSCPanel) DrawOptions() {
	for _, option := range F.options {
		F.DrawOption(*option)
	}
}
