//  Copyright (c) 2014 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package plan

import (
	"encoding/json"
)

type IntersectAll struct {
	readonly
	first  Operator
	second Operator
}

func NewIntersectAll(first, second Operator) *IntersectAll {
	return &IntersectAll{
		first:  first,
		second: second,
	}
}

func (this *IntersectAll) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitIntersectAll(this)
}

func (this *IntersectAll) New() Operator {
	return &IntersectAll{}
}

func (this *IntersectAll) First() Operator {
	return this.first
}

func (this *IntersectAll) Second() Operator {
	return this.second
}

func (this *IntersectAll) MarshalJSON() ([]byte, error) {
	r := map[string]interface{}{"#operator": "IntersectAll"}
	r["first"] = this.first
	r["second"] = this.second
	return json.Marshal(r)
}
