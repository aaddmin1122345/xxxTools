// fyneGuiInterface.go

package fyneGui

// TextUpdater 是更新 GUI 文本的接口
type TextUpdater interface {
	UpdateText(text string) error
}
