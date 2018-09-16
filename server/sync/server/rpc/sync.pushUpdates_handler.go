/*
 *  Copyright (c) 2018, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpc

import (
    "github.com/golang/glog"
    "github.com/nebulaim/telegramd/proto/mtproto"
    "golang.org/x/net/context"
    "github.com/nebulaim/telegramd/baselib/logger"
)

// sync.pushUpdates#5c612649 user_id:int updates:Updates = Bool;
func (s *SyncServiceImpl) SyncPushUpdates(ctx context.Context, request *mtproto.TLSyncPushUpdates) (*mtproto.Bool, error) {
    glog.Infof("sync.pushUpdates#5c612649 - request: {%s}", logger.JsonDebugData(request))
    pushData := &mtproto.PushData{
        Constructor: mtproto.TLConstructor_CRC32_sync_pushUpdatesData,
        Data2:       &mtproto.PushData_Data{Updates: request.GetUpdates()},
    }

    err := s.processUpdatesRequest(request.GetUserId(), request.GetUpdates())
    if err == nil {
        s.pushUpdatesToSession(syncTypeUser, request.GetUserId(), pushData, 0)
    } else {
        glog.Error(err)
        return mtproto.ToBool(false), nil
    }
    return mtproto.ToBool(true), nil
}
