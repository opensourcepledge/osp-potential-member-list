// © 2024 Vlad-Stefan Harbuz <vlad@vladh.net>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
        "flag"
        "html/template"
        "log"
        "net/http"
        "strconv"
        "strings"
        "time"
)

const SHEET_ID = "1hzRN_ZJjlh4979HAp6wkhebXmlgXVX6i44tpQihJEdQ"
const SHEET_RANGE = "A2:G500"
const SHEET_REFRESH_N_SECONDS = 60 * 10 // 10 minutes

const (
        PledgedYes = iota
        PledgedRefused
        PledgedUnknown
)

type Member struct {
        Name string
        Url string
        Pledged int
}

func getMembers() []Member {
        cells := getSheetRange(SHEET_ID, SHEET_RANGE)
        members := make([]Member, 0, len(cells))
        for _, row := range cells {
                if len(row) < 7 {
                        continue
                }
                pledgeStatus := PledgedUnknown
                putativeStatus := strings.ToLower(row[1].(string))
                if strings.Contains(putativeStatus, "yes") {
                        pledgeStatus = PledgedYes
                } else if strings.Contains(putativeStatus, "refused") {
                        pledgeStatus = PledgedRefused
                }
                members = append(members, Member {
                        Name: row[0].(string),
                        Url: row[6].(string),
                        Pledged: pledgeStatus,
                })
        }
        return members
}

func main() {
        port := flag.Int("port", 3003, "http port")

        members := getMembers()
        refreshTime := time.Now()

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                timeSinceRefresh := time.Now().Sub(refreshTime)
                if timeSinceRefresh.Seconds() >= SHEET_REFRESH_N_SECONDS {
                        members = getMembers()
                        refreshTime = time.Now()
                }
                tplIndex := template.Must(template.ParseFiles("templates/index.html"))
                tplIndex.Execute(w, members)
        })

        log.Printf("Starting HTTP server on port %d", *port)
        err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)
        if err != nil {
                log.Fatal(err.Error())
        }
}
