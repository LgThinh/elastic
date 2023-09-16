package repo

import (
	"bytes"
	"es-tranform/pkg/model"
	"es-tranform/pkg/utils"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"log"
	"strings"
	"text/template"
)

type Repo struct {
	//db *gorm.DB
	es *elasticsearch.Client
}

func NewRepo(es *elasticsearch.Client) *Repo {
	return &Repo{
		//db: db,
		es: es}
}

func (r *Repo) CreateQuery(params model.RequestQuery) (string, error) {
	tmpl, err := r.ParseTemplate(utils.QuerryTemplate)
	if err != nil {
		log.Println("Cannot Parse Template", err)
	}

	//err = tmpl.Execute(os.Stdout, params)
	//if err != nil {
	//	fmt.Println("Error executing template:", err)
	//	os.Exit(1)
	//}

	var doc bytes.Buffer

	tmpl.Execute(&doc, params)
	fmt.Println(tmpl)
	s := doc.String()

	return s, nil
}

func (r *Repo) CreateQueryFlex(req model.Request) (string, error) {
	tmpl, err := r.ParseTemplate(req.Body)
	if err != nil {
		log.Println("Cannot Parse Template", err)
	}

	//err = tmpl.Execute(os.Stdout, params)
	//if err != nil {
	//	fmt.Println("Error executing template:", err)
	//	os.Exit(1)
	//}

	var doc bytes.Buffer

	tmpl.Execute(&doc, req.Params)
	s := doc.String()

	return s, nil

}

func (r *Repo) ParseTemplate(strTemplate string) (*template.Template, error) {
	t := new(template.Template)

	// custom func
	t = t.Funcs(template.FuncMap{"param_array_string": func(arr []string, sep string) string {
		str := strings.Join(arr, fmt.Sprintf(`"%s"`, sep))
		str = fmt.Sprintf(`["%s"]`, str)
		return str
	}})

	t, err := t.Parse(strTemplate)
	if err != nil {
		return nil, err
	}

	return t, nil
}
