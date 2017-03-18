package report

import (
	"time"
	"os"
	"strings"
	"strconv"
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
)

type ReportItem struct {
	algorithm string
	fileSample string
	sampleLines int
	time int64
	sorted bool
}

type Report struct {
	name string
	filename string
	items []ReportItem
	dateTime time.Time
}

//noinspection GoUnusedExportedFunction
func NewReport(name string, filename string) *Report {
	return &Report{name, filename, make([]ReportItem, 0), time.Now()}
}

func (r *Report) AddItem(algorithm string, fileSample string, sampleLines int, time int64, sorted bool) {
	var item ReportItem
	item.algorithm = algorithm
	item.fileSample = fileSample
	item.sampleLines = sampleLines
	item.time = time
	item.sorted = sorted
	r.items = append(r.items, item)
}

func (i *ReportItem) toString() string {
	parts:=[5]string{i.algorithm, i.fileSample, strconv.Itoa(i.sampleLines), strconv.FormatInt(i.time, 10), strconv.FormatBool(i.sorted)}
	return strings.Join(parts[:], ";")
}

func (r *Report) WriteToFile() {
	f, err := os.OpenFile(r.filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	check(err)
	w := bufio.NewWriter(f)
	defer f.Close()
	defer w.Flush()
	for i:= 0; i <  len(r.items); i++ {
		fmt.Fprintf(w, "%s\n", r.items[i].toString())
	}
	f.Sync()
}

func check(err error) {
	if err != nil {
		log.Errorf("%s", err)
	}
}
