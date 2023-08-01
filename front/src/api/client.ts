import {
    Client,
    ClientConfig,
    CreateClientConfig,
    User,
    UploadRequestOptions,
    OperationMetadata,
    OperationsDefinition,
    OperationRequestOptions,
    SubscriptionRequestOptions,
    SubscriptionEventHandler,
    FetchUserRequestOptions,
    UploadValidationOptions,
    ExtractProfileName,
    ExtractMeta
} from "fireboom-wundersdk/client";

import type { CustomClaims, Role } from './claims'
import type { Casdoor__GetRolesByIdInput,Casdoor__GetRolesByIdResponse,            Casdoor__GetRolesByIdResponseData,Casdoor__LoginInput,Casdoor__LoginResponse,            Casdoor__LoginResponseData,Casdoor__RefreshTokenInput,Casdoor__RefreshTokenResponse,            Casdoor__RefreshTokenResponseData,Casdoor__SendCodeInput,Casdoor__SendCodeResponse,            Casdoor__SendCodeResponseData,Casdoor__UpdateSMSProviderInput,Casdoor__UpdateSMSProviderResponse,            Casdoor__UpdateSMSProviderResponseData,Post__CreateOneInput,Post__CreateOneResponse,            Post__CreateOneResponseData,Post__DeleteManyInput,Post__DeleteManyResponse,            Post__DeleteManyResponseData,Post__DeleteOneInput,Post__DeleteOneResponse,            Post__DeleteOneResponseData,Post__GetListInput,Post__GetListResponse,            Post__GetListResponseData,Post__GetOneInput,Post__GetOneResponse,            Post__GetOneResponseData,Post__UpdateOneInput,Post__UpdateOneResponse,            Post__UpdateOneResponseData,Statistics__MonthlySalesResponse,            Statistics__MonthlySalesResponseData,Statistics__QueryRawInput,Statistics__QueryRawResponse,            Statistics__QueryRawResponseData,Statistics__SaleTypePercentResponse,            Statistics__SaleTypePercentResponseData,Statistics__SalesTop10Response,            Statistics__SalesTop10ResponseData,Statistics__VisitTrendingResponse,            Statistics__VisitTrendingResponseData,System__GetMenusInput,System__GetMenusResponse,            System__GetMenusResponseData,System__Log__ChangeOpenInput,System__Log__ChangeOpenResponse,            System__Log__ChangeOpenResponseData,System__Log__CreateLogInput,System__Log__CreateLogResponse,            System__Log__CreateLogResponseData,System__Log__GetAllLogResponse,            System__Log__GetAllLogResponseData,System__Log__GetIsOpenResponse,            System__Log__GetIsOpenResponseData,System__Log__GetLogInput,System__Log__GetLogResponse,            System__Log__GetLogResponseData,System__Menu__CreateOneInput,System__Menu__CreateOneResponse,            System__Menu__CreateOneResponseData,System__Menu__DeleteManyInput,System__Menu__DeleteManyResponse,            System__Menu__DeleteManyResponseData,System__Menu__DeleteOneInput,System__Menu__DeleteOneResponse,            System__Menu__DeleteOneResponseData,System__Menu__GetApiListResponse,            System__Menu__GetApiListResponseData,System__Menu__GetApisByMenusInput,System__Menu__GetApisByMenusResponse,            System__Menu__GetApisByMenusResponseData,System__Menu__GetChildrenMenusInput,System__Menu__GetChildrenMenusResponse,            System__Menu__GetChildrenMenusResponseData,System__Menu__GetListInput,System__Menu__GetListResponse,            System__Menu__GetListResponseData,System__Menu__GetManyResponse,            System__Menu__GetManyResponseData,System__Menu__GetMenuByLevelOrPidInput,System__Menu__GetMenuByLevelOrPidResponse,            System__Menu__GetMenuByLevelOrPidResponseData,System__Menu__GetOneInput,System__Menu__GetOneResponse,            System__Menu__GetOneResponseData,System__Menu__UpdateOneInput,System__Menu__UpdateOneResponse,            System__Menu__UpdateOneResponseData,System__Operation__GetManyResponse,            System__Operation__GetManyResponseData,System__Perm__CreateManyInput,System__Perm__CreateManyResponse,            System__Perm__CreateManyResponseData,System__Perm__GetBindPermsResponse,            System__Perm__GetBindPermsResponseData,System__Perm__GetRolePermsInput,System__Perm__GetRolePermsResponse,            System__Perm__GetRolePermsResponseData,System__Role__AddOneInput,System__Role__AddOneResponse,            System__Role__AddOneResponseData,System__Role__BindMenusInput,System__Role__BindMenusResponse,            System__Role__BindMenusResponseData,System__Role__BindRoleApisInput,System__Role__BindRoleApisResponse,            System__Role__BindRoleApisResponseData,System__Role__DeleteManyInput,System__Role__DeleteManyResponse,            System__Role__DeleteManyResponseData,System__Role__DeleteOneInput,System__Role__DeleteOneResponse,            System__Role__DeleteOneResponseData,System__Role__GetListInput,System__Role__GetListResponse,            System__Role__GetListResponseData,System__Role__GetManyResponse,            System__Role__GetManyResponseData,System__Role__GetOneInput,System__Role__GetOneResponse,            System__Role__GetOneResponseData,System__Role__GetRoleMenuIdInput,System__Role__GetRoleMenuIdResponse,            System__Role__GetRoleMenuIdResponseData,System__Role__UnBindMenusInput,System__Role__UnBindMenusResponse,            System__Role__UnBindMenusResponseData,System__Role__UnBindRoleApisInput,System__Role__UnBindRoleApisResponse,            System__Role__UnBindRoleApisResponseData,System__Role__UpdateOneInput,System__Role__UpdateOneResponse,            System__Role__UpdateOneResponseData,System__Sub__CreateSubInput,System__Sub__CreateSubResponse,            System__Sub__CreateSubResponseData,System__Sub__SSEInput,System__Sub__SSEResponse,            System__Sub__SSEResponseData,System__User__ConnectRoleInput,System__User__ConnectRoleResponse,            System__User__ConnectRoleResponseData,System__User__CountUsersResponse,            System__User__CountUsersResponseData,System__User__CreateOneInput,System__User__CreateOneResponse,            System__User__CreateOneResponseData,System__User__DeleteOneInput,System__User__DeleteOneResponse,            System__User__DeleteOneResponseData,System__User__DisconnectRoleInput,System__User__DisconnectRoleResponse,            System__User__DisconnectRoleResponseData,System__User__GetAllListResponse,            System__User__GetAllListResponseData,System__User__GetLikeUserInput,System__User__GetLikeUserResponse,            System__User__GetLikeUserResponseData,System__User__GetListInput,System__User__GetListResponse,            System__User__GetListResponseData,System__User__GetOneInput,System__User__GetOneResponse,            System__User__GetOneResponseData,System__User__GetRoleUsersInput,System__User__GetRoleUsersResponse,            System__User__GetRoleUsersResponseData,System__User__GetUserRoleInput,System__User__GetUserRoleResponse,            System__User__GetUserRoleResponseData,System__User__UpdateOneInput,System__User__UpdateOneResponse,            System__User__UpdateOneResponseData, } from './models'

export const WUNDERGRAPH_S3_ENABLED = true
export const WUNDERGRAPH_AUTH_ENABLED = false

export interface UploadResponse { key: string }

// TODO: missing upload profiles

type S3Providers ={
    tengxunyun: {
    hasProfiles: false,
    profiles: {
    }
    }
}

const S3UploadProviderData: { [provider: string]: { [profile: string]: UploadValidationOptions } } = {
    tengxunyun: {
    },
}


export const defaultClientConfig: ClientConfig = {
    applicationHash: "2109afdc",
    baseURL: "http://localhost:9991",
    sdkVersion: "",
}

export const operationMetadata: OperationMetadata = {
    "Casdoor/GetRolesById": {
        requiresAuthentication: false
		}
    ,
    "Casdoor/Login": {
        requiresAuthentication: false
		}
    ,
    "Casdoor/RefreshToken": {
        requiresAuthentication: false
		}
    ,
    "Casdoor/SendCode": {
        requiresAuthentication: false
		}
    ,
    "Casdoor/UpdateSMSProvider": {
        requiresAuthentication: true
		}
    ,
    "Post/CreateOne": {
        requiresAuthentication: true
		}
    ,
    "Post/DeleteMany": {
        requiresAuthentication: true
		}
    ,
    "Post/DeleteOne": {
        requiresAuthentication: true
		}
    ,
    "Post/GetList": {
        requiresAuthentication: true
		}
    ,
    "Post/GetOne": {
        requiresAuthentication: true
		}
    ,
    "Post/UpdateOne": {
        requiresAuthentication: true
		}
    ,
    "Statistics/MonthlySales": {
        requiresAuthentication: false
		}
    ,
    "Statistics/QueryRaw": {
        requiresAuthentication: false
		}
    ,
    "Statistics/SaleTypePercent": {
        requiresAuthentication: false
		}
    ,
    "Statistics/SalesTop10": {
        requiresAuthentication: false
		}
    ,
    "Statistics/VisitTrending": {
        requiresAuthentication: false
		}
    ,
    "System/GetMenus": {
        requiresAuthentication: true
		}
    ,
    "System/Log/ChangeOpen": {
        requiresAuthentication: true
		}
    ,
    "System/Log/CreateLog": {
        requiresAuthentication: false
		}
    ,
    "System/Log/GetAllLog": {
        requiresAuthentication: false
		}
    ,
    "System/Log/GetIsOpen": {
        requiresAuthentication: true
		}
    ,
    "System/Log/GetLog": {
        requiresAuthentication: false
		}
    ,
    "System/Menu/CreateOne": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/DeleteMany": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/DeleteOne": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetApiList": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetApisByMenus": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetChildrenMenus": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetList": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetMany": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/GetMenuByLevelOrPid": {
        requiresAuthentication: false
		}
    ,
    "System/Menu/GetOne": {
        requiresAuthentication: true
		}
    ,
    "System/Menu/UpdateOne": {
        requiresAuthentication: true
		}
    ,
    "System/Operation/GetMany": {
        requiresAuthentication: false
		}
    ,
    "System/Perm/CreateMany": {
        requiresAuthentication: true
		}
    ,
    "System/Perm/GetBindPerms": {
        requiresAuthentication: false
		}
    ,
    "System/Perm/GetRolePerms": {
        requiresAuthentication: false
		}
    ,
    "System/Role/AddOne": {
        requiresAuthentication: true
		}
    ,
    "System/Role/BindMenus": {
        requiresAuthentication: false
		}
    ,
    "System/Role/BindRoleApis": {
        requiresAuthentication: false
		}
    ,
    "System/Role/DeleteMany": {
        requiresAuthentication: false
		}
    ,
    "System/Role/DeleteOne": {
        requiresAuthentication: true
		}
    ,
    "System/Role/GetList": {
        requiresAuthentication: true
		}
    ,
    "System/Role/GetMany": {
        requiresAuthentication: false
		}
    ,
    "System/Role/GetOne": {
        requiresAuthentication: true
		}
    ,
    "System/Role/GetRoleMenuId": {
        requiresAuthentication: false
		}
    ,
    "System/Role/UnBindMenus": {
        requiresAuthentication: false
		}
    ,
    "System/Role/UnBindRoleApis": {
        requiresAuthentication: false
		}
    ,
    "System/Role/UpdateOne": {
        requiresAuthentication: true
		}
    ,
    "System/Sub/CreateSub": {
        requiresAuthentication: true
		}
    ,
    "System/Sub/SSE": {
        requiresAuthentication: false
		}
    ,
    "System/User/ConnectRole": {
        requiresAuthentication: false
		}
    ,
    "System/User/CountUsers": {
        requiresAuthentication: true
		}
    ,
    "System/User/CreateOne": {
        requiresAuthentication: true
		}
    ,
    "System/User/DeleteOne": {
        requiresAuthentication: true
		}
    ,
    "System/User/DisconnectRole": {
        requiresAuthentication: false
		}
    ,
    "System/User/GetAllList": {
        requiresAuthentication: true
		}
    ,
    "System/User/GetLikeUser": {
        requiresAuthentication: true
		}
    ,
    "System/User/GetList": {
        requiresAuthentication: true
		}
    ,
    "System/User/GetOne": {
        requiresAuthentication: true
		}
    ,
    "System/User/GetRoleUsers": {
        requiresAuthentication: false
		}
    ,
    "System/User/GetUserRole": {
        requiresAuthentication: false
		}
    ,
    "System/User/UpdateOne": {
        requiresAuthentication: false
		}
    
}

export class WunderGraphClient extends Client {
	query<
		OperationName extends Extract<keyof Operations['queries'], string>,
		Input extends Operations['queries'][OperationName]['input'] = Operations['queries'][OperationName]['input'],
		Data extends Operations['queries'][OperationName]['data'] = Operations['queries'][OperationName]['data']
	>(options: OperationName extends string ? OperationRequestOptions<OperationName, Input> : OperationRequestOptions) {
		return super.query<OperationRequestOptions, Data>(options);
	}
	mutate<
		OperationName extends Extract<keyof Operations['mutations'], string>,
		Input extends Operations['mutations'][OperationName]['input'] = Operations['mutations'][OperationName]['input'],
		Data extends Operations['mutations'][OperationName]['data'] = Operations['mutations'][OperationName]['data']
	>(options: OperationName extends string ? OperationRequestOptions<OperationName, Input> : OperationRequestOptions) {
		return super.mutate<OperationRequestOptions, Data>(options);
	}
	subscribe<
		OperationName extends Extract<keyof Operations['subscriptions'], string>,
		Input extends Operations['subscriptions'][OperationName]['input'] = Operations['subscriptions'][OperationName]['input'],
		Data extends Operations['subscriptions'][OperationName]['data'] = Operations['subscriptions'][OperationName]['data']
	>(
		options: OperationName extends string
			? SubscriptionRequestOptions<OperationName, Input>
			: SubscriptionRequestOptions,
		cb: SubscriptionEventHandler<Data>
	) {
		return super.subscribe(options, cb);
	}
    public async uploadFiles<
        ProviderName extends Extract<keyof S3Providers, string>,
        ProfileName extends ExtractProfileName<S3Providers[ProviderName]['profiles']> = ExtractProfileName<
        S3Providers[ProviderName]['profiles']
        >,
        Meta extends ExtractMeta<S3Providers[ProviderName]['profiles'], ProfileName> = ExtractMeta<
        S3Providers[ProviderName]['profiles'],
        ProfileName
        >
    >(config: UploadRequestOptions<ProviderName, ProfileName, Meta>) {
        const profile = config.profile ? S3UploadProviderData[config.provider][config.profile as string] : undefined;
        return super.uploadFiles(config, profile);
    }
	public login(authProviderID: Operations['authProvider'], redirectURI?: string) {
		return super.login(authProviderID, redirectURI);
	}
	public async fetchUser<TUser extends User = User<Role, CustomClaims>>(options?: FetchUserRequestOptions) {
		return super.fetchUser<TUser>(options);
	}
}

export const createClient = (config?: CreateClientConfig) => {
	return new WunderGraphClient({
		...defaultClientConfig,
		...config,
		operationMetadata,
		csrfEnabled: false,
	});
};

export type Queries = {
    'Casdoor/GetRolesById': {
        input: Casdoor__GetRolesByIdInput
        data: Casdoor__GetRolesByIdResponseData
        requiresAuthentication: false
        
    }
    'Post/GetList': {
        input: Post__GetListInput
        data: Post__GetListResponseData
        requiresAuthentication: true
        
    }
    'Post/GetOne': {
        input: Post__GetOneInput
        data: Post__GetOneResponseData
        requiresAuthentication: true
        
    }
    'Statistics/MonthlySales': {
        input?: undefined
        data: Statistics__MonthlySalesResponseData
        requiresAuthentication: false
        
    }
    'Statistics/SaleTypePercent': {
        input?: undefined
        data: Statistics__SaleTypePercentResponseData
        requiresAuthentication: false
        
    }
    'Statistics/SalesTop10': {
        input?: undefined
        data: Statistics__SalesTop10ResponseData
        requiresAuthentication: false
        
    }
    'Statistics/VisitTrending': {
        input?: undefined
        data: Statistics__VisitTrendingResponseData
        requiresAuthentication: false
        
    }
    'System/GetMenus': {
        input: System__GetMenusInput
        data: System__GetMenusResponseData
        requiresAuthentication: true
        
    }
    'System/Log/GetAllLog': {
        input?: undefined
        data: System__Log__GetAllLogResponseData
        requiresAuthentication: false
        
    }
    'System/Log/GetIsOpen': {
        input?: undefined
        data: System__Log__GetIsOpenResponseData
        requiresAuthentication: true
        
    }
    'System/Log/GetLog': {
        input: System__Log__GetLogInput
        data: System__Log__GetLogResponseData
        requiresAuthentication: false
        
    }
    'System/Menu/GetApiList': {
        input?: undefined
        data: System__Menu__GetApiListResponseData
        requiresAuthentication: true
        
    }
    'System/Menu/GetApisByMenus': {
        input: System__Menu__GetApisByMenusInput
        data: System__Menu__GetApisByMenusResponseData
        requiresAuthentication: true
        
    }
    'System/Menu/GetChildrenMenus': {
        input: System__Menu__GetChildrenMenusInput
        data: System__Menu__GetChildrenMenusResponseData
        requiresAuthentication: true
        
    }
    'System/Menu/GetList': {
        input: System__Menu__GetListInput
        data: System__Menu__GetListResponseData
        requiresAuthentication: true
        
    }
    'System/Menu/GetMany': {
        input?: undefined
        data: System__Menu__GetManyResponseData
        requiresAuthentication: true
        
    }
    'System/Menu/GetMenuByLevelOrPid': {
        input: System__Menu__GetMenuByLevelOrPidInput
        data: System__Menu__GetMenuByLevelOrPidResponseData
        requiresAuthentication: false
        
    }
    'System/Menu/GetOne': {
        input: System__Menu__GetOneInput
        data: System__Menu__GetOneResponseData
        requiresAuthentication: true
        
    }
    'System/Operation/GetMany': {
        input?: undefined
        data: System__Operation__GetManyResponseData
        requiresAuthentication: false
        
    }
    'System/Perm/GetBindPerms': {
        input?: undefined
        data: System__Perm__GetBindPermsResponseData
        requiresAuthentication: false
        
    }
    'System/Perm/GetRolePerms': {
        input: System__Perm__GetRolePermsInput
        data: System__Perm__GetRolePermsResponseData
        requiresAuthentication: false
        
    }
    'System/Role/GetList': {
        input: System__Role__GetListInput
        data: System__Role__GetListResponseData
        requiresAuthentication: true
        
    }
    'System/Role/GetMany': {
        input?: undefined
        data: System__Role__GetManyResponseData
        requiresAuthentication: false
        
    }
    'System/Role/GetOne': {
        input: System__Role__GetOneInput
        data: System__Role__GetOneResponseData
        requiresAuthentication: true
        
    }
    'System/Role/GetRoleMenuId': {
        input: System__Role__GetRoleMenuIdInput
        data: System__Role__GetRoleMenuIdResponseData
        requiresAuthentication: false
        
    }
    'System/Sub/SSE': {
        input: System__Sub__SSEInput
        data: System__Sub__SSEResponseData
        requiresAuthentication: false
        liveQuery: boolean
    }
    'System/User/CountUsers': {
        input?: undefined
        data: System__User__CountUsersResponseData
        requiresAuthentication: true
        
    }
    'System/User/GetAllList': {
        input?: undefined
        data: System__User__GetAllListResponseData
        requiresAuthentication: true
        
    }
    'System/User/GetLikeUser': {
        input: System__User__GetLikeUserInput
        data: System__User__GetLikeUserResponseData
        requiresAuthentication: true
        
    }
    'System/User/GetList': {
        input: System__User__GetListInput
        data: System__User__GetListResponseData
        requiresAuthentication: true
        
    }
    'System/User/GetOne': {
        input: System__User__GetOneInput
        data: System__User__GetOneResponseData
        requiresAuthentication: true
        
    }
    'System/User/GetRoleUsers': {
        input: System__User__GetRoleUsersInput
        data: System__User__GetRoleUsersResponseData
        requiresAuthentication: false
        
    }
    'System/User/GetUserRole': {
        input: System__User__GetUserRoleInput
        data: System__User__GetUserRoleResponseData
        requiresAuthentication: false
        
    }
}

export type Mutations = {
    'Casdoor/Login': {
        input: Casdoor__LoginInput
        data: Casdoor__LoginResponseData
        requiresAuthentication: false
    }
    'Casdoor/RefreshToken': {
        input: Casdoor__RefreshTokenInput
        data: Casdoor__RefreshTokenResponseData
        requiresAuthentication: false
    }
    'Casdoor/SendCode': {
        input: Casdoor__SendCodeInput
        data: Casdoor__SendCodeResponseData
        requiresAuthentication: false
    }
    'Casdoor/UpdateSMSProvider': {
        input: Casdoor__UpdateSMSProviderInput
        data: Casdoor__UpdateSMSProviderResponseData
        requiresAuthentication: true
    }
    'Post/CreateOne': {
        input: Post__CreateOneInput
        data: Post__CreateOneResponseData
        requiresAuthentication: true
    }
    'Post/DeleteMany': {
        input: Post__DeleteManyInput
        data: Post__DeleteManyResponseData
        requiresAuthentication: true
    }
    'Post/DeleteOne': {
        input: Post__DeleteOneInput
        data: Post__DeleteOneResponseData
        requiresAuthentication: true
    }
    'Post/UpdateOne': {
        input: Post__UpdateOneInput
        data: Post__UpdateOneResponseData
        requiresAuthentication: true
    }
    'Statistics/QueryRaw': {
        input: Statistics__QueryRawInput
        data: Statistics__QueryRawResponseData
        requiresAuthentication: false
    }
    'System/Log/ChangeOpen': {
        input: System__Log__ChangeOpenInput
        data: System__Log__ChangeOpenResponseData
        requiresAuthentication: true
    }
    'System/Log/CreateLog': {
        input: System__Log__CreateLogInput
        data: System__Log__CreateLogResponseData
        requiresAuthentication: false
    }
    'System/Menu/CreateOne': {
        input: System__Menu__CreateOneInput
        data: System__Menu__CreateOneResponseData
        requiresAuthentication: true
    }
    'System/Menu/DeleteMany': {
        input: System__Menu__DeleteManyInput
        data: System__Menu__DeleteManyResponseData
        requiresAuthentication: true
    }
    'System/Menu/DeleteOne': {
        input: System__Menu__DeleteOneInput
        data: System__Menu__DeleteOneResponseData
        requiresAuthentication: true
    }
    'System/Menu/UpdateOne': {
        input: System__Menu__UpdateOneInput
        data: System__Menu__UpdateOneResponseData
        requiresAuthentication: true
    }
    'System/Perm/CreateMany': {
        input: System__Perm__CreateManyInput
        data: System__Perm__CreateManyResponseData
        requiresAuthentication: true
    }
    'System/Role/AddOne': {
        input: System__Role__AddOneInput
        data: System__Role__AddOneResponseData
        requiresAuthentication: true
    }
    'System/Role/BindMenus': {
        input: System__Role__BindMenusInput
        data: System__Role__BindMenusResponseData
        requiresAuthentication: false
    }
    'System/Role/BindRoleApis': {
        input: System__Role__BindRoleApisInput
        data: System__Role__BindRoleApisResponseData
        requiresAuthentication: false
    }
    'System/Role/DeleteMany': {
        input: System__Role__DeleteManyInput
        data: System__Role__DeleteManyResponseData
        requiresAuthentication: false
    }
    'System/Role/DeleteOne': {
        input: System__Role__DeleteOneInput
        data: System__Role__DeleteOneResponseData
        requiresAuthentication: true
    }
    'System/Role/UnBindMenus': {
        input: System__Role__UnBindMenusInput
        data: System__Role__UnBindMenusResponseData
        requiresAuthentication: false
    }
    'System/Role/UnBindRoleApis': {
        input: System__Role__UnBindRoleApisInput
        data: System__Role__UnBindRoleApisResponseData
        requiresAuthentication: false
    }
    'System/Role/UpdateOne': {
        input: System__Role__UpdateOneInput
        data: System__Role__UpdateOneResponseData
        requiresAuthentication: true
    }
    'System/Sub/CreateSub': {
        input: System__Sub__CreateSubInput
        data: System__Sub__CreateSubResponseData
        requiresAuthentication: true
    }
    'System/User/ConnectRole': {
        input: System__User__ConnectRoleInput
        data: System__User__ConnectRoleResponseData
        requiresAuthentication: false
    }
    'System/User/CreateOne': {
        input: System__User__CreateOneInput
        data: System__User__CreateOneResponseData
        requiresAuthentication: true
    }
    'System/User/DeleteOne': {
        input: System__User__DeleteOneInput
        data: System__User__DeleteOneResponseData
        requiresAuthentication: true
    }
    'System/User/DisconnectRole': {
        input: System__User__DisconnectRoleInput
        data: System__User__DisconnectRoleResponseData
        requiresAuthentication: false
    }
    'System/User/UpdateOne': {
        input: System__User__UpdateOneInput
        data: System__User__UpdateOneResponseData
        requiresAuthentication: false
    }
}

export type Subscriptions = {
}

export type LiveQueries = {
}

export interface Operations extends OperationsDefinition<Queries, Mutations, Subscriptions, Role,S3Providers> {}
