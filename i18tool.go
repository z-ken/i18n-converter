package main

import (
	"encoding/json"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func main() {
	err := ui.Main(func() {

		input := ui.NewMultilineEntry()
		locale := ui.NewCombobox()
		locale.Append("en_US")
		locale.Append("zh_CN")
		locale.SetSelected(0)
		output := ui.NewMultilineEntry()

		button := ui.NewButton("Convert...")
		echo := ui.NewLabel("")

		tab := ui.NewTab()

		form1 := ui.NewForm()
		form2 := ui.NewForm()

		form1.Append("i18n json:", input, false)
		form1.Append("Locale:", locale, false)
		form1.Append("", button, false)
		form1.Append("", echo, false)

		form2.Append("i18n sql:", output, false)

		tab.Append("json", form1)
		tab.Append("sql", form2)

		window := ui.NewWindow("i18n tool", 700, 500, false)
		window.SetMargined(true)
		window.SetChild(tab)

		input.OnChanged(func(entry *ui.MultilineEntry) {
			echo.SetText("")
		})


		button.OnClicked(func(*ui.Button) {

			output.SetText("")

			var maps map[string]string
			if err := json.Unmarshal([]byte(input.Text()), &maps); err != nil {
				echo.SetText("Invalid json format !")
				return
			}

			output.Append("INSERT INTO core_sys.`sys_i18n_appearance` \n" +
			"(`key`, `content`, `locale`, `system_code`, `created_time`, `created_by`) \n" +
			"VALUES \n")


			var localeStr string
			if locale.Selected() == 0 {
				localeStr = "en_US"
			} else {
				localeStr = "zh_CN"
			}

			ml := len(maps)
			i := 0
			for k,v := range maps {
				output.Append("(\"" + k +"\", \"" + v + "\", \"" + localeStr + "\", \"CSL\", NOW(), 0)")
				if i++; i < ml {
					output.Append(",\n")
				}
			}

			echo.SetText("SQL generated successfully !")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
