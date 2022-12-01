package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func MapsToDoubleList(id int, name string, maps map[string]interface{}) [][]string {
	records := [][]string{}

	f, err := os.Create("some.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()

	for title, clusters := range maps {
		switch title {
		case "oceanstores":
			w.Write([]string{title})
			total := []string{"total"}
			used := []string{"used"}
			free := []string{"free"}
			names := []string{" "}
			for name, osdata := range clusters.(map[string]interface{}) {
				names = append(names, name)
				for key, val := range osdata.(map[string]interface{}) {
					switch key {
					case "total":
						total = append(total, val.(string))
					case "used":
						used = append(used, val.(string))
					case "free":
						free = append(free, val.(string))
					}
				}
			}
			w.Write(names)
			w.Write(total)
			w.Write(used)
			w.Write(free)
		default:
			for name, _ := range clusters.(map[string]interface{}) {
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
					// total := []string{"total"}
					// used := []string{"used"}
					// free := []string{"free"}
					// nameCl := name
					// for key, val := range osdata.(map[string]interface{}) {
					// 	switch key {
					// 	case "total":
					// 		total = append(total, val.(string))
					// 	case "used":
					// 		used = append(used, val.(string))
					// 	case "free":
					// 		free = append(free, val.(string))
					// 	}
					// }
					// w.Write([]string{nameCl})
					// w.Write(total)
					// w.Write(used)
					// w.Write(free)
				}
			}
		}
	}
	return records
}

// if err := w.Write(keys); err != nil {
// 	fmt.Println(err)
// }
