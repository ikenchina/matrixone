// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fz

import (
	"constraints"
	"fmt"
	"math/rand"
)

func RandBetween[T constraints.Integer, T2 constraints.Integer](target *T, lower T2, upper T2) {
	if lower == upper {
		*target = T(lower)
		return
	}
	if lower > upper {
		panic(fmt.Errorf("bad argument: %d %d", lower, upper))
	}
	gap := upper - lower
	i := rand.Intn(int(gap + 1))
	*target = T(lower) + T(i)
	return
}

func RandBool[T ~bool](target *T) {
	*target = rand.Intn(2) == 0
}
