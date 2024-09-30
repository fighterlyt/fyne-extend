package fyne_extend

import (
	"sync"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// FormItem 表单元素
type FormItem struct {
	Text             string         // 文字
	Kind             InputKind      // 类型
	input            *widget.Entry  // 底层的entry
	data             binding.String // 底层的数据
	Key              string         // key
	*widget.FormItem                // 底层的表单元素
}

func NewFormItem(text string, kind InputKind, key string) *FormItem {
	return &FormItem{
		Text: text,
		Kind: kind,
		Key:  key,
	}
}

func (f *FormItem) Build() {
	f.FormItem = &widget.FormItem{
		Text: f.Text,
	}

	f.input, f.data = f.Kind.Build()
}

func (f *FormItem) Base() *widget.FormItem {
	return f.FormItem
}

type InputKind int

const (
	InputText InputKind = iota + 1
	InputPassword
)

func (i InputKind) Build() (input *widget.Entry, data binding.String) {
	switch i {
	case InputText:
		input = widget.NewEntry()
	case InputPassword:
		input = widget.NewPasswordEntry()
	default:
		input = widget.NewPasswordEntry()
	}

	data = binding.NewString()

	input.Bind(data)

	return input, data
}

// Form 表单
type Form struct {
	*widget.Form
	items map[string]*FormItem
	lock  *sync.RWMutex
}

func NewForm(items []*FormItem, submit func()) *Form {
	var (
		result = &Form{
			items: make(map[string]*FormItem, len(items)),
			lock:  &sync.RWMutex{},
		}

		formItems = make([]*widget.FormItem, 0, len(items))
	)

	for i := range items {
		items[i].Build()
		result.items[items[i].Key] = items[i]
		formItems = append(formItems, items[i].Base())
	}

	result.Form = &widget.Form{
		OnSubmit:   submit,
		SubmitText: `保存`,
		CancelText: `取消`,
		Items:      formItems,
	}

	return result
}

func (f *Form) GetString(key string) string {
	f.lock.RLock()
	defer f.lock.RUnlock()

	item, _ := f.items[key]

	if item == nil {
		return ``
	}

	result, _ := item.data.Get() // 注意:binding.String 是没有错误

	return result
}

func (f *Form) SetString(key, value string) {
	f.lock.Lock()
	defer f.lock.Unlock()

	item, _ := f.items[key]

	if item == nil {
		return
	}

	_ = item.data.Set(value) // 注意:binding.String 是没有错误
}
