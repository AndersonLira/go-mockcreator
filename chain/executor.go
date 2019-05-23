package chain 

type Executor interface {
	Get(xml string) (string,error)
	GetNext() (Executor)
}