export interface BindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
  roleType?: string;
}
export interface InternalBindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
  roleType?: string;
}
export interface InjectedBindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
  roleType?: string;
}

export interface BindRoleApisResponse {
  data?: BindRoleApisResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface BindRoleApisResponseData {
  open_BindRoleApis?: {
    count?: number;
  };
}

export interface GetAllApisResponse {
  data?: GetAllApisResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface GetAllApisResponseData {
  open_GetAllApis?: {
    content?: string;
    createTime?: string;
    deleteTime?: string;
    id?: string;
    method?: string;
    operationType?: string;
    remark?: string;
    restUrl?: string;
    roleType?: string;
    roles?: string;
    title?: string;
    updateTime?: string;
  }[];
}

export interface GetAllRolesResponse {
  data?: GetAllRolesResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface GetAllRolesResponseData {
  open_GetAllRoles?: {
    code?: string;
    remark?: string;
  }[];
}
export interface GetRoleBindApisInput {
  code: string;
  roleType?: string;
}
export interface InternalGetRoleBindApisInput {
  code: string;
  roleType?: string;
}
export interface InjectedGetRoleBindApisInput {
  code: string;
  roleType?: string;
}

export interface GetRoleBindApisResponse {
  data?: GetRoleBindApisResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface GetRoleBindApisResponseData {
  open_GetRoleBindApis?: {
    content?: string;
    createTime?: string;
    deleteTime?: string;
    id?: string;
  }[];
}
export interface Post__CreateOneInput {
  content?: string;
  poster?: string;
  publishedAt?: string;
  title: string;
}
export interface InternalPost__CreateOneInput {
  content?: string;
  poster?: string;
  publishedAt?: string;
  title: string;
  userId?: string;
}
export interface InjectedPost__CreateOneInput {
  content?: string;
  poster?: string;
  publishedAt?: string;
  title: string;
  userId: string;
}

export interface Post__CreateOneResponse {
  data?: Post__CreateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__CreateOneResponseData {
  data?: {
    User?: {
      avatarUrl?: string;
      id?: string;
      name?: string;
    };
    authorId?: string;
    id?: number;
    poster?: string;
    publishedAt?: string;
    title?: string;
  };
}
export interface Post__DeleteManyInput {
  ids: {}[];
}
export interface InternalPost__DeleteManyInput {
  ids: {}[];
}
export interface InjectedPost__DeleteManyInput {
  ids: {}[];
}

export interface Post__DeleteManyResponse {
  data?: Post__DeleteManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__DeleteManyResponseData {
  data?: {
    count?: number;
  };
}
export interface Post__DeleteOneInput {
  id: number;
}
export interface InternalPost__DeleteOneInput {
  id: number;
}
export interface InjectedPost__DeleteOneInput {
  id: number;
}

export interface Post__DeleteOneResponse {
  data?: Post__DeleteOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__DeleteOneResponseData {
  data?: {
    id?: number;
  };
}
export interface Post__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface InternalPost__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface main_DateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_DateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_IntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_IntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_MenuListRelationFilter {}
export interface main_MenuRelationFilter {}
export interface main_MenuWhereInput {
  OR?: {}[];
}
export interface main_NestedDateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedDateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedIntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedIntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedStringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NestedStringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_PostListRelationFilter {}
export interface main_PostOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_PostOrderByWithRelationInput {
  authorId?: string;
  content?: string;
  id?: string;
  poster?: string;
  publishedAt?: string;
  title?: string;
}
export interface main_PostWhereInput {
  OR?: {}[];
}
export interface main_RoleListRelationFilter {}
export interface main_RoleOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_RoleWhereInput {
  OR?: {}[];
}
export interface main_StringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_StringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_UserListRelationFilter {}
export interface main_UserOrderByWithRelationInput {
  avatarUrl?: string;
  createdAt?: string;
  id?: string;
  name?: string;
}
export interface main_UserRelationFilter {}
export interface main_UserWhereInput {
  OR?: {}[];
}
export interface InjectedPost__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}

export interface Post__GetListResponse {
  data?: Post__GetListResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__GetListResponseData {
  data?: {
    User?: {
      avatarUrl?: string;
      name?: string;
    };
    authorId?: string;
    content?: string;
    id?: number;
    poster?: string;
    publishedAt?: string;
    title?: string;
  }[];
  total?: number;
}
export interface Post__GetOneInput {
  id: number;
}
export interface InternalPost__GetOneInput {
  id: number;
}
export interface InjectedPost__GetOneInput {
  id: number;
}

export interface Post__GetOneResponse {
  data?: Post__GetOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__GetOneResponseData {
  data?: {
    User?: {
      avatarUrl?: string;
      id?: string;
      name?: string;
    };
    authorId?: string;
    content?: string;
    id?: number;
    poster?: string;
    publishedAt?: string;
    title?: string;
  };
}
export interface Post__UpdateOneInput {
  content?: string;
  id: number;
  poster?: string;
  publishedAt?: string;
  title?: string;
}
export interface InternalPost__UpdateOneInput {
  content?: string;
  id: number;
  poster?: string;
  publishedAt?: string;
  title?: string;
}
export interface InjectedPost__UpdateOneInput {
  content?: string;
  id: number;
  poster?: string;
  publishedAt?: string;
  title?: string;
}

export interface Post__UpdateOneResponse {
  data?: Post__UpdateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Post__UpdateOneResponseData {
  data?: {
    User?: {
      avatarUrl?: string;
      id?: string;
      name?: string;
    };
    authorId?: string;
    content?: string;
    id?: number;
    poster?: string;
    publishedAt?: string;
    title?: string;
  };
}

export interface Statistics__MonthlySalesResponse {
  data?: Statistics__MonthlySalesResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__MonthlySalesResponseData {
  data?: {
    months?: string;
    totalSales?: number;
  }[];
}

export interface Statistics__SaleTypePercentResponse {
  data?: Statistics__SaleTypePercentResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__SaleTypePercentResponseData {
  data?: {
    typeId?: number;
    typeName?: string;
    totalSales?: number;
  }[];
}

export interface Statistics__SalesTop10Response {
  data?: Statistics__SalesTop10ResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__SalesTop10ResponseData {
  data?: {
    shopName?: string;
  }[];
}

export interface Statistics__VisitTrendingResponse {
  data?: Statistics__VisitTrendingResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface Statistics__VisitTrendingResponseData {
  data?: {
    count?: number;
    days?: string;
  }[];
}
export interface System__Menu__CreateOneInput {
  icon?: string;
  label: string;
  level?: number;
  path: string;
  sort?: number;
}
export interface InternalSystem__Menu__CreateOneInput {
  icon?: string;
  label: string;
  level?: number;
  path: string;
  sort?: number;
}
export interface InjectedSystem__Menu__CreateOneInput {
  icon?: string;
  label: string;
  level?: number;
  path: string;
  sort?: number;
}

export interface System__Menu__CreateOneResponse {
  data?: System__Menu__CreateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__CreateOneResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  };
}
export interface System__Menu__DeleteManyInput {
  ids: {}[];
}
export interface InternalSystem__Menu__DeleteManyInput {
  ids: {}[];
}
export interface InjectedSystem__Menu__DeleteManyInput {
  ids: {}[];
}

export interface System__Menu__DeleteManyResponse {
  data?: System__Menu__DeleteManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__DeleteManyResponseData {
  data?: {
    count?: number;
  };
}
export interface System__Menu__DeleteOneInput {
  id: number;
}
export interface InternalSystem__Menu__DeleteOneInput {
  id: number;
}
export interface InjectedSystem__Menu__DeleteOneInput {
  id: number;
}

export interface System__Menu__DeleteOneResponse {
  data?: System__Menu__DeleteOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__DeleteOneResponseData {
  data?: {
    id?: number;
  };
}
export interface System__Menu__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface InternalSystem__Menu__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface main_DateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_DateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_IntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_IntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_MenuListRelationFilter {}
export interface main_MenuOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_MenuOrderByWithRelationInput {
  icon?: string;
  id?: string;
  label?: string;
  level?: string;
  parentId?: string;
  path?: string;
  sort?: string;
}
export interface main_MenuRelationFilter {}
export interface main_MenuWhereInput {
  OR?: {}[];
}
export interface main_NestedDateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedDateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedIntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedIntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedStringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NestedStringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_PostListRelationFilter {}
export interface main_PostWhereInput {
  OR?: {}[];
}
export interface main_RoleListRelationFilter {}
export interface main_RoleOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_RoleWhereInput {
  OR?: {}[];
}
export interface main_StringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_StringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_UserListRelationFilter {}
export interface main_UserRelationFilter {}
export interface main_UserWhereInput {
  OR?: {}[];
}
export interface InjectedSystem__Menu__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}

export interface System__Menu__GetListResponse {
  data?: System__Menu__GetListResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetListResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  }[];
  total?: number;
}

export interface System__Menu__GetManyResponse {
  data?: System__Menu__GetManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetManyResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  }[];
}
export interface System__Menu__GetOneInput {
  id: number;
}
export interface InternalSystem__Menu__GetOneInput {
  id: number;
}
export interface InjectedSystem__Menu__GetOneInput {
  id: number;
}

export interface System__Menu__GetOneResponse {
  data?: System__Menu__GetOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__GetOneResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  };
}
export interface System__Menu__UpdateOneInput {
  icon?: string;
  id: number;
  label?: string;
  level?: number;
  path?: string;
  sort?: number;
}
export interface InternalSystem__Menu__UpdateOneInput {
  icon?: string;
  id: number;
  label?: string;
  level?: number;
  path?: string;
  sort?: number;
}
export interface main_DateTimeFieldUpdateOperationsInput {
  set?: string;
}
export interface main_DateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_DateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_IntFieldUpdateOperationsInput {
  decrement?: number;
  divide?: number;
  increment?: number;
  multiply?: number;
  set?: number;
}
export interface main_IntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_MenuCreateNestedOneWithoutOther_MenuInput {}
export interface main_MenuCreateOrConnectWithoutOther_MenuInput {}
export interface main_MenuCreateWithoutOther_MenuInput {
  icon?: string;
  label?: string;
  level?: number;
  path?: string;
  sort?: number;
}
export interface main_MenuUpdateOneWithoutOther_MenuNestedInput {}
export interface main_MenuUpdateWithoutOther_MenuInput {}
export interface main_MenuUpsertWithoutOther_MenuInput {}
export interface main_MenuWhereUniqueInput {
  id: number;
}
export interface main_NestedDateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedDateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedIntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedStringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NestedStringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NullableDateTimeFieldUpdateOperationsInput {
  set?: string;
}
export interface main_NullableStringFieldUpdateOperationsInput {
  set?: string;
}
export interface main_PostCreateNestedManyWithoutUserInput {}
export interface main_PostCreateOrConnectWithoutUserInput {}
export interface main_PostCreateWithoutUserInput {
  content?: string;
  poster?: string;
  publishedAt?: string;
  title?: string;
}
export interface main_PostScalarWhereInput {
  OR?: {}[];
}
export interface main_PostUpdateManyMutationInput {}
export interface main_PostUpdateManyWithWhereWithoutUserInput {}
export interface main_PostUpdateManyWithoutUserNestedInput {}
export interface main_PostUpdateWithWhereUniqueWithoutUserInput {}
export interface main_PostUpdateWithoutUserInput {}
export interface main_PostUpsertWithWhereUniqueWithoutUserInput {}
export interface main_PostWhereUniqueInput {
  id: number;
}
export interface main_RoleCreateNestedManyWithoutMenuInput {}
export interface main_RoleCreateOrConnectWithoutMenuInput {}
export interface main_RoleCreateWithoutMenuInput {
  code?: string;
  remark?: string;
}
export interface main_RoleScalarWhereInput {
  OR?: {}[];
}
export interface main_RoleUpdateManyMutationInput {}
export interface main_RoleUpdateManyWithWhereWithoutMenuInput {}
export interface main_RoleUpdateManyWithoutMenuNestedInput {}
export interface main_RoleUpdateWithWhereUniqueWithoutMenuInput {}
export interface main_RoleUpdateWithoutMenuInput {}
export interface main_RoleUpsertWithWhereUniqueWithoutMenuInput {}
export interface main_RoleWhereUniqueInput {
  code?: string;
  id: number;
}
export interface main_StringFieldUpdateOperationsInput {
  set?: string;
}
export interface main_StringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_StringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_UserCreateNestedManyWithoutRoleInput {}
export interface main_UserCreateOrConnectWithoutRoleInput {}
export interface main_UserCreateWithoutRoleInput {
  avatarUrl?: string;
  createdAt?: string;
  id: string;
  name?: string;
}
export interface main_UserScalarWhereInput {
  OR?: {}[];
}
export interface main_UserUpdateManyMutationInput {}
export interface main_UserUpdateManyWithWhereWithoutRoleInput {}
export interface main_UserUpdateManyWithoutRoleNestedInput {}
export interface main_UserUpdateWithWhereUniqueWithoutRoleInput {}
export interface main_UserUpdateWithoutRoleInput {}
export interface main_UserUpsertWithWhereUniqueWithoutRoleInput {}
export interface main_UserWhereUniqueInput {
  id: string;
}
export interface InjectedSystem__Menu__UpdateOneInput {
  icon?: string;
  id: number;
  label?: string;
  level?: number;
  path?: string;
  sort?: number;
}

export interface System__Menu__UpdateOneResponse {
  data?: System__Menu__UpdateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Menu__UpdateOneResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  };
}

export interface System__Operation__GetManyResponse {
  data?: System__Operation__GetManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Operation__GetManyResponseData {
  data?: {
    createTime?: string;
    id?: string;
    method?: string;
    operationType?: string;
    title?: string;
  }[];
}
export interface System__Role__AddOneInput {
  code: string;
  remark?: string;
}
export interface InternalSystem__Role__AddOneInput {
  code: string;
  remark?: string;
}
export interface InjectedSystem__Role__AddOneInput {
  code: string;
  remark?: string;
}

export interface System__Role__AddOneResponse {
  data?: System__Role__AddOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__AddOneResponseData {
  data?: {
    code?: string;
    id?: number;
    remark?: string;
  };
}
export interface System__Role__BindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
}
export interface InternalSystem__Role__BindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
}
export interface InjectedSystem__Role__BindRoleApisInput {
  allRoles: {}[];
  apis: {}[];
  roleCode: string;
}

export interface System__Role__BindRoleApisResponse {
  data?: System__Role__BindRoleApisResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__BindRoleApisResponseData {
  open_BindRoleApis?: {
    count?: number;
  };
}
export interface System__Role__ConnectOneMenuInput {
  id?: number;
  menuId?: number;
}
export interface InternalSystem__Role__ConnectOneMenuInput {
  id?: number;
  menuId?: number;
}
export interface InjectedSystem__Role__ConnectOneMenuInput {
  id?: number;
  menuId?: number;
}

export interface System__Role__ConnectOneMenuResponse {
  data?: System__Role__ConnectOneMenuResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__ConnectOneMenuResponseData {
  data?: {
    id?: number;
  };
}
export interface System__Role__DeleteManyInput {
  ids?: {}[];
}
export interface InternalSystem__Role__DeleteManyInput {
  ids?: {}[];
}
export interface InjectedSystem__Role__DeleteManyInput {
  ids?: {}[];
}

export interface System__Role__DeleteManyResponse {
  data?: System__Role__DeleteManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__DeleteManyResponseData {
  data?: {
    count?: number;
  };
}
export interface System__Role__DeleteOneInput {
  code: string;
}
export interface InternalSystem__Role__DeleteOneInput {
  code: string;
}
export interface InjectedSystem__Role__DeleteOneInput {
  code: string;
}

export interface System__Role__DeleteOneResponse {
  data?: System__Role__DeleteOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__DeleteOneResponseData {
  data?: {
    id?: number;
  };
}
export interface System__Role__DisconnectOneMenuInput {
  id: number;
  menuId: number;
}
export interface InternalSystem__Role__DisconnectOneMenuInput {
  id: number;
  menuId: number;
}
export interface InjectedSystem__Role__DisconnectOneMenuInput {
  id: number;
  menuId: number;
}

export interface System__Role__DisconnectOneMenuResponse {
  data?: System__Role__DisconnectOneMenuResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__DisconnectOneMenuResponseData {
  data?: {
    id?: number;
  };
}
export interface System__Role__GetBindMenusInput {
  roleId: number;
}
export interface InternalSystem__Role__GetBindMenusInput {
  roleId: number;
}
export interface InjectedSystem__Role__GetBindMenusInput {
  roleId: number;
}

export interface System__Role__GetBindMenusResponse {
  data?: System__Role__GetBindMenusResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetBindMenusResponseData {
  data?: {
    icon?: string;
    id?: number;
    label?: string;
    level?: number;
    parentId?: number;
    path?: string;
    sort?: number;
  }[];
}
export interface System__Role__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface InternalSystem__Role__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}
export interface main_DateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_DateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_IntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_IntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_MenuListRelationFilter {}
export interface main_MenuOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_MenuRelationFilter {}
export interface main_MenuWhereInput {
  OR?: {}[];
}
export interface main_NestedDateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedDateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedIntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedIntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedStringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NestedStringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_PostListRelationFilter {}
export interface main_PostWhereInput {
  OR?: {}[];
}
export interface main_RoleListRelationFilter {}
export interface main_RoleOrderByWithRelationInput {
  code?: string;
  id?: string;
  remark?: string;
}
export interface main_RoleWhereInput {
  OR?: {}[];
}
export interface main_StringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_StringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_UserListRelationFilter {}
export interface main_UserOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_UserRelationFilter {}
export interface main_UserWhereInput {
  OR?: {}[];
}
export interface InjectedSystem__Role__GetListInput {
  orderBy?: {}[];
  skip?: number;
  take?: number;
}

export interface System__Role__GetListResponse {
  data?: System__Role__GetListResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetListResponseData {
  data?: {
    code?: string;
    id?: number;
    remark?: string;
  }[];
  total?: number;
}

export interface System__Role__GetManyResponse {
  data?: System__Role__GetManyResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetManyResponseData {
  data?: {
    code?: string;
    remark?: string;
  }[];
}
export interface System__Role__GetRoleBindApisInput {
  code: string;
}
export interface InternalSystem__Role__GetRoleBindApisInput {
  code: string;
}
export interface InjectedSystem__Role__GetRoleBindApisInput {
  code: string;
}

export interface System__Role__GetRoleBindApisResponse {
  data?: System__Role__GetRoleBindApisResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__GetRoleBindApisResponseData {
  data?: {
    content?: string;
    createTime?: string;
    deleteTime?: string;
    id?: number;
    method?: string;
    operationType?: string;
    remark?: string;
    restUrl?: string;
    roleType?: string;
    roles?: string;
    title?: string;
    updateTime?: string;
  }[];
}
export interface System__Role__UpdateOneInput {
  id: number;
  remark?: string;
}
export interface InternalSystem__Role__UpdateOneInput {
  id: number;
  remark?: string;
}
export interface InjectedSystem__Role__UpdateOneInput {
  id: number;
  remark?: string;
}

export interface System__Role__UpdateOneResponse {
  data?: System__Role__UpdateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__Role__UpdateOneResponseData {
  data?: {
    code?: string;
    id?: number;
    remark?: string;
  };
}
export interface System__User__ConnectRoleInput {
  code: string;
  userId: string;
}
export interface InternalSystem__User__ConnectRoleInput {
  code: string;
  userId: string;
}
export interface InjectedSystem__User__ConnectRoleInput {
  code: string;
  userId: string;
}

export interface System__User__ConnectRoleResponse {
  data?: System__User__ConnectRoleResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__ConnectRoleResponseData {
  data?: {
    avatarUrl?: string;
    createdAt?: string;
    id?: string;
    name?: string;
  };
}

export interface System__User__CountUsersResponse {
  data?: System__User__CountUsersResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__CountUsersResponseData {
  data?: {
    count?: {
      id?: number;
    };
  };
}
export interface System__User__CreateOneInput {
  avatarUrl?: string;
  id: string;
  name: string;
  role?: string;
}
export interface InternalSystem__User__CreateOneInput {
  avatarUrl?: string;
  id: string;
  name: string;
  role?: string;
}
export interface InjectedSystem__User__CreateOneInput {
  avatarUrl?: string;
  id: string;
  name: string;
  role?: string;
}

export interface System__User__CreateOneResponse {
  data?: System__User__CreateOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__CreateOneResponseData {
  data?: {
    id?: string;
  };
}
export interface System__User__DisconnectRoleInput {
  code: string;
  userId: string;
}
export interface InternalSystem__User__DisconnectRoleInput {
  code: string;
  userId: string;
}
export interface InjectedSystem__User__DisconnectRoleInput {
  code: string;
  userId: string;
}

export interface System__User__DisconnectRoleResponse {
  data?: System__User__DisconnectRoleResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__DisconnectRoleResponseData {
  data?: {
    avatarUrl?: string;
    createdAt?: string;
    id?: string;
    name?: string;
  };
}
export interface System__User__GetListInput {
  orderBy?: {}[];
  skip: number;
  take: number;
}
export interface InternalSystem__User__GetListInput {
  orderBy?: {}[];
  skip: number;
  take: number;
}
export interface main_DateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_DateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_IntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_IntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_MenuListRelationFilter {}
export interface main_MenuRelationFilter {}
export interface main_MenuWhereInput {
  OR?: {}[];
}
export interface main_NestedDateTimeFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedDateTimeNullableFilter {
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
}
export interface main_NestedIntFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedIntNullableFilter {
  equals?: number;
  gt?: number;
  gte?: number;
  in?: {}[];
  lt?: number;
  lte?: number;
  notIn?: {}[];
}
export interface main_NestedStringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_NestedStringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_PostListRelationFilter {}
export interface main_PostOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_PostWhereInput {
  OR?: {}[];
}
export interface main_RoleListRelationFilter {}
export interface main_RoleOrderByRelationAggregateInput {
  _count?: string;
}
export interface main_RoleWhereInput {
  OR?: {}[];
}
export interface main_StringFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_StringNullableFilter {
  contains?: string;
  endsWith?: string;
  equals?: string;
  gt?: string;
  gte?: string;
  in?: {}[];
  lt?: string;
  lte?: string;
  notIn?: {}[];
  startsWith?: string;
}
export interface main_UserListRelationFilter {}
export interface main_UserOrderByWithRelationInput {
  avatarUrl?: string;
  createdAt?: string;
  id?: string;
  name?: string;
}
export interface main_UserRelationFilter {}
export interface main_UserWhereInput {
  OR?: {}[];
}
export interface InjectedSystem__User__GetListInput {
  orderBy?: {}[];
  skip: number;
  take: number;
}

export interface System__User__GetListResponse {
  data?: System__User__GetListResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetListResponseData {
  data?: {
    avatar?: string;
    createdAt?: string;
    id?: string;
    name?: string;
  }[];
  total?: number;
}
export interface System__User__GetOneInput {
  id: string;
}
export interface InternalSystem__User__GetOneInput {
  id: string;
}
export interface InjectedSystem__User__GetOneInput {
  id: string;
}

export interface System__User__GetOneResponse {
  data?: System__User__GetOneResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetOneResponseData {
  data?: {
    id?: string;
  };
}
export interface System__User__GetUserRoleInput {
  userId: string;
}
export interface InternalSystem__User__GetUserRoleInput {
  userId: string;
}
export interface InjectedSystem__User__GetUserRoleInput {
  userId: string;
}

export interface System__User__GetUserRoleResponse {
  data?: System__User__GetUserRoleResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface System__User__GetUserRoleResponseData {
  data?: {
    Role?: {
      code?: string;
      remark?: string;
    }[];
  };
}
export interface InternalUser__MeInput {
  equals?: string;
}
export interface InjectedUser__MeInput {
  equals: string;
}

export interface User__MeResponse {
  data?: User__MeResponseData;
  errors?: ReadonlyArray<GraphQLError>;
}
export interface User__MeResponseData {
  data?: {
    avatarUrl?: string;
    id?: string;
    name?: string;
    roles?: {
      code?: string;
    }[];
  };
}
export type JSONValue =
  | string
  | number
  | boolean
  | JSONObject
  | Array<JSONValue>;

export type JSONObject = { [key: string]: JSONValue };

export interface GraphQLError {
  message: string;
  path?: ReadonlyArray<string | number>;
}
