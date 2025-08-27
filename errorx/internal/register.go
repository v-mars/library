package internal

const (
	DefaultErrorMsg          = "Service Internal Error"
	DefaultIsAffectStability = true
)

var (
	ServiceInternalErrorCode int32 = 1
	CodeDefinitions                = make(map[int32]*CodeDefinition)
)

type CodeDefinition struct {
	// Code 错误码，用于唯一标识错误类型
	Code int32
	// Message 错误消息，描述错误的具体内容
	Message string
	// IsAffectStability 是否影响系统稳定性，用于标识错误的严重程度
	IsAffectStability bool
}

type RegisterOption func(definition *CodeDefinition)

func WithAffectStability(affectStability bool) RegisterOption {
	return func(definition *CodeDefinition) {
		definition.IsAffectStability = affectStability
	}
}

func Register(code int32, msg string, opts ...RegisterOption) {
	definition := &CodeDefinition{
		Code:              code,
		Message:           msg,
		IsAffectStability: DefaultIsAffectStability,
	}

	for _, opt := range opts {
		opt(definition)
	}

	CodeDefinitions[code] = definition
}

func SetDefaultErrorCode(code int32) {
	ServiceInternalErrorCode = code
}
