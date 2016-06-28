package log

type logLevel struct {
	Level     int
	Prefix    string
	ColorFunc func(...interface{}) string
}

type logLevels []*logLevel

func (l *logLevels) getFunc(Level int) func(...interface{}) string {
	level := l.getLevel(Level)
	if level != nil {
		return level.ColorFunc
	}
	return nil
}

func (l *logLevels) getLevel(Level int) *logLevel {
	for _, item := range *l {
		if item.Level == Level {
			return item
		}
	}
	return nil
}
