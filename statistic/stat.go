package statistic

import (
	"fmt"
	"log"
	"os"

	"github.com/montanaflynn/stats"
	"github.com/olekukonko/tablewriter"

	"rabbitmq-hammer/process"
)

type Statistic struct {
	procTimesWriter stats.Float64Data
	procTimesReader stats.Float64Data

	tableHeader    []string
	processesNames []string

	rData *Data
	wData *Data
}

func NewStatistic(reader, writer process.Processor) *Statistic {
	return &Statistic{
		tableHeader:    []string{"PROCESS", "MIN", "MEAN", "MEDIAN", "MAX", "RPS (MEAN)", "ALL TIME"},
		processesNames: []string{"PUBLISH", "CONSUME"},
		wData:          CalcData(writer.GetStatistic()),
		rData:          CalcData(reader.GetStatistic()),
	}
}

func (o *Statistic) Show() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(o.tableHeader)

	for i, data := range []*Data{o.wData, o.rData} {
		table.Append([]string{
			fmt.Sprintf(o.processesNames[i]),
			fmt.Sprintf("%8.2f μs", data.Min),
			fmt.Sprintf("%8.2f μs", data.Mean),
			fmt.Sprintf("%8.2f μs", data.Median),
			fmt.Sprintf("%8.2f μs", data.Max),
			fmt.Sprintf("%6.2f", data.RPS),
			fmt.Sprintf("%v", data.Time),
		})
	}

	table.Render()

	log.Println("Publish\t\t\tConsume")
	for i := 0; i < len(o.rData.Percentiles); i++ {
		log.Printf("%v\t\t%v\n", o.wData.Percentiles[i], o.rData.Percentiles[i])
	}
}
