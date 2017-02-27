package report

import "time"

type ReportItem struct {
	algorithm string
	fileSample string
	sampleLines uint32
	time uint64
	sorted bool
}

type Report struct {
	name string
	items []ReportItem
	dateTime time.Time
}

func NewReport() *Report {
	return &Report{"", make([]ReportItem, 0), time.Now()}
}

func NewReportItem() *ReportItem {
	return &ReportItem{"", "", 0, 0, false}
}

func (r *Report) WriteItem() {

}

