// @ts-check

const protobuf = require("protobufjs");

const OTLP_HTTP_PROTOCOL_PROTOBUF = "http/protobuf";
const OTLP_HTTP_PROTOCOL_JSON = "http/json";

const OTLP_TRACE_PROTO = `
syntax = "proto3";
package ghaw.otlp;

message ExportTraceServiceRequest {
  repeated ResourceSpans resource_spans = 1;
}

message ResourceSpans {
  Resource resource = 1;
  repeated ScopeSpans scope_spans = 2;
  string schema_url = 3;
}

message Resource {
  repeated KeyValue attributes = 1;
  uint32 dropped_attributes_count = 2;
}

message ScopeSpans {
  InstrumentationScope scope = 1;
  repeated Span spans = 2;
  string schema_url = 3;
}

message InstrumentationScope {
  string name = 1;
  string version = 2;
  repeated KeyValue attributes = 3;
  uint32 dropped_attributes_count = 4;
}

message Span {
  enum SpanKind {
    SPAN_KIND_UNSPECIFIED = 0;
    SPAN_KIND_INTERNAL = 1;
    SPAN_KIND_SERVER = 2;
    SPAN_KIND_CLIENT = 3;
    SPAN_KIND_PRODUCER = 4;
    SPAN_KIND_CONSUMER = 5;
  }

  bytes trace_id = 1;
  bytes span_id = 2;
  string trace_state = 3;
  bytes parent_span_id = 4;
  string name = 5;
  SpanKind kind = 6;
  fixed64 start_time_unix_nano = 7;
  fixed64 end_time_unix_nano = 8;
  repeated KeyValue attributes = 9;
  uint32 dropped_attributes_count = 10;
  repeated Event events = 11;
  uint32 dropped_events_count = 12;
  repeated Link links = 13;
  uint32 dropped_links_count = 14;
  Status status = 15;
  fixed32 flags = 16;
}

message Event {
  fixed64 time_unix_nano = 1;
  string name = 2;
  repeated KeyValue attributes = 3;
  uint32 dropped_attributes_count = 4;
}

message Link {
  bytes trace_id = 1;
  bytes span_id = 2;
  string trace_state = 3;
  repeated KeyValue attributes = 4;
  uint32 dropped_attributes_count = 5;
  fixed32 flags = 6;
}

message Status {
  reserved 1;

  enum StatusCode {
    STATUS_CODE_UNSET = 0;
    STATUS_CODE_OK = 1;
    STATUS_CODE_ERROR = 2;
  }

  StatusCode code = 2;
  string message = 3;
}

message KeyValue {
  string key = 1;
  AnyValue value = 2;
}

message AnyValue {
  oneof value {
    string string_value = 1;
    bool bool_value = 2;
    int64 int_value = 3;
    double double_value = 4;
    ArrayValue array_value = 5;
    KeyValueList kvlist_value = 6;
    bytes bytes_value = 7;
  }
}

message ArrayValue {
  repeated AnyValue values = 1;
}

message KeyValueList {
  repeated KeyValue values = 1;
}
`;

const root = protobuf.parse(OTLP_TRACE_PROTO).root;
const ExportTraceServiceRequest = root.lookupType("ghaw.otlp.ExportTraceServiceRequest");

/**
 * @param {string} hex
 * @returns {Uint8Array}
 */
function hexToBytes(hex) {
  if (!hex) return new Uint8Array();
  return Buffer.from(hex, "hex");
}

/**
 * @param {Uint8Array | Buffer | string | undefined} bytes
 * @returns {string}
 */
function bytesToHex(bytes) {
  if (!bytes) return "";
  if (typeof bytes === "string") return bytes;
  return Buffer.from(bytes).toString("hex");
}

/**
 * @param {unknown} value
 * @returns {number | string | boolean | undefined | object}
 */
function normalizeNumericString(value) {
  if (typeof value === "string" && /^-?\d+$/.test(value)) {
    const parsed = Number(value);
    if (Number.isSafeInteger(parsed)) {
      return parsed;
    }
  }
  return value;
}

/**
 * @param {any} value
 * @returns {any}
 */
function encodeAnyValue(value) {
  if (!value || typeof value !== "object") {
    return undefined;
  }
  if (Object.prototype.hasOwnProperty.call(value, "stringValue")) {
    return { stringValue: value.stringValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "boolValue")) {
    return { boolValue: value.boolValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "intValue")) {
    return { intValue: value.intValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "doubleValue")) {
    return { doubleValue: value.doubleValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "bytesValue")) {
    return { bytesValue: Buffer.from(value.bytesValue) };
  }
  if (value.arrayValue && Array.isArray(value.arrayValue.values)) {
    return {
      arrayValue: {
        values: value.arrayValue.values.map(encodeAnyValue),
      },
    };
  }
  if (value.kvlistValue && Array.isArray(value.kvlistValue.values)) {
    return {
      kvlistValue: {
        values: value.kvlistValue.values.map(encodeKeyValue),
      },
    };
  }
  return undefined;
}

/**
 * @param {any} attr
 * @returns {{ key: string, value?: any }}
 */
function encodeKeyValue(attr) {
  return {
    key: attr.key,
    ...(attr.value ? { value: encodeAnyValue(attr.value) } : {}),
  };
}

/**
 * @param {any[]} attrs
 * @returns {any[]}
 */
function encodeKeyValues(attrs) {
  if (!Array.isArray(attrs)) return [];
  return attrs.map(encodeKeyValue);
}

/**
 * @param {any} payload
 * @returns {Uint8Array}
 */
function encodeExportTraceServiceRequest(payload) {
  const message = ExportTraceServiceRequest.fromObject({
    resourceSpans: Array.isArray(payload.resourceSpans)
      ? payload.resourceSpans.map(resourceSpan => ({
          ...(resourceSpan.resource
            ? {
                resource: {
                  attributes: encodeKeyValues(resourceSpan.resource.attributes),
                  ...(resourceSpan.resource.droppedAttributesCount !== undefined ? { droppedAttributesCount: resourceSpan.resource.droppedAttributesCount } : {}),
                },
              }
            : {}),
          scopeSpans: Array.isArray(resourceSpan.scopeSpans)
            ? resourceSpan.scopeSpans.map(scopeSpan => ({
                ...(scopeSpan.scope
                  ? {
                      scope: {
                        ...(scopeSpan.scope.name ? { name: scopeSpan.scope.name } : {}),
                        ...(scopeSpan.scope.version ? { version: scopeSpan.scope.version } : {}),
                        ...(scopeSpan.scope.attributes ? { attributes: encodeKeyValues(scopeSpan.scope.attributes) } : {}),
                        ...(scopeSpan.scope.droppedAttributesCount !== undefined ? { droppedAttributesCount: scopeSpan.scope.droppedAttributesCount } : {}),
                      },
                    }
                  : {}),
                spans: Array.isArray(scopeSpan.spans)
                  ? scopeSpan.spans.map(span => ({
                      traceId: hexToBytes(span.traceId),
                      spanId: hexToBytes(span.spanId),
                      ...(span.parentSpanId ? { parentSpanId: hexToBytes(span.parentSpanId) } : {}),
                      ...(span.traceState ? { traceState: span.traceState } : {}),
                      ...(span.flags !== undefined ? { flags: span.flags } : {}),
                      name: span.name,
                      kind: span.kind,
                      startTimeUnixNano: span.startTimeUnixNano,
                      endTimeUnixNano: span.endTimeUnixNano,
                      attributes: encodeKeyValues(span.attributes),
                      ...(span.droppedAttributesCount !== undefined ? { droppedAttributesCount: span.droppedAttributesCount } : {}),
                      ...(Array.isArray(span.events)
                        ? {
                            events: span.events.map(event => ({
                              timeUnixNano: event.timeUnixNano,
                              name: event.name,
                              attributes: encodeKeyValues(event.attributes),
                              ...(event.droppedAttributesCount !== undefined ? { droppedAttributesCount: event.droppedAttributesCount } : {}),
                            })),
                          }
                        : {}),
                      ...(span.droppedEventsCount !== undefined ? { droppedEventsCount: span.droppedEventsCount } : {}),
                      ...(Array.isArray(span.links)
                        ? {
                            links: span.links.map(link => ({
                              traceId: hexToBytes(link.traceId),
                              spanId: hexToBytes(link.spanId),
                              ...(link.traceState ? { traceState: link.traceState } : {}),
                              attributes: encodeKeyValues(link.attributes),
                              ...(link.droppedAttributesCount !== undefined ? { droppedAttributesCount: link.droppedAttributesCount } : {}),
                              ...(link.flags !== undefined ? { flags: link.flags } : {}),
                            })),
                          }
                        : {}),
                      ...(span.droppedLinksCount !== undefined ? { droppedLinksCount: span.droppedLinksCount } : {}),
                      ...(span.status
                        ? {
                            status: {
                              ...(span.status.message ? { message: span.status.message } : {}),
                              ...(span.status.code !== undefined ? { code: span.status.code } : {}),
                            },
                          }
                        : {}),
                    }))
                  : [],
                ...(resourceSpan.schemaUrl ? { schemaUrl: resourceSpan.schemaUrl } : {}),
              }))
            : [],
          ...(resourceSpan.schemaUrl ? { schemaUrl: resourceSpan.schemaUrl } : {}),
        }))
      : [],
  });

  const err = ExportTraceServiceRequest.verify(message);
  if (err) {
    throw new Error(`Invalid OTLP protobuf payload: ${err}`);
  }

  return ExportTraceServiceRequest.encode(message).finish();
}

/**
 * @param {any} value
 * @returns {any}
 */
function decodeAnyValue(value) {
  if (!value || typeof value !== "object") return value;
  if (Object.prototype.hasOwnProperty.call(value, "stringValue")) {
    return { stringValue: value.stringValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "boolValue")) {
    return { boolValue: value.boolValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "intValue")) {
    return { intValue: normalizeNumericString(value.intValue) };
  }
  if (Object.prototype.hasOwnProperty.call(value, "doubleValue")) {
    return { doubleValue: value.doubleValue };
  }
  if (Object.prototype.hasOwnProperty.call(value, "bytesValue")) {
    return { bytesValue: Buffer.from(value.bytesValue) };
  }
  if (value.arrayValue && Array.isArray(value.arrayValue.values)) {
    return {
      arrayValue: {
        values: value.arrayValue.values.map(decodeAnyValue),
      },
    };
  }
  if (value.kvlistValue && Array.isArray(value.kvlistValue.values)) {
    return {
      kvlistValue: {
        values: value.kvlistValue.values.map(decodeKeyValue),
      },
    };
  }
  return value;
}

/**
 * @param {any} attr
 * @returns {{ key: string, value?: any }}
 */
function decodeKeyValue(attr) {
  return {
    key: attr.key,
    ...(attr.value ? { value: decodeAnyValue(attr.value) } : {}),
  };
}

/**
 * @param {Uint8Array | Buffer} buffer
 * @returns {any}
 */
function decodeExportTraceServiceRequest(buffer) {
  const decoded = ExportTraceServiceRequest.decode(buffer);
  const object = ExportTraceServiceRequest.toObject(decoded, {
    longs: String,
    enums: Number,
    bytes: Buffer,
    defaults: false,
  });

  return {
    resourceSpans: Array.isArray(object.resourceSpans)
      ? object.resourceSpans.map(resourceSpan => ({
          ...(resourceSpan.resource
            ? {
                resource: {
                  ...(Array.isArray(resourceSpan.resource.attributes) ? { attributes: resourceSpan.resource.attributes.map(decodeKeyValue) } : {}),
                  ...(resourceSpan.resource.droppedAttributesCount !== undefined ? { droppedAttributesCount: resourceSpan.resource.droppedAttributesCount } : {}),
                },
              }
            : {}),
          ...(Array.isArray(resourceSpan.scopeSpans)
            ? {
                scopeSpans: resourceSpan.scopeSpans.map(scopeSpan => ({
                  ...(scopeSpan.scope
                    ? {
                        scope: {
                          ...(scopeSpan.scope.name ? { name: scopeSpan.scope.name } : {}),
                          ...(scopeSpan.scope.version ? { version: scopeSpan.scope.version } : {}),
                          ...(Array.isArray(scopeSpan.scope.attributes) ? { attributes: scopeSpan.scope.attributes.map(decodeKeyValue) } : {}),
                          ...(scopeSpan.scope.droppedAttributesCount !== undefined ? { droppedAttributesCount: scopeSpan.scope.droppedAttributesCount } : {}),
                        },
                      }
                    : {}),
                  spans: Array.isArray(scopeSpan.spans)
                    ? scopeSpan.spans.map(span => ({
                        traceId: bytesToHex(span.traceId),
                        spanId: bytesToHex(span.spanId),
                        ...(span.parentSpanId ? { parentSpanId: bytesToHex(span.parentSpanId) } : {}),
                        ...(span.traceState ? { traceState: span.traceState } : {}),
                        ...(span.flags !== undefined ? { flags: span.flags } : {}),
                        name: span.name,
                        kind: span.kind,
                        startTimeUnixNano: span.startTimeUnixNano,
                        endTimeUnixNano: span.endTimeUnixNano,
                        ...(Array.isArray(span.attributes) ? { attributes: span.attributes.map(decodeKeyValue) } : { attributes: [] }),
                        ...(span.droppedAttributesCount !== undefined ? { droppedAttributesCount: span.droppedAttributesCount } : {}),
                        ...(Array.isArray(span.events)
                          ? {
                              events: span.events.map(event => ({
                                ...(event.timeUnixNano !== undefined ? { timeUnixNano: event.timeUnixNano } : {}),
                                name: event.name,
                                ...(Array.isArray(event.attributes) ? { attributes: event.attributes.map(decodeKeyValue) } : {}),
                                ...(event.droppedAttributesCount !== undefined ? { droppedAttributesCount: event.droppedAttributesCount } : {}),
                              })),
                            }
                          : {}),
                        ...(span.droppedEventsCount !== undefined ? { droppedEventsCount: span.droppedEventsCount } : {}),
                        ...(Array.isArray(span.links)
                          ? {
                              links: span.links.map(link => ({
                                traceId: bytesToHex(link.traceId),
                                spanId: bytesToHex(link.spanId),
                                ...(link.traceState ? { traceState: link.traceState } : {}),
                                ...(Array.isArray(link.attributes) ? { attributes: link.attributes.map(decodeKeyValue) } : {}),
                                ...(link.droppedAttributesCount !== undefined ? { droppedAttributesCount: link.droppedAttributesCount } : {}),
                                ...(link.flags !== undefined ? { flags: link.flags } : {}),
                              })),
                            }
                          : {}),
                        ...(span.droppedLinksCount !== undefined ? { droppedLinksCount: span.droppedLinksCount } : {}),
                        ...(span.status
                          ? {
                              status: {
                                ...(span.status.message ? { message: span.status.message } : {}),
                                ...(span.status.code !== undefined ? { code: span.status.code } : {}),
                              },
                            }
                          : {}),
                      }))
                    : [],
                  ...(scopeSpan.schemaUrl ? { schemaUrl: scopeSpan.schemaUrl } : {}),
                })),
              }
            : { scopeSpans: [] }),
          ...(resourceSpan.schemaUrl ? { schemaUrl: resourceSpan.schemaUrl } : {}),
        }))
      : [],
  };
}

module.exports = {
  OTLP_HTTP_PROTOCOL_PROTOBUF,
  OTLP_HTTP_PROTOCOL_JSON,
  encodeExportTraceServiceRequest,
  decodeExportTraceServiceRequest,
};
