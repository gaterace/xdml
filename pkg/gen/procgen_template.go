package gen

const ProcMysqlTemplate = `DELIMITER $$

DROP PROCEDURE IF EXISTS {{ .ProcName }} $$

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
    {{ if .IsOutput }}OUT {{ end }}{{ .ParamName }} {{ .ParamType }}{{ if .HasComma }},{{ end }}{{ end }}
)
{{ if eq .Pattern "SELECT"}}READS SQL DATA{{ else }}MODIFIES SQL DATA{{ end }}
SQL SECURITY INVOKER
BEGIN
{{ range .Declares}}    DECLARE {{ .Name }} {{ .Val }};
{{ end }}
{{ if .ErrCodeParam }}    DECLARE CONTINUE HANDLER FOR SQLEXCEPTION
    BEGIN
        SET {{ .ErrCodeParam }} = 500;
    END;
    
    SET {{ .ErrCodeParam }} = 0;
{{ end }}{{ range .Sets }}    SET {{ .Name }} = {{ .Val }};
{{ end }}
{{ if eq .Pattern "SELECT" }}    SELECT
{{ range .Body }}        {{ .Name }} AS {{ .Val }}{{ if .HasComma }},{{ end }}
{{ end }}    FROM {{ .TableAlias.Name }} AS {{ .TableAlias.Val }}
    WHERE
{{ range .Where }}        {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ else }};{{ end }}
{{ end }}
{{ end }}{{/* SELECT */}}{{ if eq .Pattern "INSERT" }}    START TRANSACTION
    INSERT INTO {{ .TableAlias.Name }}
    (
    {{ range .Body }}    {{ .Name }}{{ if .HasComma }},{{ end }}
    {{ end }})
    VALUES
    (
    {{ range .Body }}    {{ .Val }}{{ if .HasComma }},{{ end }}
    {{ end }});

    {{ range .PostSets }}SET {{ .Name }} = {{ .Val }};{{ end }}

    IF out_intErrorCode = 0 THEN
        COMMIT;
    ELSE 
        ROLLBACK;
    END IF;


{{ end }}{{/* INSERT */}}{{ if eq .Pattern "UPDATE" }}    START TRANSACTION
    UPDATE {{ .TableAlias.Name }} SET
    {{ range .Body }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }},{{ end }}
    {{ end }}
    WHERE
    {{ range .Where }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ else }};{{ end }}
    {{ end }}

    IF out_intErrorCode = 0 THEN
        COMMIT;
    ELSE 
        ROLLBACK;
    END IF;

{{ end }}{{/* UPDATE */}}{{ if eq .Pattern "DELETE" }}    START TRANSACTION
{{ if .Body }}
    UPDATE {{ .TableAlias.Name }} SET
{{ range .Body }}        {{ .Name }} = {{ .Val }}{{ if .HasComma }},{{ end }}
{{ end }}{{ else }}    DELETE FROM {{ .TableAlias.Name }}{{ end }}
    WHERE
    {{ range .Where }}    {{ .Name }} = {{ .Val }}{{ if .HasComma }} AND{{ else }};{{ end }}
    {{ end }}

    IF out_intErrorCode = 0 THEN
        COMMIT;
    ELSE 
        ROLLBACK;
    END IF;
{{ end }}{{/* DELETE */}}
END$$
DELIMITER ;

`
