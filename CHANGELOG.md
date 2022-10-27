# Change Log

## [7.15.0](https://github.com/plivo/plivo-go/tree/v7.15.0) (2022-10-27)
**Audio Streaming**
- API support for starting, deleting, getting streams on a live call
- XML creation support for stream element

## [7.14.0](https://github.com/plivo/plivo-go/tree/v7.14.0) (2022-10-17)
**Feature - Brandusecase API, 10DLC api enhancements**
- Added Brandusecase API, 10DLC api enhancements

## [7.13.0](https://github.com/plivo/plivo-go/tree/v7.13.0) (2022-10-14)
**Adding new attributes to Account PhoneNumber object**
-Added 3 new keys to AccountPhoneNumber object:`tendlc_registration_status`, `tendlc_campaign_id` and `toll_free_sms_verification` (https://www.plivo.com/docs/numbers/api/account-phone-number#the-accountphonenumber-object)
-Added 3 new filters to AccountPhoneNumber - list all my numbers API:`tendlc_registration_status`, `tendlc_campaign_id` and `toll_free_sms_verification` (https://www.plivo.com/docs/numbers/api/account-phone-number#list-all-my-numbers)

## [7.12.1](https://github.com/plivo/plivo-go/tree/v7.12.1) (2022-09-28)
**Adding more attributes to campaign creation**
- Adding more attributes to campaign creation request

## [7.12.0](https://github.com/plivo/plivo-go/tree/v7.12.0) (2022-08-30)
**Feature - 10DLC api updates**
- Updated 10dlc api with total 15 apis now such as campaign, brand, profile and number link

## [7.11.0](https://github.com/plivo/plivo-go/tree/v7.11.0) (2022-08-01)
**Feature - Token creation**
- `JWT Token Creation API` added functionality to create a new JWT token.

## [7.10.0](https://github.com/plivo/plivo-go/tree/v7.10.0) (2022-07-11)
**Feature - STIR Attestation**
- Add stir attestation param as part of Get CDR and Get live call APIs Response

## [7.9.0](https://github.com/plivo/plivo-go/tree/v7.9.0) (2022-05-05)
**Features - List all recordings and The MultiPartyCall element**
- `fromNumber` and `toNumber` added to filtering param [List all recordings](https://www.plivo.com/docs/voice/api/recording#list-all-recordings)
- `recordMinMemberCount` param added in [Add a participant to a multiparty call using API](https://www.plivo.com/docs/voice/api/multiparty-call/participants#add-a-participant)

## [7.8.0](https://github.com/plivo/plivo-go/tree/v7.8.0) (2022-03-25)
**Features - DialElement**
- `confirmTimeout` parameter added to [The Dial element](https://www.plivo.com/docs/voice/xml/dial/)

## [7.7.2](https://github.com/plivo/plivo-go/tree/v7.7.2) (2022-03-23)
**Bug Fix - Voice**
- Added `Polly.Marlene` to [SSML voices](https://www.plivo.com/docs/voice/concepts/ssml#ssml-voices)

## [7.7.1](https://github.com/plivo/plivo-go/tree/v7.7.1) (2022-03-17)
**Bug Fix - Voice**
- Added `machine_detection_url` and `machine_detection_method` in [Make a call API](https://www.plivo.com/docs/voice/api/call#make-a-call)

## [7.7.0](https://github.com/plivo/plivo-go/tree/v7.7.0) (2022-03-02)
**Bug Fix - Fix go modules**
- Fix the import path for go modules to work

## [7.6.1](https://github.com/plivo/plivo-go/tree/v7.6.1) (2022-02-22)
**Features - ListParticipants**
- Parameter added as member_address in response and mock

## [7.6.0](https://github.com/plivo/plivo-go/tree/v7.6.0) (2022-01-27)
**Features - MPCStartRecording**
- Parameter name change from statusCallBack to recordingCallback

## [7.5.0](https://github.com/plivo/plivo-go/tree/v7.5.0) (2021-12-14)
**Features - Voice**
- Routing SDK traffic through Akamai endpoints for all the [Voice APIs](https://www.plivo.com/docs/voice/api/overview/)

## [7.4.0](https://github.com/plivo/plivo-go/tree/v7.4.0) (2021-12-02)
**Features - Messaging: 10 DLC**
- 10DLC API's for brand and campaign support.

## [7.3.0](https://github.com/plivo/plivo-go/tree/v7.3.0) (2021-11-23)
**Features - Voice: Multiparty calls**
- The [Add Multiparty Call API](https://www.plivo.com/docs/voice/api/multiparty-call/participants#add-a-participant) allows for greater functionality by accepting options like `start recording audio`, `stop recording audio`, and their HTTP methods.
- [Multiparty Calls](https://www.plivo.com/docs/voice/api/multiparty-call/) now has new APIs to `stop` and `play` audio.


## [7.2.2](https://github.com/plivo/plivo-go/tree/v7.2.2) (2021-07-29)
- Removed validation for `ringtimeout` and `delaydial` params in [Start a multi party call](https://www.plivo.com/docs/voice/api/multiparty-call#start-a-new-multiparty-call).

## [7.2.1](https://github.com/plivo/plivo-go/tree/v7.2.1) (2021-07-22)
- Updated default HTTP client request timeout to 5 seconds.

## [7.2.0](https://github.com/plivo/plivo-go/tree/v7.2.0) (2021-07-14)
- Add SDK support for Voice MultiPartyCall APIs and XML.

## [7.1.0](https://github.com/plivo/plivo-go/tree/v7.1.0) (2021-07-13)
- Power pack ID has been included to the response for the [list all messages API](https://www.plivo.com/docs/sms/api/message/list-all-messages/) and the [get message details API](https://www.plivo.com/docs/sms/api/message#retrieve-a-message).
- Support for filtering messages by Power pack ID has been added to the [list all messages API](https://www.plivo.com/docs/sms/api/message#list-all-messages).


## [7.0.0](https://github.com/plivo/plivo-go/tree/v7.0.0) (2021-07-05)
- **BREAKING**: Remove the total_count parameter in meta data for list MDR response

## [6.0.1](https://github.com/plivo/plivo-go/tree/v6.0.1) (2021-07-02)
- Read voice network group from voice pricing
- Fix GetCDR and ListCDR response to include all fields

## [6.0.0](https://github.com/plivo/plivo-go/tree/v6.0.0) (2021-06-29)
- **BREAKING**: Update AddSpeak method signature: remove optional parameters
- Add methods to set SpeakElement attributes

## [5.6.0](https://github.com/plivo/plivo-go/tree/v5.6.0) (2021-06-15)
- Add stir verification param as part of Get CDR and live call APIs

## [5.5.2](https://github.com/plivo/plivo-go/tree/v5.5.2) (2021-04-08)
- Read origination prefix from voice pricing

## [5.5.1](https://github.com/plivo/plivo-go/tree/v5.5.1) (2020-12-16)
- Add SSML utilities

## [5.5.0](https://github.com/plivo/plivo-go/tree/v5.5.0) (2020-11-17)
- Add number_priority support for Powerpack API.

## [5.4.1](https://github.com/plivo/plivo-go/tree/v5.4.1) (2020-11-11)
- Fix send SMS json payload.

## [5.4.0](https://github.com/plivo/plivo-go/tree/v5.4.0) (2020-11-05)
- Add Regulatory Compliance API support.

## [5.3.0](https://github.com/plivo/plivo-go/tree/v5.3.0) (2020-10-31)
- Change lookup API endpoint and response.

## [5.2.0](https://github.com/plivo/plivo-go/tree/v5.2.0) (2020-09-25)
- Add Lookup API support.

## [5.1.0](https://github.com/plivo/plivo-go/tree/v5.1.0) (2020-09-24)
- Add "publicURI" optional param support for Application API.

## [5.0.0](https://github.com/plivo/plivo-go/tree/v5.0.0) (2020-08-19)
- Internal changes in Phlo for MultiPartyCall component.
- **BREAKING**: Rename MultiPartyCall struct to PhloMultiPartyCall.

## [4.9.1](https://github.com/plivo/plivo-go/tree/v4.9.1) (2020-08-10)
- Fix Get Details of a Call API response.

## [4.9.0](https://github.com/plivo/plivo-go/tree/v4.9.0) (2020-08-04)
- Add service type support(SMS/MMS) for Powerpack API.

## [4.8.0](https://github.com/plivo/plivo-go/tree/v4.8.0) (2020-07-23)
- Add retries to multiple regions for voice requests.

## [4.7.1](https://github.com/plivo/plivo-go/tree/v4.7.1) (2020-07-13)
- Fix Call Create & Retrieve Call details API responses.

## [4.7.0](https://github.com/plivo/plivo-go/tree/v4.7.0) (2020-05-28)
- Add JWT helper functions.

## [4.6.0](https://github.com/plivo/plivo-go/tree/v4.6.0) (2020-04-29)
- Add V3 signature helper functions.

## [4.5.0](https://github.com/plivo/plivo-go/tree/v4.5.0) (2020-04-24)
- Add city and mms filter support for Number Search API
- Add city, country and mms into List Number and Number Search API's Response
- Fix for TotalCount in Number API's Response

## [4.4.0](https://github.com/plivo/plivo-go/tree/v4.4.0) (2020-03-31)
- Add application cascade delete support.

## [4.3.0](https://github.com/plivo/plivo-go/tree/v4.3.0) (2020-03-30)
- Add Tollfree support for Powerpack

## [4.2.0](https://github.com/plivo/plivo-go/tree/v4.2.0) (2020-03-27)
- Add post call quality feedback API support.

## [4.1.6](https://github.com/plivo/plivo-go/tree/v4.1.6) (2020-02-25)
- Add Media support

## [4.1.5](https://github.com/plivo/plivo-go/tree/v4.1.5) (2020-01-24)
- Hot fix for go build

## [4.1.4](https://github.com/plivo/plivo-go/tree/v4.1.4) (2019-12-20)
- Add Powerpack support

## [4.1.3](https://github.com/plivo/plivo-go/tree/v4.1.3) (2019-12-04)
- Add MMS support

## [4.1.2](https://github.com/plivo/plivo-go/tree/v4.1.2) (2019-11-13)
- Add GetInput XML support

## [4.1.1](https://github.com/plivo/plivo-go/tree/v4.1.1) (2019-11-05)
- Add SSML support

## [4.1.0](https://github.com/plivo/plivo-go/tree/v4.1.0) (2019-03-12)
- Add PHLO support
- Add Multi-party call triggers

## [4.0.6-beta1](https://github.com/plivo/plivo-go/tree/v4.0.6-beta1) (2019-02-05)
- Add PHLO support in beta
- Add Multi-party call triggers

## [4.0.5](https://github.com/plivo/plivo-go/tree/v4.0.5) (2018-11-19)
- Add hangup party details to CDR. CDR filtering allowed by hangup_source and hangup_cause_code.
- Add sub-account cascade delete support.
- Add call status to GET calls and live-calls methods.

## [4.0.4](https://github.com/plivo/plivo-go/tree/v4.0.4) (2018-10-31)
- Add live calls filtering by to, from numbers and call direction.

## [4.0.3](https://github.com/plivo/plivo-go/tree/v4.0.3) (2018-10-01)
- Added Trackable parameter in messages.

## [4.0.2](https://github.com/plivo/plivo-go/tree/v4.0.2) (2018-09-18)
- Added parent_call_uuid parameter to filter calls.
- Queued status added for filtering calls in queued status.
- Added log_incoming_messages parameter to application create and update.
- Added powerpack support.

## [4.0.1](https://github.com/plivo/plivo-go/tree/v4.0.1) (2018-07-25)
- Fixed caller id retrieval

## [4.0.0](https://github.com/plivo/plivo-go/tree/v4.0.0) (2018-01-18)
- Major restructuring of the repo to bring all go files to repo's root
- Supports v2 signature validation
- A few fixes (#2 & #3)

## [v4.0.0-beta.1](https://github.com/plivo/plivo-go/releases/tag/v4.0.0-beta.1) (2017-10-25)
- The official SDK of Plivo
- Supports all Go versions >= 1.0.x
