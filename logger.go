package ghttp

type Logger struct {
	Infof func(format string, args ...interface{})
	Errorf func(format string, args ...interface{})


}

var logger = Logger{
	Infof: func(format string, args ...interface{}) {},
	Errorf: func(format string, args ...interface{}) {},
}