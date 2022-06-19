package structstoschema

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Event struct {
	ClientIP  string  `avro:"clientIP"`
	SessionId string  `avro:"sessionId"`
	Payload   Payload `avro:"payload"`
}

type Payload struct {
	Type     string        `avro:"type"`
	Target   EventTarget   `avro:"target"`
	Position EventPosition `avro:"position"`
}

type EventTarget struct {
	Type        string `avro:"type"`
	TagName     string `avro:"tagName"`
	TextContent string `avro:"textContent"`
	Link        string `avro:"link"`
}

type EventPosition struct {
	PageX int32 `avro:"pageX"`
	PageY int32 `avro:"pageY"`
}

var _ = Describe("Test Avro", func() {
	It("Parse Avro should work", func() {
		schema, err := ParseAvro(Event{})
		Expect(err).To(BeNil())
		Expect(schema).To(Equal(`{"name":"Event","type":"record","fields":[{"name":"ClientIP","type":"string"},{"name":"SessionId","type":"string"},{"name":"Payload","type":"record","fields":[{"name":"Type","type":"string"},{"name":"Target","type":"record","fields":[{"name":"Type","type":"string"},{"name":"TagName","type":"string"},{"name":"TextContent","type":"string"},{"name":"Link","type":"string"}]},{"name":"Position","type":"record","fields":[{"name":"PageX","type":"int"},{"name":"PageY","type":"int"}]}]}]}`))
	})
})
