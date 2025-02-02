// +-------------------------------------------------------------------------
// | Copyright (C) 2024 PEG Tech, Inc.
// +-------------------------------------------------------------------------
// | Licensed under the Apache License, Version 2.0 (the "License");
// | you may not use this work except in compliance with the License.
// | You may obtain a copy of the License in the LICENSE file, or at:
// |
// | http://www.apache.org/licenses/LICENSE-2.0
// |
// | Unless required by applicable law or agreed to in writing, software
// | distributed under the License is distributed on an "AS IS" BASIS,
// | WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// | See the License for the specific language governing permissions and
// | limitations under the License.
// +-------------------------------------------------------------------------

package errors

import (
	"fmt"
)

// PetaExpressError stores information of a PetaExpress error response.
type PetaExpressError struct {
	RetCode int    `json:"ret_code"`
	Message string `json:"message"`
}

// Error returns the description of PetaExpress error response.
func (ise PetaExpressError) Error() string {
	return fmt.Sprintf("PetaExpress Error: Code (%d), Message (%s)", ise.RetCode, ise.Message)
}
