export interface Casdoor__GetRolesByIdInput {
    userId: string

}
export interface InternalCasdoor__GetRolesByIdInput {
    id?: number
    roleId?: number
    userId: string
}
export interface InjectedCasdoor__GetRolesByIdInput {
    id?: number
    roleId?: number
    userId: string

}



export interface Casdoor__GetRolesByIdResponse {
    data?: Casdoor__GetRolesByIdResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Casdoor__GetRolesByIdResponseData {
    data?: {
    }[],
}export interface Casdoor__LoginInput {
    code?: string
    countryCode?: string
    loginType?: string
    password?: string
    phone?: string
    username?: string

}
export interface InternalCasdoor__LoginInput {
    code?: string
    countryCode?: string
    loginType?: string
    password?: string
    phone?: string
    username?: string
}
export interface InjectedCasdoor__LoginInput {
    code?: string
    countryCode?: string
    loginType?: string
    password?: string
    phone?: string
    username?: string

}



export interface Casdoor__LoginResponse {
    data?: Casdoor__LoginResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Casdoor__LoginResponseData {
    casdoor_apiLogin?: {
        data?: {
            accessToken?: string
            expireIn?: number
            refreshToken?: string
        },
        msg?: string
    },
}export interface Casdoor__RefreshTokenInput {
    refreshToken: string

}
export interface InternalCasdoor__RefreshTokenInput {
    refreshToken: string
}
export interface InjectedCasdoor__RefreshTokenInput {
    refreshToken: string

}



export interface Casdoor__RefreshTokenResponse {
    data?: Casdoor__RefreshTokenResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Casdoor__RefreshTokenResponseData {
    data?: {
        data?: {
            accessToken?: string
            expireIn?: number
            refreshToken?: string
        },
        msg?: string
    },
}export interface Casdoor__SendCodeInput {
    countryCode?: string
    dest: string

}
export interface InternalCasdoor__SendCodeInput {
    countryCode?: string
    dest: string
}
export interface InjectedCasdoor__SendCodeInput {
    countryCode?: string
    dest: string

}



export interface Casdoor__SendCodeResponse {
    data?: Casdoor__SendCodeResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Casdoor__SendCodeResponseData {
    data?: {
        msg?: string
    },
}export interface Casdoor__UpdateSMSProviderInput {
    clientId: string
    clientSecret: string
    signName: string
    templateCode: string

}
export interface InternalCasdoor__UpdateSMSProviderInput {
    clientId: string
    clientSecret: string
    signName: string
    templateCode: string
}
export interface InjectedCasdoor__UpdateSMSProviderInput {
    clientId: string
    clientSecret: string
    signName: string
    templateCode: string

}



export interface Casdoor__UpdateSMSProviderResponse {
    data?: Casdoor__UpdateSMSProviderResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Casdoor__UpdateSMSProviderResponseData {
    casdoor_apiUpdateProvider?: {
        msg?: string
    },
}export interface Post__CreateOneInput {
    content?: string
    poster?: string
    publishedAt?: string
    title: string
    username: string
    cateId: number
}
export interface InternalPost__CreateOneInput {
    content?: string
    poster?: string
    publishedAt?: string
    title: string
    username: string
}
export interface InjectedPost__CreateOneInput {
    content?: string
    poster?: string
    publishedAt?: string
    title: string
    username: string

}



export interface Post__CreateOneResponse {
    data?: Post__CreateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__CreateOneResponseData {
    data?: {
        author?: string
        id?: number
        poster?: string
        published_at?: string
        title?: string
    },
}export interface Post__DeleteManyInput {
    ids: {
    }[],

}
export interface InternalPost__DeleteManyInput {
    ids: {
    }[],
}
export interface InjectedPost__DeleteManyInput {
    ids: {
    }[],

}



export interface Post__DeleteManyResponse {
    data?: Post__DeleteManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__DeleteManyResponseData {
    data?: {
        count?: number
    },
}export interface Post__DeleteOneInput {
    id: number

}
export interface InternalPost__DeleteOneInput {
    id: number
}
export interface InjectedPost__DeleteOneInput {
    id: number

}



export interface Post__DeleteOneResponse {
    data?: Post__DeleteOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__DeleteOneResponseData {
    data?: {
        id?: number
    },
}export interface Post__GetListInput {
    skip?: number
    take?: number

}
export interface InternalPost__GetListInput {
    skip?: number
    take?: number
}
export interface main_DateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_IntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedDateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_NestedIntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedStringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_NestedStringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_postWhereInput {
    OR?: {
    }[],

}
export interface InjectedPost__GetListInput {
    skip?: number
    take?: number

}



export interface Post__GetListResponse {
    data?: Post__GetListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__GetListResponseData {
    data?: {
        author?: string
        content?: string
        id?: number
        poster?: string
        published_at?: string
        title?: string
    }[],
    total?: number
}export interface Post__GetOneInput {
    id: number

}
export interface InternalPost__GetOneInput {
    id: number
}
export interface InjectedPost__GetOneInput {
    id: number

}



export interface Post__GetOneResponse {
    data?: Post__GetOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__GetOneResponseData {
    data?: {
        author?: string
        content?: string
        id?: number
        poster?: string
        published_at?: string
        title?: string
    },
}export interface Post__UpdateOneInput {
    content: string
    id: number
    poster?: string
    publishedAt?: string
    title?: string

}
export interface InternalPost__UpdateOneInput {
    content: string
    id: number
    poster?: string
    publishedAt?: string
    title?: string
}
export interface InjectedPost__UpdateOneInput {
    content: string
    id: number
    poster?: string
    publishedAt?: string
    title?: string

}



export interface Post__UpdateOneResponse {
    data?: Post__UpdateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__UpdateOneResponseData {
    data?: {
        author?: string
        content?: string
        id?: number
        poster?: string
        published_at?: string
        title?: string
    },
}


export interface Statistics__MonthlySalesResponse {
    data?: Statistics__MonthlySalesResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__MonthlySalesResponseData {
    data?: {
        months?: string
    }[],
}export interface Statistics__QueryRawInput {
    query: string

}
export interface InternalStatistics__QueryRawInput {
    query: string
}
export interface InjectedStatistics__QueryRawInput {
    query: string

}



export interface Statistics__QueryRawResponse {
    data?: Statistics__QueryRawResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__QueryRawResponseData {
}


export interface Statistics__SaleTypePercentResponse {
    data?: Statistics__SaleTypePercentResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__SaleTypePercentResponseData {
    data?: {
        typeId?: number
        typeName?: string
    }[],
}


export interface Statistics__SalesTop10Response {
    data?: Statistics__SalesTop10ResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__SalesTop10ResponseData {
    data?: {
        shopName?: string
    }[],
}


export interface Statistics__VisitTrendingResponse {
    data?: Statistics__VisitTrendingResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__VisitTrendingResponseData {
    data?: {
        count?: number
        days?: string
    }[],
}export interface System__GetMenusInput {
    pid?: number

}
export interface InternalSystem__GetMenusInput {
    pid?: number
}
export interface InjectedSystem__GetMenusInput {
    pid?: number

}



export interface System__GetMenusResponse {
    data?: System__GetMenusResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__GetMenusResponseData {
    data?: {
        children?: {
            id?: number
            label?: string
            name?: string
            path?: string
            sort?: number
        }[],
        id?: number
        label?: string
        name?: string
        path?: string
        sort?: number
    }[],
}export interface System__Log__ChangeOpenInput {

}
export interface InternalSystem__Log__ChangeOpenInput {
}
export interface InjectedSystem__Log__ChangeOpenInput {

}



export interface System__Log__ChangeOpenResponse {
    data?: System__Log__ChangeOpenResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Log__ChangeOpenResponseData {
    main_updateOnedic?: {
        id?: number
    },
}export interface System__Log__CreateLogInput {
    code?: string
    ip?: string
    method?: string
    path?: string
    ua?: string
    updatedAt?: string

}
export interface InternalSystem__Log__CreateLogInput {
    code?: string
    ip?: string
    method?: string
    path?: string
    ua?: string
    updatedAt?: string
}
export interface InjectedSystem__Log__CreateLogInput {
    code?: string
    ip?: string
    method?: string
    path?: string
    ua?: string
    updatedAt?: string

}



export interface System__Log__CreateLogResponse {
    data?: System__Log__CreateLogResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Log__CreateLogResponseData {
    main_createOneapilog?: {
        id?: number
        ip?: string
        method?: string
        path?: string
        ua?: string
        updatedAt?: string
    },
}


export interface System__Log__GetAllLogResponse {
    data?: System__Log__GetAllLogResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Log__GetAllLogResponseData {
    logNum?: {
        _count?: {
            id?: number
        },
    },
}


export interface System__Log__GetIsOpenResponse {
    data?: System__Log__GetIsOpenResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Log__GetIsOpenResponseData {
    main_findFirstdic?: {
        id?: number
    },
}export interface System__Log__GetLogInput {
    skip?: number
    take?: number

}
export interface InternalSystem__Log__GetLogInput {
    skip?: number
    take?: number
}
export interface InjectedSystem__Log__GetLogInput {
    skip?: number
    take?: number

}



export interface System__Log__GetLogResponse {
    data?: System__Log__GetLogResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Log__GetLogResponseData {
    data?: {
        id?: number
        ip?: string
        method?: string
        path?: string
        ua?: string
        updatedAt?: string
    }[],
}export interface System__Menu__CreateOneInput {
    apiId?: string
    icon?: string
    label: string
    level: number
    menuType: string
    parentId?: number
    path: string
    perms?: string
    sort: number

}
export interface InternalSystem__Menu__CreateOneInput {
    apiId?: string
    create_time?: string
    icon?: string
    label: string
    level: number
    menuType: string
    parentId?: number
    path: string
    perms?: string
    sort: number
}
export interface InjectedSystem__Menu__CreateOneInput {
    apiId?: string
    create_time: string
    icon?: string
    label: string
    level: number
    menuType: string
    parentId?: number
    path: string
    perms?: string
    sort: number

}



export interface System__Menu__CreateOneResponse {
    data?: System__Menu__CreateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__CreateOneResponseData {
    data?: {
        id?: number
    },
}export interface System__Menu__DeleteManyInput {
    ids: {
    }[],

}
export interface InternalSystem__Menu__DeleteManyInput {
    ids: {
    }[],
}
export interface InjectedSystem__Menu__DeleteManyInput {
    ids: {
    }[],

}



export interface System__Menu__DeleteManyResponse {
    data?: System__Menu__DeleteManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__DeleteManyResponseData {
    data?: {
        count?: number
    },
}export interface System__Menu__DeleteOneInput {
    id: number

}
export interface InternalSystem__Menu__DeleteOneInput {
    id: number
}
export interface InjectedSystem__Menu__DeleteOneInput {
    id: number

}



export interface System__Menu__DeleteOneResponse {
    data?: System__Menu__DeleteOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__DeleteOneResponseData {
    data?: {
        id?: number
    },
}


export interface System__Menu__GetApiListResponse {
    data?: System__Menu__GetApiListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetApiListResponseData {
    data?: {
        id?: number
        path?: string
    }[],
}export interface System__Menu__GetApisByMenusInput {
    menuIds: {
    }[],

}
export interface InternalSystem__Menu__GetApisByMenusInput {
    menuIds: {
    }[],
}
export interface InjectedSystem__Menu__GetApisByMenusInput {
    menuIds: {
    }[],

}



export interface System__Menu__GetApisByMenusResponse {
    data?: System__Menu__GetApisByMenusResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetApisByMenusResponseData {
    data?: {
    }[],
}export interface System__Menu__GetChildrenMenusInput {
    pid: number

}
export interface InternalSystem__Menu__GetChildrenMenusInput {
    pid: number
}
export interface InjectedSystem__Menu__GetChildrenMenusInput {
    pid: number

}



export interface System__Menu__GetChildrenMenusResponse {
    data?: System__Menu__GetChildrenMenusResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetChildrenMenusResponseData {
    data?: {
        id?: number
        label?: string
        level?: number
        name?: string
        path?: string
        sort?: number
    }[],
}export interface System__Menu__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number

}
export interface InternalSystem__Menu__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number
}
export interface main_DateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_IntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_IntNullableFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedDateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_NestedIntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedIntNullableFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedStringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_NestedStringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_menuOrderByWithRelationInput {
    api_id?: string
    create_time?: string
    icon?: string
    id?: string
    is_bottom?: string
    label?: string
    level?: string
    menu_type?: string
    name?: string
    parentId?: string
    path?: string
    perms?: string
    sort?: string

}
export interface main_menuWhereInput {
    OR?: {
    }[],

}
export interface InjectedSystem__Menu__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number

}



export interface System__Menu__GetListResponse {
    data?: System__Menu__GetListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetListResponseData {
    data?: {
        icon?: string
        id?: number
        label?: string
        level?: number
        parentId?: number
        path?: string
        sort?: number
    }[],
    total?: number
}


export interface System__Menu__GetManyResponse {
    data?: System__Menu__GetManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetManyResponseData {
    data?: {
        api_id?: string
        create_time?: string
        icon?: string
        id?: number
        is_bottom?: number
        label?: string
        level?: number
        menu_type?: string
        parentId?: number
        path?: string
        perms?: string
        sort?: number
    }[],
}export interface System__Menu__GetMenuByLevelOrPidInput {
    level?: number
    pid?: number

}
export interface InternalSystem__Menu__GetMenuByLevelOrPidInput {
    level?: number
    pid?: number
}
export interface InjectedSystem__Menu__GetMenuByLevelOrPidInput {
    level?: number
    pid?: number

}



export interface System__Menu__GetMenuByLevelOrPidResponse {
    data?: System__Menu__GetMenuByLevelOrPidResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetMenuByLevelOrPidResponseData {
    data?: {
        id?: number
        is_bottom?: number
        label?: string
    }[],
}export interface System__Menu__GetOneInput {
    id: number

}
export interface InternalSystem__Menu__GetOneInput {
    id: number
}
export interface InjectedSystem__Menu__GetOneInput {
    id: number

}



export interface System__Menu__GetOneResponse {
    data?: System__Menu__GetOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetOneResponseData {
    data?: {
        icon?: string
        id?: number
        label?: string
        level?: number
        parentId?: number
        path?: string
        sort?: number
    },
}export interface System__Menu__UpdateOneInput {
    icon?: string
    id: number
    label?: string
    level?: number
    path?: string
    sort?: number

}
export interface InternalSystem__Menu__UpdateOneInput {
    icon?: string
    id: number
    label?: string
    level?: number
    path?: string
    sort?: number
}
export interface InjectedSystem__Menu__UpdateOneInput {
    icon?: string
    id: number
    label?: string
    level?: number
    path?: string
    sort?: number

}



export interface System__Menu__UpdateOneResponse {
    data?: System__Menu__UpdateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__UpdateOneResponseData {
    data?: {
        icon?: string
        id?: number
        label?: string
        level?: number
        parentId?: number
        path?: string
        sort?: number
    },
}


export interface System__Operation__GetManyResponse {
    data?: System__Operation__GetManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Operation__GetManyResponseData {
    data?: {
        createTime?: string
        id?: number
        method?: string
        operationType?: string
        title?: string
    }[],
}export interface System__Perm__CreateManyInput {
    data: {
    }[],

}
export interface InternalSystem__Perm__CreateManyInput {
    data: {
    }[],
}
export interface main_permissionCreateManyInput {
    createdAt?: string
    enabled?: number
    id?: string
    method?: string
    path?: string
    updatedAt?: string

}
export interface InjectedSystem__Perm__CreateManyInput {
    data: {
    }[],

}



export interface System__Perm__CreateManyResponse {
    data?: System__Perm__CreateManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Perm__CreateManyResponseData {
    data?: {
        count?: number
    },
}


export interface System__Perm__GetBindPermsResponse {
    data?: System__Perm__GetBindPermsResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Perm__GetBindPermsResponseData {
    data?: {
        createdAt?: string
        enabled?: number
        method?: string
        path?: string
    }[],
}export interface System__Perm__GetRolePermsInput {
    code: {
    }[],

}
export interface InternalSystem__Perm__GetRolePermsInput {
    code: {
    }[],
    menuId?: number
    roleId?: number
}
export interface InjectedSystem__Perm__GetRolePermsInput {
    code: {
    }[],
    menuId?: number
    roleId?: number

}



export interface System__Perm__GetRolePermsResponse {
    data?: System__Perm__GetRolePermsResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Perm__GetRolePermsResponseData {
    data?: {
        _join?: {
            data?: {
            }[],
        },
        id?: number
    }[],
}export interface System__Role__AddOneInput {
    code: string
    remark: string

}
export interface InternalSystem__Role__AddOneInput {
    code: string
    remark: string
}
export interface InjectedSystem__Role__AddOneInput {
    code: string
    remark: string

}



export interface System__Role__AddOneResponse {
    data?: System__Role__AddOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__AddOneResponseData {
    data?: {
        code?: string
        id?: number
        remark?: string
    },
}export interface System__Role__BindMenusInput {
    data: {
    }[],

}
export interface InternalSystem__Role__BindMenusInput {
    data: {
    }[],
}
export interface main_menu_roleCreateManyInput {
    id?: number
    menu_id?: number
    role_id?: number

}
export interface InjectedSystem__Role__BindMenusInput {
    data: {
    }[],

}



export interface System__Role__BindMenusResponse {
    data?: System__Role__BindMenusResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__BindMenusResponseData {
    data?: {
        count?: number
    },
}export interface System__Role__BindRoleApisInput {
    allRoles: {
    }[],
    apis: {
    }[],
    roleCode: string

}
export interface InternalSystem__Role__BindRoleApisInput {
    allRoles: {
    }[],
    apis: {
    }[],
    roleCode: string
}
export interface InjectedSystem__Role__BindRoleApisInput {
    allRoles: {
    }[],
    apis: {
    }[],
    roleCode: string

}



export interface System__Role__BindRoleApisResponse {
    data?: System__Role__BindRoleApisResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__BindRoleApisResponseData {
    system_bindRoleApis?: {
        count?: number
    },
}export interface System__Role__DeleteManyInput {
    ids?: {
    }[],

}
export interface InternalSystem__Role__DeleteManyInput {
    ids?: {
    }[],
}
export interface InjectedSystem__Role__DeleteManyInput {
    ids?: {
    }[],

}



export interface System__Role__DeleteManyResponse {
    data?: System__Role__DeleteManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__DeleteManyResponseData {
    data?: {
        count?: number
    },
}export interface System__Role__DeleteOneInput {
    code: string

}
export interface InternalSystem__Role__DeleteOneInput {
    code: string
}
export interface InjectedSystem__Role__DeleteOneInput {
    code: string

}



export interface System__Role__DeleteOneResponse {
    data?: System__Role__DeleteOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__DeleteOneResponseData {
    data?: {
        id?: number
    },
}export interface System__Role__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number

}
export interface InternalSystem__Role__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number
}
export interface main_IntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedIntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedStringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_roleOrderByWithRelationInput {
    code?: string
    id?: string
    remark?: string

}
export interface main_roleWhereInput {
    OR?: {
    }[],

}
export interface InjectedSystem__Role__GetListInput {
    orderBy?: {
    }[],
    skip?: number
    take?: number

}



export interface System__Role__GetListResponse {
    data?: System__Role__GetListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetListResponseData {
    data?: {
        code?: string
        id?: number
        remark?: string
    }[],
    total?: number
}


export interface System__Role__GetManyResponse {
    data?: System__Role__GetManyResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetManyResponseData {
    data?: {
        code?: string
        id?: number
        remark?: string
    }[],
}export interface System__Role__GetOneInput {
    id: number

}
export interface InternalSystem__Role__GetOneInput {
    id: number
}
export interface InjectedSystem__Role__GetOneInput {
    id: number

}



export interface System__Role__GetOneResponse {
    data?: System__Role__GetOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetOneResponseData {
    data?: {
        code?: string
    },
}export interface System__Role__GetRoleMenuIdInput {
    roleId: number

}
export interface InternalSystem__Role__GetRoleMenuIdInput {
    roleId: number
}
export interface InjectedSystem__Role__GetRoleMenuIdInput {
    roleId: number

}



export interface System__Role__GetRoleMenuIdResponse {
    data?: System__Role__GetRoleMenuIdResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetRoleMenuIdResponseData {
    data?: {
    }[],
}export interface System__Role__UnBindMenusInput {
    menuIds: {
    }[],
    roleId: number

}
export interface InternalSystem__Role__UnBindMenusInput {
    menuIds: {
    }[],
    roleId: number
}
export interface InjectedSystem__Role__UnBindMenusInput {
    menuIds: {
    }[],
    roleId: number

}



export interface System__Role__UnBindMenusResponse {
    data?: System__Role__UnBindMenusResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__UnBindMenusResponseData {
    main_deleteManymenu_role?: {
        count?: number
    },
}export interface System__Role__UnBindRoleApisInput {
    apis: {
    }[],
    roleCode: string

}
export interface InternalSystem__Role__UnBindRoleApisInput {
    apis: {
    }[],
    roleCode: string
}
export interface InjectedSystem__Role__UnBindRoleApisInput {
    apis: {
    }[],
    roleCode: string

}



export interface System__Role__UnBindRoleApisResponse {
    data?: System__Role__UnBindRoleApisResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__UnBindRoleApisResponseData {
    system_unBindRoleApis?: {
        count?: number
    },
}export interface System__Role__UpdateOneInput {
    id: number
    remark?: string

}
export interface InternalSystem__Role__UpdateOneInput {
    id: number
    remark?: string
}
export interface InjectedSystem__Role__UpdateOneInput {
    id: number
    remark?: string

}



export interface System__Role__UpdateOneResponse {
    data?: System__Role__UpdateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__UpdateOneResponseData {
    data?: {
        code?: string
        id?: number
        remark?: string
    },
}export interface System__Sub__CreateSubInput {
    create_role?: string
    message?: string
    target_role?: string
    updatedAt?: string

}
export interface InternalSystem__Sub__CreateSubInput {
    create_role?: string
    message?: string
    target_role?: string
    updatedAt?: string
}
export interface InjectedSystem__Sub__CreateSubInput {
    create_role?: string
    message?: string
    target_role?: string
    updatedAt?: string

}



export interface System__Sub__CreateSubResponse {
    data?: System__Sub__CreateSubResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Sub__CreateSubResponseData {
    main_createOnesub?: {
        create_role?: string
        id?: number
        message?: string
        target_role?: string
        updatedAt?: string
    },
}export interface System__Sub__SSEInput {
    roles: string,

}
export interface InternalSystem__Sub__SSEInput {
    roles: {
    }[],
}
export interface InjectedSystem__Sub__SSEInput {
    roles: {
    }[],

}



export interface System__Sub__SSEResponse {
    data?: System__Sub__SSEResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Sub__SSEResponseData {
    data?: {
        create_role?: string
        message?: string
        target_role?: string
        updatedAt?: string
    }[],
}export interface System__User__ConnectRoleInput {
    roleId: number
    userId: number

}
export interface InternalSystem__User__ConnectRoleInput {
    roleId: number
    userId: number
}
export interface InjectedSystem__User__ConnectRoleInput {
    roleId: number
    userId: number

}



export interface System__User__ConnectRoleResponse {
    data?: System__User__ConnectRoleResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__ConnectRoleResponseData {
    data?: {
        id?: number
    },
}


export interface System__User__CountUsersResponse {
    data?: System__User__CountUsersResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__CountUsersResponseData {
    data?: {
        count?: {
            id?: number
        },
    },
}export interface System__User__CreateOneInput {
    countryCode?: string
    name: string
    password?: string
    passwordType?: string
    phone: string

}
export interface InternalSystem__User__CreateOneInput {
    countryCode?: string
    name: string
    password?: string
    passwordType?: string
    phone: string
}
export interface InjectedSystem__User__CreateOneInput {
    countryCode?: string
    name: string
    password?: string
    passwordType?: string
    phone: string

}



export interface System__User__CreateOneResponse {
    data?: System__User__CreateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__CreateOneResponseData {
    data?: {
        msg?: string
        status?: number
    },
}export interface System__User__DeleteOneInput {
    id?: number

}
export interface InternalSystem__User__DeleteOneInput {
    id?: number
}
export interface InjectedSystem__User__DeleteOneInput {
    id?: number

}



export interface System__User__DeleteOneResponse {
    data?: System__User__DeleteOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__DeleteOneResponseData {
    main_deleteOneuser?: {
        id?: number
        name?: string
        password?: string
        password_salt?: string
        password_type?: string
        phone?: string
        user_id?: string
    },
}export interface System__User__DisconnectRoleInput {
    roleId: number
    userId?: number

}
export interface InternalSystem__User__DisconnectRoleInput {
    roleId: number
    userId?: number
}
export interface InjectedSystem__User__DisconnectRoleInput {
    roleId: number
    userId?: number

}



export interface System__User__DisconnectRoleResponse {
    data?: System__User__DisconnectRoleResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__DisconnectRoleResponseData {
    data?: {
        id?: number
    },
}


export interface System__User__GetAllListResponse {
    data?: System__User__GetAllListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetAllListResponseData {
    main_aggregateuser?: {
        _count?: {
            id?: number
        },
    },
}export interface System__User__GetLikeUserInput {
    name?: string

}
export interface InternalSystem__User__GetLikeUserInput {
    name?: string
}
export interface InjectedSystem__User__GetLikeUserInput {
    name?: string

}



export interface System__User__GetLikeUserResponse {
    data?: System__User__GetLikeUserResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetLikeUserResponseData {
    main_findManyuser?: {
        id?: number
        name?: string
        phone?: string
    }[],
}export interface System__User__GetListInput {
    orderBy?: {
    }[],
    skip: number
    take: number

}
export interface InternalSystem__User__GetListInput {
    orderBy?: {
    }[],
    roleId?: number
    skip: number
    take: number
    userId?: number
}
export interface main_DateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_IntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedDateTimeNullableFilter {
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],

}
export interface main_NestedIntFilter {
    equals?: number
    gt?: number
    gte?: number
    in?: {
    }[],
    lt?: number
    lte?: number
    notIn?: {
    }[],

}
export interface main_NestedStringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_NestedStringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_StringNullableFilter {
    contains?: string
    endsWith?: string
    equals?: string
    gt?: string
    gte?: string
    in?: {
    }[],
    lt?: string
    lte?: string
    notIn?: {
    }[],
    startsWith?: string

}
export interface main_userOrderByWithRelationInput {
    avatar?: string
    country_code?: string
    created_at?: string
    id?: string
    name?: string
    password?: string
    password_salt?: string
    password_type?: string
    phone?: string
    user_id?: string

}
export interface main_userWhereInput {
    OR?: {
    }[],

}
export interface InjectedSystem__User__GetListInput {
    orderBy?: {
    }[],
    roleId?: number
    skip: number
    take: number
    userId?: number

}



export interface System__User__GetListResponse {
    data?: System__User__GetListResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetListResponseData {
    data?: {
        _join?: {
            main_findManyrole_user?: {
                _join?: {
                    main_findManyrole?: {
                        code?: string
                        id?: number
                        remark?: string
                    }[],
                },
                role_id?: number
                user_id?: number
            }[],
        },
        avatar?: string
        createdAt?: string
        id?: number
        name?: string
        phone?: string
    }[],
}export interface System__User__GetOneInput {
    name?: string
    phone?: string
    roleId?: number
    userId?: number

}
export interface InternalSystem__User__GetOneInput {
    name?: string
    phone?: string
    roleId?: number
    userId?: number
}
export interface InjectedSystem__User__GetOneInput {
    name?: string
    phone?: string
    roleId?: number
    userId?: number

}



export interface System__User__GetOneResponse {
    data?: System__User__GetOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetOneResponseData {
    data?: {
        avatar?: string
        id?: number
        name?: string
        roles?: {
        }[],
    },
}export interface System__User__GetRoleUsersInput {
    code: string

}
export interface InternalSystem__User__GetRoleUsersInput {
    code: string
    roleId?: number
    userId?: number
}
export interface InjectedSystem__User__GetRoleUsersInput {
    code: string
    roleId?: number
    userId?: number

}



export interface System__User__GetRoleUsersResponse {
    data?: System__User__GetRoleUsersResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetRoleUsersResponseData {
    main_findManyrole?: {
        _join?: {
            main_findManyrole_user?: {
                _join?: {
                    main_findManyuser?: {
                        id?: number
                        name?: string
                    }[],
                },
                role_id?: number
                user_id?: number
            }[],
        },
        code?: string
        id?: number
        remark?: string
    }[],
}export interface System__User__GetUserRoleInput {
    userId: number

}
export interface InternalSystem__User__GetUserRoleInput {
    roleId?: number
    userId: number
}
export interface InjectedSystem__User__GetUserRoleInput {
    roleId?: number
    userId: number

}



export interface System__User__GetUserRoleResponse {
    data?: System__User__GetUserRoleResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetUserRoleResponseData {
    data?: {
        code?: string
        id?: number
        remark?: string
    }[],
}export interface System__User__UpdateOneInput {
    avatar?: string
    country_code?: string
    id?: number
    name?: string
    password?: string
    phone?: string

}
export interface InternalSystem__User__UpdateOneInput {
    avatar?: string
    country_code?: string
    id?: number
    name?: string
    password?: string
    phone?: string
}
export interface InjectedSystem__User__UpdateOneInput {
    avatar?: string
    country_code?: string
    id?: number
    name?: string
    password?: string
    phone?: string

}



export interface System__User__UpdateOneResponse {
    data?: System__User__UpdateOneResponseData
    errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__UpdateOneResponseData {
    main_updateOneuser?: {
        avatar?: string
        name?: string
        password?: string
        phone?: string
    },
}
export type JSONValue = string | number | boolean | JSONObject | Array<JSONValue>;

export type JSONObject = { [key: string]: JSONValue };

export interface GraphQLError {
    message: string;
    path?: ReadonlyArray<string | number>;
}