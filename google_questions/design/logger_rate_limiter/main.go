package loggerratelimiter

type Logger struct {
	msgMap map[string]int
}

func Constructor() Logger {
	return Logger{
		msgMap: make(map[string]int),
	}
}

func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if timestamp >= this.msgMap[message] {
		this.msgMap[message] = timestamp + 10
		return true
	}
	return false
}

/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */
