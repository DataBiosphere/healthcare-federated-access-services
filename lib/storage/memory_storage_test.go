// Copyright 2019 Google LLC
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

package storage

import (
	"testing"

	"github.com/golang/protobuf/proto" /* copybara-comment */
	cpb "github.com/GoogleCloudPlatform/healthcare-federated-access-services/proto/common/v1" /* copybara-comment: go_proto */
)

const (
	testStoreFileType = "store"
	testFileID        = "main"
)

func TestMemoryStorageDelete(t *testing.T) {
	store := NewMemoryStorage("storage", "testdata")
	content := &cpb.TestPersona{}
	if err := store.Read(testStoreFileType, DefaultRealm, DefaultUser, testFileID, LatestRev, content); err != nil {
		t.Errorf("reading default file: want success, got error: %v", err)
	}
	if err := store.Delete(testStoreFileType, DefaultRealm, DefaultUser, testFileID, LatestRev); err != nil {
		t.Errorf("deleting default file: want success, got error: %v", err)
	}
	if err := store.Read(testStoreFileType, DefaultRealm, DefaultUser, testFileID, LatestRev, content); err == nil {
		t.Errorf("reading deleted file: want error, got success")
	}
}

func TestMemoryStorageMultiRead(t *testing.T) {
	store := NewMemoryStorage("ic-min", "testdata/config")
	// content is map[<user>]map[<key>]*cpb.Account.
	content := make(map[string]map[string]proto.Message)
	count, err := store.MultiReadTx(AccountDatatype, "test", DefaultUser, nil, 0, 100, content, &cpb.Account{}, nil)
	if err != nil {
		t.Fatalf("MultiReadTx() failed: %v", err)
	}
	want := 4
	if count != want {
		t.Errorf("MultiReadTx() count results mismatch: got %d, want %d", count, want)
	}
	got := 0
	for user, ucontent := range content {
		for key, acct := range ucontent {
			if acct == nil {
				t.Fatalf("MultiReadTx() invalid results for user %q: key %q content is nil", user, key)
			}
			if _, ok := acct.(*cpb.Account); !ok {
				t.Fatalf("MultiReadTx() invalid results for user %q: key %q content is not an account", user, key)
			}
			got++
		}
	}
	if got != want {
		t.Errorf("MultiReadTx() count results mismatch with count %d: got %d, want %d", count, got, want)
	}
}
