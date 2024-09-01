package hello

import "LearingGo/gen"

type Service struct{}

func (s *Service) Hello(request *gen.String, response *gen.String) error {
	response.Value = "Hello" + request.GetValue()
	return nil
}
