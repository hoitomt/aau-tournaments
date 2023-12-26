// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: ee_tournament_list_response.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ExposureEventTournamentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*ExposureEventTournamentListDetail `protobuf:"bytes,1,rep,name=results,json=Results,proto3" json:"results,omitempty"`
}

func (x *ExposureEventTournamentListResponse) Reset() {
	*x = ExposureEventTournamentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ee_tournament_list_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureEventTournamentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureEventTournamentListResponse) ProtoMessage() {}

func (x *ExposureEventTournamentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ee_tournament_list_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureEventTournamentListResponse.ProtoReflect.Descriptor instead.
func (*ExposureEventTournamentListResponse) Descriptor() ([]byte, []int) {
	return file_ee_tournament_list_response_proto_rawDescGZIP(), []int{0}
}

func (x *ExposureEventTournamentListResponse) GetResults() []*ExposureEventTournamentListDetail {
	if x != nil {
		return x.Results
	}
	return nil
}

type ExposureEventTournamentListDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          int32   `protobuf:"varint,1,opt,name=id,json=Id,proto3" json:"id,omitempty"`                                                                                // 196894,
	Name                        string  `protobuf:"bytes,2,opt,name=name,json=Name,proto3" json:"name,omitempty"`                                                                           // "NRYBL",
	OrganizationId              int32   `protobuf:"varint,3,opt,name=organization_id,json=OrganizationId,proto3" json:"organization_id,omitempty"`                                          // 21437,
	OrganizationWebsite         string  `protobuf:"bytes,4,opt,name=organization_website,json=OrganizationWebsite,proto3" json:"organization_website,omitempty"`                            // "http://www.idahoprospects.com",
	OrganizationName            string  `protobuf:"bytes,5,opt,name=organization_name,json=OrganizationName,proto3" json:"organization_name,omitempty"`                                     // "Idaho Prospects Basketball",
	Gender                      int32   `protobuf:"varint,6,opt,name=gender,json=Gender,proto3" json:"gender,omitempty"`                                                                    // 0,
	StartDate                   string  `protobuf:"bytes,7,opt,name=start_date,json=StartDate,proto3" json:"start_date,omitempty"`                                                          // "2023-01-02T00:00:00.000",
	EndDate                     string  `protobuf:"bytes,8,opt,name=end_date,json=EndDate,proto3" json:"end_date,omitempty"`                                                                // "2023-02-15T00:00:00.000",
	Location                    string  `protobuf:"bytes,9,opt,name=location,json=Location,proto3" json:"location,omitempty"`                                                               // null,
	StreetAddress               string  `protobuf:"bytes,10,opt,name=street_address,json=StreetAddress,proto3" json:"street_address,omitempty"`                                             // null,
	ExtendedAddress             string  `protobuf:"bytes,11,opt,name=extended_address,json=ExtendedAddress,proto3" json:"extended_address,omitempty"`                                       // null,
	City                        string  `protobuf:"bytes,12,opt,name=city,json=City,proto3" json:"city,omitempty"`                                                                          // "Laketown",
	StateRegion                 string  `protobuf:"bytes,13,opt,name=state_region,json=StateRegion,proto3" json:"state_region,omitempty"`                                                   // "Utah",
	StateRegionAbbr             string  `protobuf:"bytes,14,opt,name=state_region_abbr,json=StateRegionAbbr,proto3" json:"state_region_abbr,omitempty"`                                     // "UT",
	PostalCode                  string  `protobuf:"bytes,15,opt,name=postal_code,json=PostalCode,proto3" json:"postal_code,omitempty"`                                                      // null,
	Longitude                   float32 `protobuf:"fixed32,16,opt,name=longitude,json=Longitude,proto3" json:"longitude,omitempty"`                                                         // -111.3224256,
	Latitude                    float32 `protobuf:"fixed32,17,opt,name=latitude,json=Latitude,proto3" json:"latitude,omitempty"`                                                            // 41.8254935,
	LogoId                      int32   `protobuf:"varint,18,opt,name=logo_id,json=LogoId,proto3" json:"logo_id,omitempty"`                                                                 // null,
	OrganizationLogoId          int32   `protobuf:"varint,19,opt,name=organization_logo_id,json=OrganizationLogoId,proto3" json:"organization_logo_id,omitempty"`                           // 81728,
	DateCreated                 string  `protobuf:"bytes,20,opt,name=date_created,json=DateCreated,proto3" json:"date_created,omitempty"`                                                   // "2022-12-06T11:02:41.250",
	MarketingState              int32   `protobuf:"varint,21,opt,name=marketing_state,json=MarketingState,proto3" json:"marketing_state,omitempty"`                                         // 2,
	EventType                   int32   `protobuf:"varint,22,opt,name=event_type,json=EventType,proto3" json:"event_type,omitempty"`                                                        // 1,
	Type                        string  `protobuf:"bytes,23,opt,name=type,json=Type,proto3" json:"type,omitempty"`                                                                          // "League",
	Featured                    int32   `protobuf:"varint,24,opt,name=featured,json=Featured,proto3" json:"featured,omitempty"`                                                             // 0,
	ExposureCertified           bool    `protobuf:"varint,25,opt,name=exposure_certified,json=ExposureCertified,proto3" json:"exposure_certified,omitempty"`                                // true,
	ContactName                 string  `protobuf:"bytes,26,opt,name=contact_name,json=ContactName,proto3" json:"contact_name,omitempty"`                                                   // "Matt Jolley",
	ContactEmail                string  `protobuf:"bytes,27,opt,name=contact_email,json=ContactEmail,proto3" json:"contact_email,omitempty"`                                                // "mjolley23@gmail.com",
	ContactPhone                string  `protobuf:"bytes,28,opt,name=contact_phone,json=ContactPhone,proto3" json:"contact_phone,omitempty"`                                                // null,
	SportType                   int32   `protobuf:"varint,29,opt,name=sport_type,json=SportType,proto3" json:"sport_type,omitempty"`                                                        // 1,
	SportHost                   string  `protobuf:"bytes,30,opt,name=sport_host,json=SportHost,proto3" json:"sport_host,omitempty"`                                                         // "basketball.exposureevents.com",
	SportName                   string  `protobuf:"bytes,31,opt,name=sport_name,json=SportName,proto3" json:"sport_name,omitempty"`                                                         // "Basketball",
	Logo                        string  `protobuf:"bytes,32,opt,name=logo,json=Logo,proto3" json:"logo,omitempty"`                                                                          // "https://cdn.exposureevents.com/assets/files/png/81728?v=11097745",
	EnableRegistration          bool    `protobuf:"varint,33,opt,name=enable_registration,json=EnableRegistration,proto3" json:"enable_registration,omitempty"`                             // false,
	GamesGuaranteed             int32   `protobuf:"varint,34,opt,name=games_guaranteed,json=GamesGuaranteed,proto3" json:"games_guaranteed,omitempty"`                                      // null,
	MinCost                     int32   `protobuf:"varint,35,opt,name=min_cost,json=MinCost,proto3" json:"min_cost,omitempty"`                                                              // null,
	MaxCost                     int32   `protobuf:"varint,36,opt,name=max_cost,json=MaxCost,proto3" json:"max_cost,omitempty"`                                                              // null,
	Ability                     string  `protobuf:"bytes,37,opt,name=ability,json=Ability,proto3" json:"ability,omitempty"`                                                                 // null,
	YouthAgeGradesBoth          string  `protobuf:"bytes,39,opt,name=youth_age_grades_both,json=YouthAgeGradesBoth,proto3" json:"youth_age_grades_both,omitempty"`                          // "Boys & Girls: 6th-K",
	DateFormatted               string  `protobuf:"bytes,40,opt,name=date_formatted,json=DateFormatted,proto3" json:"date_formatted,omitempty"`                                             // "Jan 2 - Feb 15, 2023",
	PriceFormatted              string  `protobuf:"bytes,41,opt,name=price_formatted,json=PriceFormatted,proto3" json:"price_formatted,omitempty"`                                          // null,
	RegistrationEnded           bool    `protobuf:"varint,42,opt,name=registration_ended,json=RegistrationEnded,proto3" json:"registration_ended,omitempty"`                                // true,
	StateRegionLink             string  `protobuf:"bytes,43,opt,name=state_region_link,json=StateRegionLink,proto3" json:"state_region_link,omitempty"`                                     // "https://basketball.exposureevents.com/youth-basketball-events/utah",
	CalendarLink                string  `protobuf:"bytes,44,opt,name=calendar_link,json=CalendarLink,proto3" json:"calendar_link,omitempty"`                                                // "https://basketball.exposureevents.com/calendar/event?eventid=196894&expired=1676419200",
	Slug                        string  `protobuf:"bytes,45,opt,name=slug,json=Slug,proto3" json:"slug,omitempty"`                                                                          // "nrybl",
	Link                        string  `protobuf:"bytes,46,opt,name=link,json=Link,proto3" json:"link,omitempty"`                                                                          // "https://basketball.exposureevents.com/196894/nrybl",
	CityState                   string  `protobuf:"bytes,47,opt,name=city_state,json=CityState,proto3" json:"city_state,omitempty"`                                                         // "Laketown, Utah",
	Website                     string  `protobuf:"bytes,48,opt,name=website,json=Website,proto3" json:"website,omitempty"`                                                                 // null,
	RegistrationLink            string  `protobuf:"bytes,49,opt,name=registration_link,json=RegistrationLink,proto3" json:"registration_link,omitempty"`                                    // "https://basketball.exposureevents.com/196894/nrybl/registration",
	ExternalRegistrationWebsite string  `protobuf:"bytes,50,opt,name=external_registration_website,json=ExternalRegistrationWebsite,proto3" json:"external_registration_website,omitempty"` // null,
	ExternalScheduleWebsite     string  `protobuf:"bytes,51,opt,name=external_schedule_website,json=ExternalScheduleWebsite,proto3" json:"external_schedule_website,omitempty"`             // null,
	ShowSchedule                bool    `protobuf:"varint,52,opt,name=show_schedule,json=ShowSchedule,proto3" json:"show_schedule,omitempty"`                                               // true,
	ScheduleLink                string  `protobuf:"bytes,53,opt,name=schedule_link,json=ScheduleLink,proto3" json:"schedule_link,omitempty"`                                                // "https://basketball.exposureevents.com/196894/e/schedule",
	IsSiteUrl                   bool    `protobuf:"varint,54,opt,name=is_site_url,json=IsSiteUrl,proto3" json:"is_site_url,omitempty"`                                                      // false,
	CollegeCoachLink            string  `protobuf:"bytes,55,opt,name=college_coach_link,json=CollegeCoachLink,proto3" json:"college_coach_link,omitempty"`                                  // null,
	GateLink                    string  `protobuf:"bytes,56,opt,name=gate_link,json=GateLink,proto3" json:"gate_link,omitempty"`                                                            // null,
	PaymentsLink                string  `protobuf:"bytes,57,opt,name=payments_link,json=PaymentsLink,proto3" json:"payments_link,omitempty"`                                                // null
}

func (x *ExposureEventTournamentListDetail) Reset() {
	*x = ExposureEventTournamentListDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ee_tournament_list_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExposureEventTournamentListDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExposureEventTournamentListDetail) ProtoMessage() {}

func (x *ExposureEventTournamentListDetail) ProtoReflect() protoreflect.Message {
	mi := &file_ee_tournament_list_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExposureEventTournamentListDetail.ProtoReflect.Descriptor instead.
func (*ExposureEventTournamentListDetail) Descriptor() ([]byte, []int) {
	return file_ee_tournament_list_response_proto_rawDescGZIP(), []int{1}
}

func (x *ExposureEventTournamentListDetail) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetOrganizationId() int32 {
	if x != nil {
		return x.OrganizationId
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetOrganizationWebsite() string {
	if x != nil {
		return x.OrganizationWebsite
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetOrganizationName() string {
	if x != nil {
		return x.OrganizationName
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetExtendedAddress() string {
	if x != nil {
		return x.ExtendedAddress
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetStateRegion() string {
	if x != nil {
		return x.StateRegion
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetStateRegionAbbr() string {
	if x != nil {
		return x.StateRegionAbbr
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetLogoId() int32 {
	if x != nil {
		return x.LogoId
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetOrganizationLogoId() int32 {
	if x != nil {
		return x.OrganizationLogoId
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetDateCreated() string {
	if x != nil {
		return x.DateCreated
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetMarketingState() int32 {
	if x != nil {
		return x.MarketingState
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetFeatured() int32 {
	if x != nil {
		return x.Featured
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetExposureCertified() bool {
	if x != nil {
		return x.ExposureCertified
	}
	return false
}

func (x *ExposureEventTournamentListDetail) GetContactName() string {
	if x != nil {
		return x.ContactName
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetContactEmail() string {
	if x != nil {
		return x.ContactEmail
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetContactPhone() string {
	if x != nil {
		return x.ContactPhone
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetSportType() int32 {
	if x != nil {
		return x.SportType
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetSportHost() string {
	if x != nil {
		return x.SportHost
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetSportName() string {
	if x != nil {
		return x.SportName
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetEnableRegistration() bool {
	if x != nil {
		return x.EnableRegistration
	}
	return false
}

func (x *ExposureEventTournamentListDetail) GetGamesGuaranteed() int32 {
	if x != nil {
		return x.GamesGuaranteed
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetMinCost() int32 {
	if x != nil {
		return x.MinCost
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetMaxCost() int32 {
	if x != nil {
		return x.MaxCost
	}
	return 0
}

func (x *ExposureEventTournamentListDetail) GetAbility() string {
	if x != nil {
		return x.Ability
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetYouthAgeGradesBoth() string {
	if x != nil {
		return x.YouthAgeGradesBoth
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetDateFormatted() string {
	if x != nil {
		return x.DateFormatted
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetPriceFormatted() string {
	if x != nil {
		return x.PriceFormatted
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetRegistrationEnded() bool {
	if x != nil {
		return x.RegistrationEnded
	}
	return false
}

func (x *ExposureEventTournamentListDetail) GetStateRegionLink() string {
	if x != nil {
		return x.StateRegionLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetCalendarLink() string {
	if x != nil {
		return x.CalendarLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetCityState() string {
	if x != nil {
		return x.CityState
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetRegistrationLink() string {
	if x != nil {
		return x.RegistrationLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetExternalRegistrationWebsite() string {
	if x != nil {
		return x.ExternalRegistrationWebsite
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetExternalScheduleWebsite() string {
	if x != nil {
		return x.ExternalScheduleWebsite
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetShowSchedule() bool {
	if x != nil {
		return x.ShowSchedule
	}
	return false
}

func (x *ExposureEventTournamentListDetail) GetScheduleLink() string {
	if x != nil {
		return x.ScheduleLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetIsSiteUrl() bool {
	if x != nil {
		return x.IsSiteUrl
	}
	return false
}

func (x *ExposureEventTournamentListDetail) GetCollegeCoachLink() string {
	if x != nil {
		return x.CollegeCoachLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetGateLink() string {
	if x != nil {
		return x.GateLink
	}
	return ""
}

func (x *ExposureEventTournamentListDetail) GetPaymentsLink() string {
	if x != nil {
		return x.PaymentsLink
	}
	return ""
}

var File_ee_tournament_list_response_proto protoreflect.FileDescriptor

var file_ee_tournament_list_response_proto_rawDesc = []byte{
	0x0a, 0x21, 0x65, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x6a, 0x0a, 0x23, 0x45,
	0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x45, 0x78, 0x70,
	0x6f, 0x73, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x07,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xd8, 0x0f, 0x0a, 0x21, 0x45, 0x78, 0x70, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x14, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2b, 0x0a,
	0x11, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65,
	0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x29, 0x0a, 0x10, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x21,
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x62, 0x62, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x41, 0x62, 0x62, 0x72, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x09, 0x4c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x4c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x12, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x49,
	0x64, 0x12, 0x30, 0x0a, 0x14, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x6f, 0x67, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x12, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67,
	0x6f, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x61, 0x74, 0x65, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x64, 0x12, 0x2d,
	0x0a, 0x12, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x19, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x45, 0x78, 0x70, 0x6f,
	0x73, 0x75, 0x72, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x53, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53,
	0x70, 0x6f, 0x72, 0x74, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x70,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18,
	0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f, 0x12, 0x2f, 0x0a, 0x13, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x21, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10,
	0x67, 0x61, 0x6d, 0x65, 0x73, 0x5f, 0x67, 0x75, 0x61, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64,
	0x18, 0x22, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x73, 0x47, 0x75, 0x61,
	0x72, 0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x63,
	0x6f, 0x73, 0x74, 0x18, 0x23, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x69, 0x6e, 0x43, 0x6f,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x24,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4d, 0x61, 0x78, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x31, 0x0a, 0x15, 0x79, 0x6f, 0x75, 0x74, 0x68,
	0x5f, 0x61, 0x67, 0x65, 0x5f, 0x67, 0x72, 0x61, 0x64, 0x65, 0x73, 0x5f, 0x62, 0x6f, 0x74, 0x68,
	0x18, 0x27, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x59, 0x6f, 0x75, 0x74, 0x68, 0x41, 0x67, 0x65,
	0x47, 0x72, 0x61, 0x64, 0x65, 0x73, 0x42, 0x6f, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x44, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65,
	0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x74, 0x65, 0x64, 0x18, 0x29, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x64, 0x65, 0x64,
	0x18, 0x2a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x64, 0x65, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x2b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x65, 0x6e, 0x64, 0x61,
	0x72, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x43, 0x61,
	0x6c, 0x65, 0x6e, 0x64, 0x61, 0x72, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6c,
	0x75, 0x67, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x12,
	0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x30, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x6e, 0x6b,
	0x18, 0x31, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x42, 0x0a, 0x1d, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1b, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x19,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x17, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x57, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x6f, 0x77,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x53, 0x68, 0x6f, 0x77, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x35,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x4c, 0x69,
	0x6e, 0x6b, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x36, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x53, 0x69, 0x74, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x61, 0x63, 0x68, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x37, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x67, 0x65, 0x43, 0x6f, 0x61, 0x63, 0x68, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x38, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x39,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69,
	0x6e, 0x6b, 0x42, 0x14, 0x5a, 0x12, 0x65, 0x78, 0x70, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x2d, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ee_tournament_list_response_proto_rawDescOnce sync.Once
	file_ee_tournament_list_response_proto_rawDescData = file_ee_tournament_list_response_proto_rawDesc
)

func file_ee_tournament_list_response_proto_rawDescGZIP() []byte {
	file_ee_tournament_list_response_proto_rawDescOnce.Do(func() {
		file_ee_tournament_list_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_ee_tournament_list_response_proto_rawDescData)
	})
	return file_ee_tournament_list_response_proto_rawDescData
}

var file_ee_tournament_list_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ee_tournament_list_response_proto_goTypes = []interface{}{
	(*ExposureEventTournamentListResponse)(nil), // 0: protos.ExposureEventTournamentListResponse
	(*ExposureEventTournamentListDetail)(nil),   // 1: protos.ExposureEventTournamentListDetail
}
var file_ee_tournament_list_response_proto_depIdxs = []int32{
	1, // 0: protos.ExposureEventTournamentListResponse.results:type_name -> protos.ExposureEventTournamentListDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ee_tournament_list_response_proto_init() }
func file_ee_tournament_list_response_proto_init() {
	if File_ee_tournament_list_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ee_tournament_list_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureEventTournamentListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ee_tournament_list_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExposureEventTournamentListDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ee_tournament_list_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ee_tournament_list_response_proto_goTypes,
		DependencyIndexes: file_ee_tournament_list_response_proto_depIdxs,
		MessageInfos:      file_ee_tournament_list_response_proto_msgTypes,
	}.Build()
	File_ee_tournament_list_response_proto = out.File
	file_ee_tournament_list_response_proto_rawDesc = nil
	file_ee_tournament_list_response_proto_goTypes = nil
	file_ee_tournament_list_response_proto_depIdxs = nil
}
