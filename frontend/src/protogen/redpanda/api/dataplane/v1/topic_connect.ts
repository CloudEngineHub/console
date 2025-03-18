// @generated by protoc-gen-connect-es v1.2.0 with parameter "target=ts,import_extension="
// @generated from file redpanda/api/dataplane/v1/topic.proto (package redpanda.api.dataplane.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { AddPartitionsToTopicsRequest, AddPartitionsToTopicsResponse, AddTopicPartitionsRequest, AddTopicPartitionsResponse, CreateTopicRequest, CreateTopicResponse, DeleteTopicRequest, DeleteTopicResponse, GetTopicConfigurationsRequest, GetTopicConfigurationsResponse, ListTopicsRequest, ListTopicsResponse, SetPartitionsToTopicsRequest, SetPartitionsToTopicsResponse, SetTopicConfigurationsRequest, SetTopicConfigurationsResponse, SetTopicPartitionsRequest, SetTopicPartitionsResponse, UpdateTopicConfigurationsRequest, UpdateTopicConfigurationsResponse } from "./topic_pb";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service redpanda.api.dataplane.v1.TopicService
 */
export const TopicService = {
  typeName: "redpanda.api.dataplane.v1.TopicService",
  methods: {
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.CreateTopic
     */
    createTopic: {
      name: "CreateTopic",
      I: CreateTopicRequest,
      O: CreateTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.ListTopics
     */
    listTopics: {
      name: "ListTopics",
      I: ListTopicsRequest,
      O: ListTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.DeleteTopic
     */
    deleteTopic: {
      name: "DeleteTopic",
      I: DeleteTopicRequest,
      O: DeleteTopicResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.GetTopicConfigurations
     */
    getTopicConfigurations: {
      name: "GetTopicConfigurations",
      I: GetTopicConfigurationsRequest,
      O: GetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.UpdateTopicConfigurations
     */
    updateTopicConfigurations: {
      name: "UpdateTopicConfigurations",
      I: UpdateTopicConfigurationsRequest,
      O: UpdateTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.SetTopicConfigurations
     */
    setTopicConfigurations: {
      name: "SetTopicConfigurations",
      I: SetTopicConfigurationsRequest,
      O: SetTopicConfigurationsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.AddTopicPartitions
     */
    addTopicPartitions: {
      name: "AddTopicPartitions",
      I: AddTopicPartitionsRequest,
      O: AddTopicPartitionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.SetTopicPartitions
     */
    setTopicPartitions: {
      name: "SetTopicPartitions",
      I: SetTopicPartitionsRequest,
      O: SetTopicPartitionsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.AddPartitionsToTopics
     */
    addPartitionsToTopics: {
      name: "AddPartitionsToTopics",
      I: AddPartitionsToTopicsRequest,
      O: AddPartitionsToTopicsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc redpanda.api.dataplane.v1.TopicService.SetPartitionsToTopics
     */
    setPartitionsToTopics: {
      name: "SetPartitionsToTopics",
      I: SetPartitionsToTopicsRequest,
      O: SetPartitionsToTopicsResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

