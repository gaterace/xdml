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
    WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
) ON [PRIMARY]

GO

{{ range .Indices }}{{ if not .IsPrimary }}    
CREATE {{ if .IsUnique }}UNIQUE NONCLUSTERED{{ else }}CLUSTERED{{ end }} INDEX {{ .IndexName }} ON {{ .TableName }}
(
	{{ .FieldList }}
) WITH (PAD_INDEX  = OFF, STATISTICS_NORECOMPUTE  = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, DROP_EXISTING = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS  = ON, ALLOW_PAGE_LOCKS  = ON) ON [PRIMARY]
GO
{{ end }}{{ end }}


SET ANSI_PADDING ON
GO
`
