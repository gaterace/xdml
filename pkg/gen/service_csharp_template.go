// Copyright 2017-2019 Demian Harvill
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gen

const ServiceCsharpTemplate = `{{ if eq  .Action "itf" }}using System;
 
{{ range .ItfImports }}using {{ . }};
{{ end }}

{{ range .Documentation }}// {{ . }}
{{ end }}namespace  {{ .ItfPackage }}
{
    public interface I{{ .Name }} 
    {
{{ range .Methods }}{{ range .Documentation }}        // {{ . }}
{{ end }}        {{ .ResponseClass }}  {{ .MethodName }}({{ .RequestClass }} request);

{{ end }}    }
}
{{ else }}using System;
using {{ .ItfPackage }};
{{ range .ItfImports }}using {{ . }};
{{ end }}{{ range .ImplImports }}using {{ . }};
{{ end }}using System.Data.SqlClient;
using System.Data;
using Dml;

namespace {{ .ImplPackage}}
{
    {{ range .Documentation }}// {{ . }}{{ end }}
    public class {{ .Name }}  : I{{ .Name }} 
    {
        private static readonly log4net.ILog log = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType);

        private string connectionString;

        public {{ .Name }}(string connectionString)
        {
            this.connectionString = connectionString;
        } 

{{ range .Methods }}
        public {{ .ResponseClass }}  {{ .MethodName }}({{ .RequestClass }} request)
        {
            {{ .ResponseClass }} response = new {{ .ResponseClass }}();

            using (SqlConnection conn = new SqlConnection(this.connectionString))
            {
                using (SqlCommand cmd = new SqlCommand("{{ .ProcName }}", conn))
                {
                    SqlParameter tempParam = null;
                    cmd.CommandType = System.Data.CommandType.StoredProcedure;
{{ range .Params }}{{ if .IsOutput }}                    tempParam = new SqlParameter();
                    tempParam.ParameterName = "{{ .SqlName }}";
                    tempParam.SqlDbType = SqlDbType.{{ .SqlType }};
                    tempParam.Direction = ParameterDirection.Output;
                    cmd.Parameters.Add(tempParam);
                    tempParam = null;{{ else }}                    cmd.Parameters.Add("{{ .SqlName }}", SqlDbType.{{ .SqlType }}).Value = {{ if ne .Converter "" }}{{ .Converter }}({{ end }}request.{{ .ParamName }}{{ if ne .Converter "" }}){{ end }};{{ end }}                    
{{ end }}
                    try
                    {
                        conn.Open();
{{ if .HasResultSet }}
                        SqlDataReader rdr = cmd.ExecuteReader();
{{ if .HasRepeated }}
                        while (rdr.Read())
{{ else }}
                        if (rdr.Read())
{{ end }} 
                        {
                            {{ .ResultClass }} {{ .ResultObj }} = new {{ .ResultClass }}();                            
{{ range .Results }}
                            {{ .ResultObj }}.{{ .ResultField }} = {{ if ne .Converter "" }}{{ .Converter }}({{ end }}rdr.{{ .SqlGetter }}({{ .ColumnIndex }}){{ if ne .Converter "" }}){{ end }};  // {{ .ColumnName }}{{ end }}
{{ if .HasRepeated }}
                            response.{{ .ResultSetter }}.Add({{ .ResultObj }});
{{ else }}
                            response.{{ .ResultSetter }} = {{ .ResultObj }};
{{ end }}                            
                        }

                        if (rdr != null)
                        {
                            rdr.Close();
                        }
{{ else }}                        
                        int numRows = cmd.ExecuteNonQuery();
{{ end}}
{{ range .Params}}{{ if .IsOutput }}                        response.{{ .ParamName }} =  {{ if ne .Converter "" }}{{ .Converter }}({{ end }} ({{ .ParamType }}) cmd.Parameters["{{ .SqlName }} "].Value{{ if ne .Converter "" }}){{ end }};
{{ end }}{{ end}} 
                    }
                    catch (Exception ex)
                    {
                        log.Error("unable to {{ .MethodName }}", ex);
{{ if .HasErrorCode }}                        response.ErrorCode = 500;{{end}}
                    }
                }
            }

            return response;    
        }
{{ end }}        
    } 
}


{{ end }}

`
