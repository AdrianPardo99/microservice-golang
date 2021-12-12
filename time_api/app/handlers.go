package app

import (
	"encoding/json"
	"net/http"
    "time"
	"strings"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query().Get("tz")
	var jsonMap map[string]interface{}
	var jsonString string
	if vars == "" {
		jsonString = "{\"current_time\":\""+time.Now().UTC().String()+"\"}"
	}else {
		array_timezones := strings.Split(vars,",")
		jsonString = "{"
		for _, timezone := range array_timezones {
			loc, err := time.LoadLocation(timezone)
			if err !=nil {
				/* If timezone doesnt exist show this message */
				jsonString += "\""+timezone+"\":\"Unable to display this timezone\","
				continue
			}
			jsonString +="\""+ timezone +"\":\""+time.Now().In(loc).String()+"\","
		}
		jsonString= jsonString[:(len(jsonString)-1)]
		jsonString+="}"
	}
	json.Unmarshal([]byte(jsonString ), &jsonMap)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonMap)
}
