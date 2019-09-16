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

const ProcSqlServerTemplate = `USE {{ .DatabaseName }}
GO

SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO
/*****************************************************************************
* {{ .ProcName }}
* 
{{ range .Documentation }}* {{ . }}
{{ end }}*
* Parameters:
{{ range .Params }}* {{ .ParamName | rpadText 30 }}  {{ range .Documentation }}{{ . }}{{ end }}{{ if .IsOutput}} (OUTPUT){{ end }}
{{ end }}*****************************************************************************/
CREATE PROCEDURE {{ .ProcName }}
({{ range .Params }}
    {{ .ParamName }} {{ .ParamType }}{{ if .IsOutput }} OUTPUT{{ end }}{{ if .HasComma }},{{ end }}{{ end }}
)
AS
BEGIN
    DECLARE @chvUspName NVARCHAR(128)
    DECLARE @intError INT
    DECLARE @intRowCount INT
{{ range .Declares}}    DECLARE {{ .Name }} {{ .Val }}
{{ end }}
    SET NOCOUNT ON

    SET @chvUspName = OBJECT_NAME(@@PROCID)
{{ range .Sets }}    SET {{ .Name }} = {{ .Val }}
{{ end }}
{{ if .ErrCodeParam }}    SET {{ .ErrCodeParam }} = 0{{ end }}
{{ if eq .Pattern "SELECT" }}    SELECT
{{ range .Body }}        {{ .Name }} AS {{ .Val }}{{ if .HasComma }},{{ end }}
{{ end }}    FROM {{ .TableAlias.Name }} AS {{ .TableAlias.Val }}
    WHERE
{{ range .Where }}        {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ end }}
{{ end }}

    SELECT @intError = @@ERROR

    IF (@intError <> 0)
    BEGIN
        RAISERROR(N'USP=''%s''; AXN=''%s''; ERR=''%d''', 14, 1, @chvUspName, N'SELECT FROM {{ .TableAlias.Name }}', @intError)
        {{ if .ErrCodeParam }}SET {{ .ErrCodeParam }} = @intError{{ end }}
        GOTO ERROR
    END
    SET NOCOUNT OFF

    RETURN(0)

ERROR:
    SET NOCOUNT OFF

    RETURN(-100)

{{ end }}{{/* SELECT */}}{{ if eq .Pattern "INSERT" }}    BEGIN TRANSACTION
    {{ if .ErrCodeParam }}    SET {{ .ErrCodeParam }} = 0{{ end }}
    INSERT INTO {{ .TableAlias.Name }}
    (
    {{ range .Body }}    {{ .Name }}{{ if .HasComma }},{{ end }}
    {{ end }})
    VALUES
    (
    {{ range .Body }}    {{ .Val }}{{ if .HasComma }},{{ end }}
    {{ end }})


    {{ range .PostSets }}SET {{ .Name }} = {{ .Val }}{{ end }}
    SELECT @intError = @@ERROR

    IF (@intError <> 0)
    BEGIN
        RAISERROR(N'USP=''%s''; AXN=''%s''; ERR=''%d''', 14, 1, @chvUspName, N'INSERT INTO {{ .TableAlias.Name }}', @intError)
        {{ if .ErrCodeParam }}SET {{ .ErrCodeParam }} = @intError{{ end }}
        GOTO ERROR
    END

    COMMIT TRANSACTION
    SET NOCOUNT OFF

    RETURN(0)

ERROR:
    IF (@@TRANCOUNT > 0)
        ROLLBACK TRANSACTION
    SET NOCOUNT OFF

    RETURN(-100)
{{ end }}{{/* INSERT */}}{{ if eq .Pattern "UPDATE" }}    BEGIN TRANSACTION
    UPDATE {{ .TableAlias.Name }} SET
    {{ range .Body }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }},{{ end }}
    {{ end }}
    WHERE
    {{ range .Where }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ end }}
    {{ end }}
    SELECT @intError = @@ERROR, @intRowcount = @@ROWCOUNT

    IF ((@intError <> 0 OR @intRowcount = 0) AND @intError <> 2627)
    BEGIN
        RAISERROR(N'USP=''%s''; AXN=''%s''; ERR=''%d''', 14, 1, @chvUspName, N'UPDATE dbo.tb_MakerAddress', @intError) 
        {{ if .ErrCodeParam }}SET {{ .ErrCodeParam }} = @intError{{ end }}
        GOTO ERROR
    END

    COMMIT TRANSACTION
-->
    SET NOCOUNT OFF

    RETURN(0)

ERROR:
    IF (@@TRANCOUNT > 0)
        ROLLBACK TRANSACTION

    SET NOCOUNT OFF

    RETURN(-100)

{{ end }}{{/* UPDATE */}}{{ if eq .Pattern "DELETE" }}    BEGIN TRANSACTION
{{ if .Body }}
    UPDATE {{ .TableAlias.Name }} SET
{{ range .Body }}        {{ .Name }} = {{ .Val }}{{ if .HasComma }},{{ end }}
{{ end }}{{ else }}    DELETE FROM {{ .TableAlias.Name }}{{ end }}
    WHERE
    {{ range .Where }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ end }}
    {{ end }}
    SELECT @intError = @@ERROR, @intRowcount = @@ROWCOUNT

    IF ((@intError <> 0 OR @intRowcount = 0) AND @intError <> 2627)
    BEGIN
        RAISERROR(N'USP=''%s''; AXN=''%s''; ERR=''%d''', 14, 1, @chvUspName, N'UPDATE dbo.tb_MakerAddress', @intError) 
        {{ if .ErrCodeParam }}SET {{ .ErrCodeParam }} = @intError{{ end }}
        GOTO ERROR
    END

    COMMIT TRANSACTION
-->
    SET NOCOUNT OFF

    RETURN(0)

ERROR:
    IF (@@TRANCOUNT > 0)
        ROLLBACK TRANSACTION

    SET NOCOUNT OFF

    RETURN(-100)    

{{ end }}{{/* DELETE*/}}
END

GO
`
