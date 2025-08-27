package logs

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"testing"
)

// MockLogger 是一个模拟的FullLogger实现，用于测试
type MockLogger struct {
	calls []struct {
		format string
		args   []interface{}
	}
}

// 实现FullLogger接口的所有方法
func (m *MockLogger) Trace(v ...interface{})                                      {}
func (m *MockLogger) Debug(v ...interface{})                                      {}
func (m *MockLogger) Info(v ...interface{})                                       {}
func (m *MockLogger) Notice(v ...interface{})                                     {}
func (m *MockLogger) Warn(v ...interface{})                                       {}
func (m *MockLogger) Error(v ...interface{})                                      {}
func (m *MockLogger) Fatal(v ...interface{})                                      { os.Exit(1) }
func (m *MockLogger) Tracef(format string, v ...interface{})                      {}
func (m *MockLogger) Debugf(format string, v ...interface{})                      {}
func (m *MockLogger) Infof(format string, v ...interface{})                       {}
func (m *MockLogger) Noticef(format string, v ...interface{})                     {}
func (m *MockLogger) Warnf(format string, v ...interface{})                       {}
func (m *MockLogger) Errorf(format string, v ...interface{})                      { m.recordCall(format, v...) }
func (m *MockLogger) Fatalf(format string, v ...interface{})                      { os.Exit(1) }
func (m *MockLogger) CtxTracef(ctx interface{}, format string, v ...interface{})  {}
func (m *MockLogger) CtxDebugf(ctx interface{}, format string, v ...interface{})  {}
func (m *MockLogger) CtxInfof(ctx interface{}, format string, v ...interface{})   {}
func (m *MockLogger) CtxNoticef(ctx interface{}, format string, v ...interface{}) {}
func (m *MockLogger) CtxWarnf(ctx interface{}, format string, v ...interface{})   {}
func (m *MockLogger) CtxErrorf(ctx interface{}, format string, v ...interface{})  {}
func (m *MockLogger) CtxFatalf(ctx interface{}, format string, v ...interface{})  { os.Exit(1) }
func (m *MockLogger) SetLevel(lv Level)                                           {}
func (m *MockLogger) SetOutput(w io.Writer)                                       {}

// 记录调用信息
func (m *MockLogger) recordCall(format string, v ...interface{}) {
	m.calls = append(m.calls, struct {
		format string
		args   []interface{}
	}{format: format, args: v})
}

// GetCalls 返回所有记录的调用
func (m *MockLogger) GetCalls() []struct {
	format string
	args   []interface{}
} {
	return m.calls
}

// TestErrorf 测试Errorf函数是否正确调用logger的Errorf方法
func TestErrorf(t *testing.T) {
	// 创建mock logger
	mockLogger := &MockLogger{}

	// 保存原始logger
	originalLogger := logger
	// 替换为mock logger
	//SetLogger(mockLogger)
	// 测试结束后恢复原始logger
	defer SetLogger(originalLogger)

	// 调用Errorf函数
	Errorf("test error: %s", "something went wrong")
	Errorf("error code: %d, message: %s", 500, "internal server error")

	// 验证调用次数
	calls := mockLogger.GetCalls()
	if len(calls) != 2 {
		t.Errorf("Expected 2 calls to Errorf, got %d", len(calls))
	}

	// 验证第一次调用的参数
	if calls[0].format != "test error: %s" {
		t.Errorf("Expected format 'test error: %%s', got '%s'", calls[0].format)
	}

	if len(calls[0].args) != 1 || calls[0].args[0] != "something went wrong" {
		t.Errorf("Expected args ['something went wrong'], got %v", calls[0].args)
	}

	// 验证第二次调用的参数
	if calls[1].format != "error code: %d, message: %s" {
		t.Errorf("Expected format 'error code: %%d, message: %%s', got '%s'", calls[1].format)
	}

	if len(calls[1].args) != 2 || calls[1].args[0] != 500 || calls[1].args[1] != "internal server error" {
		t.Errorf("Expected args [500, 'internal server error'], got %v", calls[1].args)
	}
}

// TestErrorfWithDefaultLogger 测试使用默认logger时Errorf函数的行为
func TestErrorfWithDefaultLogger(t *testing.T) {
	// 保存原始输出
	originalOutput := logger.(*defaultLogger).stdlog.Writer()

	// 创建buffer来捕获日志输出
	var buf bytes.Buffer

	// 设置logger输出到buffer
	SetOutput(&buf)

	// 确保日志级别允许输出Error级别的日志
	SetLevel(LevelError)

	// 测试结束后恢复原始设置
	defer SetOutput(originalOutput)
	defer SetLevel(LevelInfo)

	// 调用Errorf函数
	Errorf("test error: %s", "something went wrong")

	// 验证输出内容
	output := buf.String()

	// 检查是否包含预期的内容
	if !regexp.MustCompile(`\[Error\] test error: something went wrong`).MatchString(output) {
		t.Errorf("Expected output to contain '[Error] test error: something went wrong', got '%s'", output)
	}

	// 检查是否包含时间戳格式
	if !regexp.MustCompile(`\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}`).MatchString(output) {
		t.Errorf("Expected output to contain timestamp, got '%s'", output)
	}
}

// TestErrorfWithDifferentLevels 测试在不同日志级别下Errorf函数的行为
func TestErrorfWithDifferentLevels(t *testing.T) {
	// 保存原始输出
	originalOutput := logger.(*defaultLogger).stdlog.Writer()

	// 创建buffer来捕获日志输出
	var buf bytes.Buffer

	// 设置logger输出到buffer
	SetOutput(&buf)

	// 测试结束后恢复原始设置
	defer SetOutput(originalOutput)

	// 设置日志级别为Warn，这应该阻止Error级别的日志输出
	SetLevel(LevelWarn)

	// 调用Errorf函数
	Errorf("test error: %s", "this should not appear")

	// 验证没有输出内容（因为LevelWarn > LevelError）
	output := buf.String()
	if output != "" {
		t.Errorf("Expected no output when log level is Warn, got '%s'", output)
	}

	// 清空buffer
	buf.Reset()

	// 设置日志级别为Error，这应该允许Error级别的日志输出
	SetLevel(LevelError)

	// 调用Errorf函数
	Errorf("test error: %s", "this should appear")

	// 验证有输出内容
	output = buf.String()
	if output == "" {
		t.Error("Expected output when log level is Error, got no output")
	}
}
