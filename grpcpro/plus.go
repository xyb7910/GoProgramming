package main

import (
	"bytes"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/protobuf/types/descriptorpb"
	"html/template"
	"log"
)

const tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
    {{- range $_, $m := .MethodList}}
    {{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
    {{- end}}
}

func Register{{.ServiceName}}(
    srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
    if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
        return err
    }
    return nil
}

type {{.ServiceName}}Client struct {
    *rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
    *{{.ServiceName}}Client, error,
) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
    in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
    return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`

type netRpcPlugin struct {
	*generator.Generator
}

type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

type ServerSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}

func (p *netRpcPlugin) buildServiceSpec(svc *descriptorpb.ServiceDescriptorProto) *ServerSpec {
	spec := &ServerSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  generator.CamelCase(m.GetInputType()),
			OutputTypeName: generator.CamelCase(m.GetOutputType()),
		})
	}
	return spec
}

func (p *netRpcPlugin) Name() string {
	return "net_rpc_plugin"
}

func (p *netRpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

func (p *netRpcPlugin) Generate(svc *descriptorpb.ServiceDescriptorProto) {
	p.genServerCode(svc)
}

func (p *netRpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.generateImportCode(file)
	}
}

func (p *netRpcPlugin) genServerCode(svc *descriptorpb.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}
	p.P(buf.String())
}

func (p *netRpcPlugin) generateImportCode(file *generator.FileDescriptor) {

}

//func init() {
//	generator.RegisterPlugin(new(netRpcPlugin))
//}
