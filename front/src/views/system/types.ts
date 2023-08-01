import {
  System__Menu__GetListResponseData,
  System__Role__GetListResponseData,
  System__Operation__GetManyResponseData,
  System__User__GetListResponseData
} from "@/api/models";

export type User = NonNullable<
  System__User__GetListResponseData["data"]
>[number];

export type Role = NonNullable<
  System__Role__GetListResponseData["data"]
>[number];

export type Menu = NonNullable<
  System__Menu__GetListResponseData["data"]
>[number];

export type API = NonNullable<
  System__Operation__GetManyResponseData["data"]
>[number];
