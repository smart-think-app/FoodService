package consume

type IBasicConsume interface {
	Run([]byte) error
}
