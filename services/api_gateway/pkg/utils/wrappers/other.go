package wrappers

import "github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"

// --------------------------------------------------------- InfoResponse ---------------------------------------------------------
type InfoResponse struct {
	Response *proto_files.InfoResponse
}

func (ir *InfoResponse) IsResponseNil() bool {
	return ir.Response == nil
}

func (ir *InfoResponse) IsInfoNil() bool {
	return ir.Response != nil && ir.Response.Info == nil
}

// --------------------------------------------------------- EnhancedInfoResponse ---------------------------------------------------------
type EnhancedInfoResponse struct {
	Response *proto_files.EnhancedInfoResponse
}

func (eir *EnhancedInfoResponse) IsResponseNil() bool {
	return eir.Response == nil
}

func (eir *EnhancedInfoResponse) IsInfoNil() bool {
	return eir.Response != nil && eir.Response.Info == nil
}

// --------------------------------------------------------- IDInfoResponse ---------------------------------------------------------

type IDInfoResponse struct {
	Response *proto_files.IDInfoResponse
}

func (eir *IDInfoResponse) IsResponseNil() bool {
	return eir.Response == nil
}

func (eir *IDInfoResponse) IsInfoNil() bool {
	return eir.Response != nil && eir.Response.Info == nil
}
