package framwork

import "Frame/framework"

type BaseService struct {

}

//regist config
func (s *BaseService) RegisterCfg() (int, error) {
	return 0, nil
}
//
func (s *BaseService) UpdateCfg() (int, error) {
	return 0, nil
}

func (s *BaseService) OnNewMessage(buff []byte,isFromClient bool)  error{
	return nil
}

func (s *BaseService) ProcessMessage(buff []byte,isFromClient bool) error{
	return nil
}

func(s *BaseService) OnExit() error{
	return nil
}

func (s *BaseService) MainLoop() error {
	for  {
		break
	}
	return nil
}

func (s *BaseService) Init (fw *framework.FrameWork) error{
	return nil
}