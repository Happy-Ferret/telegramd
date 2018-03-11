/*
 *  Copyright (c) 2017, https://github.com/nebulaim
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
	"fmt"
	"github.com/golang/glog"
	"github.com/nebulaim/telegramd/baselib/logger"
	"github.com/nebulaim/telegramd/grpc_util"
	"github.com/nebulaim/telegramd/mtproto"
	"golang.org/x/net/context"
)

// upload.reuploadCdnFile#1af91c09 file_token:bytes request_token:bytes = Vector<CdnFileHash>;
func (s *UploadServiceImpl) UploadReuploadCdnFile(ctx context.Context, request *mtproto.TLUploadReuploadCdnFile) (*mtproto.Vector_CdnFileHash, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	glog.Infof("UploadReuploadCdnFile - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	// TODO(@benqi): Impl UploadReuploadCdnFile logic

	return nil, fmt.Errorf("Not impl UploadReuploadCdnFile")
}
