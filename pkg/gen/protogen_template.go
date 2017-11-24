package gen

const ProtoTemplate = `syntax = "proto3";

package {{ .JavaPackage }};
option csharp_namespace = "{{ .CsharpPackage}}";
option go_package = "{{ .GoPackage }}";

{{ range .Imports }}import "{{ . }}";
{{ end }}

{{ range .Enums }}{{ range .Documentation }}// {{ . }}
{{ end }}enum {{ .EnumName }} {
{{ range .Fields}}    {{ range .Documentation }}// {{ . }}
{{ end }}    {{ .FieldName }} = {{ .FieldVal }};
{{ end }}}

{{ end }}
{{ range .Services }}{{ range .Documentation }}// {{ . }}
{{ end }}service {{ .ServiceName }} {
{{ range .Methods }}    {{ range .Documentation }}// {{ . }}
{{ end }}    rpc {{ .MethodName }} ({{ .RequestClass }}) returns ({{ .ResponseClass }});
{{ end }}  
}

{{ end }}

{{ range .Classes }}{{ range .Documentation }}// {{ . }}
{{ end }}message {{ .ClassName }} {
{{ range .Fields}}    {{ range .Documentation }}// {{ . }}
{{ end }}    {{ .FieldModifier }}{{ .FieldType }} {{ .FieldName }} = {{ .FieldVal }};
{{ end }}
}

{{ end }}

`
