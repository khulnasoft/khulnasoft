# Copyright 2016-2021, Pulumi Corporation.
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


import khulnasoft
import pytest
from unittest.mock import patch

import khulnasoft_plant


@pytest.fixture
def my_mocks():
    old_settings = khulnasoft.runtime.settings.SETTINGS
    try:
        mocks = MyMocks()
        khulnasoft.runtime.mocks.set_mocks(mocks)
        yield mocks
    finally:
        khulnasoft.runtime.settings.configure(old_settings)


class MyMocks(khulnasoft.runtime.Mocks):
    def call(self, args):
        return {}
    def new_resource(self, args):
        return 'foo', args.inputs


@khulnasoft.runtime.test
def test_default_value_does_not_trigger_deprecation_warning(my_mocks):
    """
    Constructs a resource with deprecated inputs with a default value
    and checks that the supplied default values don't trigger a deprecation warning.
    """
    with patch("warnings.warn") as mock_warn:
        khulnasoft_plant.tree.v1.RubberTree("my-tree", khulnasoft_plant.tree.v1.RubberTreeArgs())
        mock_warn.assert_not_called()
