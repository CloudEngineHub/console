// @generated by protoc-gen-connect-query v2.0.1 with parameter "target=ts,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/console/v1alpha1/transform.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */

import { TransformService } from "./transform_pb";

/**
 * @generated from rpc redpanda.api.console.v1alpha1.TransformService.ListTransforms
 */
export const listTransforms = TransformService.method.listTransforms;

/**
 * @generated from rpc redpanda.api.console.v1alpha1.TransformService.GetTransform
 */
export const getTransform = TransformService.method.getTransform;

/**
 * @generated from rpc redpanda.api.console.v1alpha1.TransformService.DeleteTransform
 */
export const deleteTransform = TransformService.method.deleteTransform;
