package fyneGui

import (
	"awesomeProject/cve"
	"awesomeProject/http"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func GuiInit() {
	myApp := app.New()
	myWindow := myApp.NewWindow("")

	// 固定窗口大小
	myWindow.Resize(fyne.NewSize(400, 300))

	var urlEntry = widget.NewEntry() // 输入框

	// 创建复选框组
	checkGroup := widget.NewCheckGroup([]string{"CVE_2017_8917", "2"}, func(selected []string) {
		// 处理复选框状态变化的逻辑
		fmt.Println("Selected checkboxes:", selected)
	})

	var outputLabel = widget.NewLabel("")

	runButton := widget.NewButton("exploit", func() {
		// 从输入框中获取网址信息
		url := urlEntry.Text

		// 检查复选框的状态，并执行相应的函数
		var _ string
		if len(checkGroup.Selected) > 0 {
			// 当至少有一个复选框被选中时
			for _, selected := range checkGroup.Selected {
				switch selected {

				case "CVE_2017_8917":
					_, err := cve.CVE_2017_8917(url)
					if err != nil {
						fmt.Println("CVE_2017_8917出错!", err)
					}

				case "2":
					_, err := http.SendGetRequest(url)
					if err != nil {
						fmt.Println("CVE_2017_8917出错!", err)
					}
				}

			}
		} else {
			// 当没有复选框被选中时，默认执行 sendGetRequest 函数
			fmt.Println("没有选择要使用的poc!")
		}
		var err *error

		if err != nil {
			outputLabel.SetText(fmt.Sprintf("Error: %v", err))
			return
		}
		//outputLabel.SetText(_)
	})

	content := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
		container.NewVBox(
			urlEntry,   // 添加输入框
			checkGroup, // 添加复选框组
			runButton,
			outputLabel,
		),
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
