package context

import "github.com/buingoctai/book-chapters-summary/pkg/templates"

type ContentAdaptor struct {
    Template templates.Template
}

func NewContentAdaptor(templateType string) *ContentAdaptor {
    var adaptor templates.Template
    switch templateType {
    case "word":
        adaptor = &templates.WordBasedTemplate{}
    case "line":
        adaptor = &templates.LineBasedTemplate{}
    default:
        // Handle unknown template types
    }
    return &ContentAdaptor{Template: adaptor}
}

func (c *ContentAdaptor) AdaptContent(content string) []string {
    return c.Template.AdaptContent(content)
}
