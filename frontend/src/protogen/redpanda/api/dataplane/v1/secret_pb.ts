// @generated by protoc-gen-es v2.2.5 with parameter "target=ts"
// @generated from file redpanda/api/dataplane/v1/secret.proto (package redpanda.api.dataplane.v1, syntax proto3)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage, GenService } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_buf_validate_validate } from "../../../../buf/validate/validate_pb";
import { file_google_api_annotations } from "../../../../google/api/annotations_pb";
import { file_google_api_field_behavior } from "../../../../google/api/field_behavior_pb";
import { file_protoc_gen_openapiv2_options_annotations } from "../../../../protoc-gen-openapiv2/options/annotations_pb";
import { file_redpanda_api_auth_v1_authorization } from "../../auth/v1/authorization_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file redpanda/api/dataplane/v1/secret.proto.
 */
export const file_redpanda_api_dataplane_v1_secret: GenFile = /*@__PURE__*/
  fileDesc("CiZyZWRwYW5kYS9hcGkvZGF0YXBsYW5lL3YxL3NlY3JldC5wcm90bxIZcmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MSLsAQoGU2VjcmV0EhIKAmlkGAEgASgJQgbgQQXgQQMSbQoGbGFiZWxzGAIgAygLMi0ucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TZWNyZXQuTGFiZWxzRW50cnlCLuBBBbpIKJoBJSojciEyH14oW1xwe0x9XHB7Wn1ccHtOfV8uOi89K1wtQF0qKSQSMAoGc2NvcGVzGAMgAygOMiAucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TY29wZRotCgtMYWJlbHNFbnRyeRILCgNrZXkYASABKAkSDQoFdmFsdWUYAiABKAk6AjgBImIKE0xpc3RTZWNyZXRzUmVzcG9uc2USMgoHc2VjcmV0cxgBIAMoCzIhLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuU2VjcmV0EhcKD25leHRfcGFnZV90b2tlbhgCIAEoCSKgAgoRTGlzdFNlY3JldHNGaWx0ZXISMwoNbmFtZV9jb250YWlucxgBIAEoCUIcukgZchcYgAEyEl4oW2EtekEtWjAtOS1fXSopJBJ1CgZsYWJlbHMYAiADKAsyOC5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkxpc3RTZWNyZXRzRmlsdGVyLkxhYmVsc0VudHJ5Qiu6SCiaASUqI3IhMh9eKFtccHtMfVxwe1p9XHB7Tn1fLjovPStcLUBdKikkEjAKBnNjb3BlcxgDIAMoDjIgLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuU2NvcGUaLQoLTGFiZWxzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ASLEAQoSTGlzdFNlY3JldHNSZXF1ZXN0EjwKBmZpbHRlchgBIAEoCzIsLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuTGlzdFNlY3JldHNGaWx0ZXISEgoKcGFnZV90b2tlbhgCIAEoCRJcCglwYWdlX3NpemUYAyABKAVCSZJBRjIyTGltaXQgdGhlIHBhZ2luYXRlZCByZXNwb25zZSB0byBhIG51bWJlciBvZiBpdGVtcy5ZAAAAAAAASUBpAAAAAAAA8D8iPQoQR2V0U2VjcmV0UmVxdWVzdBIpCgJpZBgBIAEoCUIdukgachgQARj/ATIRXltBLVpdW0EtWjAtOV9dKiQiRgoRR2V0U2VjcmV0UmVzcG9uc2USMQoGc2VjcmV0GAEgASgLMiEucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TZWNyZXQi2gIKE0NyZWF0ZVNlY3JldFJlcXVlc3QSKQoCaWQYASABKAlCHbpIGnIYEAEY/wEyEV5bQS1aXVtBLVowLTlfXSokEnoKBmxhYmVscxgCIAMoCzI6LnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuQ3JlYXRlU2VjcmV0UmVxdWVzdC5MYWJlbHNFbnRyeUIu4EEFukgomgElKiNyITIfXihbXHB7TH1ccHtafVxwe059Xy46Lz0rXC1AXSopJBJFCgZzY29wZXMYAyADKA4yIC5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlNjb3BlQhO6SBCSAQ0IARgBIgeCAQQQASAAEiYKC3NlY3JldF9kYXRhGAQgASgMQhHgQQTgQQK6SAh6BhABGICABBotCgtMYWJlbHNFbnRyeRILCgNrZXkYASABKAkSDQoFdmFsdWUYAiABKAk6AjgBIkkKFENyZWF0ZVNlY3JldFJlc3BvbnNlEjEKBnNlY3JldBgBIAEoCzIhLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuU2VjcmV0ItoCChNVcGRhdGVTZWNyZXRSZXF1ZXN0EikKAmlkGAEgASgJQh26SBpyGBABGP8BMhFeW0EtWl1bQS1aMC05X10qJBJ6CgZsYWJlbHMYAiADKAsyOi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZVNlY3JldFJlcXVlc3QuTGFiZWxzRW50cnlCLuBBBbpIKJoBJSojciEyH14oW1xwe0x9XHB7Wn1ccHtOfV8uOi89K1wtQF0qKSQSRQoGc2NvcGVzGAMgAygOMiAucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TY29wZUITukgQkgENCAEYASIHggEEEAEgABImCgtzZWNyZXRfZGF0YRgEIAEoDEIR4EEE4EECukgIegYQARiAgAQaLQoLTGFiZWxzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ASJJChRVcGRhdGVTZWNyZXRSZXNwb25zZRIxCgZzZWNyZXQYASABKAsyIS5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlNlY3JldCJAChNEZWxldGVTZWNyZXRSZXF1ZXN0EikKAmlkGAEgASgJQh26SBpyGBABGP8BMhFeW0EtWl1bQS1aMC05X10qJCIWChREZWxldGVTZWNyZXRSZXNwb25zZSIZChdMaXN0U2VjcmV0U2NvcGVzUmVxdWVzdCJMChhMaXN0U2VjcmV0U2NvcGVzUmVzcG9uc2USMAoGc2NvcGVzGAEgAygOMiAucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TY29wZSL3AQocR2V0S2Fma2FDb25uZWN0U2VjcmV0UmVxdWVzdBKmAQoMY2x1c3Rlcl9uYW1lGAEgASgJQo8BkkFqMkpVbmlxdWUgbmFtZSBvZiB0YXJnZXQgY29ubmVjdCBjbHVzdGVyLiBGb3IgUmVkcGFuZGEgQ2xvdWQsIHVzZSBgcmVkcGFuZGFgLkoKInJlZHBhbmRhIso+D/oCDGNsdXN0ZXJfbmFtZeBBArpIHMgBAXIXEAEYgAEyEF5bYS16QS1aMC05LV9dKyQSLgoCaWQYAiABKAlCIrpIH3IdEAEY/wEyFl5bYS16QS1aMC05L18rPS5AJS1dKyQiUgodR2V0S2Fma2FDb25uZWN0U2VjcmV0UmVzcG9uc2USMQoGc2VjcmV0GAEgASgLMiEucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5TZWNyZXQiyQMKH0NyZWF0ZUthZmthQ29ubmVjdFNlY3JldFJlcXVlc3QSkwEKDGNsdXN0ZXJfbmFtZRgBIAEoCUJ9kkFYMkpVbmlxdWUgbmFtZSBvZiB0YXJnZXQgY29ubmVjdCBjbHVzdGVyLiBGb3IgUmVkcGFuZGEgQ2xvdWQsIHVzZSBgcmVkcGFuZGFgLkoKInJlZHBhbmRhIuBBArpIHMgBAXIXEAEYgAEyEF5bYS16QS1aMC05LV9dKyQSMAoEbmFtZRgCIAEoCUIi4EECukgcyAEBchcQARiAATIQXlthLXpBLVowLTktX10rJBKGAQoGbGFiZWxzGAMgAygLMkYucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5DcmVhdGVLYWZrYUNvbm5lY3RTZWNyZXRSZXF1ZXN0LkxhYmVsc0VudHJ5Qi7gQQW6SCiaASUqI3IhMh9eKFtccHtMfVxwe1p9XHB7Tn1fLjovPStcLUBdKikkEiYKC3NlY3JldF9kYXRhGAQgASgMQhHgQQTgQQK6SAh6BhABGICABBotCgtMYWJlbHNFbnRyeRILCgNrZXkYASABKAkSDQoFdmFsdWUYAiABKAk6AjgBIlUKIENyZWF0ZUthZmthQ29ubmVjdFNlY3JldFJlc3BvbnNlEjEKBnNlY3JldBgBIAEoCzIhLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuU2VjcmV0Ir0DCh5MaXN0S2Fma2FDb25uZWN0U2VjcmV0c1JlcXVlc3QSpgEKDGNsdXN0ZXJfbmFtZRgBIAEoCUKPAZJBajJKVW5pcXVlIG5hbWUgb2YgdGFyZ2V0IGNvbm5lY3QgY2x1c3Rlci4gRm9yIFJlZHBhbmRhIENsb3VkLCB1c2UgYHJlZHBhbmRhYC5KCiJyZWRwYW5kYSLKPg/6AgxjbHVzdGVyX25hbWXgQQK6SBzIAQFyFxABGIABMhBeW2EtekEtWjAtOS1fXSskEjwKBmZpbHRlchgCIAEoCzIsLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuTGlzdFNlY3JldHNGaWx0ZXISEgoKcGFnZV90b2tlbhgDIAEoCRKfAQoJcGFnZV9zaXplGAQgASgFQosBkkF1MmFMaW1pdCB0aGUgcGFnaW5hdGVkIHJlc3BvbnNlIHRvIGEgbnVtYmVyIG9mIGl0ZW1zLiBEZWZhdWx0cyB0byAxMDAuIFVzZSAtMSB0byBkaXNhYmxlIHBhZ2luYXRpb24uWQAAAAAAQI9AaQAAAAAAAPC/ukgQGg4Y6Aco////////////ASJuCh9MaXN0S2Fma2FDb25uZWN0U2VjcmV0c1Jlc3BvbnNlEjIKB3NlY3JldHMYASADKAsyIS5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlNlY3JldBIXCg9uZXh0X3BhZ2VfdG9rZW4YAiABKAkixwMKH1VwZGF0ZUthZmthQ29ubmVjdFNlY3JldFJlcXVlc3QSkwEKDGNsdXN0ZXJfbmFtZRgBIAEoCUJ9kkFYMkpVbmlxdWUgbmFtZSBvZiB0YXJnZXQgY29ubmVjdCBjbHVzdGVyLiBGb3IgUmVkcGFuZGEgQ2xvdWQsIHVzZSBgcmVkcGFuZGFgLkoKInJlZHBhbmRhIuBBArpIHMgBAXIXEAEYgAEyEF5bYS16QS1aMC05LV9dKyQSLgoCaWQYAiABKAlCIrpIH3IdEAEY/wEyFl5bYS16QS1aMC05L18rPS5AJS1dKyQShgEKBmxhYmVscxgDIAMoCzJGLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuVXBkYXRlS2Fma2FDb25uZWN0U2VjcmV0UmVxdWVzdC5MYWJlbHNFbnRyeUIu4EEFukgomgElKiNyITIfXihbXHB7TH1ccHtafVxwe059Xy46Lz0rXC1AXSopJBImCgtzZWNyZXRfZGF0YRgEIAEoDEIR4EEE4EECukgIegYQARiAgAQaLQoLTGFiZWxzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ASJVCiBVcGRhdGVLYWZrYUNvbm5lY3RTZWNyZXRSZXNwb25zZRIxCgZzZWNyZXQYASABKAsyIS5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlNlY3JldCL6AQofRGVsZXRlS2Fma2FDb25uZWN0U2VjcmV0UmVxdWVzdBKmAQoMY2x1c3Rlcl9uYW1lGAEgASgJQo8BkkFqMkpVbmlxdWUgbmFtZSBvZiB0YXJnZXQgY29ubmVjdCBjbHVzdGVyLiBGb3IgUmVkcGFuZGEgQ2xvdWQsIHVzZSBgcmVkcGFuZGFgLkoKInJlZHBhbmRhIso+D/oCDGNsdXN0ZXJfbmFtZeBBArpIHMgBAXIXEAEYgAEyEF5bYS16QS1aMC05LV9dKyQSLgoCaWQYAiABKAlCIrpIH3IdEAEY/wEyFl5bYS16QS1aMC05L18rPS5AJS1dKyQiIgogRGVsZXRlS2Fma2FDb25uZWN0U2VjcmV0UmVzcG9uc2UqVgoFU2NvcGUSFQoRU0NPUEVfVU5TUEVDSUZJRUQQABIaChZTQ09QRV9SRURQQU5EQV9DT05ORUNUEAESGgoWU0NPUEVfUkVEUEFOREFfQ0xVU1RFUhACMuAdCg1TZWNyZXRTZXJ2aWNlEpMCCglHZXRTZWNyZXQSKy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkdldFNlY3JldFJlcXVlc3QaLC5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkdldFNlY3JldFJlc3BvbnNlIqoBkkGGARIKR2V0IFNlY3JldBoNR2V0IGEgc2VjcmV0Lko9CgMyMDASNgoCT0sSMAouGiwucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5HZXRTZWNyZXRSZXNwb25zZUoqCgM0MDQSIwoJTm90IEZvdW5kEhYKFBoSLmdvb2dsZS5ycGMuU3RhdHVziqYdBAgBEAaC0+STAhISEC92MS9zZWNyZXRzL3tpZH0SngIKC0xpc3RTZWNyZXRzEi0ucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5MaXN0U2VjcmV0c1JlcXVlc3QaLi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkxpc3RTZWNyZXRzUmVzcG9uc2UirwGSQZABEgxMaXN0IFNlY3JldHMaP0xpc3Qgc2VjcmV0cy4gT3B0aW9uYWw6IGZpbHRlciBiYXNlZCBvbiBzZWNyZXQgbmFtZSBhbmQgbGFiZWxzLko/CgMyMDASOAoCT0sSMgowGi4ucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5MaXN0U2VjcmV0c1Jlc3BvbnNliqYdBAgBEAaC0+STAg0SCy92MS9zZWNyZXRzEoICCgxDcmVhdGVTZWNyZXQSLi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkNyZWF0ZVNlY3JldFJlcXVlc3QaLy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkNyZWF0ZVNlY3JldFJlc3BvbnNlIpABkkFvEg1DcmVhdGUgU2VjcmV0GhBDcmVhdGUgYSBzZWNyZXQuSkwKAzIwMRJFCg5TZWNyZXQgY3JlYXRlZBIzCjEaLy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkNyZWF0ZVNlY3JldFJlc3BvbnNliqYdBAgCEAaC0+STAhA6ASoiCy92MS9zZWNyZXRzEqgCCgxVcGRhdGVTZWNyZXQSLi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZVNlY3JldFJlcXVlc3QaLy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZVNlY3JldFJlc3BvbnNlIrYBkkGPARINVXBkYXRlIFNlY3JldBoQVXBkYXRlIGEgc2VjcmV0LkpACgMyMDASOQoCT0sSMwoxGi8ucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5VcGRhdGVTZWNyZXRSZXNwb25zZUoqCgM0MDQSIwoJTm90IEZvdW5kEhYKFBoSLmdvb2dsZS5ycGMuU3RhdHVziqYdBAgCEAaC0+STAhU6ASoaEC92MS9zZWNyZXRzL3tpZH0SigIKDERlbGV0ZVNlY3JldBIuLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuRGVsZXRlU2VjcmV0UmVxdWVzdBovLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuRGVsZXRlU2VjcmV0UmVzcG9uc2UimAGSQXUSDURlbGV0ZSBTZWNyZXQaEERlbGV0ZSBhIHNlY3JldC5KJgoDMjA0Eh8KG1NlY3JldCBkZWxldGVkIHN1Y2Nlc3NmdWxseRIASioKAzQwNBIjCglOb3QgRm91bmQSFgoUGhIuZ29vZ2xlLnJwYy5TdGF0dXOKph0ECAIQBoLT5JMCEioQL3YxL3NlY3JldHMve2lkfRKbAgoQTGlzdFNlY3JldFNjb3BlcxIyLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuTGlzdFNlY3JldFNjb3Blc1JlcXVlc3QaMy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkxpc3RTZWNyZXRTY29wZXNSZXNwb25zZSKdAZJBeRISTGlzdCBTZWNyZXQgU2NvcGVzGh1MaXN0IHN1cHBvcnRlZCBzZWNyZXQgc2NvcGVzLkpECgMyMDASPQoCT0sSNwo1GjMucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5MaXN0U2VjcmV0U2NvcGVzUmVzcG9uc2WKph0ECAEQBoLT5JMCExIRL3YxL3NlY3JldC1zY29wZXMSmAMKFUdldEthZmthQ29ubmVjdFNlY3JldBI3LnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuR2V0S2Fma2FDb25uZWN0U2VjcmV0UmVxdWVzdBo4LnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuR2V0S2Fma2FDb25uZWN0U2VjcmV0UmVzcG9uc2UiiwKSQcEBEhpHZXQgQ29ubmVjdCBDbHVzdGVyIFNlY3JldBosR2V0IGEgc3BlY2lmaWMgS2Fma2EgQ29ubmVjdCBjbHVzdGVyIHNlY3JldC5KSQoDMjAwEkIKAk9LEjwKOho4LnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjEuR2V0S2Fma2FDb25uZWN0U2VjcmV0UmVzcG9uc2VKKgoDNDA0EiMKCU5vdCBGb3VuZBIWChQaEi5nb29nbGUucnBjLlN0YXR1c4qmHQQIARAGgtPkkwI4EjYvdjEva2Fma2EtY29ubmVjdC9jbHVzdGVycy97Y2x1c3Rlcl9uYW1lfS9zZWNyZXRzL3tpZH0SjgMKF0xpc3RLYWZrYUNvbm5lY3RTZWNyZXRzEjkucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5MaXN0S2Fma2FDb25uZWN0U2VjcmV0c1JlcXVlc3QaOi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkxpc3RLYWZrYUNvbm5lY3RTZWNyZXRzUmVzcG9uc2Ui+wGSQbYBEhxMaXN0IENvbm5lY3QgQ2x1c3RlciBTZWNyZXRzGlVMaXN0IEthZmthIENvbm5lY3QgY2x1c3RlciBzZWNyZXRzLiBPcHRpb25hbDogZmlsdGVyIGJhc2VkIG9uIHNlY3JldCBuYW1lIGFuZCBsYWJlbHMuSj8KAzIwMBI4CgJPSxIyCjAaLi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkxpc3RTZWNyZXRzUmVzcG9uc2WKph0ECAEQBoLT5JMCMxIxL3YxL2thZmthLWNvbm5lY3QvY2x1c3RlcnMve2NsdXN0ZXJfbmFtZX0vc2VjcmV0cxL8AgoYQ3JlYXRlS2Fma2FDb25uZWN0U2VjcmV0EjoucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5DcmVhdGVLYWZrYUNvbm5lY3RTZWNyZXRSZXF1ZXN0GjsucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5DcmVhdGVLYWZrYUNvbm5lY3RTZWNyZXRSZXNwb25zZSLmAZJBngESHUNyZWF0ZSBDb25uZWN0IENsdXN0ZXIgU2VjcmV0GiZDcmVhdGUgYSBLYWZrYSBDb25uZWN0IGNsdXN0ZXIgc2VjcmV0LkpVCgMyMDESTgoOU2VjcmV0IGNyZWF0ZWQSPAo6GjgucmVkcGFuZGEuYXBpLmRhdGFwbGFuZS52MS5HZXRLYWZrYUNvbm5lY3RTZWNyZXRSZXNwb25zZYqmHQQIAhAGgtPkkwI2OgEqIjEvdjEva2Fma2EtY29ubmVjdC9jbHVzdGVycy97Y2x1c3Rlcl9uYW1lfS9zZWNyZXRzEqQDChhVcGRhdGVLYWZrYUNvbm5lY3RTZWNyZXQSOi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZUthZmthQ29ubmVjdFNlY3JldFJlcXVlc3QaOy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZUthZmthQ29ubmVjdFNlY3JldFJlc3BvbnNlIo4CkkHBARIdVXBkYXRlIENvbm5lY3QgQ2x1c3RlciBTZWNyZXQaJlVwZGF0ZSBhIEthZmthIENvbm5lY3QgY2x1c3RlciBzZWNyZXQuSkwKAzIwMBJFCgJPSxI/Cj0aOy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLlVwZGF0ZUthZmthQ29ubmVjdFNlY3JldFJlc3BvbnNlSioKAzQwNBIjCglOb3QgRm91bmQSFgoUGhIuZ29vZ2xlLnJwYy5TdGF0dXOKph0ECAIQBoLT5JMCOzoBKho2L3YxL2thZmthLWNvbm5lY3QvY2x1c3RlcnMve2NsdXN0ZXJfbmFtZX0vc2VjcmV0cy97aWR9EvsCChhEZWxldGVLYWZrYUNvbm5lY3RTZWNyZXQSOi5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkRlbGV0ZUthZmthQ29ubmVjdFNlY3JldFJlcXVlc3QaOy5yZWRwYW5kYS5hcGkuZGF0YXBsYW5lLnYxLkRlbGV0ZUthZmthQ29ubmVjdFNlY3JldFJlc3BvbnNlIuUBkkGbARIdRGVsZXRlIENvbm5lY3QgQ2x1c3RlciBTZWNyZXQaJkRlbGV0ZSBhIEthZmthIENvbm5lY3QgY2x1c3RlciBzZWNyZXQuSiYKAzIwNBIfChtTZWNyZXQgZGVsZXRlZCBzdWNjZXNzZnVsbHkSAEoqCgM0MDQSIwoJTm90IEZvdW5kEhYKFBoSLmdvb2dsZS5ycGMuU3RhdHVziqYdBAgCEAaC0+STAjgqNi92MS9rYWZrYS1jb25uZWN0L2NsdXN0ZXJzL3tjbHVzdGVyX25hbWV9L3NlY3JldHMve2lkfRptkkFqCgdTZWNyZXRzEl9NYW5hZ2UgW3NlY3JldHNdKGh0dHBzOi8vZG9jcy5yZWRwYW5kYS5jb20vcmVkcGFuZGEtY2xvdWQvc2VjdXJpdHkvc2VjcmV0cykgZm9yIFJlZHBhbmRhIENsb3VkLkKQAgodY29tLnJlZHBhbmRhLmFwaS5kYXRhcGxhbmUudjFCC1NlY3JldFByb3RvUAFaW2dpdGh1Yi5jb20vcmVkcGFuZGEtZGF0YS9jb25zb2xlL2JhY2tlbmQvcGtnL3Byb3RvZ2VuL3JlZHBhbmRhL2FwaS9kYXRhcGxhbmUvdjE7ZGF0YXBsYW5ldjGiAgNSQUSqAhlSZWRwYW5kYS5BcGkuRGF0YXBsYW5lLlYxygIZUmVkcGFuZGFcQXBpXERhdGFwbGFuZVxWMeICJVJlZHBhbmRhXEFwaVxEYXRhcGxhbmVcVjFcR1BCTWV0YWRhdGHqAhxSZWRwYW5kYTo6QXBpOjpEYXRhcGxhbmU6OlYxYgZwcm90bzM", [file_buf_validate_validate, file_google_api_annotations, file_google_api_field_behavior, file_protoc_gen_openapiv2_options_annotations, file_redpanda_api_auth_v1_authorization]);

/**
 * Defines the secret resource.
 *
 * @generated from message redpanda.api.dataplane.v1.Secret
 */
export type Secret = Message<"redpanda.api.dataplane.v1.Secret"> & {
  /**
   * Secret identifier.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Secret labels.
   *
   * @generated from field: map<string, string> labels = 2;
   */
  labels: { [key: string]: string };

  /**
   * Secret scopes
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Scope scopes = 3;
   */
  scopes: Scope[];
};

/**
 * Describes the message redpanda.api.dataplane.v1.Secret.
 * Use `create(SecretSchema)` to create a new message.
 */
export const SecretSchema: GenMessage<Secret> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 0);

/**
 * ListSecretsResponse is the response of ListSecrets.
 *
 * @generated from message redpanda.api.dataplane.v1.ListSecretsResponse
 */
export type ListSecretsResponse = Message<"redpanda.api.dataplane.v1.ListSecretsResponse"> & {
  /**
   * Secrets retrieved.
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Secret secrets = 1;
   */
  secrets: Secret[];

  /**
   * Token to retrieve the next page.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListSecretsResponse.
 * Use `create(ListSecretsResponseSchema)` to create a new message.
 */
export const ListSecretsResponseSchema: GenMessage<ListSecretsResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 1);

/**
 * ListSecretsFilter are the filter options for listing secrets.
 *
 * @generated from message redpanda.api.dataplane.v1.ListSecretsFilter
 */
export type ListSecretsFilter = Message<"redpanda.api.dataplane.v1.ListSecretsFilter"> & {
  /**
   * Substring match on secret name. Case-sensitive.
   *
   * @generated from field: string name_contains = 1;
   */
  nameContains: string;

  /**
   * The secret labels to search for.
   *
   * @generated from field: map<string, string> labels = 2;
   */
  labels: { [key: string]: string };

  /**
   * Secret scopes to search for
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Scope scopes = 3;
   */
  scopes: Scope[];
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListSecretsFilter.
 * Use `create(ListSecretsFilterSchema)` to create a new message.
 */
export const ListSecretsFilterSchema: GenMessage<ListSecretsFilter> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 2);

/**
 * ListSecretsRequest is the request of ListSecrets.
 *
 * @generated from message redpanda.api.dataplane.v1.ListSecretsRequest
 */
export type ListSecretsRequest = Message<"redpanda.api.dataplane.v1.ListSecretsRequest"> & {
  /**
   * List filter.
   *
   * @generated from field: redpanda.api.dataplane.v1.ListSecretsFilter filter = 1;
   */
  filter?: ListSecretsFilter;

  /**
   * Value of the next_page_token field returned by the previous response.
   * If not provided, the system assumes the first page is requested.
   *
   * @generated from field: string page_token = 2;
   */
  pageToken: string;

  /**
   * Limit the paginated response to a number of items.
   *
   * @generated from field: int32 page_size = 3;
   */
  pageSize: number;
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListSecretsRequest.
 * Use `create(ListSecretsRequestSchema)` to create a new message.
 */
export const ListSecretsRequestSchema: GenMessage<ListSecretsRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 3);

/**
 * GetSecretRequest is the request of GetSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.GetSecretRequest
 */
export type GetSecretRequest = Message<"redpanda.api.dataplane.v1.GetSecretRequest"> & {
  /**
   * The id of the secret to retrieve.
   *
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.GetSecretRequest.
 * Use `create(GetSecretRequestSchema)` to create a new message.
 */
export const GetSecretRequestSchema: GenMessage<GetSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 4);

/**
 * GetSecretResponse is the response of GetSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.GetSecretResponse
 */
export type GetSecretResponse = Message<"redpanda.api.dataplane.v1.GetSecretResponse"> & {
  /**
   * The created secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.GetSecretResponse.
 * Use `create(GetSecretResponseSchema)` to create a new message.
 */
export const GetSecretResponseSchema: GenMessage<GetSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 5);

/**
 * CreateSecretRequest is the request of CreateSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.CreateSecretRequest
 */
export type CreateSecretRequest = Message<"redpanda.api.dataplane.v1.CreateSecretRequest"> & {
  /**
   * Secret identifier.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Secret labels.
   *
   * @generated from field: map<string, string> labels = 2;
   */
  labels: { [key: string]: string };

  /**
   * Secret scopes
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Scope scopes = 3;
   */
  scopes: Scope[];

  /**
   * The secret data. Must be Base64-encoded.
   *
   * @generated from field: bytes secret_data = 4;
   */
  secretData: Uint8Array;
};

/**
 * Describes the message redpanda.api.dataplane.v1.CreateSecretRequest.
 * Use `create(CreateSecretRequestSchema)` to create a new message.
 */
export const CreateSecretRequestSchema: GenMessage<CreateSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 6);

/**
 * CreateSecretResponse is the response of CreateSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.CreateSecretResponse
 */
export type CreateSecretResponse = Message<"redpanda.api.dataplane.v1.CreateSecretResponse"> & {
  /**
   * The created secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.CreateSecretResponse.
 * Use `create(CreateSecretResponseSchema)` to create a new message.
 */
export const CreateSecretResponseSchema: GenMessage<CreateSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 7);

/**
 * UpdateSecretRequest is the request of UpdateSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.UpdateSecretRequest
 */
export type UpdateSecretRequest = Message<"redpanda.api.dataplane.v1.UpdateSecretRequest"> & {
  /**
   * Secret identifier.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * Secret labels.
   *
   * @generated from field: map<string, string> labels = 2;
   */
  labels: { [key: string]: string };

  /**
   * Secret scopes
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Scope scopes = 3;
   */
  scopes: Scope[];

  /**
   * The secret data. Must be Base64-encoded.
   *
   * @generated from field: bytes secret_data = 4;
   */
  secretData: Uint8Array;
};

/**
 * Describes the message redpanda.api.dataplane.v1.UpdateSecretRequest.
 * Use `create(UpdateSecretRequestSchema)` to create a new message.
 */
export const UpdateSecretRequestSchema: GenMessage<UpdateSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 8);

/**
 * UpdateSecretResponse is the response of UpdateSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.UpdateSecretResponse
 */
export type UpdateSecretResponse = Message<"redpanda.api.dataplane.v1.UpdateSecretResponse"> & {
  /**
   * The updated secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.UpdateSecretResponse.
 * Use `create(UpdateSecretResponseSchema)` to create a new message.
 */
export const UpdateSecretResponseSchema: GenMessage<UpdateSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 9);

/**
 * DeleteSecretRequest is the request of DeleteSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.DeleteSecretRequest
 */
export type DeleteSecretRequest = Message<"redpanda.api.dataplane.v1.DeleteSecretRequest"> & {
  /**
   * The id of the secret to delete.
   *
   * @generated from field: string id = 1;
   */
  id: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.DeleteSecretRequest.
 * Use `create(DeleteSecretRequestSchema)` to create a new message.
 */
export const DeleteSecretRequestSchema: GenMessage<DeleteSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 10);

/**
 * DeleteSecretResponse is the response of DeleteSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.DeleteSecretResponse
 */
export type DeleteSecretResponse = Message<"redpanda.api.dataplane.v1.DeleteSecretResponse"> & {
};

/**
 * Describes the message redpanda.api.dataplane.v1.DeleteSecretResponse.
 * Use `create(DeleteSecretResponseSchema)` to create a new message.
 */
export const DeleteSecretResponseSchema: GenMessage<DeleteSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 11);

/**
 * ListSecretScopesRequest is the request of ListSecretScopes.
 *
 * @generated from message redpanda.api.dataplane.v1.ListSecretScopesRequest
 */
export type ListSecretScopesRequest = Message<"redpanda.api.dataplane.v1.ListSecretScopesRequest"> & {
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListSecretScopesRequest.
 * Use `create(ListSecretScopesRequestSchema)` to create a new message.
 */
export const ListSecretScopesRequestSchema: GenMessage<ListSecretScopesRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 12);

/**
 * ListSecretScopesResponse is the response of ListSecretScopes.
 *
 * @generated from message redpanda.api.dataplane.v1.ListSecretScopesResponse
 */
export type ListSecretScopesResponse = Message<"redpanda.api.dataplane.v1.ListSecretScopesResponse"> & {
  /**
   * @generated from field: repeated redpanda.api.dataplane.v1.Scope scopes = 1;
   */
  scopes: Scope[];
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListSecretScopesResponse.
 * Use `create(ListSecretScopesResponseSchema)` to create a new message.
 */
export const ListSecretScopesResponseSchema: GenMessage<ListSecretScopesResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 13);

/**
 * GetKafkaConnectSecretRequest is the request of GetKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.GetKafkaConnectSecretRequest
 */
export type GetKafkaConnectSecretRequest = Message<"redpanda.api.dataplane.v1.GetKafkaConnectSecretRequest"> & {
  /**
   * Unique name of the target connect cluster.
   *
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * The ID of the secret to retrieve.
   *
   * @generated from field: string id = 2;
   */
  id: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.GetKafkaConnectSecretRequest.
 * Use `create(GetKafkaConnectSecretRequestSchema)` to create a new message.
 */
export const GetKafkaConnectSecretRequestSchema: GenMessage<GetKafkaConnectSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 14);

/**
 * GetKafkaConnectSecretResponse is the response of GetKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.GetKafkaConnectSecretResponse
 */
export type GetKafkaConnectSecretResponse = Message<"redpanda.api.dataplane.v1.GetKafkaConnectSecretResponse"> & {
  /**
   * The retrieved secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.GetKafkaConnectSecretResponse.
 * Use `create(GetKafkaConnectSecretResponseSchema)` to create a new message.
 */
export const GetKafkaConnectSecretResponseSchema: GenMessage<GetKafkaConnectSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 15);

/**
 * CreateKafkaConnectSecretRequest is the request of CreateKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.CreateKafkaConnectSecretRequest
 */
export type CreateKafkaConnectSecretRequest = Message<"redpanda.api.dataplane.v1.CreateKafkaConnectSecretRequest"> & {
  /**
   * Unique name of the target connect cluster.
   *
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * Name of connector.
   *
   * @generated from field: string name = 2;
   */
  name: string;

  /**
   * Secret labels.
   *
   * @generated from field: map<string, string> labels = 3;
   */
  labels: { [key: string]: string };

  /**
   * The secret data. Must be Base64-encoded.
   *
   * @generated from field: bytes secret_data = 4;
   */
  secretData: Uint8Array;
};

/**
 * Describes the message redpanda.api.dataplane.v1.CreateKafkaConnectSecretRequest.
 * Use `create(CreateKafkaConnectSecretRequestSchema)` to create a new message.
 */
export const CreateKafkaConnectSecretRequestSchema: GenMessage<CreateKafkaConnectSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 16);

/**
 * CreateKafkaConnectSecretResponse is the response of CreateKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.CreateKafkaConnectSecretResponse
 */
export type CreateKafkaConnectSecretResponse = Message<"redpanda.api.dataplane.v1.CreateKafkaConnectSecretResponse"> & {
  /**
   * The created secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.CreateKafkaConnectSecretResponse.
 * Use `create(CreateKafkaConnectSecretResponseSchema)` to create a new message.
 */
export const CreateKafkaConnectSecretResponseSchema: GenMessage<CreateKafkaConnectSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 17);

/**
 * ListKafkaConnectSecretRequest is the request of ListKafkaConnectSecrets.
 *
 * @generated from message redpanda.api.dataplane.v1.ListKafkaConnectSecretsRequest
 */
export type ListKafkaConnectSecretsRequest = Message<"redpanda.api.dataplane.v1.ListKafkaConnectSecretsRequest"> & {
  /**
   * Unique name of the target connect cluster.
   *
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * List filter.
   *
   * @generated from field: redpanda.api.dataplane.v1.ListSecretsFilter filter = 2;
   */
  filter?: ListSecretsFilter;

  /**
   * Value of the next_page_token field returned by the previous response.
   * If not provided, the system assumes the first page is requested.
   *
   * @generated from field: string page_token = 3;
   */
  pageToken: string;

  /**
   * Limit the paginated response to a number of items.
   *
   * @generated from field: int32 page_size = 4;
   */
  pageSize: number;
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListKafkaConnectSecretsRequest.
 * Use `create(ListKafkaConnectSecretsRequestSchema)` to create a new message.
 */
export const ListKafkaConnectSecretsRequestSchema: GenMessage<ListKafkaConnectSecretsRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 18);

/**
 * ListKafkaConnectSecretsResponse is the response of ListKafkaConnectSecrets.
 *
 * @generated from message redpanda.api.dataplane.v1.ListKafkaConnectSecretsResponse
 */
export type ListKafkaConnectSecretsResponse = Message<"redpanda.api.dataplane.v1.ListKafkaConnectSecretsResponse"> & {
  /**
   * Secrets retrieved.
   *
   * @generated from field: repeated redpanda.api.dataplane.v1.Secret secrets = 1;
   */
  secrets: Secret[];

  /**
   * Token to retrieve the next page.
   *
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.ListKafkaConnectSecretsResponse.
 * Use `create(ListKafkaConnectSecretsResponseSchema)` to create a new message.
 */
export const ListKafkaConnectSecretsResponseSchema: GenMessage<ListKafkaConnectSecretsResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 19);

/**
 * UpdateKafkaConnectSecretRequest is the request of UpdateKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.UpdateKafkaConnectSecretRequest
 */
export type UpdateKafkaConnectSecretRequest = Message<"redpanda.api.dataplane.v1.UpdateKafkaConnectSecretRequest"> & {
  /**
   * Unique name of the target connect cluster.
   *
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * ID of the secret to update.
   *
   * @generated from field: string id = 2;
   */
  id: string;

  /**
   * Secret labels.
   *
   * @generated from field: map<string, string> labels = 3;
   */
  labels: { [key: string]: string };

  /**
   * The secret data. Must be Base64-encoded.
   *
   * @generated from field: bytes secret_data = 4;
   */
  secretData: Uint8Array;
};

/**
 * Describes the message redpanda.api.dataplane.v1.UpdateKafkaConnectSecretRequest.
 * Use `create(UpdateKafkaConnectSecretRequestSchema)` to create a new message.
 */
export const UpdateKafkaConnectSecretRequestSchema: GenMessage<UpdateKafkaConnectSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 20);

/**
 * UpdateKafkaConnectSecretResponse is the response of UpdateKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.UpdateKafkaConnectSecretResponse
 */
export type UpdateKafkaConnectSecretResponse = Message<"redpanda.api.dataplane.v1.UpdateKafkaConnectSecretResponse"> & {
  /**
   * The updated secret.
   *
   * @generated from field: redpanda.api.dataplane.v1.Secret secret = 1;
   */
  secret?: Secret;
};

/**
 * Describes the message redpanda.api.dataplane.v1.UpdateKafkaConnectSecretResponse.
 * Use `create(UpdateKafkaConnectSecretResponseSchema)` to create a new message.
 */
export const UpdateKafkaConnectSecretResponseSchema: GenMessage<UpdateKafkaConnectSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 21);

/**
 * DeleteKafkaConnectSecretRequest is the request of DeleteKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.DeleteKafkaConnectSecretRequest
 */
export type DeleteKafkaConnectSecretRequest = Message<"redpanda.api.dataplane.v1.DeleteKafkaConnectSecretRequest"> & {
  /**
   * Unique name of the target connect cluster.
   *
   * @generated from field: string cluster_name = 1;
   */
  clusterName: string;

  /**
   * ID of the secret to delete.
   *
   * @generated from field: string id = 2;
   */
  id: string;
};

/**
 * Describes the message redpanda.api.dataplane.v1.DeleteKafkaConnectSecretRequest.
 * Use `create(DeleteKafkaConnectSecretRequestSchema)` to create a new message.
 */
export const DeleteKafkaConnectSecretRequestSchema: GenMessage<DeleteKafkaConnectSecretRequest> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 22);

/**
 * DeleteKafkaConnectSecretResponse is the response of DeleteKafkaConnectSecret.
 *
 * @generated from message redpanda.api.dataplane.v1.DeleteKafkaConnectSecretResponse
 */
export type DeleteKafkaConnectSecretResponse = Message<"redpanda.api.dataplane.v1.DeleteKafkaConnectSecretResponse"> & {
};

/**
 * Describes the message redpanda.api.dataplane.v1.DeleteKafkaConnectSecretResponse.
 * Use `create(DeleteKafkaConnectSecretResponseSchema)` to create a new message.
 */
export const DeleteKafkaConnectSecretResponseSchema: GenMessage<DeleteKafkaConnectSecretResponse> = /*@__PURE__*/
  messageDesc(file_redpanda_api_dataplane_v1_secret, 23);

/**
 * Defines the scope of a secret.
 *
 * @generated from enum redpanda.api.dataplane.v1.Scope
 */
export enum Scope {
  /**
   * @generated from enum value: SCOPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SCOPE_REDPANDA_CONNECT = 1;
   */
  REDPANDA_CONNECT = 1,

  /**
   * @generated from enum value: SCOPE_REDPANDA_CLUSTER = 2;
   */
  REDPANDA_CLUSTER = 2,
}

/**
 * Describes the enum redpanda.api.dataplane.v1.Scope.
 */
export const ScopeSchema: GenEnum<Scope> = /*@__PURE__*/
  enumDesc(file_redpanda_api_dataplane_v1_secret, 0);

/**
 * @generated from service redpanda.api.dataplane.v1.SecretService
 */
export const SecretService: GenService<{
  /**
   * GetSecret retrieves the specific secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.GetSecret
   */
  getSecret: {
    methodKind: "unary";
    input: typeof GetSecretRequestSchema;
    output: typeof GetSecretResponseSchema;
  },
  /**
   * ListSecrets lists the secrets based on optional filter.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.ListSecrets
   */
  listSecrets: {
    methodKind: "unary";
    input: typeof ListSecretsRequestSchema;
    output: typeof ListSecretsResponseSchema;
  },
  /**
   * CreateSecret creates the secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.CreateSecret
   */
  createSecret: {
    methodKind: "unary";
    input: typeof CreateSecretRequestSchema;
    output: typeof CreateSecretResponseSchema;
  },
  /**
   * UpdateSecret updates the secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.UpdateSecret
   */
  updateSecret: {
    methodKind: "unary";
    input: typeof UpdateSecretRequestSchema;
    output: typeof UpdateSecretResponseSchema;
  },
  /**
   * DeleteSecret deletes the secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.DeleteSecret
   */
  deleteSecret: {
    methodKind: "unary";
    input: typeof DeleteSecretRequestSchema;
    output: typeof DeleteSecretResponseSchema;
  },
  /**
   * ListSecretScopes lists the supported secret scopes.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.ListSecretScopes
   */
  listSecretScopes: {
    methodKind: "unary";
    input: typeof ListSecretScopesRequestSchema;
    output: typeof ListSecretScopesResponseSchema;
  },
  /**
   * GetKafkaConnectSecret retrieves the specific secret for a specific Connect.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.GetKafkaConnectSecret
   */
  getKafkaConnectSecret: {
    methodKind: "unary";
    input: typeof GetKafkaConnectSecretRequestSchema;
    output: typeof GetKafkaConnectSecretResponseSchema;
  },
  /**
   * ListKafkaConnectSecrets lists the Connect secrets based on optional filter.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.ListKafkaConnectSecrets
   */
  listKafkaConnectSecrets: {
    methodKind: "unary";
    input: typeof ListKafkaConnectSecretsRequestSchema;
    output: typeof ListKafkaConnectSecretsResponseSchema;
  },
  /**
   * CreateKafkaConnectSecret creates the secret for a Connect.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.CreateKafkaConnectSecret
   */
  createKafkaConnectSecret: {
    methodKind: "unary";
    input: typeof CreateKafkaConnectSecretRequestSchema;
    output: typeof CreateKafkaConnectSecretResponseSchema;
  },
  /**
   * UpdateKafkaConnectSecret updates the Connect secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.UpdateKafkaConnectSecret
   */
  updateKafkaConnectSecret: {
    methodKind: "unary";
    input: typeof UpdateKafkaConnectSecretRequestSchema;
    output: typeof UpdateKafkaConnectSecretResponseSchema;
  },
  /**
   * DeleteKafkaConnectSecret deletes the secret.
   *
   * @generated from rpc redpanda.api.dataplane.v1.SecretService.DeleteKafkaConnectSecret
   */
  deleteKafkaConnectSecret: {
    methodKind: "unary";
    input: typeof DeleteKafkaConnectSecretRequestSchema;
    output: typeof DeleteKafkaConnectSecretResponseSchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_redpanda_api_dataplane_v1_secret, 0);

