package gssc

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"gssc/pkg/gssc"
	"strconv"
	"strings"
)

var todaysHeadersCmd = &cobra.Command{
	Use:     "bugun",
	Aliases: []string{"today"},
	Short:   "Gunun basliklarini (yeniden eskiye) listeler",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		res := gssc.HeadersToday()

		fmt.Println("Gunun basliklari (" + res.Date + ")")
		fmt.Println("Baslik sayisi : " + strconv.Itoa(res.TotalCount))

		var headerNamesSlice = make([]string, len(res.Headers))
		for i := 0; i < len(res.Headers); i++ {
			headerNamesSlice[i] = res.Headers[i].HeaderText + " (" + strconv.Itoa(res.Headers[i].Count) + ")"
		}

		/*templates := &promptui.SelectTemplates{
			Label:  "Bas",
			Active: "{{ . | red }}",
			//Inactive: "{{ . | red }}",
			//Selected: "{{ . | yellow }}",
			Details: "Bit",
			//Help:     "",
			//FuncMap: nil,
		}*/

		prompt := promptui.Select{
			Label:    "Okunacak basligi secin",
			Items:    headerNamesSlice, //[]string
			Size:     10,
			HideHelp: true,
			//Templates: templates,
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Bir seyler ters gitti muhtemelen\n")
			return
		}

		selected := result[0:strings.LastIndex(result, " (")]
		var selectedHeader = -1
		for _, v := range res.Headers {
			if v.HeaderText == selected {
				selectedHeader = v.HeaderID
				break
			}
		}
		headerEntriesToday := gssc.EntriesToday(selectedHeader)
		for _, v := range headerEntriesToday.Entryler {
			fmt.Println(v.Mesaj)
			fmt.Println(v.Tarih)
			fmt.Println(v.Yazar)
		}
	},
}

func init() {
	rootCmd.AddCommand(todaysHeadersCmd)
}
