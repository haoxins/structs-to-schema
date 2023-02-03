[![Build Status Badge]][Build Status]
[![Go Docs Badge]][Go Docs]

[Build Status Badge]: https://github.com/haoxins/structs-to-schema/actions/workflows/test.yaml/badge.svg
[Build Status]: https://github.com/haoxins/structs-to-schema/actions/workflows/test.yaml
[Go Docs Badge]: https://pkg.go.dev/badge/github.com/haoxins/structs-to-schema
[Go Docs]: https://pkg.go.dev/github.com/haoxins/structs-to-schema

## From Golang Structs to Schema

- Avro (Avsc)

```go
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

schema, err := ParseAvro(Event{})
```

- The `schema` will be

```js
{
    "name":"Event",
    "type":"record",
    "fields":[
        {
            "name":"clientIP",
            "type":"string"
        },
        {
            "name":"sessionId",
            "type":"string"
        },
        {
            "name":"payload",
            "type":"record",
            "fields":[
                {
                    "name":"type",
                    "type":"string"
                },
                {
                    "name":"target",
                    "type":"record",
                    "fields":[
                        {
                            "name":"type",
                            "type":"string"
                        },
                        {
                            "name":"tagName",
                            "type":"string"
                        },
                        {
                            "name":"textContent",
                            "type":"string"
                        },
                        {
                            "name":"link",
                            "type":"string"
                        }
                    ]
                },
                {
                    "name":"position",
                    "type":"record",
                    "fields":[
                        {
                            "name":"pageX",
                            "type":"int"
                        },
                        {
                            "name":"pageY",
                            "type":"int"
                        }
                    ]
                }
            ]
        }
    ]
}
```
