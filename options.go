package main

// Options type
type Options struct {
	PartialSuffix         string
	JobsPartial           bool
	JobsPartialBackground int
	JobsPartialForeground int
	PathPartial           bool
	PathPartialBackground int
	PathPartialForeground int
	GitPartial            bool
	GitPartialBackground  int
	GitPartialForeground  int
}

func (o *Options) setDefaults() {
	o.PartialSuffix = ""

	o.JobsPartial = true
	o.JobsPartialBackground = 2 // 244
	o.JobsPartialForeground = 0 // 21
	o.PathPartial = true
	o.PathPartialBackground = 8
	o.PathPartialForeground = 21
	o.GitPartial = true
	o.GitPartialBackground = 18 // 236
	o.GitPartialForeground = 21
}
