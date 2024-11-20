// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import * as khulnasoft from "@khulnasoft/khulnasoft";

export class MyRandom extends khulnasoft.ComponentResource {
  public readonly randomID: khulnasoft.Output<string>;

  constructor(name: string, opts: khulnasoft.ResourceOptions) {
    super("pkg:index:MyRandom", name, {}, opts);
    this.randomID = khulnasoft.output(`${name}-${Math.floor(Math.random() * 1000)}`);
    this.registerOutputs({ randomID: this.randomID });
  }
}