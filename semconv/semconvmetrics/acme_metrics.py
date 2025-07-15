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

from opentelemetry.metrics import Meter
from opentelemetry.metrics import Counter

ACME_SALES_COUNT: Final = "acme.sales.count"
"""
Deprecated in favor of stable :py:const:`opentelemetry.semconv.metrics.acme_metrics.ACME_SALES_COUNT`.
"""


def create_acme_sales_count(meter: Meter) -> Counter:
    """Count of all sales"""
    return meter.create_counter(
        name=ACME_SALES_COUNT,
        description="Count of all sales",
        unit="{sale}",
    )
