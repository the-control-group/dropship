// Copyright (c) 2016 "ChrisMcKenzie"
// This file is part of Dropship.
//
// Dropship is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License v3 as
// published by the Free Software Foundation
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
package manager

import (
	"net/url"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/libkv/store"
)

var kvstore store.Store

func Start(storePath string) {
	storeUrl, err := url.Parse(storePath)
	if err != nil {
		log.Fatal(err)
		return
	}

	kvstore, err = initStore(storeUrl)
	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		log.Fatal(serveRpc(3000))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		log.Fatal(serveInterface(3001))
		wg.Done()
	}()
	wg.Wait()
}
