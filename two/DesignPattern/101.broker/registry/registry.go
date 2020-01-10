package registry

type Register interface {
	Execute()
}

type Executor func() Register

var PackageManager = map[string]Executor{}

func Add(name string, executor Executor) {
	PackageManager[name] = executor
}
