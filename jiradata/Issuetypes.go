package jiradata

/////////////////////////////////////////////////////////////////////////
// This Code is Generated by SlipScheme Project:
// https://github.com/coryb/slipscheme
//
// Generated with command:
// slipscheme -dir jiradata -pkg jiradata -overwrite schemas/Project.json
/////////////////////////////////////////////////////////////////////////
//                            DO NOT EDIT                              //
/////////////////////////////////////////////////////////////////////////

// IssueTypes defined from schema:
// {
//   "title": "issueTypes",
//   "type": "array",
//   "items": {
//     "title": "Issue Type",
//     "type": "object",
//     "properties": {
//       "avatarId": {
//         "title": "avatarId",
//         "type": "integer"
//       },
//       "description": {
//         "title": "description",
//         "type": "string"
//       },
//       "iconUrl": {
//         "title": "iconUrl",
//         "type": "string"
//       },
//       "id": {
//         "title": "id",
//         "type": "string"
//       },
//       "name": {
//         "title": "name",
//         "type": "string"
//       },
//       "self": {
//         "title": "self",
//         "type": "string"
//       },
//       "subtask": {
//         "title": "subtask",
//         "type": "boolean"
//       }
//     }
//   }
// }
type IssueTypes []*IssueType
