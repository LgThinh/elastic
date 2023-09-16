package utils

const QuerryTemplate = `{
  "from": {{ if .From }} {{ .From }} {{ else }} 0 {{ end }},
  "size": {{ if .Size }} {{ .Size }} {{ else }} 10 {{ end }},
  "query": {
    "bool": {
      "must": [
        {
          "query_string": {
            "query": "{{ if .SearchKeyword }}{{ .SearchKeyword }}{{ else }}*{{ end }}",
            "fields": {{ if .SearchFields }}{{  param_array_string .SearchFields ", " }}{{ else }}["*"]{{ end }}
          }
        }
      ],
      "filter": [
        {{ if .AuthorName }}
        {
          "terms": {
            "author_info.first_name.keyword": {{ param_array_string .AuthorName ", "}}
          }
        },
        {{ end }}
        {{ if .AuthorName }}
        {
          "terms": {
            "tag.id.keyword": {{ param_array_string .TagId ", "}}
          }
        },
        {{ end }}
        {{ if .PostTermTaxonomyId }}
        {
          "terms": {
            "post_term.post_term_taxonomy_id.keyword": {{ param_array_string .PostTermTaxonomyId ", "}}
          }
        },
        {{ end }}
        {{ if .EntityTermTaxonomyId }}
        {
          "terms": {
            "entity_relationship.entity_term_taxonomy_id.keyword": "{{ .EntityTermTaxonomyId }}"
          }
        },
        {{ end }}
        {
          "range": {
            "available_to": {
              "gte": "now/d"
            }
          }
        },
        {
          "range": {
            "available_from": {
              "lte": "now/d"
            }
          }
        }
      ]
    }
  },
  "sort": [
    {
      {{ if .Sort }}
      {{$first := true}}
      {{ range $col, $order := .Sort }}
      {{if $first}}
      {{$first = false}}
      {{ else }},{{ end }}
      "{{ $col }}": {
        "order": "{{ $order }}"
      }
      {{ end }}
      {{ else }}
      "created_at": {
        "order": "desc"
      }
      {{ end }}
    }
  ],
  "aggs": {
    "total": {
      "cardinality": {
        "field": "post_id.keyword"
      }
    }
  }
}`
