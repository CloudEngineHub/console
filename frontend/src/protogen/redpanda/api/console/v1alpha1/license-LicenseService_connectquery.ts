// @generated by protoc-gen-connect-query v2.0.1 with parameter "target=ts,js_import_style=legacy_commonjs"
// @generated from file redpanda/api/console/v1alpha1/license.proto (package redpanda.api.console.v1alpha1, syntax proto3)
/* eslint-disable */

import { LicenseService } from "./license_pb";

/**
 * ListLicenses lists all the roles based on optional filter.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.ListLicenses
 */
export const listLicenses = LicenseService.method.listLicenses;

/**
 * SetLicense installs a new license on the Redpanda cluster.
 * This endpoint only works if the Redpanda Admin API is configured.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.SetLicense
 */
export const setLicense = LicenseService.method.setLicense;

/**
 * ListEnterpriseFeatures reports the license status and Redpanda enterprise features in use.
 * This can only be reported if the Redpanda Admin API is configured and supports this request.
 *
 * @generated from rpc redpanda.api.console.v1alpha1.LicenseService.ListEnterpriseFeatures
 */
export const listEnterpriseFeatures = LicenseService.method.listEnterpriseFeatures;
