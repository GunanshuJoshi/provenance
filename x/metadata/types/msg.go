package types

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
	"github.com/provenance-io/provenance/x/metadata/types/p8e"
	yaml "gopkg.in/yaml.v2"
)

const (
	TypeMsgWriteScopeRequest                      = "write_scope_request"
	TypeMsgDeleteScopeRequest                     = "delete_scope_request"
	TypeMsgAddScopeDataAccessRequest              = "add_scope_data_access_request"
	TypeMsgDeleteScopeDataAccessRequest           = "delete_scope_data_access_request"
	TypeMsgAddScopeOwnerRequest                   = "add_scope_owner_request"
	TypeMsgDeleteScopeOwnerRequest                = "delete_scope_owner_request"
	TypeMsgWriteSessionRequest                    = "write_session_request"
	TypeMsgWriteRecordRequest                     = "write_record_request"
	TypeMsgDeleteRecordRequest                    = "delete_record_request"
	TypeMsgWriteScopeSpecificationRequest         = "write_scope_specification_request"
	TypeMsgDeleteScopeSpecificationRequest        = "delete_scope_specification_request"
	TypeMsgWriteContractSpecificationRequest      = "write_contract_specification_request"
	TypeMsgDeleteContractSpecificationRequest     = "delete_contract_specification_request"
	TypeMsgAddContractSpecToScopeSpecRequest      = "add_contract_spec_to_scope_spec_request"
	TypeMsgDeleteContractSpecFromScopeSpecRequest = "delete_contract_spec_from_scope_spec_request"
	TypeMsgWriteRecordSpecificationRequest        = "write_record_specification_request"
	TypeMsgDeleteRecordSpecificationRequest       = "delete_record_specification_request"
	TypeMsgWriteP8EContractSpecRequest            = "write_p8e_contract_spec_request"
	TypeMsgP8eMemorializeContractRequest          = "p8e_memorialize_contract_request"
	TypeMsgBindOSLocatorRequest                   = "write_os_locator_request"
	TypeMsgDeleteOSLocatorRequest                 = "delete_os_locator_request"
	TypeMsgModifyOSLocatorRequest                 = "modify_os_locator_request"
)

// Compile time interface checks.
var (
	_ sdk.Msg = &MsgWriteScopeRequest{}
	_ sdk.Msg = &MsgDeleteScopeRequest{}
	_ sdk.Msg = &MsgAddScopeDataAccessRequest{}
	_ sdk.Msg = &MsgDeleteScopeDataAccessRequest{}
	_ sdk.Msg = &MsgAddScopeOwnerRequest{}
	_ sdk.Msg = &MsgDeleteScopeOwnerRequest{}
	_ sdk.Msg = &MsgWriteSessionRequest{}
	_ sdk.Msg = &MsgWriteRecordRequest{}
	_ sdk.Msg = &MsgDeleteRecordRequest{}
	_ sdk.Msg = &MsgWriteScopeSpecificationRequest{}
	_ sdk.Msg = &MsgDeleteScopeSpecificationRequest{}
	_ sdk.Msg = &MsgWriteContractSpecificationRequest{}
	_ sdk.Msg = &MsgDeleteContractSpecificationRequest{}
	_ sdk.Msg = &MsgAddContractSpecToScopeSpecRequest{}
	_ sdk.Msg = &MsgDeleteContractSpecFromScopeSpecRequest{}
	_ sdk.Msg = &MsgWriteRecordSpecificationRequest{}
	_ sdk.Msg = &MsgDeleteRecordSpecificationRequest{}
	_ sdk.Msg = &MsgBindOSLocatorRequest{}
	_ sdk.Msg = &MsgDeleteOSLocatorRequest{}
	_ sdk.Msg = &MsgModifyOSLocatorRequest{}
	_ sdk.Msg = &MsgWriteP8EContractSpecRequest{}
	_ sdk.Msg = &MsgP8EMemorializeContractRequest{}
)

// private method to convert an array of strings into an array of Acc Addresses.
func stringsToAccAddresses(strings []string) []sdk.AccAddress {
	retval := make([]sdk.AccAddress, len(strings))

	for i, str := range strings {
		retval[i] = stringToAccAddress(str)
	}

	return retval
}

func stringToAccAddress(s string) sdk.AccAddress {
	accAddress, err := sdk.AccAddressFromBech32(s)
	if err != nil {
		panic(err)
	}
	return accAddress
}

// ------------------  MsgWriteScopeRequest  ------------------

// NewMsgWriteScopeRequest creates a new msg instance
func NewMsgWriteScopeRequest(scope Scope, signers []string) *MsgWriteScopeRequest {
	return &MsgWriteScopeRequest{
		Scope:   scope,
		Signers: signers,
	}
}

func (msg MsgWriteScopeRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteScopeRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteScopeRequest) Type() string {
	return TypeMsgWriteScopeRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteScopeRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteScopeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteScopeRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Scope.ValidateBasic()
}

// ConvertOptionalFields will look at the ScopeUuid and SpecUuid fields in the message.
// For each, if present, it will be converted to a MetadataAddress and set in the Scope appropriately.
// Once used, those uuid fields will be set to empty strings so that calling this again has no effect.
func (msg *MsgWriteScopeRequest) ConvertOptionalFields() error {
	if len(msg.ScopeUuid) > 0 {
		uid, err := uuid.Parse(msg.ScopeUuid)
		if err != nil {
			return fmt.Errorf("invalid scope uuid: %w", err)
		}
		scopeAddr := ScopeMetadataAddress(uid)
		if !msg.Scope.ScopeId.Empty() && !msg.Scope.ScopeId.Equals(scopeAddr) {
			return fmt.Errorf("msg.Scope.ScopeId [%s] is different from the one created from msg.ScopeUuid [%s]",
				msg.Scope.ScopeId, msg.ScopeUuid)
		}
		msg.Scope.ScopeId = scopeAddr
		msg.ScopeUuid = ""
	}
	if len(msg.SpecUuid) > 0 {
		uid, err := uuid.Parse(msg.SpecUuid)
		if err != nil {
			return fmt.Errorf("invalid spec uuid: %w", err)
		}
		specAddr := ScopeSpecMetadataAddress(uid)
		if !msg.Scope.SpecificationId.Empty() && !msg.Scope.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Scope.SpecificationId [%s] is different from the one created from msg.SpecUuid [%s]",
				msg.Scope.SpecificationId, msg.SpecUuid)
		}
		msg.Scope.SpecificationId = specAddr
		msg.SpecUuid = ""
	}
	return nil
}

// ------------------  NewMsgDeleteScopeRequest  ------------------

// NewMsgDeleteScopeRequest creates a new msg instance
func NewMsgDeleteScopeRequest(scopeID MetadataAddress, signers []string) *MsgDeleteScopeRequest {
	return &MsgDeleteScopeRequest{
		ScopeId: scopeID,
		Signers: signers,
	}
}

func (msg MsgDeleteScopeRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteScopeRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteScopeRequest) Type() string {
	return TypeMsgDeleteScopeRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteScopeRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteScopeRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteScopeRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if !msg.ScopeId.IsScopeAddress() {
		return fmt.Errorf("invalid scope address")
	}
	return nil
}

// ------------------  MsgAddScopeDataAccessRequest  ------------------

// NewMsgAddScopeDataAccessRequest creates a new msg instance
func NewMsgAddScopeDataAccessRequest(scopeID MetadataAddress, dataAccessAddrs []string, signers []string) *MsgAddScopeDataAccessRequest {
	return &MsgAddScopeDataAccessRequest{
		ScopeId:    scopeID,
		DataAccess: dataAccessAddrs,
		Signers:    signers,
	}
}

func (msg MsgAddScopeDataAccessRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgAddScopeDataAccessRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgAddScopeDataAccessRequest) Type() string {
	return TypeMsgAddScopeDataAccessRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddScopeDataAccessRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAddScopeDataAccessRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgAddScopeDataAccessRequest) ValidateBasic() error {
	if !msg.ScopeId.IsScopeAddress() {
		return fmt.Errorf("address is not a scope id: %v", msg.ScopeId.String())
	}
	if len(msg.DataAccess) < 1 {
		return fmt.Errorf("at least one data access address is required")
	}
	for _, da := range msg.DataAccess {
		_, err := sdk.AccAddressFromBech32(da)
		if err != nil {
			return fmt.Errorf("data access address is invalid: %s", da)
		}
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgDeleteScopeDataAccessRequest  ------------------

// NewMsgDeleteScopeDataAccessRequest creates a new msg instance
func NewMsgDeleteScopeDataAccessRequest(scopeID MetadataAddress, dataAccessAddrs []string, signers []string) *MsgDeleteScopeDataAccessRequest {
	return &MsgDeleteScopeDataAccessRequest{
		ScopeId:    scopeID,
		DataAccess: dataAccessAddrs,
		Signers:    signers,
	}
}

func (msg MsgDeleteScopeDataAccessRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteScopeDataAccessRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteScopeDataAccessRequest) Type() string {
	return TypeMsgDeleteScopeDataAccessRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteScopeDataAccessRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteScopeDataAccessRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteScopeDataAccessRequest) ValidateBasic() error {
	if !msg.ScopeId.IsScopeAddress() {
		return fmt.Errorf("address is not a scope id: %v", msg.ScopeId.String())
	}
	if len(msg.DataAccess) < 1 {
		return fmt.Errorf("at least one data access address is required")
	}
	for _, da := range msg.DataAccess {
		_, err := sdk.AccAddressFromBech32(da)
		if err != nil {
			return fmt.Errorf("data access address is invalid: %s", da)
		}
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgAddScopeOwnerRequest  ------------------

// NewMsgAddScopeOwnerRequest creates a new msg instance
func NewMsgAddScopeOwnerRequest(scopeID MetadataAddress, owners []Party, signers []string) *MsgAddScopeOwnerRequest {
	return &MsgAddScopeOwnerRequest{
		ScopeId: scopeID,
		Owners:  owners,
		Signers: signers,
	}
}

func (msg MsgAddScopeOwnerRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgAddScopeOwnerRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgAddScopeOwnerRequest) Type() string {
	return TypeMsgAddScopeOwnerRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddScopeOwnerRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAddScopeOwnerRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgAddScopeOwnerRequest) ValidateBasic() error {
	if !msg.ScopeId.IsScopeAddress() {
		return fmt.Errorf("address is not a scope id: %v", msg.ScopeId.String())
	}
	if err := ValidatePartiesBasic(msg.Owners); err != nil {
		return fmt.Errorf("invalid owners: %w", err)
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgDeleteScopeOwnerRequest  ------------------

// NewMsgDeleteScopeOwnerRequest creates a new msg instance
func NewMsgDeleteScopeOwnerRequest(scopeID MetadataAddress, owners []string, signers []string) *MsgDeleteScopeOwnerRequest {
	return &MsgDeleteScopeOwnerRequest{
		ScopeId: scopeID,
		Owners:  owners,
		Signers: signers,
	}
}

func (msg MsgDeleteScopeOwnerRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteScopeOwnerRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteScopeOwnerRequest) Type() string {
	return TypeMsgDeleteScopeOwnerRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteScopeOwnerRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteScopeOwnerRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteScopeOwnerRequest) ValidateBasic() error {
	if !msg.ScopeId.IsScopeAddress() {
		return fmt.Errorf("address is not a scope id: %v", msg.ScopeId.String())
	}
	if len(msg.Owners) < 1 {
		return fmt.Errorf("at least one owner address is required")
	}
	for _, owner := range msg.Owners {
		_, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return fmt.Errorf("owner address is invalid: %s", owner)
		}
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgWriteSessionRequest  ------------------

// NewMsgWriteSessionRequest creates a new msg instance
func NewMsgWriteSessionRequest(session Session, signers []string) *MsgWriteSessionRequest {
	return &MsgWriteSessionRequest{Session: session, Signers: signers}
}

func (msg MsgWriteSessionRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteSessionRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteSessionRequest) Type() string {
	return TypeMsgWriteSessionRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteSessionRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteSessionRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteSessionRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Session.ValidateBasic()
}

// ConvertOptionalFields will look at the SessionIdComponents and SpecUuid fields in the message.
// For each, if present, it will be converted to a MetadataAddress and set in the Session appropriately.
// Once used, those fields will be emptied so that calling this again has no effect.
func (msg *MsgWriteSessionRequest) ConvertOptionalFields() error {
	if msg.SessionIdComponents != nil {
		sessionAddr, err := msg.SessionIdComponents.GetSessionAddr()
		if err != nil {
			return fmt.Errorf("invalid session id components: %w", err)
		}
		if sessionAddr != nil {
			if !msg.Session.SessionId.Empty() && !msg.Session.SessionId.Equals(*sessionAddr) {
				return fmt.Errorf("msg.Session.SessionId [%s] is different from the one created from msg.SessionIdComponents %v",
					msg.Session.SessionId, msg.SessionIdComponents)
			}
			msg.Session.SessionId = *sessionAddr
		}
		msg.SessionIdComponents = nil
	}
	if len(msg.SpecUuid) > 0 {
		uid, err := uuid.Parse(msg.SpecUuid)
		if err != nil {
			return fmt.Errorf("invalid spec uuid: %w", err)
		}
		specAddr := ContractSpecMetadataAddress(uid)
		if !msg.Session.SpecificationId.Empty() && !msg.Session.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Session.SpecificationId [%s] is different from the one created from msg.SpecUuid [%s]",
				msg.Session.SpecificationId, msg.SpecUuid)
		}
		msg.Session.SpecificationId = specAddr
		msg.SpecUuid = ""
	}
	return nil
}

// ------------------  MsgWriteRecordRequest  ------------------

// NewMsgWriteRecordRequest creates a new msg instance
func NewMsgWriteRecordRequest(record Record, sessionIDComponents *SessionIdComponents, contractSpecUUID string, signers []string, parties []Party) *MsgWriteRecordRequest {
	return &MsgWriteRecordRequest{Record: record, Parties: parties, Signers: signers, SessionIdComponents: sessionIDComponents, ContractSpecUuid: contractSpecUUID}
}

func (msg MsgWriteRecordRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteRecordRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteRecordRequest) Type() string {
	return TypeMsgWriteRecordRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteRecordRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteRecordRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteRecordRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Record.ValidateBasic()
}

// ConvertOptionalFields will look at the SessionIdComponents and ContractSpecUuid fields in the message.
// For each, if present, it will be converted to a MetadataAddress and set in the Record appropriately.
// Once used, those fields will be emptied so that calling this again has no effect.
func (msg *MsgWriteRecordRequest) ConvertOptionalFields() error {
	if msg.SessionIdComponents != nil {
		sessionAddr, err := msg.SessionIdComponents.GetSessionAddr()
		if err != nil {
			return fmt.Errorf("invalid session id components: %w", err)
		}
		if sessionAddr != nil {
			if !msg.Record.SessionId.Empty() && !msg.Record.SessionId.Equals(*sessionAddr) {
				return fmt.Errorf("msg.Record.SessionId [%s] is different from the one created from msg.SessionIdComponents %v",
					msg.Record.SessionId, msg.SessionIdComponents)
			}
			msg.Record.SessionId = *sessionAddr
			msg.SessionIdComponents = nil
		}
	}
	if len(msg.ContractSpecUuid) > 0 {
		uid, err := uuid.Parse(msg.ContractSpecUuid)
		if err != nil {
			return fmt.Errorf("invalid contract spec uuid: %w", err)
		}
		if len(strings.TrimSpace(msg.Record.Name)) == 0 {
			return errors.New("empty record name")
		}
		specAddr := RecordSpecMetadataAddress(uid, msg.Record.Name)
		if !msg.Record.SpecificationId.Empty() && !msg.Record.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Record.SpecificationId [%s] is different from the one created from msg.ContractSpecUuid [%s] and msg.Record.Name [%s]",
				msg.Record.SpecificationId, msg.ContractSpecUuid, msg.Record.Name)
		}
		msg.Record.SpecificationId = specAddr
		msg.ContractSpecUuid = ""
	}
	return nil
}

// ------------------  MsgDeleteRecordRequest  ------------------

// NewMsgDeleteScopeSpecificationRequest creates a new msg instance
func NewMsgDeleteRecordRequest(recordID MetadataAddress, signers []string) *MsgDeleteRecordRequest {
	return &MsgDeleteRecordRequest{RecordId: recordID, Signers: signers}
}

func (msg MsgDeleteRecordRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteRecordRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteRecordRequest) Type() string {
	return TypeMsgDeleteRecordRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteRecordRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteRecordRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteRecordRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgWriteScopeSpecificationRequest  ------------------

// NewMsgAddScopeSpecificationRequest creates a new msg instance
func NewMsgWriteScopeSpecificationRequest(specification ScopeSpecification, signers []string) *MsgWriteScopeSpecificationRequest {
	return &MsgWriteScopeSpecificationRequest{Specification: specification, Signers: signers}
}

func (msg MsgWriteScopeSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteScopeSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteScopeSpecificationRequest) Type() string {
	return TypeMsgWriteScopeSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteScopeSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteScopeSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteScopeSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Specification.ValidateBasic()
}

// ConvertOptionalFields will look at the SpecUuid field in the message.
// If present, it will be converted to a MetadataAddress and set in the Specification appropriately.
// Once used, it will be emptied so that calling this again has no effect.
func (msg *MsgWriteScopeSpecificationRequest) ConvertOptionalFields() error {
	if len(msg.SpecUuid) > 0 {
		uid, err := uuid.Parse(msg.SpecUuid)
		if err != nil {
			return fmt.Errorf("invalid spec uuid: %w", err)
		}
		specAddr := ScopeSpecMetadataAddress(uid)
		if !msg.Specification.SpecificationId.Empty() && !msg.Specification.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Specification.SpecificationId [%s] is different from the one created from msg.SpecUuid [%s]",
				msg.Specification.SpecificationId, msg.SpecUuid)
		}
		msg.Specification.SpecificationId = specAddr
		msg.SpecUuid = ""
	}
	return nil
}

// ------------------  MsgWriteP8EContractSpecRequest  ------------------

// NewMsgWriteContractSpecRequest creates a new msg instance
func NewMsgWriteP8EContractSpecRequest(contractSpec p8e.ContractSpec, signers []string) *MsgWriteP8EContractSpecRequest {
	return &MsgWriteP8EContractSpecRequest{
		Contractspec: contractSpec,
		Signers:      signers,
	}
}

func (msg MsgWriteP8EContractSpecRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteP8EContractSpecRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteP8EContractSpecRequest) Type() string {
	return TypeMsgWriteP8EContractSpecRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteP8EContractSpecRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteP8EContractSpecRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteP8EContractSpecRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	_, _, err := ConvertP8eContractSpec(&msg.Contractspec, msg.Signers)
	if err != nil {
		return fmt.Errorf("failed to convert p8e ContractSpec %s", err)
	}
	return nil
}

// ------------------  MsgDeleteScopeSpecificationRequest  ------------------

// NewMsgDeleteScopeSpecificationRequest creates a new msg instance
func NewMsgDeleteScopeSpecificationRequest(specificationID MetadataAddress, signers []string) *MsgDeleteScopeSpecificationRequest {
	return &MsgDeleteScopeSpecificationRequest{SpecificationId: specificationID, Signers: signers}
}

func (msg MsgDeleteScopeSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteScopeSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteScopeSpecificationRequest) Type() string {
	return TypeMsgDeleteScopeSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteScopeSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteScopeSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteScopeSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgWriteContractSpecificationRequest  ------------------

// NewMsgWriteContractSpecificationRequest creates a new msg instance
func NewMsgWriteContractSpecificationRequest(specification ContractSpecification, signers []string) *MsgWriteContractSpecificationRequest {
	return &MsgWriteContractSpecificationRequest{Specification: specification, Signers: signers}
}

func (msg MsgWriteContractSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteContractSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteContractSpecificationRequest) Type() string {
	return TypeMsgWriteContractSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteContractSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteContractSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteContractSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Specification.ValidateBasic()
}

// ConvertOptionalFields will look at the SpecUuid field in the message.
// If present, it will be converted to a MetadataAddress and set in the Specification appropriately.
// Once used, it will be emptied so that calling this again has no effect.
func (msg *MsgWriteContractSpecificationRequest) ConvertOptionalFields() error {
	if len(msg.SpecUuid) > 0 {
		uid, err := uuid.Parse(msg.SpecUuid)
		if err != nil {
			return fmt.Errorf("invalid spec uuid: %w", err)
		}
		specAddr := ContractSpecMetadataAddress(uid)
		if !msg.Specification.SpecificationId.Empty() && !msg.Specification.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Specification.SpecificationId [%s] is different from the one created from msg.SpecUuid [%s]",
				msg.Specification.SpecificationId, msg.SpecUuid)
		}
		msg.Specification.SpecificationId = specAddr
		msg.SpecUuid = ""
	}
	return nil
}

// ------------------  MsgDeleteContractSpecificationRequest  ------------------

// NewMsgDeleteContractSpecificationRequest creates a new msg instance
func NewMsgDeleteContractSpecificationRequest(specificationID MetadataAddress, signers []string) *MsgDeleteContractSpecificationRequest {
	return &MsgDeleteContractSpecificationRequest{SpecificationId: specificationID, Signers: signers}
}

func (msg MsgDeleteContractSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteContractSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteContractSpecificationRequest) Type() string {
	return TypeMsgDeleteContractSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteContractSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteContractSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteContractSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgAddContractSpecToScopeSpecRequest  ------------------

// NewMsgAddContractSpecToScopeSpecRequest creates a new msg instance
func NewMsgAddContractSpecToScopeSpecRequest(contractSpecID MetadataAddress, scopeSpecID MetadataAddress, signers []string) *MsgAddContractSpecToScopeSpecRequest {
	return &MsgAddContractSpecToScopeSpecRequest{ContractSpecificationId: contractSpecID, ScopeSpecificationId: scopeSpecID, Signers: signers}
}

func (msg MsgAddContractSpecToScopeSpecRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgAddContractSpecToScopeSpecRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgAddContractSpecToScopeSpecRequest) Type() string {
	return TypeMsgAddContractSpecToScopeSpecRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgAddContractSpecToScopeSpecRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgAddContractSpecToScopeSpecRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgAddContractSpecToScopeSpecRequest) ValidateBasic() error {
	if !msg.ContractSpecificationId.IsContractSpecificationAddress() {
		return fmt.Errorf("address is not a contract specification id: %s", msg.ContractSpecificationId.String())
	}
	if !msg.ScopeSpecificationId.IsScopeSpecificationAddress() {
		return fmt.Errorf("address is not a scope specification id: %s", msg.ScopeSpecificationId.String())
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgDeleteContractSpecFromScopeSpecRequest  ------------------

// NewMsgDeleteContractSpecFromScopeSpecRequest creates a new msg instance
func NewMsgDeleteContractSpecFromScopeSpecRequest(contractSpecID MetadataAddress, scopeSpecID MetadataAddress, signers []string) *MsgDeleteContractSpecFromScopeSpecRequest {
	return &MsgDeleteContractSpecFromScopeSpecRequest{ContractSpecificationId: contractSpecID, ScopeSpecificationId: scopeSpecID, Signers: signers}
}

func (msg MsgDeleteContractSpecFromScopeSpecRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteContractSpecFromScopeSpecRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteContractSpecFromScopeSpecRequest) Type() string {
	return TypeMsgDeleteContractSpecFromScopeSpecRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteContractSpecFromScopeSpecRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteContractSpecFromScopeSpecRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteContractSpecFromScopeSpecRequest) ValidateBasic() error {
	if !msg.ContractSpecificationId.IsContractSpecificationAddress() {
		return fmt.Errorf("address is not a contract specification id: %s", msg.ContractSpecificationId.String())
	}
	if !msg.ScopeSpecificationId.IsScopeSpecificationAddress() {
		return fmt.Errorf("address is not a scope specification id: %s", msg.ScopeSpecificationId.String())
	}
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgWriteRecordSpecificationRequest  ------------------

// NewMsgAddRecordSpecificationRequest creates a new msg instance
func NewMsgWriteRecordSpecificationRequest(recordSpecification RecordSpecification, signers []string) *MsgWriteRecordSpecificationRequest {
	return &MsgWriteRecordSpecificationRequest{Specification: recordSpecification, Signers: signers}
}

func (msg MsgWriteRecordSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgWriteRecordSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgWriteRecordSpecificationRequest) Type() string {
	return TypeMsgWriteRecordSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgWriteRecordSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgWriteRecordSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgWriteRecordSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	if err := msg.ConvertOptionalFields(); err != nil {
		return err
	}
	return msg.Specification.ValidateBasic()
}

// ConvertOptionalFields will look at the ContractSpecUuid field in the message.
// If present, it will be converted to a MetadataAddress and set in the Specification appropriately.
// Once used, it will be emptied so that calling this again has no effect.
func (msg *MsgWriteRecordSpecificationRequest) ConvertOptionalFields() error {
	if len(msg.ContractSpecUuid) > 0 {
		uid, err := uuid.Parse(msg.ContractSpecUuid)
		if err != nil {
			return fmt.Errorf("invalid spec uuid: %w", err)
		}
		if len(strings.TrimSpace(msg.Specification.Name)) == 0 {
			return errors.New("empty specification name")
		}
		specAddr := RecordSpecMetadataAddress(uid, msg.Specification.Name)
		if !msg.Specification.SpecificationId.Empty() && !msg.Specification.SpecificationId.Equals(specAddr) {
			return fmt.Errorf("msg.Specification.SpecificationId [%s] is different from the one created from msg.ContractSpecUuid [%s] and msg.Specification.Name [%s]",
				msg.Specification.SpecificationId, msg.ContractSpecUuid, msg.Specification.Name)
		}
		msg.Specification.SpecificationId = specAddr
		msg.ContractSpecUuid = ""
	}
	return nil
}

// ------------------  MsgDeleteRecordSpecificationRequest  ------------------

// NewMsgDeleteRecordSpecificationRequest creates a new msg instance
func NewMsgDeleteRecordSpecificationRequest(specificationID MetadataAddress, signers []string) *MsgDeleteRecordSpecificationRequest {
	return &MsgDeleteRecordSpecificationRequest{SpecificationId: specificationID, Signers: signers}
}

func (msg MsgDeleteRecordSpecificationRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgDeleteRecordSpecificationRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgDeleteRecordSpecificationRequest) Type() string {
	return TypeMsgDeleteRecordSpecificationRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgDeleteRecordSpecificationRequest) GetSigners() []sdk.AccAddress {
	return stringsToAccAddresses(msg.Signers)
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgDeleteRecordSpecificationRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgDeleteRecordSpecificationRequest) ValidateBasic() error {
	if len(msg.Signers) < 1 {
		return fmt.Errorf("at least one signer is required")
	}
	return nil
}

// ------------------  MsgP8EMemorializeContractRequest  ------------------

// NewMsgP8EMemorializeContractRequest creates a new msg instance
func NewMsgP8EMemorializeContractRequest() *MsgP8EMemorializeContractRequest {
	return &MsgP8EMemorializeContractRequest{}
}

func (msg MsgP8EMemorializeContractRequest) String() string {
	out, _ := yaml.Marshal(msg)
	return string(out)
}

// Route returns the module route
func (msg MsgP8EMemorializeContractRequest) Route() string {
	return ModuleName
}

// Type returns the type name for this msg
func (msg MsgP8EMemorializeContractRequest) Type() string {
	return TypeMsgP8eMemorializeContractRequest
}

// GetSigners returns the address(es) that must sign over msg.GetSignBytes()
func (msg MsgP8EMemorializeContractRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{stringToAccAddress(msg.Invoker)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgP8EMemorializeContractRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// ValidateBasic performs a quick validity check
func (msg MsgP8EMemorializeContractRequest) ValidateBasic() error {
	_, err := ConvertP8eMemorializeContractRequest(&msg)
	return err
}

// ------------------  MsgBindOSLocatorRequest  ------------------

// NewMsgBindOSLocatorRequest creates a new msg instance
func NewMsgBindOSLocatorRequest(obj ObjectStoreLocator) *MsgBindOSLocatorRequest {
	return &MsgBindOSLocatorRequest{
		Locator: obj,
	}
}

func (msg MsgBindOSLocatorRequest) Route() string {
	return ModuleName
}

func (msg MsgBindOSLocatorRequest) Type() string {
	return TypeMsgBindOSLocatorRequest
}

func (msg MsgBindOSLocatorRequest) ValidateBasic() error {
	err := ValidateOSLocatorObj(msg.Locator.Owner, msg.Locator.EncryptionKey, msg.Locator.LocatorUri)
	if err != nil {
		return err
	}
	return nil
}

func (msg MsgBindOSLocatorRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgBindOSLocatorRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{stringToAccAddress(msg.Locator.Owner)}
}

// ------------------  MsgDeleteOSLocatorRequest  ------------------

func NewMsgDeleteOSLocatorRequest(obj ObjectStoreLocator) *MsgDeleteOSLocatorRequest {
	return &MsgDeleteOSLocatorRequest{
		Locator: obj,
	}
}
func (msg MsgDeleteOSLocatorRequest) Route() string {
	return ModuleName
}

func (msg MsgDeleteOSLocatorRequest) Type() string {
	return TypeMsgDeleteOSLocatorRequest
}

func (msg MsgDeleteOSLocatorRequest) ValidateBasic() error {
	err := ValidateOSLocatorObj(msg.Locator.Owner, msg.Locator.EncryptionKey, msg.Locator.LocatorUri)
	if err != nil {
		return err
	}

	return nil
}

func (msg MsgDeleteOSLocatorRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

// Signers returns the addrs of signers that must sign.
// CONTRACT: All signatures must be present to be valid.
// CONTRACT: Returns addrs in some deterministic order.
// here we assume msg for delete request has the right address
// should be verified later in the keeper?
func (msg MsgDeleteOSLocatorRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{stringToAccAddress(msg.Locator.Owner)}
}

// ValidateOSLocatorObj Validates OSLocatorObj data
func ValidateOSLocatorObj(ownerAddr, encryptionKey string, uri string) error {
	if strings.TrimSpace(ownerAddr) == "" {
		return fmt.Errorf("owner address cannot be empty")
	}

	if _, err := sdk.AccAddressFromBech32(ownerAddr); err != nil {
		return fmt.Errorf("failed to add locator for a given owner address,"+
			" invalid address: %s", ownerAddr)
	}

	if strings.TrimSpace(uri) == "" {
		return fmt.Errorf("uri cannot be empty")
	}

	if _, err := url.Parse(uri); err != nil {
		return fmt.Errorf("failed to add locator for a given"+
			" owner address, invalid uri: %s", uri)
	}

	if strings.TrimSpace(encryptionKey) != "" {
		if _, err := sdk.AccAddressFromBech32(encryptionKey); err != nil {
			return fmt.Errorf("failed to add locator for a given owner address: %s,"+
				" invalid encryption key address: %s", ownerAddr, encryptionKey)
		}
	}
	return nil
}

// ------------------  MsgModifyOSLocatorRequest  ------------------

func NewMsgModifyOSLocatorRequest(obj ObjectStoreLocator) *MsgModifyOSLocatorRequest {
	return &MsgModifyOSLocatorRequest{
		Locator: obj,
	}
}

func (msg MsgModifyOSLocatorRequest) Route() string {
	return ModuleName
}

func (msg MsgModifyOSLocatorRequest) Type() string {
	return TypeMsgModifyOSLocatorRequest
}

func (msg MsgModifyOSLocatorRequest) ValidateBasic() error {
	err := ValidateOSLocatorObj(msg.Locator.Owner, msg.Locator.EncryptionKey, msg.Locator.LocatorUri)
	if err != nil {
		return err
	}

	return nil
}

func (msg MsgModifyOSLocatorRequest) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&msg))
}

func (msg MsgModifyOSLocatorRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{stringToAccAddress(msg.Locator.Owner)}
}

// ------------------  SessionIdComponents  ------------------

func (msg *SessionIdComponents) GetSessionAddr() (*MetadataAddress, error) {
	var scopeUUID, sessionUUID *uuid.UUID
	if len(msg.SessionUuid) > 0 {
		uid, err := uuid.Parse(msg.SessionUuid)
		if err != nil {
			return nil, fmt.Errorf("invalid session uuid: %w", err)
		}
		sessionUUID = &uid
	}
	if msgScopeUUID := msg.GetScopeUuid(); len(msgScopeUUID) > 0 {
		uid, err := uuid.Parse(msgScopeUUID)
		if err != nil {
			return nil, fmt.Errorf("invalid scope uuid: %w", err)
		}
		scopeUUID = &uid
	} else if msgScopeAddr := msg.GetScopeAddr(); len(msgScopeAddr) > 0 {
		addr, addrErr := MetadataAddressFromBech32(msgScopeAddr)
		if addrErr != nil {
			return nil, fmt.Errorf("invalid scope addr: %w", addrErr)
		}
		uid, err := addr.ScopeUUID()
		if err != nil {
			return nil, fmt.Errorf("invalid scope addr: %w", err)
		}
		scopeUUID = &uid
	}

	if scopeUUID == nil && sessionUUID == nil {
		return nil, nil
	}
	if scopeUUID == nil {
		return nil, errors.New("session uuid provided but missing scope uuid or addr")
	}
	if sessionUUID == nil {
		return nil, errors.New("scope uuid or addr provided but missing session uuid")
	}
	ma := SessionMetadataAddress(*scopeUUID, *sessionUUID)
	return &ma, nil
}

// ------------------  Response Message Constructors  ------------------

func NewMsgWriteScopeResponse(scopeID MetadataAddress) *MsgWriteScopeResponse {
	return &MsgWriteScopeResponse{
		ScopeIdInfo: GetScopeIDInfo(scopeID),
	}
}

func NewMsgDeleteScopeResponse() *MsgDeleteScopeResponse {
	return &MsgDeleteScopeResponse{}
}

func NewMsgAddScopeDataAccessResponse() *MsgAddScopeDataAccessResponse {
	return &MsgAddScopeDataAccessResponse{}
}

func NewMsgDeleteScopeDataAccessResponse() *MsgDeleteScopeDataAccessResponse {
	return &MsgDeleteScopeDataAccessResponse{}
}

func NewMsgAddScopeOwnerResponse() *MsgAddScopeOwnerResponse {
	return &MsgAddScopeOwnerResponse{}
}

func NewMsgDeleteScopeOwnerResponse() *MsgDeleteScopeOwnerResponse {
	return &MsgDeleteScopeOwnerResponse{}
}

func NewMsgWriteSessionResponse(sessionID MetadataAddress) *MsgWriteSessionResponse {
	return &MsgWriteSessionResponse{
		SessionIdInfo: GetSessionIDInfo(sessionID),
	}
}

func NewMsgWriteRecordResponse(recordID MetadataAddress) *MsgWriteRecordResponse {
	return &MsgWriteRecordResponse{
		RecordIdInfo: GetRecordIDInfo(recordID),
	}
}

func NewMsgDeleteRecordResponse() *MsgDeleteRecordResponse {
	return &MsgDeleteRecordResponse{}
}

func NewMsgWriteScopeSpecificationResponse(scopeSpecID MetadataAddress) *MsgWriteScopeSpecificationResponse {
	return &MsgWriteScopeSpecificationResponse{
		ScopeSpecIdInfo: GetScopeSpecIDInfo(scopeSpecID),
	}
}

func NewMsgDeleteScopeSpecificationResponse() *MsgDeleteScopeSpecificationResponse {
	return &MsgDeleteScopeSpecificationResponse{}
}

func NewMsgWriteContractSpecificationResponse(contractSpecID MetadataAddress) *MsgWriteContractSpecificationResponse {
	return &MsgWriteContractSpecificationResponse{
		ContractSpecIdInfo: GetContractSpecIDInfo(contractSpecID),
	}
}

func NewMsgDeleteContractSpecificationResponse() *MsgDeleteContractSpecificationResponse {
	return &MsgDeleteContractSpecificationResponse{}
}

func NewMsgAddContractSpecToScopeSpecResponse() *MsgAddContractSpecToScopeSpecResponse {
	return &MsgAddContractSpecToScopeSpecResponse{}
}

func NewMsgDeleteContractSpecFromScopeSpecResponse() *MsgDeleteContractSpecFromScopeSpecResponse {
	return &MsgDeleteContractSpecFromScopeSpecResponse{}
}

func NewMsgWriteRecordSpecificationResponse(recordSpecID MetadataAddress) *MsgWriteRecordSpecificationResponse {
	return &MsgWriteRecordSpecificationResponse{
		RecordSpecIdInfo: GetRecordSpecIDInfo(recordSpecID),
	}
}

func NewMsgDeleteRecordSpecificationResponse() *MsgDeleteRecordSpecificationResponse {
	return &MsgDeleteRecordSpecificationResponse{}
}

func NewMsgWriteP8EContractSpecResponse(
	contractSpecID MetadataAddress,
	recordSpecIDs ...MetadataAddress,
) *MsgWriteP8EContractSpecResponse {
	retval := &MsgWriteP8EContractSpecResponse{
		ContractSpecIdInfo: GetContractSpecIDInfo(contractSpecID),
		RecordSpecIdInfos:  make([]*RecordSpecIdInfo, len(recordSpecIDs)),
	}
	for i, rid := range recordSpecIDs {
		retval.RecordSpecIdInfos[i] = GetRecordSpecIDInfo(rid)
	}
	return retval
}

func NewMsgP8EMemorializeContractResponse(
	scopeIDInfo *ScopeIdInfo,
	sessionIDInfo *SessionIdInfo,
	recordIDInfos []*RecordIdInfo,
) *MsgP8EMemorializeContractResponse {
	return &MsgP8EMemorializeContractResponse{
		ScopeIdInfo:   scopeIDInfo,
		SessionIdInfo: sessionIDInfo,
		RecordIdInfos: recordIDInfos,
	}
}

func NewMsgBindOSLocatorResponse(objectStoreLocator ObjectStoreLocator) *MsgBindOSLocatorResponse {
	return &MsgBindOSLocatorResponse{
		Locator: objectStoreLocator,
	}
}

func NewMsgDeleteOSLocatorResponse(objectStoreLocator ObjectStoreLocator) *MsgDeleteOSLocatorResponse {
	return &MsgDeleteOSLocatorResponse{
		Locator: objectStoreLocator,
	}
}

func NewMsgModifyOSLocatorResponse(objectStoreLocator ObjectStoreLocator) *MsgModifyOSLocatorResponse {
	return &MsgModifyOSLocatorResponse{
		Locator: objectStoreLocator,
	}
}
