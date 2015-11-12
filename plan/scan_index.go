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

	"github.com/couchbase/query/algebra"
	"github.com/couchbase/query/datastore"
	"github.com/couchbase/query/expression"
	"github.com/couchbase/query/expression/parser"
)

type IndexScan struct {
	readonly
	index    datastore.Index
	term     *algebra.KeyspaceTerm
	spans    Spans
	distinct bool
	limit    expression.Expression
	covers   expression.Covers
}

func NewIndexScan(index datastore.Index, term *algebra.KeyspaceTerm, spans Spans,
	distinct bool, limit expression.Expression, covers expression.Covers) *IndexScan {
	return &IndexScan{
		index:    index,
		term:     term,
		spans:    spans,
		distinct: distinct,
		limit:    limit,
		covers:   covers,
	}
}

func (this *IndexScan) Accept(visitor Visitor) (interface{}, error) {
	return visitor.VisitIndexScan(this)
}

func (this *IndexScan) New() Operator {
	return &IndexScan{}
}

func (this *IndexScan) Index() datastore.Index {
	return this.index
}

func (this *IndexScan) Term() *algebra.KeyspaceTerm {
	return this.term
}

func (this *IndexScan) Spans() Spans {
	return this.spans
}

func (this *IndexScan) Distinct() bool {
	return this.distinct
}

func (this *IndexScan) Limit() expression.Expression {
	return this.limit
}

func (this *IndexScan) Covers() expression.Covers {
	return this.covers
}

func (this *IndexScan) Covering() bool {
	return len(this.covers) > 0
}

func (this *IndexScan) MarshalJSON() ([]byte, error) {
	r := map[string]interface{}{"#operator": "IndexScan"}
	r["index"] = this.index.Name()
	r["namespace"] = this.term.Namespace()
	r["keyspace"] = this.term.Keyspace()
	r["using"] = this.index.Type()
	r["spans"] = this.spans

	if this.distinct {
		r["distinct"] = this.distinct
	}

	if this.limit != nil {
		r["limit"] = expression.NewStringer().Visit(this.limit)
	}

	if this.covers != nil {
		r["covers"] = this.covers
	}

	return json.Marshal(r)
}

func (this *IndexScan) UnmarshalJSON(body []byte) error {
	var _unmarshalled struct {
		_         string              `json:"#operator"`
		Index     string              `json:"index"`
		Namespace string              `json:"namespace"`
		Keyspace  string              `json:"keyspace"`
		Using     datastore.IndexType `json:"using"`
		Spans     Spans               `json:"spans"`
		Distinct  bool                `json:"distinct"`
		Limit     string              `json:"limit"`
		Covers    []string            `json:"covers"`
	}

	err := json.Unmarshal(body, &_unmarshalled)
	if err != nil {
		return err
	}

	k, err := datastore.GetKeyspace(_unmarshalled.Namespace, _unmarshalled.Keyspace)
	if err != nil {
		return err
	}

	this.term = algebra.NewKeyspaceTerm(
		_unmarshalled.Namespace, _unmarshalled.Keyspace,
		nil, "", nil, nil)

	this.spans = _unmarshalled.Spans
	this.distinct = _unmarshalled.Distinct

	if _unmarshalled.Limit != "" {
		this.limit, err = parser.Parse(_unmarshalled.Limit)
		if err != nil {
			return err
		}
	}

	if _unmarshalled.Covers != nil {
		this.covers = make(expression.Covers, len(_unmarshalled.Covers))
		for i, c := range _unmarshalled.Covers {
			expr, err := parser.Parse(c)
			if err != nil {
				return err
			}

			this.covers[i] = expression.NewCover(expr)
		}
	}

	indexer, err := k.Indexer(_unmarshalled.Using)
	if err != nil {
		return err
	}

	this.index, err = indexer.IndexByName(_unmarshalled.Index)
	return err
}