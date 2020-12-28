// BSD 3-Clause License
//
// Copyright (c) 2020, Kingsgroup
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package werr

import (
	"errors"
	"fmt"
	"testing"
)

func TestRetriable_Unwrap(t *testing.T) {
	err1 := errors.New("err1")
	err2 := NewRetriable(err1)
	var err3 error = err2
	if errors.Unwrap(err3) != err1 {
		t.Fatal("errors.Unwrap(err3) != err1")
	}
	if !errors.Is(err3, err1) {
		t.Fatal("!errors.Is(err3, err1)")
	}

	err4 := fmt.Errorf("err4: %w", err3)
	if !errors.Is(err4, err1) {
		t.Fatal("!errors.Is(err4, err1)")
	}

	if err5, ok := AsRetriable(err4); !ok || err5 != err2 {
		t.Fatal("!ok || err5 != err2")
	}
	if err6, ok := AsRetriable(err3); !ok || err6 != err2 {
		t.Fatal("!ok || err6 != err2")
	}
}
