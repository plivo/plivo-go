# Change Log

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
