// @generated by protoc-gen-connect-query v1.4.0 with parameter "target=ts,import_extension=,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/dataplane/v1/cloud_storage.proto (package redpanda.api.dataplane.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { DeleteMountTaskRequest, DeleteMountTaskResponse, GetMountTaskRequest, GetMountTaskResponse, ListMountableTopicsRequest, ListMountableTopicsResponse, ListMountTasksRequest, ListMountTasksResponse, MountTopicsRequest, MountTopicsResponse, UnmountTopicsRequest, UnmountTopicsResponse, UpdateMountTaskRequest, UpdateMountTaskResponse } from "./cloud_storage_pb";

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.MountTopics
 */
export const mountTopics = {
  localName: "mountTopics",
  name: "MountTopics",
  kind: MethodKind.Unary,
  I: MountTopicsRequest,
  O: MountTopicsResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.UnmountTopics
 */
export const unmountTopics = {
  localName: "unmountTopics",
  name: "UnmountTopics",
  kind: MethodKind.Unary,
  I: UnmountTopicsRequest,
  O: UnmountTopicsResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.ListMountableTopics
 */
export const listMountableTopics = {
  localName: "listMountableTopics",
  name: "ListMountableTopics",
  kind: MethodKind.Unary,
  I: ListMountableTopicsRequest,
  O: ListMountableTopicsResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.ListMountTasks
 */
export const listMountTasks = {
  localName: "listMountTasks",
  name: "ListMountTasks",
  kind: MethodKind.Unary,
  I: ListMountTasksRequest,
  O: ListMountTasksResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.GetMountTask
 */
export const getMountTask = {
  localName: "getMountTask",
  name: "GetMountTask",
  kind: MethodKind.Unary,
  I: GetMountTaskRequest,
  O: GetMountTaskResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.DeleteMountTask
 */
export const deleteMountTask = {
  localName: "deleteMountTask",
  name: "DeleteMountTask",
  kind: MethodKind.Unary,
  I: DeleteMountTaskRequest,
  O: DeleteMountTaskResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;

/**
 * @generated from rpc redpanda.api.dataplane.v1.CloudStorageService.UpdateMountTask
 */
export const updateMountTask = {
  localName: "updateMountTask",
  name: "UpdateMountTask",
  kind: MethodKind.Unary,
  I: UpdateMountTaskRequest,
  O: UpdateMountTaskResponse,
  service: {
    typeName: "redpanda.api.dataplane.v1.CloudStorageService"
  }
} as const;
