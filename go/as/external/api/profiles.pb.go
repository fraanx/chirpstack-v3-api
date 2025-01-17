// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/profiles.proto

package api

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type RatePolicy int32

const (
	// Drop
	RatePolicy_DROP RatePolicy = 0
	// Mark
	RatePolicy_MARK RatePolicy = 1
)

var RatePolicy_name = map[int32]string{
	0: "DROP",
	1: "MARK",
}

var RatePolicy_value = map[string]int32{
	"DROP": 0,
	"MARK": 1,
}

func (x RatePolicy) String() string {
	return proto.EnumName(RatePolicy_name, int32(x))
}

func (RatePolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{0}
}

type ServiceProfile struct {
	// Service-profile ID (UUID string).
	// This will be automatically set on create.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Service-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	UlRate uint32 `protobuf:"varint,2,opt,name=ul_rate,json=ulRate,proto3" json:"ul_rate,omitempty"`
	// Token bucket burst size.
	UlBucketSize uint32 `protobuf:"varint,3,opt,name=ul_bucket_size,json=ulBucketSize,proto3" json:"ul_bucket_size,omitempty"`
	// Drop or mark when exceeding ULRate.
	UlRatePolicy RatePolicy `protobuf:"varint,4,opt,name=ul_rate_policy,json=ulRatePolicy,proto3,enum=api.RatePolicy" json:"ul_rate_policy,omitempty"`
	// Token bucket filling rate, including ACKs (packet/h).
	DlRate uint32 `protobuf:"varint,5,opt,name=dl_rate,json=dlRate,proto3" json:"dl_rate,omitempty"`
	// Token bucket burst size.
	DlBucketSize uint32 `protobuf:"varint,6,opt,name=dl_bucket_size,json=dlBucketSize,proto3" json:"dl_bucket_size,omitempty"`
	// Drop or mark when exceeding DLRate.
	DlRatePolicy RatePolicy `protobuf:"varint,7,opt,name=dl_rate_policy,json=dlRatePolicy,proto3,enum=api.RatePolicy" json:"dl_rate_policy,omitempty"`
	// GW metadata (RSSI, SNR, GW geoloc., etc.) are added to the packet sent to AS.
	AddGwMetadata bool `protobuf:"varint,8,opt,name=add_gw_metadata,json=addGWMetaData,proto3" json:"add_gw_metadata,omitempty"`
	// Frequency to initiate an End-Device status request (request/day).
	DevStatusReqFreq uint32 `protobuf:"varint,9,opt,name=dev_status_req_freq,json=devStatusReqFreq,proto3" json:"dev_status_req_freq,omitempty"`
	// Report End-Device battery level to AS.
	ReportDevStatusBattery bool `protobuf:"varint,10,opt,name=report_dev_status_battery,json=reportDevStatusBattery,proto3" json:"report_dev_status_battery,omitempty"`
	// Report End-Device margin to AS.
	ReportDevStatusMargin bool `protobuf:"varint,11,opt,name=report_dev_status_margin,json=reportDevStatusMargin,proto3" json:"report_dev_status_margin,omitempty"`
	// Minimum allowed data rate. Used for ADR.
	DrMin uint32 `protobuf:"varint,12,opt,name=dr_min,json=drMin,proto3" json:"dr_min,omitempty"`
	// Maximum allowed data rate. Used for ADR.
	DrMax uint32 `protobuf:"varint,13,opt,name=dr_max,json=drMax,proto3" json:"dr_max,omitempty"`
	// Channel mask. sNS does not have to obey (i.e., informative).
	ChannelMask []byte `protobuf:"bytes,14,opt,name=channel_mask,json=channelMask,proto3" json:"channel_mask,omitempty"`
	// Passive Roaming allowed.
	PrAllowed bool `protobuf:"varint,15,opt,name=pr_allowed,json=prAllowed,proto3" json:"pr_allowed,omitempty"`
	// Handover Roaming allowed.
	HrAllowed bool `protobuf:"varint,16,opt,name=hr_allowed,json=hrAllowed,proto3" json:"hr_allowed,omitempty"`
	// Roaming Activation allowed.
	RaAllowed bool `protobuf:"varint,17,opt,name=ra_allowed,json=raAllowed,proto3" json:"ra_allowed,omitempty"`
	// Enable network geolocation service.
	NwkGeoLoc bool `protobuf:"varint,18,opt,name=nwk_geo_loc,json=nwkGeoLoc,proto3" json:"nwk_geo_loc,omitempty"`
	// Target Packet Error Rate.
	TargetPer uint32 `protobuf:"varint,19,opt,name=target_per,json=targetPER,proto3" json:"target_per,omitempty"`
	// Minimum number of receiving GWs (informative).
	MinGwDiversity uint32 `protobuf:"varint,20,opt,name=min_gw_diversity,json=minGWDiversity,proto3" json:"min_gw_diversity,omitempty"`
	// Gateways under this service-profile are private.
	// This means that these gateways can only be used by devices under the
	// same service-profile.
	GwsPrivate           bool     `protobuf:"varint,24,opt,name=gws_private,json=gwsPrivate,proto3" json:"gws_private,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceProfile) Reset()         { *m = ServiceProfile{} }
func (m *ServiceProfile) String() string { return proto.CompactTextString(m) }
func (*ServiceProfile) ProtoMessage()    {}
func (*ServiceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{0}
}

func (m *ServiceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceProfile.Unmarshal(m, b)
}
func (m *ServiceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceProfile.Marshal(b, m, deterministic)
}
func (m *ServiceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceProfile.Merge(m, src)
}
func (m *ServiceProfile) XXX_Size() int {
	return xxx_messageInfo_ServiceProfile.Size(m)
}
func (m *ServiceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceProfile proto.InternalMessageInfo

func (m *ServiceProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceProfile) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *ServiceProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *ServiceProfile) GetUlRate() uint32 {
	if m != nil {
		return m.UlRate
	}
	return 0
}

func (m *ServiceProfile) GetUlBucketSize() uint32 {
	if m != nil {
		return m.UlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetUlRatePolicy() RatePolicy {
	if m != nil {
		return m.UlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetDlRate() uint32 {
	if m != nil {
		return m.DlRate
	}
	return 0
}

func (m *ServiceProfile) GetDlBucketSize() uint32 {
	if m != nil {
		return m.DlBucketSize
	}
	return 0
}

func (m *ServiceProfile) GetDlRatePolicy() RatePolicy {
	if m != nil {
		return m.DlRatePolicy
	}
	return RatePolicy_DROP
}

func (m *ServiceProfile) GetAddGwMetadata() bool {
	if m != nil {
		return m.AddGwMetadata
	}
	return false
}

func (m *ServiceProfile) GetDevStatusReqFreq() uint32 {
	if m != nil {
		return m.DevStatusReqFreq
	}
	return 0
}

func (m *ServiceProfile) GetReportDevStatusBattery() bool {
	if m != nil {
		return m.ReportDevStatusBattery
	}
	return false
}

func (m *ServiceProfile) GetReportDevStatusMargin() bool {
	if m != nil {
		return m.ReportDevStatusMargin
	}
	return false
}

func (m *ServiceProfile) GetDrMin() uint32 {
	if m != nil {
		return m.DrMin
	}
	return 0
}

func (m *ServiceProfile) GetDrMax() uint32 {
	if m != nil {
		return m.DrMax
	}
	return 0
}

func (m *ServiceProfile) GetChannelMask() []byte {
	if m != nil {
		return m.ChannelMask
	}
	return nil
}

func (m *ServiceProfile) GetPrAllowed() bool {
	if m != nil {
		return m.PrAllowed
	}
	return false
}

func (m *ServiceProfile) GetHrAllowed() bool {
	if m != nil {
		return m.HrAllowed
	}
	return false
}

func (m *ServiceProfile) GetRaAllowed() bool {
	if m != nil {
		return m.RaAllowed
	}
	return false
}

func (m *ServiceProfile) GetNwkGeoLoc() bool {
	if m != nil {
		return m.NwkGeoLoc
	}
	return false
}

func (m *ServiceProfile) GetTargetPer() uint32 {
	if m != nil {
		return m.TargetPer
	}
	return 0
}

func (m *ServiceProfile) GetMinGwDiversity() uint32 {
	if m != nil {
		return m.MinGwDiversity
	}
	return 0
}

func (m *ServiceProfile) GetGwsPrivate() bool {
	if m != nil {
		return m.GwsPrivate
	}
	return false
}

type DeviceProfile struct {
	// Device-profile ID (UUID string).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Device-profile name.
	Name string `protobuf:"bytes,21,opt,name=name,proto3" json:"name,omitempty"`
	// Organization ID to which the service-profile is assigned.
	OrganizationId int64 `protobuf:"varint,22,opt,name=organization_id,json=organizationID,proto3" json:"organization_id,omitempty"`
	// Network-server ID on which the service-profile is provisioned.
	NetworkServerId int64 `protobuf:"varint,23,opt,name=network_server_id,json=networkServerID,proto3" json:"network_server_id,omitempty"`
	// End-Device supports Class B.
	SupportsClassB bool `protobuf:"varint,2,opt,name=supports_class_b,json=supportsClassB,proto3" json:"supports_class_b,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class B mode supported).
	ClassBTimeout uint32 `protobuf:"varint,3,opt,name=class_b_timeout,json=classBTimeout,proto3" json:"class_b_timeout,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotPeriod uint32 `protobuf:"varint,4,opt,name=ping_slot_period,json=pingSlotPeriod,proto3" json:"ping_slot_period,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotDr uint32 `protobuf:"varint,5,opt,name=ping_slot_dr,json=pingSlotDR,proto3" json:"ping_slot_dr,omitempty"`
	// Mandatory if class B mode supported.
	PingSlotFreq uint32 `protobuf:"varint,6,opt,name=ping_slot_freq,json=pingSlotFreq,proto3" json:"ping_slot_freq,omitempty"`
	// End-Device supports Class C.
	SupportsClassC bool `protobuf:"varint,7,opt,name=supports_class_c,json=supportsClassC,proto3" json:"supports_class_c,omitempty"`
	// Maximum delay for the End-Device to answer a MAC request or a confirmed DL frame (mandatory if class C mode supported).
	ClassCTimeout uint32 `protobuf:"varint,8,opt,name=class_c_timeout,json=classCTimeout,proto3" json:"class_c_timeout,omitempty"`
	// Version of the LoRaWAN supported by the End-Device.
	MacVersion string `protobuf:"bytes,9,opt,name=mac_version,json=macVersion,proto3" json:"mac_version,omitempty"`
	// Revision of the Regional Parameters document supported by the End-Device.
	RegParamsRevision string `protobuf:"bytes,10,opt,name=reg_params_revision,json=regParamsRevision,proto3" json:"reg_params_revision,omitempty"`
	// Class A RX1 delay (mandatory for ABP).
	RxDelay_1 uint32 `protobuf:"varint,11,opt,name=rx_delay_1,json=rxDelay1,proto3" json:"rx_delay_1,omitempty"`
	// RX1 data rate offset (mandatory for ABP).
	RxDrOffset_1 uint32 `protobuf:"varint,12,opt,name=rx_dr_offset_1,json=rxDROffset1,proto3" json:"rx_dr_offset_1,omitempty"`
	// RX2 data rate (mandatory for ABP).
	RxDatarate_2 uint32 `protobuf:"varint,13,opt,name=rx_datarate_2,json=rxDataRate2,proto3" json:"rx_datarate_2,omitempty"`
	// RX2 channel frequency (mandatory for ABP).
	RxFreq_2 uint32 `protobuf:"varint,14,opt,name=rx_freq_2,json=rxFreq2,proto3" json:"rx_freq_2,omitempty"`
	// List of factory-preset frequencies (mandatory for ABP).
	FactoryPresetFreqs []uint32 `protobuf:"varint,15,rep,packed,name=factory_preset_freqs,json=factoryPresetFreqs,proto3" json:"factory_preset_freqs,omitempty"`
	// Maximum EIRP supported by the End-Device.
	MaxEirp uint32 `protobuf:"varint,16,opt,name=max_eirp,json=maxEIRP,proto3" json:"max_eirp,omitempty"`
	// Maximum duty cycle supported by the End-Device.
	MaxDutyCycle uint32 `protobuf:"varint,17,opt,name=max_duty_cycle,json=maxDutyCycle,proto3" json:"max_duty_cycle,omitempty"`
	// End-Device supports Join (OTAA) or not (ABP).
	SupportsJoin bool `protobuf:"varint,18,opt,name=supports_join,json=supportsJoin,proto3" json:"supports_join,omitempty"`
	// RF region name.
	RfRegion string `protobuf:"bytes,19,opt,name=rf_region,json=rfRegion,proto3" json:"rf_region,omitempty"`
	// End-Device uses 32bit FCnt (mandatory for LoRaWAN 1.0 End-Device).
	Supports_32BitFCnt bool `protobuf:"varint,20,opt,name=supports_32bit_f_cnt,json=supports32BitFCnt,proto3" json:"supports_32bit_f_cnt,omitempty"`
	// Payload codec.
	// Leave blank to disable the codec feature.
	PayloadCodec string `protobuf:"bytes,24,opt,name=payload_codec,json=payloadCodec,proto3" json:"payload_codec,omitempty"`
	// Payload encoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the encoder function.
	PayloadEncoderScript string `protobuf:"bytes,25,opt,name=payload_encoder_script,json=payloadEncoderScript,proto3" json:"payload_encoder_script,omitempty"`
	// Payload decoder script.
	// Depending the codec, it is possible to provide a script which implements
	// the decoder function.
	PayloadDecoderScript string `protobuf:"bytes,26,opt,name=payload_decoder_script,json=payloadDecoderScript,proto3" json:"payload_decoder_script,omitempty"`
	// Geolocation buffer TTL (in seconds).
	// When > 0, uplink RX meta-data will be stored in a buffer so that
	// the meta-data of multiple uplinks can be used for geolocation.
	GeolocBufferTtl uint32 `protobuf:"varint,27,opt,name=geoloc_buffer_ttl,json=geolocBufferTTL,proto3" json:"geoloc_buffer_ttl,omitempty"`
	// Geolocation minimum buffer size.
	// When > 0, geolocation will only be performed when the buffer has
	// at least the given size.
	GeolocMinBufferSize uint32 `protobuf:"varint,28,opt,name=geoloc_min_buffer_size,json=geolocMinBufferSize,proto3" json:"geoloc_min_buffer_size,omitempty"`
	// User defined tags.
	Tags map[string]string `protobuf:"bytes,29,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Uplink interval.
	// This defines the expected uplink interval which the device uses for
	// communication. When the uplink interval has expired and no uplink has
	// been received, the device is considered inactive.
	UplinkInterval *duration.Duration `protobuf:"bytes,30,opt,name=uplink_interval,json=uplinkInterval,proto3" json:"uplink_interval,omitempty"`
	// ADR algorithm ID.
	// In case this is left blank, or is configured to a non-existing ADR
	// algorithm (plugin), then it falls back to 'default'.
	AdrAlgorithmId       string   `protobuf:"bytes,31,opt,name=adr_algorithm_id,json=adrAlgorithmID,proto3" json:"adr_algorithm_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceProfile) Reset()         { *m = DeviceProfile{} }
func (m *DeviceProfile) String() string { return proto.CompactTextString(m) }
func (*DeviceProfile) ProtoMessage()    {}
func (*DeviceProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1c507fa3dbc9903, []int{1}
}

func (m *DeviceProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceProfile.Unmarshal(m, b)
}
func (m *DeviceProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceProfile.Marshal(b, m, deterministic)
}
func (m *DeviceProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceProfile.Merge(m, src)
}
func (m *DeviceProfile) XXX_Size() int {
	return xxx_messageInfo_DeviceProfile.Size(m)
}
func (m *DeviceProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceProfile.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceProfile proto.InternalMessageInfo

func (m *DeviceProfile) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeviceProfile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceProfile) GetOrganizationId() int64 {
	if m != nil {
		return m.OrganizationId
	}
	return 0
}

func (m *DeviceProfile) GetNetworkServerId() int64 {
	if m != nil {
		return m.NetworkServerId
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassB() bool {
	if m != nil {
		return m.SupportsClassB
	}
	return false
}

func (m *DeviceProfile) GetClassBTimeout() uint32 {
	if m != nil {
		return m.ClassBTimeout
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotPeriod() uint32 {
	if m != nil {
		return m.PingSlotPeriod
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotDr() uint32 {
	if m != nil {
		return m.PingSlotDr
	}
	return 0
}

func (m *DeviceProfile) GetPingSlotFreq() uint32 {
	if m != nil {
		return m.PingSlotFreq
	}
	return 0
}

func (m *DeviceProfile) GetSupportsClassC() bool {
	if m != nil {
		return m.SupportsClassC
	}
	return false
}

func (m *DeviceProfile) GetClassCTimeout() uint32 {
	if m != nil {
		return m.ClassCTimeout
	}
	return 0
}

func (m *DeviceProfile) GetMacVersion() string {
	if m != nil {
		return m.MacVersion
	}
	return ""
}

func (m *DeviceProfile) GetRegParamsRevision() string {
	if m != nil {
		return m.RegParamsRevision
	}
	return ""
}

func (m *DeviceProfile) GetRxDelay_1() uint32 {
	if m != nil {
		return m.RxDelay_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDrOffset_1() uint32 {
	if m != nil {
		return m.RxDrOffset_1
	}
	return 0
}

func (m *DeviceProfile) GetRxDatarate_2() uint32 {
	if m != nil {
		return m.RxDatarate_2
	}
	return 0
}

func (m *DeviceProfile) GetRxFreq_2() uint32 {
	if m != nil {
		return m.RxFreq_2
	}
	return 0
}

func (m *DeviceProfile) GetFactoryPresetFreqs() []uint32 {
	if m != nil {
		return m.FactoryPresetFreqs
	}
	return nil
}

func (m *DeviceProfile) GetMaxEirp() uint32 {
	if m != nil {
		return m.MaxEirp
	}
	return 0
}

func (m *DeviceProfile) GetMaxDutyCycle() uint32 {
	if m != nil {
		return m.MaxDutyCycle
	}
	return 0
}

func (m *DeviceProfile) GetSupportsJoin() bool {
	if m != nil {
		return m.SupportsJoin
	}
	return false
}

func (m *DeviceProfile) GetRfRegion() string {
	if m != nil {
		return m.RfRegion
	}
	return ""
}

func (m *DeviceProfile) GetSupports_32BitFCnt() bool {
	if m != nil {
		return m.Supports_32BitFCnt
	}
	return false
}

func (m *DeviceProfile) GetPayloadCodec() string {
	if m != nil {
		return m.PayloadCodec
	}
	return ""
}

func (m *DeviceProfile) GetPayloadEncoderScript() string {
	if m != nil {
		return m.PayloadEncoderScript
	}
	return ""
}

func (m *DeviceProfile) GetPayloadDecoderScript() string {
	if m != nil {
		return m.PayloadDecoderScript
	}
	return ""
}

func (m *DeviceProfile) GetGeolocBufferTtl() uint32 {
	if m != nil {
		return m.GeolocBufferTtl
	}
	return 0
}

func (m *DeviceProfile) GetGeolocMinBufferSize() uint32 {
	if m != nil {
		return m.GeolocMinBufferSize
	}
	return 0
}

func (m *DeviceProfile) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *DeviceProfile) GetUplinkInterval() *duration.Duration {
	if m != nil {
		return m.UplinkInterval
	}
	return nil
}

func (m *DeviceProfile) GetAdrAlgorithmId() string {
	if m != nil {
		return m.AdrAlgorithmId
	}
	return ""
}

func init() {
	proto.RegisterEnum("api.RatePolicy", RatePolicy_name, RatePolicy_value)
	proto.RegisterType((*ServiceProfile)(nil), "api.ServiceProfile")
	proto.RegisterType((*DeviceProfile)(nil), "api.DeviceProfile")
	proto.RegisterMapType((map[string]string)(nil), "api.DeviceProfile.TagsEntry")
}

func init() {
	proto.RegisterFile("as/external/api/profiles.proto", fileDescriptor_a1c507fa3dbc9903)
}

var fileDescriptor_a1c507fa3dbc9903 = []byte{
	// 1258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x9e, 0x93, 0x34, 0xb1, 0xe9, 0xbf, 0x84, 0x49, 0x53, 0x25, 0x6d, 0x53, 0xb7, 0x1d, 0x36,
	0xa3, 0x40, 0xed, 0xc6, 0xd9, 0xd0, 0x75, 0x77, 0x75, 0x9c, 0x16, 0xd9, 0x1a, 0xd4, 0x60, 0x82,
	0x15, 0xd8, 0x0d, 0x71, 0x2c, 0xd2, 0x0a, 0x67, 0x49, 0x54, 0x28, 0xca, 0xb1, 0xfb, 0x3a, 0x7b,
	0xc9, 0x5d, 0x0e, 0x3c, 0x92, 0x13, 0x27, 0xc5, 0xee, 0x77, 0x27, 0x7d, 0x3f, 0x3c, 0xe2, 0xa1,
	0x3e, 0x92, 0xe4, 0x00, 0xd2, 0xae, 0x9c, 0x59, 0x69, 0x62, 0x08, 0xbb, 0x90, 0xa8, 0x6e, 0x62,
	0xf4, 0x58, 0x85, 0x32, 0xed, 0x24, 0x46, 0x5b, 0x4d, 0x57, 0x21, 0x51, 0xfb, 0x07, 0x81, 0xd6,
	0x41, 0x28, 0xbb, 0x08, 0x8d, 0xb2, 0x71, 0x57, 0x64, 0x06, 0xac, 0xd2, 0x71, 0x2e, 0x7a, 0xf1,
	0xcf, 0x3a, 0x69, 0x9c, 0x4b, 0x33, 0x55, 0xbe, 0x1c, 0xe6, 0x76, 0xda, 0x20, 0x2b, 0x4a, 0x78,
	0xa5, 0x56, 0xa9, 0x5d, 0x61, 0x2b, 0x4a, 0x50, 0x4a, 0xd6, 0x62, 0x88, 0xa4, 0xf7, 0x10, 0x11,
	0x7c, 0xa6, 0x3f, 0x92, 0xa6, 0x36, 0x01, 0xc4, 0xea, 0x2b, 0x0e, 0xc6, 0x95, 0xf0, 0x76, 0x5b,
	0xa5, 0xf6, 0x2a, 0x6b, 0x2c, 0xc3, 0xa7, 0x03, 0xfa, 0x8a, 0x6c, 0xc5, 0xd2, 0x5e, 0x6b, 0x33,
	0xe1, 0xa9, 0x34, 0x53, 0x69, 0x9c, 0xf4, 0x11, 0x4a, 0x9b, 0x05, 0x71, 0x8e, 0xf8, 0xe9, 0x80,
	0x3e, 0x22, 0x1b, 0x59, 0xc8, 0x0d, 0x58, 0xe9, 0xad, 0xb4, 0x4a, 0xed, 0x3a, 0x5b, 0xcf, 0x42,
	0x06, 0x56, 0xd2, 0xef, 0x49, 0x23, 0x0b, 0xf9, 0x28, 0xf3, 0x27, 0xd2, 0xf2, 0x54, 0x7d, 0x95,
	0xde, 0x2a, 0xf2, 0xb5, 0x2c, 0xec, 0x23, 0x78, 0xae, 0xbe, 0x4a, 0xfa, 0x33, 0xaa, 0x9c, 0x9d,
	0x27, 0x3a, 0x54, 0xfe, 0xdc, 0x5b, 0x6b, 0x95, 0xda, 0x8d, 0x5e, 0xb3, 0x03, 0x89, 0xea, 0xb8,
	0x81, 0x86, 0x08, 0x3b, 0xdb, 0xed, 0x9b, 0xab, 0x2a, 0x8a, 0xaa, 0x0f, 0xf2, 0xaa, 0xe2, 0xa6,
	0xaa, 0xb8, 0x5b, 0x75, 0x3d, 0xaf, 0x2a, 0xee, 0x55, 0x15, 0x77, 0xab, 0x6e, 0xfc, 0x47, 0x55,
	0xb1, 0x5c, 0xf5, 0x07, 0xd2, 0x04, 0x21, 0x78, 0x70, 0xcd, 0x23, 0x69, 0x41, 0x80, 0x05, 0xaf,
	0xdc, 0x2a, 0xb5, 0xcb, 0xac, 0x0e, 0x42, 0x7c, 0xfc, 0x72, 0x26, 0x2d, 0x0c, 0xc0, 0x02, 0x7d,
	0x4d, 0xb6, 0x85, 0x9c, 0xf2, 0xd4, 0x82, 0xcd, 0x52, 0x6e, 0xe4, 0x15, 0x1f, 0x1b, 0x79, 0xe5,
	0x55, 0xf0, 0x4b, 0x36, 0x85, 0x9c, 0x9e, 0x23, 0xc3, 0xe4, 0xd5, 0x07, 0x23, 0xaf, 0xe8, 0x3b,
	0xb2, 0x67, 0x64, 0xa2, 0x8d, 0xe5, 0x4b, 0xae, 0x11, 0x58, 0x2b, 0xcd, 0xdc, 0x23, 0x58, 0x60,
	0x37, 0x17, 0x0c, 0x16, 0xd6, 0x7e, 0xce, 0xd2, 0xb7, 0xc4, 0xfb, 0xd6, 0x1a, 0x81, 0x09, 0x54,
	0xec, 0x55, 0xd1, 0xf9, 0xf0, 0x9e, 0xf3, 0x0c, 0x49, 0xfa, 0x90, 0xac, 0x0b, 0xc3, 0x23, 0x15,
	0x7b, 0x35, 0xfc, 0xaa, 0x07, 0xc2, 0x9c, 0xdd, 0xc2, 0x30, 0xf3, 0xea, 0x37, 0x30, 0xcc, 0xe8,
	0x73, 0x52, 0xf3, 0x2f, 0x21, 0x8e, 0x65, 0xc8, 0x23, 0x48, 0x27, 0x5e, 0xa3, 0x55, 0x6a, 0xd7,
	0x58, 0xb5, 0xc0, 0xce, 0x20, 0x9d, 0xd0, 0xa7, 0x84, 0x24, 0x86, 0x43, 0x18, 0xea, 0x6b, 0x29,
	0xbc, 0x26, 0xd6, 0xae, 0x24, 0xe6, 0x7d, 0x0e, 0x38, 0xfa, 0xf2, 0x96, 0xde, 0xcc, 0xe9, 0xcb,
	0x65, 0xda, 0xc0, 0x0d, 0xbd, 0x95, 0xd3, 0x06, 0x16, 0xf4, 0x01, 0xa9, 0xc6, 0xd7, 0x13, 0x1e,
	0x48, 0xcd, 0x43, 0xed, 0x7b, 0x34, 0xe7, 0xe3, 0xeb, 0xc9, 0x47, 0xa9, 0x3f, 0x69, 0xdf, 0xd9,
	0x2d, 0x98, 0x40, 0x5a, 0x9e, 0x48, 0xe3, 0x6d, 0xe3, 0xa7, 0x57, 0x72, 0x64, 0x78, 0xc2, 0x68,
	0x9b, 0x6c, 0x46, 0x2a, 0x76, 0xeb, 0x26, 0xd4, 0x54, 0x9a, 0x54, 0xd9, 0xb9, 0xb7, 0x83, 0xa2,
	0x46, 0xa4, 0xe2, 0x8f, 0x5f, 0x06, 0x0b, 0x94, 0x3e, 0x23, 0xd5, 0xe0, 0x3a, 0xe5, 0x89, 0x51,
	0x53, 0xf7, 0x6f, 0x79, 0x58, 0x88, 0x04, 0xd7, 0xe9, 0x30, 0x47, 0x5e, 0xfc, 0x4d, 0x48, 0x7d,
	0x20, 0xff, 0x17, 0xc9, 0x6b, 0x93, 0xcd, 0x34, 0x4b, 0xdc, 0xe2, 0xa6, 0xdc, 0x0f, 0x21, 0x4d,
	0xf9, 0x08, 0x23, 0x58, 0x66, 0x8d, 0x05, 0x7e, 0xec, 0xe0, 0xbe, 0xfb, 0x6f, 0x0b, 0x01, 0xb7,
	0x2a, 0x92, 0x3a, 0xb3, 0x45, 0x16, 0xeb, 0x08, 0xf7, 0x2f, 0x72, 0xd0, 0x8d, 0x98, 0xa8, 0x38,
	0xe0, 0x69, 0xa8, 0xb1, 0x93, 0x4a, 0x0b, 0x8c, 0x63, 0x9d, 0x35, 0x1c, 0x7e, 0x1e, 0x6a, 0x3b,
	0x44, 0x94, 0xb6, 0x48, 0xed, 0x56, 0x29, 0x4c, 0x11, 0x42, 0xb2, 0x50, 0x0d, 0x98, 0x0b, 0xe2,
	0xad, 0x02, 0x7f, 0xff, 0x22, 0x88, 0x0b, 0x0d, 0xfe, 0xfa, 0xdf, 0xce, 0xc1, 0xc7, 0x28, 0xde,
	0x9f, 0xc3, 0xf1, 0xed, 0x1c, 0xfc, 0x9b, 0x39, 0x94, 0x97, 0xe6, 0x70, 0xbc, 0x98, 0xc3, 0x33,
	0x52, 0x8d, 0xc0, 0xe7, 0xb8, 0xa0, 0x3a, 0xc6, 0xcc, 0x55, 0x18, 0x89, 0xc0, 0xff, 0x23, 0x47,
	0x68, 0x87, 0x6c, 0x1b, 0x19, 0xf0, 0x04, 0x0c, 0x44, 0x2e, 0x9c, 0x53, 0x85, 0x42, 0x82, 0xc2,
	0x2d, 0x23, 0x83, 0x21, 0x32, 0xac, 0x20, 0xe8, 0x13, 0x42, 0xcc, 0x8c, 0x0b, 0x19, 0xc2, 0x9c,
	0x1f, 0x62, 0xa8, 0xea, 0xac, 0x6c, 0x66, 0x03, 0x07, 0x1c, 0xd2, 0x97, 0xa4, 0xe1, 0x58, 0xc3,
	0xf5, 0x78, 0x9c, 0x4a, 0xcb, 0x0f, 0x8b, 0x3c, 0x55, 0xcd, 0x6c, 0xc0, 0x3e, 0x23, 0x76, 0x48,
	0x5f, 0x90, 0xba, 0x13, 0x81, 0x05, 0xdc, 0x72, 0x7a, 0x45, 0xb8, 0x9c, 0x06, 0x2c, 0xb8, 0x0d,
	0xa6, 0x47, 0xf7, 0x49, 0xc5, 0xcc, 0xb0, 0x51, 0xbc, 0x87, 0xf9, 0xaa, 0xb3, 0x0d, 0x33, 0x73,
	0x4d, 0xea, 0xd1, 0x37, 0x64, 0x67, 0x0c, 0xbe, 0xd5, 0x66, 0xce, 0x13, 0x23, 0x5d, 0x19, 0xa7,
	0x4b, 0xbd, 0x66, 0x6b, 0xb5, 0x5d, 0x67, 0xb4, 0xe0, 0x86, 0x48, 0x39, 0x47, 0x4a, 0xf7, 0x48,
	0x39, 0x82, 0x19, 0x97, 0xca, 0x24, 0x18, 0xb6, 0x3a, 0xdb, 0x88, 0x60, 0x76, 0x72, 0xca, 0x86,
	0x6e, 0x61, 0x1c, 0x25, 0x32, 0x3b, 0xe7, 0xfe, 0xdc, 0x0f, 0x25, 0xc6, 0xad, 0xce, 0x6a, 0x11,
	0xcc, 0x06, 0x99, 0x9d, 0x1f, 0x3b, 0x8c, 0xbe, 0x24, 0xf5, 0x9b, 0x85, 0xf9, 0x4b, 0xab, 0xb8,
	0xc8, 0x5c, 0x6d, 0x01, 0xfe, 0xa6, 0x55, 0x4c, 0x1f, 0x93, 0x8a, 0x19, 0x73, 0x23, 0x03, 0xd7,
	0xc0, 0x6d, 0x6c, 0x60, 0xd9, 0x8c, 0x19, 0xbe, 0xd3, 0x2e, 0xd9, 0xb9, 0x19, 0xe1, 0xa8, 0x37,
	0x52, 0x96, 0x8f, 0xb9, 0x1f, 0x5b, 0x0c, 0x5e, 0x99, 0x6d, 0x2d, 0xb8, 0xa3, 0x5e, 0x5f, 0xd9,
	0x0f, 0xc7, 0xb1, 0x75, 0x25, 0x13, 0x98, 0x87, 0x1a, 0x04, 0xf7, 0xb5, 0x90, 0x3e, 0xa6, 0xaf,
	0xc2, 0x6a, 0x05, 0x78, 0xec, 0x30, 0xfa, 0x13, 0xd9, 0x5d, 0x88, 0x64, 0xec, 0x64, 0x86, 0xa7,
	0xbe, 0x51, 0x89, 0xf5, 0xf6, 0x50, 0xbd, 0x53, 0xb0, 0x27, 0x39, 0x79, 0x8e, 0xdc, 0xb2, 0x4b,
	0xc8, 0x3b, 0xae, 0xfd, 0x3b, 0xae, 0x81, 0x5c, 0x76, 0xbd, 0x22, 0x5b, 0x81, 0xd4, 0xa1, 0xf6,
	0xf9, 0x28, 0x1b, 0x8f, 0xa5, 0xe1, 0xd6, 0x86, 0xde, 0x63, 0x6c, 0x56, 0x33, 0x27, 0xfa, 0x88,
	0x5f, 0x5c, 0x7c, 0xa2, 0x47, 0x64, 0xb7, 0xd0, 0xba, 0x9d, 0xa6, 0xd0, 0xe3, 0xf9, 0xf3, 0x04,
	0x0d, 0xdb, 0x39, 0x7b, 0xa6, 0xe2, 0xdc, 0x83, 0xc7, 0xd0, 0x1b, 0xb2, 0x66, 0x21, 0x48, 0xbd,
	0xa7, 0xad, 0xd5, 0x76, 0xb5, 0xf7, 0x04, 0x0f, 0x9f, 0x3b, 0x9b, 0x4b, 0xe7, 0x02, 0x82, 0xf4,
	0x24, 0xb6, 0x66, 0xce, 0x50, 0x49, 0xfb, 0xa4, 0x99, 0x25, 0xa1, 0x8a, 0x27, 0x5c, 0xc5, 0x56,
	0x9a, 0x29, 0x84, 0xde, 0x41, 0xab, 0xd4, 0xae, 0xf6, 0xf6, 0x3a, 0xf9, 0x9d, 0xa1, 0xb3, 0xb8,
	0x33, 0x74, 0x06, 0xc5, 0x9d, 0x81, 0x35, 0x72, 0xc7, 0x69, 0x61, 0x70, 0x99, 0x03, 0xe1, 0xf6,
	0xe2, 0x40, 0x1b, 0x65, 0x2f, 0x23, 0xb7, 0xc5, 0x3c, 0xc3, 0x36, 0x34, 0x40, 0x98, 0xf7, 0x0b,
	0xf8, 0x74, 0xb0, 0xff, 0x96, 0x54, 0x6e, 0x3e, 0x80, 0x6e, 0x92, 0xd5, 0x89, 0x9c, 0x17, 0x1b,
	0x9d, 0x7b, 0xa4, 0x3b, 0xe4, 0xc1, 0x14, 0xc2, 0x2c, 0x3f, 0xf8, 0x2b, 0x2c, 0x7f, 0xf9, 0x75,
	0xe5, 0x97, 0xd2, 0xab, 0x16, 0x21, 0x4b, 0xc7, 0x66, 0x99, 0xac, 0x0d, 0xd8, 0xe7, 0xe1, 0xe6,
	0x77, 0xee, 0xe9, 0xec, 0x3d, 0xfb, 0x7d, 0xb3, 0xd4, 0xd7, 0xe4, 0xb9, 0xd2, 0x1d, 0xff, 0x52,
	0x99, 0x24, 0xb5, 0xe0, 0x4f, 0x70, 0xee, 0x90, 0x76, 0x16, 0x77, 0x23, 0xf7, 0xde, 0xaf, 0x17,
	0x6d, 0x48, 0x87, 0x6e, 0x52, 0xc3, 0xd2, 0x9f, 0xef, 0x02, 0x65, 0x2f, 0xb3, 0x51, 0xc7, 0xd7,
	0x51, 0x77, 0x6c, 0x00, 0xe2, 0x59, 0xf7, 0x76, 0x88, 0xd7, 0xd3, 0xa3, 0xd7, 0xee, 0x46, 0x15,
	0xe8, 0xee, 0xf4, 0xa8, 0x7b, 0xef, 0x9e, 0x35, 0x5a, 0xc7, 0xc6, 0x1c, 0xfd, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0x45, 0x0e, 0xfc, 0x4d, 0x81, 0x09, 0x00, 0x00,
}
