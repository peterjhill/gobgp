// Copyright (C) 2014, 2015 Nippon Telegraph and Telephone Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zebra

import (
	"fmt"
)

// For Cumulus Linux 3.7.7, zebra 4.0+cl3u13  (ZAPI version 5)
const (
	FRR_ZAPI5_CL_INTERFACE_ADD API_TYPE = iota
	FRR_ZAPI5_CL_INTERFACE_DELETE
	FRR_ZAPI5_CL_INTERFACE_ADDRESS_ADD
	FRR_ZAPI5_CL_INTERFACE_ADDRESS_DELETE
	FRR_ZAPI5_CL_INTERFACE_UP
	FRR_ZAPI5_CL_INTERFACE_DOWN
	FRR_ZAPI5_CL_INTERFACE_SET_MASTER
	FRR_ZAPI5_CL_INTERFACE_SET_PROTODOWN
	FRR_ZAPI5_CL_ROUTE_ADD
	FRR_ZAPI5_CL_ROUTE_DELETE
	FRR_ZAPI5_CL_ROUTE_NOTIFY_OWNER
	FRR_ZAPI5_CL_REDISTRIBUTE_ADD
	FRR_ZAPI5_CL_REDISTRIBUTE_DELETE
	FRR_ZAPI5_CL_REDISTRIBUTE_DEFAULT_ADD
	FRR_ZAPI5_CL_REDISTRIBUTE_DEFAULT_DELETE
	FRR_ZAPI5_CL_ROUTER_ID_ADD
	FRR_ZAPI5_CL_ROUTER_ID_DELETE
	FRR_ZAPI5_CL_ROUTER_ID_UPDATE
	FRR_ZAPI5_CL_HELLO
	FRR_ZAPI5_CL_NEXTHOP_REGISTER
	FRR_ZAPI5_CL_NEXTHOP_UNREGISTER
	FRR_ZAPI5_CL_NEXTHOP_UPDATE
	FRR_ZAPI5_CL_INTERFACE_NBR_ADDRESS_ADD
	FRR_ZAPI5_CL_INTERFACE_NBR_ADDRESS_DELETE
	FRR_ZAPI5_CL_INTERFACE_BFD_DEST_UPDATE
	FRR_ZAPI5_CL_IMPORT_ROUTE_REGISTER
	FRR_ZAPI5_CL_IMPORT_ROUTE_UNREGISTER
	FRR_ZAPI5_CL_IMPORT_CHECK_UPDATE
	FRR_ZAPI5_CL_BFD_DEST_REGISTER
	FRR_ZAPI5_CL_BFD_DEST_DEREGISTER
	FRR_ZAPI5_CL_BFD_DEST_UPDATE
	FRR_ZAPI5_CL_BFD_DEST_REPLAY
	FRR_ZAPI5_CL_REDISTRIBUTE_ROUTE_ADD
	FRR_ZAPI5_CL_REDISTRIBUTE_ROUTE_DEL
	FRR_ZAPI5_CL_VRF_UNREGISTER
	FRR_ZAPI5_CL_VRF_ADD
	FRR_ZAPI5_CL_VRF_DELETE
	FRR_ZAPI5_CL_VRF_LABEL
	FRR_ZAPI5_CL_INTERFACE_VRF_UPDATE
	FRR_ZAPI5_CL_BFD_CLIENT_REGISTER
	FRR_ZAPI5_CL_INTERFACE_ENABLE_RADV
	FRR_ZAPI5_CL_INTERFACE_DISABLE_RADV
	FRR_ZAPI5_CL_IPV4_NEXTHOP_LOOKUP_MRIB
	FRR_ZAPI5_CL_INTERFACE_LINK_PARAMS
	FRR_ZAPI5_CL_MPLS_LABELS_ADD
	FRR_ZAPI5_CL_MPLS_LABELS_DELETE
	FRR_ZAPI5_CL_IPMR_ROUTE_STATS
	FRR_ZAPI5_CL_LABEL_MANAGER_CONNECT
	FRR_ZAPI5_CL_GET_LABEL_CHUNK
	FRR_ZAPI5_CL_RELEASE_LABEL_CHUNK
	FRR_ZAPI5_CL_FEC_REGISTER
	FRR_ZAPI5_CL_FEC_UNREGISTER
	FRR_ZAPI5_CL_FEC_UPDATE
	FRR_ZAPI5_CL_ADVERTISE_SVI_MACSTUFF
	FRR_ZAPI5_CL_ADVERTISE_DEFAULT_GW
	FRR_ZAPI5_CL_ADVERTISE_SUBNET
	FRR_ZAPI5_CL_ADVERTISE_ALL_VNI
	FRR_ZAPI5_CL_VNI_ADD
	FRR_ZAPI5_CL_VNI_DEL
	FRR_ZAPI5_CL_L3VNI_ADD
	FRR_ZAPI5_CL_L3VNI_DEL
	FRR_ZAPI5_CL_REMOTE_VTEP_ADD
	FRR_ZAPI5_CL_REMOTE_VTEP_DEL
	FRR_ZAPI5_CL_MACIP_ADD
	FRR_ZAPI5_CL_MACIP_DEL
	FRR_ZAPI5_CL_DAD_STUFF
	FRR_ZAPI5_CL_IP_PREFIX_ROUTE_ADD
	FRR_ZAPI5_CL_IP_PREFIX_ROUTE_DEL
	FRR_ZAPI5_CL_REMOTE_MACIP_ADD
	FRR_ZAPI5_CL_REMOTE_MACIP_DEL
	FRR_ZAPI5_CL_PW_ADD
	FRR_ZAPI5_CL_PW_DELETE
	FRR_ZAPI5_CL_PW_SET
	FRR_ZAPI5_CL_PW_UNSET
	FRR_ZAPI5_CL_PW_STATUS_UPDATE
	FRR_ZAPI5_CL_RULE_ADD
	FRR_ZAPI5_CL_RULE_DELETE
	FRR_ZAPI5_CL_RULE_NOTIFY_OWNER
	FRR_ZAPI5_CL_VXLAN_FLOODINATOR
)

var cumulusZapi5CommandMap = map[API_TYPE]API_TYPE{
	FRR_ZAPI5_CL_INTERFACE_ADD:                FRR_ZAPI5_INTERFACE_ADD,
	FRR_ZAPI5_CL_INTERFACE_DELETE:             FRR_ZAPI5_INTERFACE_DELETE,
	FRR_ZAPI5_CL_INTERFACE_ADDRESS_ADD:        FRR_ZAPI5_INTERFACE_ADDRESS_ADD,
	FRR_ZAPI5_CL_INTERFACE_ADDRESS_DELETE:     FRR_ZAPI5_INTERFACE_ADDRESS_DELETE,
	FRR_ZAPI5_CL_INTERFACE_UP:                 FRR_ZAPI5_INTERFACE_UP,
	FRR_ZAPI5_CL_INTERFACE_DOWN:               FRR_ZAPI5_INTERFACE_DOWN,
	FRR_ZAPI5_CL_INTERFACE_SET_MASTER:         FRR_ZAPI5_INTERFACE_SET_MASTER,
	FRR_ZAPI5_CL_ROUTE_ADD:                    FRR_ZAPI5_ROUTE_ADD,
	FRR_ZAPI5_CL_ROUTE_DELETE:                 FRR_ZAPI5_ROUTE_DELETE,
	FRR_ZAPI5_CL_ROUTE_NOTIFY_OWNER:           FRR_ZAPI5_ROUTE_NOTIFY_OWNER,
	FRR_ZAPI5_CL_REDISTRIBUTE_ADD:             FRR_ZAPI5_REDISTRIBUTE_ADD,
	FRR_ZAPI5_CL_REDISTRIBUTE_DELETE:          FRR_ZAPI5_REDISTRIBUTE_DELETE,
	FRR_ZAPI5_CL_REDISTRIBUTE_DEFAULT_ADD:     FRR_ZAPI5_REDISTRIBUTE_DEFAULT_ADD,
	FRR_ZAPI5_CL_REDISTRIBUTE_DEFAULT_DELETE:  FRR_ZAPI5_REDISTRIBUTE_DEFAULT_DELETE,
	FRR_ZAPI5_CL_ROUTER_ID_ADD:                FRR_ZAPI5_ROUTER_ID_ADD,
	FRR_ZAPI5_CL_ROUTER_ID_DELETE:             FRR_ZAPI5_ROUTER_ID_DELETE,
	FRR_ZAPI5_CL_ROUTER_ID_UPDATE:             FRR_ZAPI5_ROUTER_ID_UPDATE,
	FRR_ZAPI5_CL_HELLO:                        FRR_ZAPI5_HELLO,
	FRR_ZAPI5_CL_NEXTHOP_REGISTER:             FRR_ZAPI5_NEXTHOP_REGISTER,
	FRR_ZAPI5_CL_NEXTHOP_UNREGISTER:           FRR_ZAPI5_NEXTHOP_UNREGISTER,
	FRR_ZAPI5_CL_NEXTHOP_UPDATE:               FRR_ZAPI5_NEXTHOP_UPDATE,
	FRR_ZAPI5_CL_INTERFACE_NBR_ADDRESS_ADD:    FRR_ZAPI5_INTERFACE_NBR_ADDRESS_ADD,
	FRR_ZAPI5_CL_INTERFACE_NBR_ADDRESS_DELETE: FRR_ZAPI5_INTERFACE_NBR_ADDRESS_DELETE,
	FRR_ZAPI5_CL_INTERFACE_BFD_DEST_UPDATE:    FRR_ZAPI5_INTERFACE_BFD_DEST_UPDATE,
	FRR_ZAPI5_CL_IMPORT_ROUTE_REGISTER:        FRR_ZAPI5_IMPORT_ROUTE_REGISTER,
	FRR_ZAPI5_CL_IMPORT_ROUTE_UNREGISTER:      FRR_ZAPI5_IMPORT_ROUTE_REGISTER,
	FRR_ZAPI5_CL_IMPORT_CHECK_UPDATE:          FRR_ZAPI5_IMPORT_CHECK_UPDATE,
	FRR_ZAPI5_CL_BFD_DEST_REGISTER:            FRR_ZAPI5_BFD_DEST_REGISTER,
	FRR_ZAPI5_CL_BFD_DEST_DEREGISTER:          FRR_ZAPI5_BFD_DEST_DEREGISTER,
	FRR_ZAPI5_CL_BFD_DEST_UPDATE:              FRR_ZAPI5_BFD_DEST_UPDATE,
	FRR_ZAPI5_CL_BFD_DEST_REPLAY:              FRR_ZAPI5_BFD_DEST_REPLAY,
	FRR_ZAPI5_CL_REDISTRIBUTE_ROUTE_ADD:       FRR_ZAPI5_REDISTRIBUTE_ROUTE_ADD,
	FRR_ZAPI5_CL_REDISTRIBUTE_ROUTE_DEL:       FRR_ZAPI5_REDISTRIBUTE_ROUTE_DEL,
	FRR_ZAPI5_CL_VRF_UNREGISTER:               FRR_ZAPI5_VRF_UNREGISTER,
	FRR_ZAPI5_CL_VRF_ADD:                      FRR_ZAPI5_VRF_ADD,
	FRR_ZAPI5_CL_VRF_DELETE:                   FRR_ZAPI5_VRF_DELETE,
	FRR_ZAPI5_CL_VRF_LABEL:                    FRR_ZAPI5_VRF_LABEL,
	FRR_ZAPI5_CL_INTERFACE_VRF_UPDATE:         FRR_ZAPI5_INTERFACE_VRF_UPDATE,
	FRR_ZAPI5_CL_BFD_CLIENT_REGISTER:          FRR_ZAPI5_BFD_CLIENT_REGISTER,
	FRR_ZAPI5_CL_INTERFACE_ENABLE_RADV:        FRR_ZAPI5_INTERFACE_ENABLE_RADV,
	FRR_ZAPI5_CL_INTERFACE_DISABLE_RADV:       FRR_ZAPI5_INTERFACE_DISABLE_RADV,
	FRR_ZAPI5_CL_IPV4_NEXTHOP_LOOKUP_MRIB:     FRR_ZAPI5_IPV4_NEXTHOP_LOOKUP_MRIB,
	FRR_ZAPI5_CL_INTERFACE_LINK_PARAMS:        FRR_ZAPI5_INTERFACE_LINK_PARAMS,
	FRR_ZAPI5_CL_MPLS_LABELS_ADD:              FRR_ZAPI5_MPLS_LABELS_ADD,
	FRR_ZAPI5_CL_MPLS_LABELS_DELETE:           FRR_ZAPI5_MPLS_LABELS_DELETE,
	FRR_ZAPI5_CL_IPMR_ROUTE_STATS:             FRR_ZAPI5_IPMR_ROUTE_STATS,
	FRR_ZAPI5_CL_LABEL_MANAGER_CONNECT:        FRR_ZAPI5_LABEL_MANAGER_CONNECT,
	FRR_ZAPI5_CL_GET_LABEL_CHUNK:              FRR_ZAPI5_GET_LABEL_CHUNK,
	FRR_ZAPI5_CL_RELEASE_LABEL_CHUNK:          FRR_ZAPI5_RELEASE_LABEL_CHUNK,
	FRR_ZAPI5_CL_FEC_REGISTER:                 FRR_ZAPI5_FEC_REGISTER,
	FRR_ZAPI5_CL_FEC_UNREGISTER:               FRR_ZAPI5_FEC_UNREGISTER,
	FRR_ZAPI5_CL_FEC_UPDATE:                   FRR_ZAPI5_FEC_UPDATE,
	FRR_ZAPI5_CL_ADVERTISE_DEFAULT_GW:         FRR_ZAPI5_ADVERTISE_DEFAULT_GW,
	FRR_ZAPI5_CL_ADVERTISE_SUBNET:             FRR_ZAPI5_ADVERTISE_SUBNET,
	FRR_ZAPI5_CL_ADVERTISE_ALL_VNI:            FRR_ZAPI5_ADVERTISE_ALL_VNI,
	FRR_ZAPI5_CL_VNI_ADD:                      FRR_ZAPI5_VNI_ADD,
	FRR_ZAPI5_CL_VNI_DEL:                      FRR_ZAPI5_VNI_DEL,
	FRR_ZAPI5_CL_L3VNI_ADD:                    FRR_ZAPI5_L3VNI_ADD,
	FRR_ZAPI5_CL_L3VNI_DEL:                    FRR_ZAPI5_L3VNI_DEL,
	FRR_ZAPI5_CL_REMOTE_VTEP_ADD:              FRR_ZAPI5_REMOTE_VTEP_ADD,
	FRR_ZAPI5_CL_REMOTE_VTEP_DEL:              FRR_ZAPI5_REMOTE_VTEP_DEL,
	FRR_ZAPI5_CL_MACIP_ADD:                    FRR_ZAPI5_MACIP_ADD,
	FRR_ZAPI5_CL_MACIP_DEL:                    FRR_ZAPI5_MACIP_DEL,
	FRR_ZAPI5_CL_IP_PREFIX_ROUTE_ADD:          FRR_ZAPI5_IP_PREFIX_ROUTE_ADD,
	FRR_ZAPI5_CL_IP_PREFIX_ROUTE_DEL:          FRR_ZAPI5_IP_PREFIX_ROUTE_DEL,
	FRR_ZAPI5_CL_REMOTE_MACIP_ADD:             FRR_ZAPI5_REMOTE_MACIP_ADD,
	FRR_ZAPI5_CL_REMOTE_MACIP_DEL:             FRR_ZAPI5_REMOTE_MACIP_DEL,
	FRR_ZAPI5_CL_PW_ADD:                       FRR_ZAPI5_PW_ADD,
	FRR_ZAPI5_CL_PW_DELETE:                    FRR_ZAPI5_PW_DELETE,
	FRR_ZAPI5_CL_PW_SET:                       FRR_ZAPI5_PW_SET,
	FRR_ZAPI5_CL_PW_UNSET:                     FRR_ZAPI5_PW_UNSET,
	FRR_ZAPI5_CL_PW_STATUS_UPDATE:             FRR_ZAPI5_PW_STATUS_UPDATE,
	FRR_ZAPI5_CL_RULE_ADD:                     FRR_ZAPI5_RULE_ADD,
	FRR_ZAPI5_CL_RULE_DELETE:                  FRR_ZAPI5_RULE_DELETE,
	FRR_ZAPI5_CL_RULE_NOTIFY_OWNER:            FRR_ZAPI5_RULE_NOTIFY_OWNER,
}

// cumulusZapi5Command adjust command (API_TYPE) between cumulus Zebra and latest Zapi5
func cumulusZapi5Command(command API_TYPE, softwareName string, from bool) (API_TYPE, error) {
	if softwareName != "cumulus" {
		return command, fmt.Errorf("softwareName %s is not supported", softwareName)
	}
	// if from is true, command will be converted from cumulus Zebra to latest Zapi5 (for parsing)
	// if from is false, command will be converted from latest Zapi5 to cumulus Zebra (for send)
	if from {
		return cumulusZapi5CommandMap[command], nil
	} else {
		for k, v := range cumulusZapi5CommandMap {
			if v == command {
				return k, nil
			}
		}
	}
	return command, fmt.Errorf("unsupported command")
}
