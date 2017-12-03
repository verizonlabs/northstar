/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package model

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	NotSet    int = 0
	Private   int = 1
	Published int = 2
)

// Defines the type that represents a template.
type Template struct {
	Id          string    `json:"id,omitempty"`
	CreatedOn   time.Time `json:"createdOn,omitempty"`
	Version     string    `json:"version,omitempty"`
	AccountId   string    `json:"accountId,omitempty"`
	UserId      string    `json:"userId,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Type        string    `json:"type"`
	Data        string    `json:"data"`
	Published   int       `json:"published"`
	Hash        string    `json:"hash"`
}

// Defines internal type for marshaling/unmarshaling.
type typeTemplate Template

// Helper method used to unmarshal notebook while validating required fields.
func (template *Template) UnmarshalJSON(data []byte) error {
	var value typeTemplate

	// Unmarshal to the internal type
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	*template = Template(value)

	// Validate template name.
	if template.Name == "" {
		return fmt.Errorf("The template type is missing.")
	}

	// Validate template type.
	if template.Type == "" {
		return fmt.Errorf("The template type is missing.")
	}

	// Validate template data.
	if template.Data == "" {
		return fmt.Errorf("The template data is missing.")
	}

	return nil
}

// Defines the type used to query templates.
type Query struct {
	AccountId string `json:"accountId,omitempty"`
	UserId    string `json:"userId,omitempty"`
	Type      string `json:"type,omitempty"`
	Published int    `json:"published,omitempty"`
	Hash      string `json:"hash,omitempty"`
}
