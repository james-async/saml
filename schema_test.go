package saml

import (
	"encoding/xml"

	"github.com/beevik/etree"
	. "gopkg.in/check.v1"
)

var _ = Suite(&SchemaTest{})

type SchemaTest struct {
}

func (test *SchemaTest) TestAttributeXMLRoundTrip(c *C) {
	expected := Attribute{
		FriendlyName: "TestFriendlyName",
		Name:         "TestName",
		NameFormat:   "urn:oasis:names:tc:SAML:2.0:attrname-format:basic",
		Values: []AttributeValue{AttributeValue{
			Type:  "xs:string",
			Value: "test",
		}},
	}

	doc := etree.NewDocument()
	doc.SetRoot(expected.Element())
	x, err := doc.WriteToBytes()
	c.Assert(err, IsNil)
	c.Assert(string(x), Equals, "<saml:Attribute FriendlyName=\"TestFriendlyName\" Name=\"TestName\" NameFormat=\"urn:oasis:names:tc:SAML:2.0:attrname-format:basic\"><saml:AttributeValue xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\" xmlns:xs=\"http://www.w3.org/2001/XMLSchema\" xsi:type=\"xs:string\">test</saml:AttributeValue></saml:Attribute>")

	var actual Attribute
	err = xml.Unmarshal(x, &actual)
	c.Assert(err, IsNil)
	c.Assert(actual, DeepEquals, expected)
}

func (test *SchemaTest) TestLogoutRequestElement(c *C) {

	logoutRequest := &LogoutRequest{
		NameID: &NameID{
			Value: "name_id_value",
		},
	}

	doc := etree.NewDocument()
	doc.SetRoot(logoutRequest.Element())
	actual, err := doc.WriteToString()
	c.Assert(err, IsNil)
	c.Assert(actual, Equals, "<samlp:LogoutRequest xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\" xmlns:samlp=\"urn:oasis:names:tc:SAML:2.0:protocol\" ID=\"\" Version=\"\" IssueInstant=\"0001-01-01T00:00:00Z\"><saml:NameID>name_id_value</saml:NameID></samlp:LogoutRequest>")
}
