package patterns

import "sync"

type singleton struct {
	Data int
}

var (
	ins *singleton
	// DirectInstance Direct
	DirectInstance singleton = singleton{}
	once           sync.Once
)

// GetInstance Lazy
//
//goland:noinspection GoExportedFuncWithUnexportedType
func GetInstance() *singleton {
	if ins == nil {
		once.Do(func() {
			ins = &singleton{}
		})
	}
	return ins
}
