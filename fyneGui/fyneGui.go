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

	// 创建下拉选择框
	selectEntry := widget.NewSelect([]string{"CVE_2017_8917", "2"}, func(selected string) {
		// 处理所选项的逻辑
		fmt.Println("Selected:", selected)
	})

	// 设置默认选项
	selectEntry.SetSelected("select a poc")

	var outputLabel = widget.NewLabel("")

	runButton := widget.NewButton("exploit", func() {
		// 从输入框中获取网址信息
		url := urlEntry.Text

		// 获取选择的值
		selected := selectEntry.Selected
		if selected == "" {
			fmt.Println("没有选择要使用的poc!")
			return
		}

		// 执行相应的函数
		switch selected {
		case "CVE_2017_8917":
			_, err := cve.CVE_2017_8917(url)
			if err != nil {
				fmt.Println("CVE_2017_8917出错!", err)
			}

		case "2":
			_, err := http.SendGetRequest(url)
			if err != nil {
				fmt.Println("2出错!", err)
			}
		}

		// 清空错误信息
		outputLabel.SetText("")
	})

	content := container.New(layout.NewBorderLayout(nil, nil, nil, nil),
		container.NewVBox(
			urlEntry,    // 添加输入框
			selectEntry, // 添加下拉选择框
			runButton,
			outputLabel,
		),
	)

	myWindow.SetContent(content)

	myWindow.ShowAndRun()
}
