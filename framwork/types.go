package framwork

//框架实例
var (
	fw *FrameWork
)

//服务框架
type FrameWork struct {
	Service Service //服务接口
}
type Service interface {
	RegisterCfg() (int, error)
	OnExit()
	Init(*FrameWork) (error)
}

func init() {
	fw = new(FrameWork)
}

func (fw *FrameWork) Run() {
	//初始化服务接口
	 e := fw.Service.Init(fw)
	if e != nil {
		panic(e.Error())
	}
}
