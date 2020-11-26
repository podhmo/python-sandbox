package main

import (
	"fmt"
	"strings"
	"text/template"
	"text/template/parse"
	_ "unsafe"
)

//go:linkname template_builtins text/template.builtins
func template_builtins() template.FuncMap

func main() {
	{
		source := `
xxx {{ .XXX }}
xxx.yyy {{ .XXX.YYY }}
{{ range $i, $x := .Items }}
- {{ $x }}
{{ end}}
`
		Run(source)
	}
	fmt.Println("----------------------------------------")
	{
		source := `// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A callback for modifying requests which are generated before sending over
	// the network.
	RequestEditor RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
    // create a client with sane default values
    client := Client{
        Server: server,
    }
    // mutate client and add all optional params
    for _, o := range opts {
        if err := o(&client); err != nil {
            return nil, err
        }
    }
    // ensure the server URL always has a trailing slash
    if !strings.HasSuffix(client.Server, "/") {
        client.Server += "/"
    }
    // create httpClient, if not already present
    if client.Client == nil {
        client.Client = http.DefaultClient
    }
    return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditor = fn
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
{{range . -}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$opid := .OperationId -}}
    // {{$opid}} request {{if .HasBody}} with any body{{end}}
    {{$opid}}{{if .HasBody}}WithBody{{end}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}) (*http.Response, error)
{{range .Bodies}}
    {{$opid}}{{.Suffix}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody) (*http.Response, error)
{{end}}{{/* range .Bodies */}}
{{end}}{{/* range . $opid := .OperationId */}}
}


{{/* Generate client methods */}}
{{range . -}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$opid := .OperationId -}}

func (c *Client) {{$opid}}{{if .HasBody}}WithBody{{end}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}) (*http.Response, error) {
    req, err := New{{$opid}}Request{{if .HasBody}}WithBody{{end}}(c.Server{{genParamNames .PathParams}}{{if $hasParams}}, params{{end}}{{if .HasBody}}, contentType, body{{end}})
    if err != nil {
        return nil, err
    }
    req = req.WithContext(ctx)
    if c.RequestEditor != nil {
        err = c.RequestEditor(ctx, req)
        if err != nil {
            return nil, err
        }
    }
    return c.Client.Do(req)
}

{{range .Bodies}}
func (c *Client) {{$opid}}{{.Suffix}}(ctx context.Context{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody) (*http.Response, error) {
    req, err := New{{$opid}}{{.Suffix}}Request(c.Server{{genParamNames $pathParams}}{{if $hasParams}}, params{{end}}, body)
    if err != nil {
        return nil, err
    }
    req = req.WithContext(ctx)
    if c.RequestEditor != nil {
        err = c.RequestEditor(ctx, req)
        if err != nil {
            return nil, err
        }
    }
    return c.Client.Do(req)
}
{{end}}{{/* range .Bodies */}}
{{end}}

{{/* Generate request builders */}}
{{range .}}
{{$hasParams := .RequiresParamObject -}}
{{$pathParams := .PathParams -}}
{{$bodyRequired := .BodyRequired -}}
{{$opid := .OperationId -}}

{{range .Bodies}}
// New{{$opid}}Request{{.Suffix}} calls the generic {{$opid}} builder with {{.ContentType}} body
func New{{$opid}}Request{{.Suffix}}(server string{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}, body {{$opid}}{{.NameTag}}RequestBody) (*http.Request, error) {
    var bodyReader io.Reader
    buf, err := json.Marshal(body)
    if err != nil {
        return nil, err
    }
    bodyReader = bytes.NewReader(buf)
    return New{{$opid}}RequestWithBody(server{{genParamNames $pathParams}}{{if $hasParams}}, params{{end}}, "{{.ContentType}}", bodyReader)
}
{{end}}

// New{{$opid}}Request{{if .HasBody}}WithBody{{end}} generates requests for {{$opid}}{{if .HasBody}} with any type of body{{end}}
func New{{$opid}}Request{{if .HasBody}}WithBody{{end}}(server string{{genParamArgs $pathParams}}{{if $hasParams}}, params *{{$opid}}Params{{end}}{{if .HasBody}}, contentType string, body io.Reader{{end}}) (*http.Request, error) {
    var err error
{{range $paramIdx, $param := .PathParams}}
    var pathParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    pathParam{{$paramIdx}} = {{.ParamName}}
    {{end}}
    {{if .IsJson}}
    var pathParamBuf{{$paramIdx}} []byte
    pathParamBuf{{$paramIdx}}, err = json.Marshal({{.ParamName}})
    if err != nil {
        return nil, err
    }
    pathParam{{$paramIdx}} = string(pathParamBuf{{$paramIdx}})
    {{end}}
    {{if .IsStyled}}
    pathParam{{$paramIdx}}, err = runtime.StyleParam("{{.Style}}", {{.Explode}}, "{{.ParamName}}", {{.GoVariableName}})
    if err != nil {
        return nil, err
    }
    {{end}}
{{end}}
    queryUrl, err := url.Parse(server)
    if err != nil {
        return nil, err
    }

    basePath := fmt.Sprintf("{{genParamFmtString .Path}}"{{range $paramIdx, $param := .PathParams}}, pathParam{{$paramIdx}}{{end}})
    if basePath[0] == '/' {
        basePath = basePath[1:]
    }

    queryUrl, err = queryUrl.Parse(basePath)
    if err != nil {
        return nil, err
    }
{{if .QueryParams}}
    queryValues := queryUrl.Query()
{{range $paramIdx, $param := .QueryParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    {{if .IsPassThrough}}
    queryValues.Add("{{.ParamName}}", {{if not .Required}}*{{end}}params.{{.GoName}})
    {{end}}
    {{if .IsJson}}
    if queryParamBuf, err := json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}}); err != nil {
        return nil, err
    } else {
        queryValues.Add("{{.ParamName}}", string(queryParamBuf))
    }

    {{end}}
    {{if .IsStyled}}
    if queryFrag, err := runtime.StyleParam("{{.Style}}", {{.Explode}}, "{{.ParamName}}", {{if not .Required}}*{{end}}params.{{.GoName}}); err != nil {
        return nil, err
    } else if parsed, err := url.ParseQuery(queryFrag); err != nil {
       return nil, err
    } else {
       for k, v := range parsed {
           for _, v2 := range v {
               queryValues.Add(k, v2)
           }
       }
    }
    {{end}}
    {{if not .Required}}}{{end}}
{{end}}
    queryUrl.RawQuery = queryValues.Encode()
{{end}}{{/* if .QueryParams */}}
    req, err := http.NewRequest("{{.Method}}", queryUrl.String(), {{if .HasBody}}body{{else}}nil{{end}})
    if err != nil {
        return nil, err
    }

{{range $paramIdx, $param := .HeaderParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    var headerParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    headerParam{{$paramIdx}} = {{if not .Required}}*{{end}}params.{{.GoName}}
    {{end}}
    {{if .IsJson}}
    var headerParamBuf{{$paramIdx}} []byte
    headerParamBuf{{$paramIdx}}, err = json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    headerParam{{$paramIdx}} = string(headerParamBuf{{$paramIdx}})
    {{end}}
    {{if .IsStyled}}
    headerParam{{$paramIdx}}, err = runtime.StyleParam("{{.Style}}", {{.Explode}}, "{{.ParamName}}", {{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    {{end}}
    req.Header.Add("{{.ParamName}}", headerParam{{$paramIdx}})
    {{if not .Required}}}{{end}}
{{end}}

{{range $paramIdx, $param := .CookieParams}}
    {{if not .Required}} if params.{{.GoName}} != nil { {{end}}
    var cookieParam{{$paramIdx}} string
    {{if .IsPassThrough}}
    cookieParam{{$paramIdx}} = {{if not .Required}}*{{end}}params.{{.GoName}}
    {{end}}
    {{if .IsJson}}
    var cookieParamBuf{{$paramIdx}} []byte
    cookieParamBuf{{$paramIdx}}, err = json.Marshal({{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    cookieParam{{$paramIdx}} = url.QueryEscape(string(cookieParamBuf{{$paramIdx}}))
    {{end}}
    {{if .IsStyled}}
    cookieParam{{$paramIdx}}, err = runtime.StyleParam("simple", {{.Explode}}, "{{.ParamName}}", {{if not .Required}}*{{end}}params.{{.GoName}})
    if err != nil {
        return nil, err
    }
    {{end}}
    cookie{{$paramIdx}} := &http.Cookie{
        Name:"{{.ParamName}}",
        Value:cookieParam{{$paramIdx}},
    }
    req.AddCookie(cookie{{$paramIdx}})
    {{if not .Required}}}{{end}}
{{end}}
    {{if .HasBody}}req.Header.Add("Content-Type", contentType){{end}}
    return req, nil
}

{{end}}{{/* Range */}}
`
		Run(source)
	}
}

func Run(source string) {
	// tmpl := template.Must(template.New("x").Parse(source))
	// fmt.Println(tmpl)
	fMap := template_builtins()
	fMap["genParamArgs"] = genParamArgs
	fMap["genParamNames"] = genParamNames
	fMap["genParamFmtString"] = genParamFmtString
	t, err := parse.Parse("X", source, "", "", fMap)
	fmt.Println(err)
	// pp.Println(t)
	for _, tree := range t {
		walkTree(tree)
	}
}

func walkTree(t *parse.Tree) {
	// TODO: with walking history (for debug)
	var walkNode func(node parse.Node, depth int)
	walkNode = func(node parse.Node, depth int) {
		switch node := node.(type) {
		case *parse.ListNode:
			for _, subnode := range node.Nodes {
				walkNode(subnode, depth+1)
			}
		case *parse.TextNode:
		case *parse.PipeNode:
			for _, x := range node.Decl {
				walkNode(x, depth+1)
			}
			for _, x := range node.Cmds {
				walkNode(x, depth+1)
			}
		case *parse.ActionNode:
			walkNode(node.Pipe, depth+1)
		case *parse.CommandNode:
			for _, x := range node.Args {
				walkNode(x, depth+1)
			}
		case *parse.IdentifierNode:
			fmt.Println(strings.Repeat("  ", depth), "identifier Ident", node.Ident)
		case *parse.VariableNode:
			fmt.Println(strings.Repeat("  ", depth), "variable Ident", node.Ident)
		case *parse.DotNode:
		case *parse.NilNode:
		case *parse.FieldNode:
			fmt.Println(strings.Repeat("  ", depth), "field Ident", node.Ident)
		case *parse.ChainNode:
		case *parse.BoolNode:
		case *parse.NumberNode:
		case *parse.StringNode:
		case *parse.BranchNode:
			walkNode(node.Pipe, depth+1)
			if node.List != nil {
				walkNode(node.List, depth+1)
			}
			if node.ElseList != nil {
				walkNode(node.ElseList, depth+1)
			}
		case *parse.IfNode:
			walkNode(&node.BranchNode, depth+1)
		case *parse.RangeNode:
			walkNode(&node.BranchNode, depth+1)
		case *parse.WithNode:
			walkNode(&node.BranchNode, depth+1)
		case *parse.TemplateNode:
			walkNode(node.Pipe, depth+1) // ??
		default:
			// skip endNode, elseNode
		}
	}
	for _, node := range t.Root.Nodes {
		walkNode(node, 0)
	}
}

// helpers
func genParamFmtString(path string) string {
	return ""
}

func genParamArgs(params []ParameterDefinition) string {
	if len(params) == 0 {
		return ""
	}
	return "<>"
}
func genParamNames(params []ParameterDefinition) string {
	if len(params) == 0 {
		return ""
	}
	return "<>"
}

type ParameterDefinition struct {
	ParamName string // The original json parameter name, eg param_name
	In        string // Where the parameter is defined - path, header, cookie, query
	Required  bool   // Is this a required parameter?
	Spec      interface{}
	Schema    interface{}
}
