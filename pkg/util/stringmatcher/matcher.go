// Copyright 2022 The jackal Authors
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

package stringmatcher

// Matcher interface is used check whether a string matches some pattern.
type Matcher interface {
	Matches(str string) bool
}

type anyMatcher struct{}

func (m *anyMatcher) Matches(_ string) bool { return true }

// Any is a Matcher implementation that matches all strings.
var Any Matcher = &anyMatcher{}
