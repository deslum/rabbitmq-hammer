package process


type Processor interface{
	Start()
	GetStatistic() []float64
}
