/*
 * Copyright The OpenTelemetry Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//----------------------------------------------------------------------------------------------------------
// DO NOT EDIT, this is an Auto-generated file from scripts/semconv/templates/registry/stable/attributes.ts.j2
//----------------------------------------------------------------------------------------------------------

/**
 * Unique identifier of an Acme corp product.
 *
 * @example my_product_id
 * @example 1234
 */
export const ATTR_ACME_PRODUCT_ID = 'acme.product.id' as const;

/**
 * Type of an Acme corp product.
 *
 * @example helicopter
 */
export const ATTR_ACME_PRODUCT_TYPE = 'acme.product.type' as const;

/**
 * Enum value "car" for attribute {@link ATTR_ACME_PRODUCT_TYPE}.
 *
 * A car.
 */
export const ACME_PRODUCT_TYPE_VALUE_CAR = "car" as const;

/**
 * Enum value "helicopter" for attribute {@link ATTR_ACME_PRODUCT_TYPE}.
 *
 * A helicopter.
 */
export const ACME_PRODUCT_TYPE_VALUE_HELICOPTER = "helicopter" as const;

/**
 * Unique identifier of the OAuth 2.0 client.
 *
 * @example example_client_id
 */
export const ATTR_OAUTH2_CLIENT_ID = 'oauth2.client.id' as const;

