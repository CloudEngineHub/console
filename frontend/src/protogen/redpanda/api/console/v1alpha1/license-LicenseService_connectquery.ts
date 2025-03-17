// @generated by protoc-gen-connect-query v1.4.0 with parameter "target=ts,import_extension=,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/console/v1alpha1/license.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { ListEnterpriseFeaturesRequest, ListEnterpriseFeaturesResponse, ListLicensesRequest, ListLicensesResponse, SetLicenseRequest, SetLicenseResponse } from "./license_pb";

/**
 * ListLicenses lists all the roles based on optional filter.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.ListLicenses
 */
export const listLicenses = {
  localName: "listLicenses",
  name: "ListLicenses",
  kind: MethodKind.Unary,
  I: ListLicensesRequest,
  O: ListLicensesResponse,
  service: {
    typeName: "redpanda.api.console.v1alpha1.LicenseService"
  }
} as const;

/**
 * SetLicense installs a new license on the Redpanda cluster.
 * This endpoint only works if the Redpanda Admin API is configured.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.SetLicense
 */
export const setLicense = {
  localName: "setLicense",
  name: "SetLicense",
  kind: MethodKind.Unary,
  I: SetLicenseRequest,
  O: SetLicenseResponse,
  service: {
    typeName: "redpanda.api.console.v1alpha1.LicenseService"
  }
} as const;

/**
 * ListEnterpriseFeatures reports the license status and Redpanda enterprise features in use.
 * This can only be reported if the Redpanda Admin API is configured and supports this request.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.ListEnterpriseFeatures
 */
export const listEnterpriseFeatures = {
  localName: "listEnterpriseFeatures",
  name: "ListEnterpriseFeatures",
  kind: MethodKind.Unary,
  I: ListEnterpriseFeaturesRequest,
  O: ListEnterpriseFeaturesResponse,
  service: {
    typeName: "redpanda.api.console.v1alpha1.LicenseService"
  }
} as const;
