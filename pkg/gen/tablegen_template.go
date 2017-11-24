package gen

const TableMysqlTemplate = `use {{ .DatabaseName }};

DROP TABLE IF EXISTS {{ .TableName }};

{{ range .Documentation }}-- {{ . }}
{{ end }}CREATE TABLE {{ .TableName }}
(
{{ range .Columns }}
{{ range .Documentation }}    -- {{ . }}
{{ end }}    {{ .ColumnName }} {{ .ColumnType }}{{ if .ColumnModifier }} {{ .ColumnModifier }}{{ end }} {{ if .IsNullable }}NULL{{ else }}NOT NULL{{ end }}{{ if .HasComma }},{{ end }}{{ end }}

{{ range .Indices }}
{{ if .IsPrimary }}    PRIMARY KEY{{ else }}{{ if .IsUnique }}    UNIQUE{{ else }}    INDEX{{ end }}{{ end }} ({{ .FieldList }}){{ if .HasComma }},{{ end }}{{ end }}
) ENGINE=InnoDB;

`

const TableSqlServerlTemplate = `USE {{ .DatabaseName }};
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO


{{ range .Documentation }}-- {{ . }}
{{ end }}CREATE TABLE {{ .TableName }}
(
{{ range .Columns }}
{{ range .Documentation }}    -- {{ . }}
{{ end }}    {{ .ColumnName }} {{ .ColumnType }}{{ if .ColumnModifier }} {{ .ColumnModifier }}{{ end }} {{ if .IsNullable }}NULL{{ else }}NOT NULL{{ end }}{{ if .HasComma }},{{ end }}{{ end }}
{{ range .Indices }}
{{ if .IsPrimary }}    CONSTRAINT [{{ .IndexName }}] PRIMARY KEY CLUSTERED
    (
{{ .FieldList }}	
){{ end }}
{{ end }}
    ) WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]

GO

{{ range .Indices }}{{ if not .IsPrimary }}  {{ end }}  
CREATE {{ if .IsUnique }}UNIQUE NONCLUSTERED{{ else }}CLUSTERED{{ end }} INDEX {{ .IndexName }} ON {{ .TableName }}
(
	{{ .FieldList }}
) WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
GO
{{ end }}


SET ANSI_PADDING ON
GO
`
