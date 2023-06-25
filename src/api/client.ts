import {
  Client,
  ClientConfig,
  CreateClientConfig,
  User,
  OperationMetadata,
  OperationsDefinition,
  OperationRequestOptions,
  SubscriptionRequestOptions,
  SubscriptionEventHandler,
  FetchUserRequestOptions
} from "fireboom-wundersdk/client";

import type { CustomClaims, Role } from "./claims";
import type {
  BindRoleApisInput,
  BindRoleApisResponse,
  BindRoleApisResponseData,
  GetAllApisResponse,
  GetAllApisResponseData,
  GetAllRolesResponse,
  GetAllRolesResponseData,
  GetRoleBindApisInput,
  GetRoleBindApisResponse,
  GetRoleBindApisResponseData,
  Post__CreateOneInput,
  Post__CreateOneResponse,
  Post__CreateOneResponseData,
  Post__DeleteManyInput,
  Post__DeleteManyResponse,
  Post__DeleteManyResponseData,
  Post__DeleteOneInput,
  Post__DeleteOneResponse,
  Post__DeleteOneResponseData,
  Post__GetListInput,
  Post__GetListResponse,
  Post__GetListResponseData,
  Post__GetOneInput,
  Post__GetOneResponse,
  Post__GetOneResponseData,
  Post__UpdateOneInput,
  Post__UpdateOneResponse,
  Post__UpdateOneResponseData,
  Statistics__MonthlySalesResponse,
  Statistics__MonthlySalesResponseData,
  Statistics__SaleTypePercentResponse,
  Statistics__SaleTypePercentResponseData,
  Statistics__SalesTop10Response,
  Statistics__SalesTop10ResponseData,
  Statistics__VisitTrendingResponse,
  Statistics__VisitTrendingResponseData,
  System__Menu__CreateOneInput,
  System__Menu__CreateOneResponse,
  System__Menu__CreateOneResponseData,
  System__Menu__DeleteManyInput,
  System__Menu__DeleteManyResponse,
  System__Menu__DeleteManyResponseData,
  System__Menu__DeleteOneInput,
  System__Menu__DeleteOneResponse,
  System__Menu__DeleteOneResponseData,
  System__Menu__GetListInput,
  System__Menu__GetListResponse,
  System__Menu__GetListResponseData,
  System__Menu__GetManyResponse,
  System__Menu__GetManyResponseData,
  System__Menu__GetOneInput,
  System__Menu__GetOneResponse,
  System__Menu__GetOneResponseData,
  System__Menu__UpdateOneInput,
  System__Menu__UpdateOneResponse,
  System__Menu__UpdateOneResponseData,
  System__Operation__GetManyResponse,
  System__Operation__GetManyResponseData,
  System__Role__AddOneInput,
  System__Role__AddOneResponse,
  System__Role__AddOneResponseData,
  System__Role__BindRoleApisInput,
  System__Role__BindRoleApisResponse,
  System__Role__BindRoleApisResponseData,
  System__Role__ConnectOneMenuInput,
  System__Role__ConnectOneMenuResponse,
  System__Role__ConnectOneMenuResponseData,
  System__Role__DeleteManyInput,
  System__Role__DeleteManyResponse,
  System__Role__DeleteManyResponseData,
  System__Role__DeleteOneInput,
  System__Role__DeleteOneResponse,
  System__Role__DeleteOneResponseData,
  System__Role__DisconnectOneMenuInput,
  System__Role__DisconnectOneMenuResponse,
  System__Role__DisconnectOneMenuResponseData,
  System__Role__GetBindMenusInput,
  System__Role__GetBindMenusResponse,
  System__Role__GetBindMenusResponseData,
  System__Role__GetListInput,
  System__Role__GetListResponse,
  System__Role__GetListResponseData,
  System__Role__GetManyResponse,
  System__Role__GetManyResponseData,
  System__Role__GetRoleBindApisInput,
  System__Role__GetRoleBindApisResponse,
  System__Role__GetRoleBindApisResponseData,
  System__Role__UpdateOneInput,
  System__Role__UpdateOneResponse,
  System__Role__UpdateOneResponseData,
  System__User__ConnectRoleInput,
  System__User__ConnectRoleResponse,
  System__User__ConnectRoleResponseData,
  System__User__CountUsersResponse,
  System__User__CountUsersResponseData,
  System__User__CreateOneInput,
  System__User__CreateOneResponse,
  System__User__CreateOneResponseData,
  System__User__DisconnectRoleInput,
  System__User__DisconnectRoleResponse,
  System__User__DisconnectRoleResponseData,
  System__User__GetListInput,
  System__User__GetListResponse,
  System__User__GetListResponseData,
  System__User__GetOneInput,
  System__User__GetOneResponse,
  System__User__GetOneResponseData,
  System__User__GetUserRoleInput,
  System__User__GetUserRoleResponse,
  System__User__GetUserRoleResponseData,
  User__MeResponse,
  User__MeResponseData
} from "./models";

export const WUNDERGRAPH_S3_ENABLED = false;
export const WUNDERGRAPH_AUTH_ENABLED = true;

export enum AuthProviderId {
  "auth0" = "auth0"
}

export interface AuthProvider {
  id: AuthProviderId;
  login: (redirectURI?: string) => void;
}

export const defaultClientConfig: ClientConfig = {
  applicationHash: "49024269",
  baseURL: "http://localhost:9991/",
  sdkVersion: ""
};

export const operationMetadata: OperationMetadata = {
  BindRoleApis: {
    requiresAuthentication: false
  },
  GetAllApis: {
    requiresAuthentication: false
  },
  GetAllRoles: {
    requiresAuthentication: false
  },
  GetRoleBindApis: {
    requiresAuthentication: false
  },
  "Post/CreateOne": {
    requiresAuthentication: true
  },
  "Post/DeleteMany": {
    requiresAuthentication: false
  },
  "Post/DeleteOne": {
    requiresAuthentication: false
  },
  "Post/GetList": {
    requiresAuthentication: false
  },
  "Post/GetOne": {
    requiresAuthentication: false
  },
  "Post/UpdateOne": {
    requiresAuthentication: false
  },
  "Statistics/MonthlySales": {
    requiresAuthentication: false
  },
  "Statistics/SaleTypePercent": {
    requiresAuthentication: false
  },
  "Statistics/SalesTop10": {
    requiresAuthentication: false
  },
  "Statistics/VisitTrending": {
    requiresAuthentication: false
  },
  "System/Menu/CreateOne": {
    requiresAuthentication: false
  },
  "System/Menu/DeleteMany": {
    requiresAuthentication: false
  },
  "System/Menu/DeleteOne": {
    requiresAuthentication: false
  },
  "System/Menu/GetList": {
    requiresAuthentication: false
  },
  "System/Menu/GetMany": {
    requiresAuthentication: false
  },
  "System/Menu/GetOne": {
    requiresAuthentication: false
  },
  "System/Menu/UpdateOne": {
    requiresAuthentication: false
  },
  "System/Operation/GetMany": {
    requiresAuthentication: false
  },
  "System/Role/AddOne": {
    requiresAuthentication: false
  },
  "System/Role/BindRoleApis": {
    requiresAuthentication: false
  },
  "System/Role/ConnectOneMenu": {
    requiresAuthentication: false
  },
  "System/Role/DeleteMany": {
    requiresAuthentication: false
  },
  "System/Role/DeleteOne": {
    requiresAuthentication: false
  },
  "System/Role/DisconnectOneMenu": {
    requiresAuthentication: false
  },
  "System/Role/GetBindMenus": {
    requiresAuthentication: false
  },
  "System/Role/GetList": {
    requiresAuthentication: false
  },
  "System/Role/GetMany": {
    requiresAuthentication: false
  },
  "System/Role/GetRoleBindApis": {
    requiresAuthentication: false
  },
  "System/Role/UpdateOne": {
    requiresAuthentication: false
  },
  "System/User/ConnectRole": {
    requiresAuthentication: false
  },
  "System/User/CountUsers": {
    requiresAuthentication: false
  },
  "System/User/CreateOne": {
    requiresAuthentication: false
  },
  "System/User/DisconnectRole": {
    requiresAuthentication: false
  },
  "System/User/GetList": {
    requiresAuthentication: false
  },
  "System/User/GetOne": {
    requiresAuthentication: false
  },
  "System/User/GetUserRole": {
    requiresAuthentication: false
  },
  "User/Me": {
    requiresAuthentication: true
  }
};

export class WunderGraphClient extends Client {
  query<
    OperationName extends Extract<keyof Operations["queries"], string>,
    Input extends Operations["queries"][OperationName]["input"] = Operations["queries"][OperationName]["input"],
    Data extends Operations["queries"][OperationName]["data"] = Operations["queries"][OperationName]["data"]
  >(
    options: OperationName extends string
      ? OperationRequestOptions<OperationName, Input>
      : OperationRequestOptions
  ) {
    return super.query<OperationRequestOptions, Data>(options);
  }
  mutate<
    OperationName extends Extract<keyof Operations["mutations"], string>,
    Input extends Operations["mutations"][OperationName]["input"] = Operations["mutations"][OperationName]["input"],
    Data extends Operations["mutations"][OperationName]["data"] = Operations["mutations"][OperationName]["data"]
  >(
    options: OperationName extends string
      ? OperationRequestOptions<OperationName, Input>
      : OperationRequestOptions
  ) {
    return super.mutate<OperationRequestOptions, Data>(options);
  }
  subscribe<
    OperationName extends Extract<keyof Operations["subscriptions"], string>,
    Input extends Operations["subscriptions"][OperationName]["input"] = Operations["subscriptions"][OperationName]["input"],
    Data extends Operations["subscriptions"][OperationName]["data"] = Operations["subscriptions"][OperationName]["data"]
  >(
    options: OperationName extends string
      ? SubscriptionRequestOptions<OperationName, Input>
      : SubscriptionRequestOptions,
    cb: SubscriptionEventHandler<Data>
  ) {
    return super.subscribe(options, cb);
  }
  public login(
    authProviderID: Operations["authProvider"],
    redirectURI?: string
  ) {
    return super.login(authProviderID, redirectURI);
  }
  public async fetchUser<TUser extends User = User<Role, CustomClaims>>(
    options?: FetchUserRequestOptions
  ) {
    return super.fetchUser<TUser>(options);
  }
}

export const createClient = (config?: CreateClientConfig) => {
  return new WunderGraphClient({
    ...defaultClientConfig,
    ...config,
    operationMetadata,
    csrfEnabled: true
  });
};

export type Queries = {
  GetAllApis: {
    input?: undefined;
    data: GetAllApisResponseData;
    requiresAuthentication: false;
  };
  GetAllRoles: {
    input?: undefined;
    data: GetAllRolesResponseData;
    requiresAuthentication: false;
  };
  GetRoleBindApis: {
    input: GetRoleBindApisInput;
    data: GetRoleBindApisResponseData;
    requiresAuthentication: false;
  };
  "Post/GetList": {
    input: Post__GetListInput;
    data: Post__GetListResponseData;
    requiresAuthentication: false;
  };
  "Post/GetOne": {
    input: Post__GetOneInput;
    data: Post__GetOneResponseData;
    requiresAuthentication: false;
  };
  "Statistics/MonthlySales": {
    input?: undefined;
    data: Statistics__MonthlySalesResponseData;
    requiresAuthentication: false;
  };
  "Statistics/SaleTypePercent": {
    input?: undefined;
    data: Statistics__SaleTypePercentResponseData;
    requiresAuthentication: false;
  };
  "Statistics/SalesTop10": {
    input?: undefined;
    data: Statistics__SalesTop10ResponseData;
    requiresAuthentication: false;
  };
  "Statistics/VisitTrending": {
    input?: undefined;
    data: Statistics__VisitTrendingResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/GetList": {
    input: System__Menu__GetListInput;
    data: System__Menu__GetListResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/GetMany": {
    input?: undefined;
    data: System__Menu__GetManyResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/GetOne": {
    input: System__Menu__GetOneInput;
    data: System__Menu__GetOneResponseData;
    requiresAuthentication: false;
  };
  "System/Operation/GetMany": {
    input?: undefined;
    data: System__Operation__GetManyResponseData;
    requiresAuthentication: false;
  };
  "System/Role/GetBindMenus": {
    input: System__Role__GetBindMenusInput;
    data: System__Role__GetBindMenusResponseData;
    requiresAuthentication: false;
  };
  "System/Role/GetList": {
    input: System__Role__GetListInput;
    data: System__Role__GetListResponseData;
    requiresAuthentication: false;
  };
  "System/Role/GetMany": {
    input?: undefined;
    data: System__Role__GetManyResponseData;
    requiresAuthentication: false;
  };
  "System/Role/GetRoleBindApis": {
    input: System__Role__GetRoleBindApisInput;
    data: System__Role__GetRoleBindApisResponseData;
    requiresAuthentication: false;
  };
  "System/User/CountUsers": {
    input?: undefined;
    data: System__User__CountUsersResponseData;
    requiresAuthentication: false;
  };
  "System/User/GetList": {
    input: System__User__GetListInput;
    data: System__User__GetListResponseData;
    requiresAuthentication: false;
  };
  "System/User/GetOne": {
    input: System__User__GetOneInput;
    data: System__User__GetOneResponseData;
    requiresAuthentication: false;
  };
  "System/User/GetUserRole": {
    input: System__User__GetUserRoleInput;
    data: System__User__GetUserRoleResponseData;
    requiresAuthentication: false;
  };
  "User/Me": {
    input?: undefined;
    data: User__MeResponseData;
    requiresAuthentication: true;
  };
};

export type Mutations = {
  BindRoleApis: {
    input: BindRoleApisInput;
    data: BindRoleApisResponseData;
    requiresAuthentication: false;
  };
  "Post/CreateOne": {
    input: Post__CreateOneInput;
    data: Post__CreateOneResponseData;
    requiresAuthentication: true;
  };
  "Post/DeleteMany": {
    input: Post__DeleteManyInput;
    data: Post__DeleteManyResponseData;
    requiresAuthentication: false;
  };
  "Post/DeleteOne": {
    input: Post__DeleteOneInput;
    data: Post__DeleteOneResponseData;
    requiresAuthentication: false;
  };
  "Post/UpdateOne": {
    input: Post__UpdateOneInput;
    data: Post__UpdateOneResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/CreateOne": {
    input: System__Menu__CreateOneInput;
    data: System__Menu__CreateOneResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/DeleteMany": {
    input: System__Menu__DeleteManyInput;
    data: System__Menu__DeleteManyResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/DeleteOne": {
    input: System__Menu__DeleteOneInput;
    data: System__Menu__DeleteOneResponseData;
    requiresAuthentication: false;
  };
  "System/Menu/UpdateOne": {
    input: System__Menu__UpdateOneInput;
    data: System__Menu__UpdateOneResponseData;
    requiresAuthentication: false;
  };
  "System/Role/AddOne": {
    input: System__Role__AddOneInput;
    data: System__Role__AddOneResponseData;
    requiresAuthentication: false;
  };
  "System/Role/BindRoleApis": {
    input: System__Role__BindRoleApisInput;
    data: System__Role__BindRoleApisResponseData;
    requiresAuthentication: false;
  };
  "System/Role/ConnectOneMenu": {
    input: System__Role__ConnectOneMenuInput;
    data: System__Role__ConnectOneMenuResponseData;
    requiresAuthentication: false;
  };
  "System/Role/DeleteMany": {
    input: System__Role__DeleteManyInput;
    data: System__Role__DeleteManyResponseData;
    requiresAuthentication: false;
  };
  "System/Role/DeleteOne": {
    input: System__Role__DeleteOneInput;
    data: System__Role__DeleteOneResponseData;
    requiresAuthentication: false;
  };
  "System/Role/DisconnectOneMenu": {
    input: System__Role__DisconnectOneMenuInput;
    data: System__Role__DisconnectOneMenuResponseData;
    requiresAuthentication: false;
  };
  "System/Role/UpdateOne": {
    input: System__Role__UpdateOneInput;
    data: System__Role__UpdateOneResponseData;
    requiresAuthentication: false;
  };
  "System/User/ConnectRole": {
    input: System__User__ConnectRoleInput;
    data: System__User__ConnectRoleResponseData;
    requiresAuthentication: false;
  };
  "System/User/CreateOne": {
    input: System__User__CreateOneInput;
    data: System__User__CreateOneResponseData;
    requiresAuthentication: false;
  };
  "System/User/DisconnectRole": {
    input: System__User__DisconnectRoleInput;
    data: System__User__DisconnectRoleResponseData;
    requiresAuthentication: false;
  };
};

export type Subscriptions = {};

export type LiveQueries = {};

export type Operations = OperationsDefinition<
  Queries,
  Mutations,
  Subscriptions,
  Role,
  {},
  keyof typeof AuthProviderId
>;
