# Copyright The OpenTelemetry Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from typing import Final


from enum import Enum
from typing_extensions import deprecated

ACME_PRODUCT_ID: Final = "acme.product.id"
"""
Deprecated in favor of stable :py:const:`opentelemetry.semconv.attributes.acme_attributes.ACME_PRODUCT_ID`.
"""

ACME_PRODUCT_TYPE: Final = "acme.product.type"
"""
Deprecated in favor of stable :py:const:`opentelemetry.semconv.attributes.acme_attributes.ACME_PRODUCT_TYPE`.
"""


@deprecated(
    "Deprecated in favor of stable :py:const:`opentelemetry.semconv.attributes.acme_attributes.AcmeProductTypeValues`."
)
class AcmeProductTypeValues(Enum):
    HELICOPTER = "helicopter"
    """Deprecated in favor of stable :py:const:`opentelemetry.semconv.attributes.acme_attributes.AcmeProductTypeValues.HELICOPTER`."""
    CAR = "car"
    """Deprecated in favor of stable :py:const:`opentelemetry.semconv.attributes.acme_attributes.AcmeProductTypeValues.CAR`."""
