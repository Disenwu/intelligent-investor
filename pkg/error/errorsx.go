package errorsx

import "fmt"

type ErrorsX struct {
	Code     int               `json:"code,omitempty"`
	Reason   string            `json:"reason,omitempty"`
	Message  string            `json:"message,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (e *ErrorsX) Error() string {
	return fmt.Sprintf("error:code = %d, reason = %s, message = %s, metadata = %v", e.Code, e.Reason, e.Message, e.Metadata)
}

func (err *ErrorsX) WithMessage(message string) *ErrorsX {
	err.Message = message
	return err
}

func (err *ErrorsX) WithMetadata(metadata map[string]string) *ErrorsX {
	err.Metadata = metadata
	return err
}

func (err *ErrorsX) KeyAndValues(keyAndValues ...string) *ErrorsX {
	if err.Metadata == nil {
		err.Metadata = make(map[string]string)
	}
	for i := 0; i < len(keyAndValues); i += 2 {
		if i+1 < len(keyAndValues) {
			err.Metadata[keyAndValues[i]] = keyAndValues[i+1]
		}
	}
	return err
}
