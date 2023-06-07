package utils

type (
	ServiceEncryptor struct {
	}
)

func (cv *ServiceEncryptor) MaskUserName(fullname string) (maskedName string, err error) {
	maskedName = fullname[:len(fullname)/2]
	for i := len(fullname) / 2; i < len(fullname); i++ {
		maskedName += "*"
	}
	return maskedName, nil
}

func NewServiceEncryptor() *ServiceEncryptor {
	return &ServiceEncryptor{}
}
