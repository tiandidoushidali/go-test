// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// ExtendedExtendedStatsAggregation is a multi-value metrics aggregation that
// computes stats over numeric values extracted from the aggregated documents.
// These values can be extracted either from specific numeric fields
// in the documents, or be generated by a provided script.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/7.0/search-aggregations-metrics-extendedstats-aggregation.html
type ExtendedStatsAggregation struct {
	field           string
	script          *Script
	format          string
	missing         interface{}
	sigma           *float64
	subAggregations map[string]Aggregation
	meta            map[string]interface{}
}

func NewExtendedStatsAggregation() *ExtendedStatsAggregation {
	return &ExtendedStatsAggregation{
		subAggregations: make(map[string]Aggregation),
	}
}

func (a *ExtendedStatsAggregation) Field(field string) *ExtendedStatsAggregation {
	a.field = field
	return a
}

func (a *ExtendedStatsAggregation) Script(script *Script) *ExtendedStatsAggregation {
	a.script = script
	return a
}

func (a *ExtendedStatsAggregation) Format(format string) *ExtendedStatsAggregation {
	a.format = format
	return a
}

func (a *ExtendedStatsAggregation) Missing(missing interface{}) *ExtendedStatsAggregation {
	a.missing = missing
	return a
}

func (a *ExtendedStatsAggregation) Sigma(sigma float64) *ExtendedStatsAggregation {
	a.sigma = &sigma
	return a
}

func (a *ExtendedStatsAggregation) SubAggregation(name string, subAggregation Aggregation) *ExtendedStatsAggregation {
	a.subAggregations[name] = subAggregation
	return a
}

// Meta sets the meta data to be included in the aggregation response.
func (a *ExtendedStatsAggregation) Meta(metaData map[string]interface{}) *ExtendedStatsAggregation {
	a.meta = metaData
	return a
}

func (a *ExtendedStatsAggregation) Source() (interface{}, error) {
	// Example:
	//	{
	//    "aggs" : {
	//      "grades_stats" : { "extended_stats" : { "field" : "grade" } }
	//    }
	//	}
	// This method returns only the { "extended_stats" : { "field" : "grade" } } part.

	source := make(map[string]interface{})
	opts := make(map[string]interface{})
	source["extended_stats"] = opts

	// ValuesSourceAggregationBuilder
	if a.field != "" {
		opts["field"] = a.field
	}
	if a.script != nil {
		src, err := a.script.Source()
		if err != nil {
			return nil, err
		}
		opts["script"] = src
	}
	if a.format != "" {
		opts["format"] = a.format
	}
	if v := a.missing; v != nil {
		opts["missing"] = v
	}
	if v := a.sigma; v != nil {
		opts["sigma"] = *v
	}

	// AggregationBuilder (SubAggregations)
	if len(a.subAggregations) > 0 {
		aggsMap := make(map[string]interface{})
		source["aggregations"] = aggsMap
		for name, aggregate := range a.subAggregations {
			src, err := aggregate.Source()
			if err != nil {
				return nil, err
			}
			aggsMap[name] = src
		}
	}

	// Add Meta data if available
	if len(a.meta) > 0 {
		source["meta"] = a.meta
	}

	return source, nil
}
