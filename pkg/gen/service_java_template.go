package gen

const ServiceJavaTemplateItf = `package {{ .ItfPackage }};
 
{{ range .ItfImports }}import {{ . }};
{{ end }}

{{ range .Documentation }}// {{ . }}
{{ end }}public interface I{{ .Name }}
{
{{ range .Methods }}{{ range .Documentation }}    // {{ . }}
{{ end }}    public {{ .ResponseClass }}  {{ .MethodName }}({{ .RequestClass }} request);

{{ end }}}

`

const ServiceJavaTemplateImpl = `package {{ .ImplPackage }};

import {{ .ItfPackage }}.I{{ .Name }};
import java.sql.*;
import javax.sql.DataSource;
import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import dml.DmlUtil;
{{ range .ItfImports }}import {{ . }};
{{ end }}
{{ range .ImplImports }}import {{ . }};
{{ end }}

{{ range .Documentation }}// {{ . }}
{{ end }}public class {{ .Name }} implements I{{ .Name }}
{
    private DataSource dataSource;
    private static final Logger log = LogManager.getLogger();

    public {{ .Name }}(DataSource dataSource) 
    {
    	this.dataSource = dataSource;
    }
{{ range .Methods }}
    public {{ .ResponseClass }}  {{ .MethodName }}({{ .RequestClass }} request)
    {
        {{ .ResponseClass }}.Builder response = {{ .ResponseClass }}.newBuilder();
        Connection conn = null;
        CallableStatement stmt = null;
{{ if .HasResultSet }}        ResultSet rs = null;{{ end }}
        try
        {
            conn = this.dataSource.getConnection();

            stmt = conn.prepareCall("{CALL {{ .ProcName }}({{ .ParamPattern }})}");

{{ range .Params }}{{ if .IsOutput }}            stmt.registerOutParameter({{ .ParamIndex }}, Types.{{ .SqlType }}); // {{ .SqlName }}{{ else }}            stmt.{{ .SqlSetter }}({{ .ParamIndex }}, {{ .JavaGetter }}); //  {{ .SqlName }}{{ end }}
{{ end }}
{{ if .HasResultSet }}            boolean hasResults = stmt.execute();
            if (hasResults) {
                rs = stmt.getResultSet();            	
{{ if .HasRepeated }}                while (rs.next()) {{ else }}                if (rs.next()){{ end }} {
{{ if .IdResult }}{{ range .Results }}                    response.{{ .JavaSetter}}(rs.{{ .SqlGetter }}("{{ .ColumnName }}");{{ end }}{{ else }}
                    {{ .ResultClass }}.Builder {{ .ResultObj }} = {{ .ResultClass }}.newBuilder();	
{{ range .Results }}                    {{ .ResultObj }}.{{ .JavaSetter}}({{ if ne .Converter "" }}{{ .Converter }}({{ end }}rs.{{ .SqlGetter }}("{{ .ColumnName }}"){{ if ne .Converter "" }}){{ end }});
{{ end  }}
                    response.{{ .ResultSetter }}({{ .ResultObj }}); {{ end }}
                }              	
            } else {
                // not found error?
            }
{{ else}}
            stmt.executeUpdate();{{ end }}
{{ range .Params }}{{ if .IsOutput }}
            {{ .JavaSetter }}({{ if ne .Converter "" }}{{ .Converter }}({{ end }}stmt.{{ .SqlGetter }}({{ .ParamIndex }}){{ if ne .Converter "" }}){{ end }}); // {{ .SqlName }}{{ end }}{{ end }}            
        }
        catch (Exception ex)
        {
            {{ if .HasErrorCode }}response.setErrorCode(-1);{{ end }}
            log.error("exception invoking stored procedure", ex);
        }
        finally
        {
{{ if .HasResultSet }}
            if (rs != null)
            {
                try
                {
                	rs.close();
                }
                catch (SqlException sqlex)
                {
                    // ignore
                }
                rs = null;

            }
{{ end }}        	
            if (stmt != null)
            {
                try
                {
                    stmt.close();
                }
                catch (SqlException sqlex)
                {
                    // ignore
                }
                stmt = null;
            }

            if (conn != null)
            {
                try
                {
                    conn.close();
                }
                catch (SqlException sqlex)
                {
                    // ignore
                }
                conn = null;
            }

        }

        return response.build();	
    }
{{ end }}
}



`