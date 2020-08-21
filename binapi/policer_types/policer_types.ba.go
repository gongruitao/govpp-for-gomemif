// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.05-release

// Package policer_types contains generated bindings for API file policer_types.api.
//
// Contents:
//   4 enums
//   1 struct
//
package policer_types

import (
	"strconv"

	api "git.fd.io/govpp.git/api"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

// Sse2QosActionType defines enum 'sse2_qos_action_type'.
type Sse2QosActionType uint8

const (
	SSE2_QOS_ACTION_API_DROP              Sse2QosActionType = 0
	SSE2_QOS_ACTION_API_TRANSMIT          Sse2QosActionType = 1
	SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT Sse2QosActionType = 2
)

var (
	Sse2QosActionType_name = map[uint8]string{
		0: "SSE2_QOS_ACTION_API_DROP",
		1: "SSE2_QOS_ACTION_API_TRANSMIT",
		2: "SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT",
	}
	Sse2QosActionType_value = map[string]uint8{
		"SSE2_QOS_ACTION_API_DROP":              0,
		"SSE2_QOS_ACTION_API_TRANSMIT":          1,
		"SSE2_QOS_ACTION_API_MARK_AND_TRANSMIT": 2,
	}
)

func (x Sse2QosActionType) String() string {
	s, ok := Sse2QosActionType_name[uint8(x)]
	if ok {
		return s
	}
	return "Sse2QosActionType(" + strconv.Itoa(int(x)) + ")"
}

// Sse2QosPolicerType defines enum 'sse2_qos_policer_type'.
type Sse2QosPolicerType uint8

const (
	SSE2_QOS_POLICER_TYPE_API_1R2C             Sse2QosPolicerType = 0
	SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697    Sse2QosPolicerType = 1
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698    Sse2QosPolicerType = 2
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115    Sse2QosPolicerType = 3
	SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1 Sse2QosPolicerType = 4
	SSE2_QOS_POLICER_TYPE_API_MAX              Sse2QosPolicerType = 5
)

var (
	Sse2QosPolicerType_name = map[uint8]string{
		0: "SSE2_QOS_POLICER_TYPE_API_1R2C",
		1: "SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697",
		2: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698",
		3: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115",
		4: "SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1",
		5: "SSE2_QOS_POLICER_TYPE_API_MAX",
	}
	Sse2QosPolicerType_value = map[string]uint8{
		"SSE2_QOS_POLICER_TYPE_API_1R2C":             0,
		"SSE2_QOS_POLICER_TYPE_API_1R3C_RFC_2697":    1,
		"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_2698":    2,
		"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_4115":    3,
		"SSE2_QOS_POLICER_TYPE_API_2R3C_RFC_MEF5CF1": 4,
		"SSE2_QOS_POLICER_TYPE_API_MAX":              5,
	}
)

func (x Sse2QosPolicerType) String() string {
	s, ok := Sse2QosPolicerType_name[uint8(x)]
	if ok {
		return s
	}
	return "Sse2QosPolicerType(" + strconv.Itoa(int(x)) + ")"
}

// Sse2QosRateType defines enum 'sse2_qos_rate_type'.
type Sse2QosRateType uint8

const (
	SSE2_QOS_RATE_API_KBPS    Sse2QosRateType = 0
	SSE2_QOS_RATE_API_PPS     Sse2QosRateType = 1
	SSE2_QOS_RATE_API_INVALID Sse2QosRateType = 2
)

var (
	Sse2QosRateType_name = map[uint8]string{
		0: "SSE2_QOS_RATE_API_KBPS",
		1: "SSE2_QOS_RATE_API_PPS",
		2: "SSE2_QOS_RATE_API_INVALID",
	}
	Sse2QosRateType_value = map[string]uint8{
		"SSE2_QOS_RATE_API_KBPS":    0,
		"SSE2_QOS_RATE_API_PPS":     1,
		"SSE2_QOS_RATE_API_INVALID": 2,
	}
)

func (x Sse2QosRateType) String() string {
	s, ok := Sse2QosRateType_name[uint8(x)]
	if ok {
		return s
	}
	return "Sse2QosRateType(" + strconv.Itoa(int(x)) + ")"
}

// Sse2QosRoundType defines enum 'sse2_qos_round_type'.
type Sse2QosRoundType uint8

const (
	SSE2_QOS_ROUND_API_TO_CLOSEST Sse2QosRoundType = 0
	SSE2_QOS_ROUND_API_TO_UP      Sse2QosRoundType = 1
	SSE2_QOS_ROUND_API_TO_DOWN    Sse2QosRoundType = 2
	SSE2_QOS_ROUND_API_INVALID    Sse2QosRoundType = 3
)

var (
	Sse2QosRoundType_name = map[uint8]string{
		0: "SSE2_QOS_ROUND_API_TO_CLOSEST",
		1: "SSE2_QOS_ROUND_API_TO_UP",
		2: "SSE2_QOS_ROUND_API_TO_DOWN",
		3: "SSE2_QOS_ROUND_API_INVALID",
	}
	Sse2QosRoundType_value = map[string]uint8{
		"SSE2_QOS_ROUND_API_TO_CLOSEST": 0,
		"SSE2_QOS_ROUND_API_TO_UP":      1,
		"SSE2_QOS_ROUND_API_TO_DOWN":    2,
		"SSE2_QOS_ROUND_API_INVALID":    3,
	}
)

func (x Sse2QosRoundType) String() string {
	s, ok := Sse2QosRoundType_name[uint8(x)]
	if ok {
		return s
	}
	return "Sse2QosRoundType(" + strconv.Itoa(int(x)) + ")"
}

// Sse2QosAction defines type 'sse2_qos_action'.
type Sse2QosAction struct {
	Type Sse2QosActionType `binapi:"sse2_qos_action_type,name=type" json:"type,omitempty"`
	Dscp uint8             `binapi:"u8,name=dscp" json:"dscp,omitempty"`
}
