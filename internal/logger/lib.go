/**
# Copyright (c) NVIDIA CORPORATION.  All rights reserved.
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
**/

package logger

import "github.com/sirupsen/logrus"

// New returns a new logger
func New() Interface {
	return logrus.StandardLogger()
}

// NullLogger is a logger that does nothing
type NullLogger struct{}

var _ Interface = (*NullLogger)(nil)

// Debugf is a no-op for the null logger
func (l *NullLogger) Debugf(string, ...interface{}) {}

// Errorf is a no-op for the null logger
func (l *NullLogger) Errorf(string, ...interface{}) {}

// Info is a no-op for the null logger
func (l *NullLogger) Info(...interface{}) {}

// Infof is a no-op for the null logger
func (l *NullLogger) Infof(string, ...interface{}) {}

// Warn is a no-op for the null logger
func (l *NullLogger) Warn(...interface{}) {}

// Warnf is a no-op for the null logger
func (l *NullLogger) Warnf(string, ...interface{}) {}

// Warningf is a no-op for the null logger
func (l *NullLogger) Warningf(string, ...interface{}) {}
