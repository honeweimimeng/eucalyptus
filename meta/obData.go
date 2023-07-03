package meta

type DataOb interface {
	Processor() DataProcessor
	Name() string
	Close() error
}

type DataObInitializer interface {
	DataOb
	InitProcessor(processor DataProcessor)
}
