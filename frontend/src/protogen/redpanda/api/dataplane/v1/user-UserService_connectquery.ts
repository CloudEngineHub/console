// @generated by protoc-gen-connect-query v2.0.1 with parameter "target=ts,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/dataplane/v1/user.proto (package redpanda.api.dataplane.v1, syntax proto3)
/* eslint-disable */

import { UserService } from "./user_pb";

/**
 * @generated from rpc redpanda.api.dataplane.v1.UserService.CreateUser
 */
export const createUser = UserService.method.createUser;

/**
 * @generated from rpc redpanda.api.dataplane.v1.UserService.UpdateUser
 */
export const updateUser = UserService.method.updateUser;

/**
 * @generated from rpc redpanda.api.dataplane.v1.UserService.ListUsers
 */
export const listUsers = UserService.method.listUsers;

/**
 * @generated from rpc redpanda.api.dataplane.v1.UserService.DeleteUser
 */
export const deleteUser = UserService.method.deleteUser;
