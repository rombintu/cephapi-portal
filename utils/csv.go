package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MapsToCSV(id int, name string, maps map[string]interface{}) *bytes.Buffer {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	defer w.Flush()

	for title, clusters := range maps {
		w.Write([]string{" "})
		switch title {
		case "oceanstores":
			w.Write([]string{title})
			total := []string{"Всего (TB)"}
			used := []string{"Используется (TB)"}
			free := []string{"Свободно (TB)"}
			names := []string{" "}
			for name, osdata := range clusters.(map[string]interface{}) {
				names = append(names, name)
				for key, val := range osdata.(map[string]interface{}) {
					switch key {
					case "total":
						total = append(total, StrToFloatStr(val.(string)))
					case "used":
						used = append(used, StrToFloatStr(val.(string)))
					case "free":
						free = append(free, StrToFloatStr(val.(string)))
					}
				}
			}
			// ADD SUMMA
			names = append(names, "Сумма (TB)")
			total = append(total, summStrFloatToStr(total))
			used = append(used, summStrFloatToStr(used))
			free = append(free, summStrFloatToStr(free))

			w.Write(names)
			w.Write(total)
			w.Write(used)
			w.Write(free)
		default:
			for name, sdsdata := range clusters.(map[string]interface{}) {
				if name == "timestamp" {
					continue
				}
				switch name {
				case "pools":
					// 	poolNames := []string{"Data"}
					// 	provis := []string{"provis"}
					// 	used := []string{"used"}
					// 	zoneID := []string{"zoneID"}
					// 	for poolName, poolData := range sdsdata.(map[string]interface{}) {

					// 		poolNames = append(poolNames, poolName)
					// 		for key, val := range poolData.(map[string]interface{}) {

					// 			switch key {
					// 			case "Provis":
					// 				provis = append(provis, val.(string))
					// 			case "Used":
					// 				used = append(used, val.(string))
					// 			case "ZoneID":
					// 				zoneID = append(zoneID, val.(string))
					// 			}
					// 		}

					// 	}
					// 	w.Write(poolNames)
					// 	w.Write(provis)
					// 	w.Write(used)
					// 	w.Write(zoneID)
					continue
				case "zones":
					w.Write([]string{title})
					total := []string{"Всего (TB)"}
					provis := []string{"Запровижено (TB)"}
					used := []string{"Используется (TB)"}
					// avail := []string{"Выделено (TB)"}
					free := []string{"Свободно (TB)"}
					names := []string{" "}
					for name, zdata := range sdsdata.(map[string]interface{}) {
						names = append(names, name)
						for key, val := range zdata.(map[string]interface{}) {
							switch key {
							case "Total":
								total = append(total, StrToFloatStr(val.(string)))
							case "Used":
								used = append(used, StrToFloatStr(val.(string)))
							case "Free":
								free = append(free, StrToFloatStr(val.(string)))
							// case "Avail":
							// 	avail = append(avail, StrToFloatStr(val.(string)))
							case "Provis":
								provis = append(provis, StrToFloatStr(val.(string)))
							}

						}
					}
					// ADD SUMMA
					names = append(names, "Сумма (TB)")
					total = append(total, summStrFloatToStr(total))
					used = append(used, summStrFloatToStr(used))
					free = append(free, summStrFloatToStr(free))
					// avail = append(avail, summStrFloatToStr(avail))
					provis = append(provis, summStrFloatToStr(provis))

					w.Write(names)
					w.Write(total)
					w.Write(provis)
					w.Write(used)
					w.Write(free)
					// w.Write(avail)
				}
			}
		}
	}
	return &buf
}

func summStrFloatToStr(strFloat []string) string {
	sum := float64(0)
	for _, fstr := range strFloat {
		f, err := strconv.ParseFloat(fstr, 64)
		if err == nil {
			sum += f
		}
	}
	return fmt.Sprintf("%.2f", sum)
}

func ToFloat(str string, K float64) string {
	fstr, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "parse-error"
	}
	return fmt.Sprintf("%.2f", fstr*K)
}

func StrToFloatStr(str string) string {
	if str == "0B" || len(str) < 3 {
		return "0"
	}
	cells := strings.Split(str, " ")
	switch cells[1] {
	case "KB":
		// return ToFloat(cells[0], math.Pow(10, -9))
		return "0"
	case "MB":
		// return ToFloat(cells[0], math.Pow(10, -6))
		return "0"
	case "GB":
		return ToFloat(cells[0], math.Pow(10, -3))
	case "TB":
		return ToFloat(cells[0], 1)
	case "PB":
		return ToFloat(cells[0], math.Pow(10, 3))
	}
	return "parse-error"
}
