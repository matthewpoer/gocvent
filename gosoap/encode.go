package gosoap

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

var tokens []xml.Token

// MarshalXML envelope the body and encode to xml
func (c Client) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {

	tokens = []xml.Token{}

	//start envelope
	if c.Definitions == nil {
		return fmt.Errorf("definitions is nil")
	}

	startEnvelope()
	if len(c.HeaderParams) > 0 {
		startHeader(c.HeaderName, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
		for k, v := range c.HeaderParams {
			t := xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: k,
				},
			}

			tokens = append(tokens, t, xml.CharData(v), xml.EndElement{Name: t.Name})
		}

		endHeader(c.HeaderName)
	}

	err := startBody(c.Method, c.Definitions.Types[0].XsdSchema[0].TargetNamespace)
	if err != nil {
		return err
	}

	recursiveEncode(c.Params)

	//end envelope
	endBody(c.Method)
	endEnvelope()

	for _, t := range tokens {
		err := e.EncodeToken(t)
		if err != nil {
			return err
		}
	}

	return e.Flush()
}

func recursiveEncode(hm interface{}) {
	v := reflect.ValueOf(hm)

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			var t xml.StartElement
			if key.String() == "ObjectTypes" || key.String() == "CvSearchObject" {
				t = xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: key.String(),
					},
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.cvent.com/api/2006-11"},
					},
				}
				if key.String() == "CvSearchObject" {
					t.Attr = append(t.Attr, xml.Attr{
						Name: xml.Name{
							Space: "",
							Local: "SearchType",
						},
						Value: "AndSearch",
					})
				}
				tokens = append(tokens, t)
				recursiveEncode(v.MapIndex(key).Interface())
				tokens = append(tokens, xml.EndElement{Name: t.Name})
			} else if key.String() == "CvObjectType" {
				recursiveEncode(v.MapIndex(key).Interface())
			} else if key.String() == "Filter" {
				recursiveEncode(v.MapIndex(key).Interface())
			} else if key.String() == "Ids" {
				IDs := xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "Ids",
					},
					Attr: []xml.Attr{
						{Name: xml.Name{Space: "", Local: "xmlns"}, Value: "http://schemas.cvent.com/api/2006-11"},
					},
				}
				ID := xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "Id",
					},
				}

				tokens = append(tokens, IDs)
				tokens = append(tokens, ID)
				tokens = append(tokens, xml.CharData(fmt.Sprintf("%s", v.MapIndex(key))))
				tokens = append(tokens, xml.EndElement{Name: ID.Name})
				tokens = append(tokens, xml.EndElement{Name: IDs.Name})

			} else {
				t = xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: key.String(),
					},
				}
				tokens = append(tokens, t)
				recursiveEncode(v.MapIndex(key).Interface())
				tokens = append(tokens, xml.EndElement{Name: t.Name})
			}
		}
	case reflect.Slice:
		elementName := "CvObjectType"
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Type().String() == "gocvent.Filter" {
				elementName = "Filter"
			}
		}
		for i := 0; i < v.Len(); i++ {
			var t xml.StartElement
			t = xml.StartElement{
				Name: xml.Name{
					Space: "",
					Local: elementName,
				},
			}
			if elementName == "CvObjectType" {
				tokens = append(tokens, t)
				tokens = append(tokens, xml.CharData(v.Index(i).String()))
				tokens = append(tokens, xml.EndElement{Name: t.Name})
			} else if elementName == "Filter" {
				tokens = append(tokens, t)

				tField := xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "Field",
					},
				}
				tokens = append(tokens, tField)
				tokens = append(tokens, xml.CharData(v.Index(i).FieldByName("Field").String()))
				tokens = append(tokens, xml.EndElement{Name: tField.Name})

				tOperator := xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "Operator",
					},
				}
				tokens = append(tokens, tOperator)
				tokens = append(tokens, xml.CharData(v.Index(i).FieldByName("Operator").String()))
				tokens = append(tokens, xml.EndElement{Name: tOperator.Name})

				if v.Index(i).FieldByName("Value").String() != "" {
					tValue := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "Value",
						},
					}
					tokens = append(tokens, tValue)
					tokens = append(tokens, xml.CharData(v.Index(i).FieldByName("Value").String()))
					tokens = append(tokens, xml.EndElement{Name: tValue.Name})
				}
				if v.Index(i).FieldByName("ValueArray").Len() > 0 {
					tValue := xml.StartElement{
						Name: xml.Name{
							Space: "",
							Local: "ValueArray",
						},
					}
					tokens = append(tokens, tValue)

					for vA := 0; vA < v.Index(i).FieldByName("ValueArray").Len(); vA++ {
						tValue := xml.StartElement{
							Name: xml.Name{
								Space: "",
								Local: "Value",
							},
						}
						tokens = append(tokens, tValue)
						tokens = append(tokens, xml.CharData(v.Index(i).FieldByName("ValueArray").Index(vA).String()))
						tokens = append(tokens, xml.EndElement{Name: tValue.Name})
					}

					tokens = append(tokens, xml.EndElement{Name: tValue.Name})
				}

				tokens = append(tokens, xml.EndElement{Name: t.Name})
			}
		}
	case reflect.String:
		content := xml.CharData(v.String())
		tokens = append(tokens, content)
	}
}

func startEnvelope() {
	e := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns:xsi"}, Value: "http://www.w3.org/2001/XMLSchema-instance"},
			{Name: xml.Name{Space: "", Local: "xmlns:xsd"}, Value: "http://www.w3.org/2001/XMLSchema"},
			{Name: xml.Name{Space: "", Local: "xmlns:soap"}, Value: "http://schemas.xmlsoap.org/soap/envelope/"},
		},
	}

	tokens = append(tokens, e)
}

func endEnvelope() {
	e := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Envelope",
		},
	}

	tokens = append(tokens, e)
}

func startHeader(m, n string) {
	h := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" || n == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, h, r)

	return
}

func endHeader(m string) {
	h := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Header",
		},
	}

	if m == "" {
		tokens = append(tokens, h)
		return
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, h)
}

// startToken initiate body of the envelope
func startBody(m, n string) error {
	b := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	if m == "" || n == "" {
		return fmt.Errorf("method or namespace is empty")
	}

	r := xml.StartElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
		Attr: []xml.Attr{
			{Name: xml.Name{Space: "", Local: "xmlns"}, Value: n},
		},
	}

	tokens = append(tokens, b, r)

	return nil
}

// endToken close body of the envelope
func endBody(m string) {
	b := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: "soap:Body",
		},
	}

	r := xml.EndElement{
		Name: xml.Name{
			Space: "",
			Local: m,
		},
	}

	tokens = append(tokens, r, b)
}
