package plugin

import (
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
)

type netRpcPlugin struct {
	*generator.Generator
}

func (n *netRpcPlugin) Name() string {
	return "netrpc"
}

func (n *netRpcPlugin) Init(g *generator.Generator) {
	n.Generator = g
}

func (n *netRpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		n.genImportCode(file)
	}
}

func (n *netRpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		n.genServiceCode(svc)
	}
}

func (n *netRpcPlugin) genImportCode(file *generator.FileDescriptor) {
	// TODO:
	panic("implement me")
}

func (n *netRpcPlugin) genServiceCode(svc *descriptorpb.ServiceDescriptorProto) {
	// TODO:
	panic("implement me")
}
