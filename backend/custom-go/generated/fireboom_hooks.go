package generated

import "custom-go/pkg/base"

type (
	Casdoor__GetRolesByIdBody = *base.OperationBody[Casdoor__GetRolesByIdInternalInput, Casdoor__GetRolesByIdResponseData]

	Casdoor__GetSMSProviderBody = *base.OperationBody[Casdoor__GetSMSProviderInternalInput, Casdoor__GetSMSProviderResponseData]

	Casdoor__LoginBody = *base.OperationBody[Casdoor__LoginInternalInput, Casdoor__LoginResponseData]

	Casdoor__RefreshTokenBody = *base.OperationBody[Casdoor__RefreshTokenInternalInput, Casdoor__RefreshTokenResponseData]

	Casdoor__SendCodeBody = *base.OperationBody[Casdoor__SendCodeInternalInput, Casdoor__SendCodeResponseData]

	Casdoor__UpdateSMSProviderBody = *base.OperationBody[Casdoor__UpdateSMSProviderInternalInput, Casdoor__UpdateSMSProviderResponseData]

	Post__CreateCategoryBody = *base.OperationBody[Post__CreateCategoryInternalInput, Post__CreateCategoryResponseData]

	Post__CreateOneBody = *base.OperationBody[Post__CreateOneInternalInput, Post__CreateOneResponseData]

	Post__DeleteCategoryBody = *base.OperationBody[Post__DeleteCategoryInternalInput, Post__DeleteCategoryResponseData]

	Post__DeleteManyBody = *base.OperationBody[Post__DeleteManyInternalInput, Post__DeleteManyResponseData]

	Post__DeleteOneBody = *base.OperationBody[Post__DeleteOneInternalInput, Post__DeleteOneResponseData]

	Post__GetCategoryBody = *base.OperationBody[Post__GetCategoryInternalInput, Post__GetCategoryResponseData]

	Post__GetLikeListBody = *base.OperationBody[Post__GetLikeListInternalInput, Post__GetLikeListResponseData]

	Post__GetListBody = *base.OperationBody[Post__GetListInternalInput, Post__GetListResponseData]

	Post__GetOneBody = *base.OperationBody[Post__GetOneInternalInput, Post__GetOneResponseData]

	Post__GetPostByAuthorBody = *base.OperationBody[Post__GetPostByAuthorInternalInput, Post__GetPostByAuthorResponseData]

	Post__GetPostByCateBody = *base.OperationBody[Post__GetPostByCateInternalInput, Post__GetPostByCateResponseData]

	Post__UpdateCategoryBody = *base.OperationBody[Post__UpdateCategoryInternalInput, Post__UpdateCategoryResponseData]

	Post__UpdateOneBody = *base.OperationBody[Post__UpdateOneInternalInput, Post__UpdateOneResponseData]

	Statistics__MonthlySalesBody = *base.OperationBody[Statistics__MonthlySalesInternalInput, Statistics__MonthlySalesResponseData]

	Statistics__QueryRawBody = *base.OperationBody[Statistics__QueryRawInternalInput, Statistics__QueryRawResponseData]

	Statistics__SaleTypePercentBody = *base.OperationBody[Statistics__SaleTypePercentInternalInput, Statistics__SaleTypePercentResponseData]

	Statistics__SalesTop10Body = *base.OperationBody[Statistics__SalesTop10InternalInput, Statistics__SalesTop10ResponseData]

	Statistics__VisitTrendingBody = *base.OperationBody[Statistics__VisitTrendingInternalInput, Statistics__VisitTrendingResponseData]

	System__GetMenusBody = *base.OperationBody[System__GetMenusInternalInput, System__GetMenusResponseData]

	System__Jwt__AddOneBody = *base.OperationBody[System__Jwt__AddOneInternalInput, System__Jwt__AddOneResponseData]

	System__Jwt__CheckBannedBody = *base.OperationBody[System__Jwt__CheckBannedInternalInput, System__Jwt__CheckBannedResponseData]

	System__Jwt__DeleteOneBody = *base.OperationBody[System__Jwt__DeleteOneInternalInput, System__Jwt__DeleteOneResponseData]

	System__Jwt__GetManyBody = *base.OperationBody[System__Jwt__GetManyInternalInput, System__Jwt__GetManyResponseData]

	System__Jwt__UpdateOneBody = *base.OperationBody[System__Jwt__UpdateOneInternalInput, System__Jwt__UpdateOneResponseData]

	System__Log__ChangeOpenBody = *base.OperationBody[System__Log__ChangeOpenInternalInput, System__Log__ChangeOpenResponseData]

	System__Log__CreateLogBody = *base.OperationBody[System__Log__CreateLogInternalInput, System__Log__CreateLogResponseData]

	System__Log__DeleteLogBody = *base.OperationBody[System__Log__DeleteLogInternalInput, System__Log__DeleteLogResponseData]

	System__Log__DeleteOneBody = *base.OperationBody[System__Log__DeleteOneInternalInput, System__Log__DeleteOneResponseData]

	System__Log__GetAllLogBody = *base.OperationBody[System__Log__GetAllLogInternalInput, System__Log__GetAllLogResponseData]

	System__Log__GetIsOpenBody = *base.OperationBody[System__Log__GetIsOpenInternalInput, System__Log__GetIsOpenResponseData]

	System__Log__GetLikeLogBody = *base.OperationBody[System__Log__GetLikeLogInternalInput, System__Log__GetLikeLogResponseData]

	System__Log__GetLogBody = *base.OperationBody[System__Log__GetLogInternalInput, System__Log__GetLogResponseData]

	System__Menu__CreateOneBody = *base.OperationBody[System__Menu__CreateOneInternalInput, System__Menu__CreateOneResponseData]

	System__Menu__DeleteManyBody = *base.OperationBody[System__Menu__DeleteManyInternalInput, System__Menu__DeleteManyResponseData]

	System__Menu__DeleteOneBody = *base.OperationBody[System__Menu__DeleteOneInternalInput, System__Menu__DeleteOneResponseData]

	System__Menu__GetApiListBody = *base.OperationBody[System__Menu__GetApiListInternalInput, System__Menu__GetApiListResponseData]

	System__Menu__GetApisByMenusBody = *base.OperationBody[System__Menu__GetApisByMenusInternalInput, System__Menu__GetApisByMenusResponseData]

	System__Menu__GetChildrenMenusBody = *base.OperationBody[System__Menu__GetChildrenMenusInternalInput, System__Menu__GetChildrenMenusResponseData]

	System__Menu__GetListBody = *base.OperationBody[System__Menu__GetListInternalInput, System__Menu__GetListResponseData]

	System__Menu__GetManyBody = *base.OperationBody[System__Menu__GetManyInternalInput, System__Menu__GetManyResponseData]

	System__Menu__GetMenuByLevelOrPidBody = *base.OperationBody[System__Menu__GetMenuByLevelOrPidInternalInput, System__Menu__GetMenuByLevelOrPidResponseData]

	System__Menu__GetOneBody = *base.OperationBody[System__Menu__GetOneInternalInput, System__Menu__GetOneResponseData]

	System__Menu__UpdateOneBody = *base.OperationBody[System__Menu__UpdateOneInternalInput, System__Menu__UpdateOneResponseData]

	System__Operation__GetManyBody = *base.OperationBody[System__Operation__GetManyInternalInput, System__Operation__GetManyResponseData]

	System__Perm__CreateManyBody = *base.OperationBody[System__Perm__CreateManyInternalInput, System__Perm__CreateManyResponseData]

	System__Perm__GetBindPermsBody = *base.OperationBody[System__Perm__GetBindPermsInternalInput, System__Perm__GetBindPermsResponseData]

	System__Perm__GetRolePermsBody = *base.OperationBody[System__Perm__GetRolePermsInternalInput, System__Perm__GetRolePermsResponseData]

	System__Role__AddOneBody = *base.OperationBody[System__Role__AddOneInternalInput, System__Role__AddOneResponseData]

	System__Role__BindMenusBody = *base.OperationBody[System__Role__BindMenusInternalInput, System__Role__BindMenusResponseData]

	System__Role__BindRoleApisBody = *base.OperationBody[System__Role__BindRoleApisInternalInput, System__Role__BindRoleApisResponseData]

	System__Role__DeleteManyBody = *base.OperationBody[System__Role__DeleteManyInternalInput, System__Role__DeleteManyResponseData]

	System__Role__DeleteOneBody = *base.OperationBody[System__Role__DeleteOneInternalInput, System__Role__DeleteOneResponseData]

	System__Role__GetListBody = *base.OperationBody[System__Role__GetListInternalInput, System__Role__GetListResponseData]

	System__Role__GetManyBody = *base.OperationBody[System__Role__GetManyInternalInput, System__Role__GetManyResponseData]

	System__Role__GetOneBody = *base.OperationBody[System__Role__GetOneInternalInput, System__Role__GetOneResponseData]

	System__Role__GetRoleByIdBody = *base.OperationBody[System__Role__GetRoleByIdInternalInput, System__Role__GetRoleByIdResponseData]

	System__Role__GetRoleMenuIdBody = *base.OperationBody[System__Role__GetRoleMenuIdInternalInput, System__Role__GetRoleMenuIdResponseData]

	System__Role__SyncDeleteRoleBody = *base.OperationBody[System__Role__SyncDeleteRoleInternalInput, System__Role__SyncDeleteRoleResponseData]

	System__Role__SyncRoleBody = *base.OperationBody[System__Role__SyncRoleInternalInput, System__Role__SyncRoleResponseData]

	System__Role__UnBindMenusBody = *base.OperationBody[System__Role__UnBindMenusInternalInput, System__Role__UnBindMenusResponseData]

	System__Role__UpdateOneBody = *base.OperationBody[System__Role__UpdateOneInternalInput, System__Role__UpdateOneResponseData]

	System__Sub__CreateSubBody = *base.OperationBody[System__Sub__CreateSubInternalInput, System__Sub__CreateSubResponseData]

	System__Sub__SSEBody = *base.OperationBody[System__Sub__SSEInternalInput, System__Sub__SSEResponseData]

	System__User__ConnectRoleBody = *base.OperationBody[System__User__ConnectRoleInternalInput, System__User__ConnectRoleResponseData]

	System__User__CountUsersBody = *base.OperationBody[System__User__CountUsersInternalInput, System__User__CountUsersResponseData]

	System__User__CreateOneBody = *base.OperationBody[System__User__CreateOneInternalInput, System__User__CreateOneResponseData]

	System__User__DeleteOneBody = *base.OperationBody[System__User__DeleteOneInternalInput, System__User__DeleteOneResponseData]

	System__User__DisconnectRoleBody = *base.OperationBody[System__User__DisconnectRoleInternalInput, System__User__DisconnectRoleResponseData]

	System__User__GetAllListBody = *base.OperationBody[System__User__GetAllListInternalInput, System__User__GetAllListResponseData]

	System__User__GetLikeUserBody = *base.OperationBody[System__User__GetLikeUserInternalInput, System__User__GetLikeUserResponseData]

	System__User__GetListBody = *base.OperationBody[System__User__GetListInternalInput, System__User__GetListResponseData]

	System__User__GetOneBody = *base.OperationBody[System__User__GetOneInternalInput, System__User__GetOneResponseData]

	System__User__GetRoleUsersBody = *base.OperationBody[System__User__GetRoleUsersInternalInput, System__User__GetRoleUsersResponseData]

	System__User__GetUserByUserIdBody = *base.OperationBody[System__User__GetUserByUserIdInternalInput, System__User__GetUserByUserIdResponseData]

	System__User__GetUserInfoBody = *base.OperationBody[System__User__GetUserInfoInternalInput, System__User__GetUserInfoResponseData]

	System__User__GetUserRoleBody = *base.OperationBody[System__User__GetUserRoleInternalInput, System__User__GetUserRoleResponseData]

	System__User__UpdateOneBody = *base.OperationBody[System__User__UpdateOneInternalInput, System__User__UpdateOneResponseData]
)

const (
	Casdoor__GetRolesById             base.OperationQueryPath = "Casdoor/GetRolesById"
	Casdoor__GetSMSProvider           base.OperationQueryPath = "Casdoor/GetSMSProvider"
	Post__GetCategory                 base.OperationQueryPath = "Post/GetCategory"
	Post__GetLikeList                 base.OperationQueryPath = "Post/GetLikeList"
	Post__GetList                     base.OperationQueryPath = "Post/GetList"
	Post__GetOne                      base.OperationQueryPath = "Post/GetOne"
	Post__GetPostByCate               base.OperationQueryPath = "Post/GetPostByCate"
	Post__UpdateCategory              base.OperationQueryPath = "Post/UpdateCategory"
	Statistics__MonthlySales          base.OperationQueryPath = "Statistics/MonthlySales"
	Statistics__SaleTypePercent       base.OperationQueryPath = "Statistics/SaleTypePercent"
	Statistics__SalesTop10            base.OperationQueryPath = "Statistics/SalesTop10"
	Statistics__VisitTrending         base.OperationQueryPath = "Statistics/VisitTrending"
	System__GetMenus                  base.OperationQueryPath = "System/GetMenus"
	System__Jwt__CheckBanned          base.OperationQueryPath = "System/Jwt/CheckBanned"
	System__Jwt__GetMany              base.OperationQueryPath = "System/Jwt/GetMany"
	System__Log__GetAllLog            base.OperationQueryPath = "System/Log/GetAllLog"
	System__Log__GetIsOpen            base.OperationQueryPath = "System/Log/GetIsOpen"
	System__Log__GetLikeLog           base.OperationQueryPath = "System/Log/GetLikeLog"
	System__Log__GetLog               base.OperationQueryPath = "System/Log/GetLog"
	System__Menu__GetApiList          base.OperationQueryPath = "System/Menu/GetApiList"
	System__Menu__GetApisByMenus      base.OperationQueryPath = "System/Menu/GetApisByMenus"
	System__Menu__GetChildrenMenus    base.OperationQueryPath = "System/Menu/GetChildrenMenus"
	System__Menu__GetList             base.OperationQueryPath = "System/Menu/GetList"
	System__Menu__GetMany             base.OperationQueryPath = "System/Menu/GetMany"
	System__Menu__GetMenuByLevelOrPid base.OperationQueryPath = "System/Menu/GetMenuByLevelOrPid"
	System__Menu__GetOne              base.OperationQueryPath = "System/Menu/GetOne"
	System__Perm__GetBindPerms        base.OperationQueryPath = "System/Perm/GetBindPerms"
	System__Perm__GetRolePerms        base.OperationQueryPath = "System/Perm/GetRolePerms"
	System__Role__GetList             base.OperationQueryPath = "System/Role/GetList"
	System__Role__GetMany             base.OperationQueryPath = "System/Role/GetMany"
	System__Role__GetOne              base.OperationQueryPath = "System/Role/GetOne"
	System__Role__GetRoleById         base.OperationQueryPath = "System/Role/GetRoleById"
	System__Role__GetRoleMenuId       base.OperationQueryPath = "System/Role/GetRoleMenuId"
	System__Sub__SSE                  base.OperationQueryPath = "System/Sub/SSE"
	System__User__CountUsers          base.OperationQueryPath = "System/User/CountUsers"
	System__User__GetAllList          base.OperationQueryPath = "System/User/GetAllList"
	System__User__GetLikeUser         base.OperationQueryPath = "System/User/GetLikeUser"
	System__User__GetList             base.OperationQueryPath = "System/User/GetList"
	System__User__GetOne              base.OperationQueryPath = "System/User/GetOne"
	System__User__GetRoleUsers        base.OperationQueryPath = "System/User/GetRoleUsers"
	System__User__GetUserByUserId     base.OperationQueryPath = "System/User/GetUserByUserId"
	System__User__GetUserInfo         base.OperationQueryPath = "System/User/GetUserInfo"
	System__User__GetUserRole         base.OperationQueryPath = "System/User/GetUserRole"
)

const (
	Casdoor__Login               base.OperationMutationPath = "Casdoor/Login"
	Casdoor__RefreshToken        base.OperationMutationPath = "Casdoor/RefreshToken"
	Casdoor__SendCode            base.OperationMutationPath = "Casdoor/SendCode"
	Casdoor__UpdateSMSProvider   base.OperationMutationPath = "Casdoor/UpdateSMSProvider"
	Post__CreateCategory         base.OperationMutationPath = "Post/CreateCategory"
	Post__CreateOne              base.OperationMutationPath = "Post/CreateOne"
	Post__DeleteCategory         base.OperationMutationPath = "Post/DeleteCategory"
	Post__DeleteMany             base.OperationMutationPath = "Post/DeleteMany"
	Post__DeleteOne              base.OperationMutationPath = "Post/DeleteOne"
	Post__GetPostByAuthor        base.OperationMutationPath = "Post/GetPostByAuthor"
	Post__UpdateOne              base.OperationMutationPath = "Post/UpdateOne"
	Statistics__QueryRaw         base.OperationMutationPath = "Statistics/QueryRaw"
	System__Jwt__AddOne          base.OperationMutationPath = "System/Jwt/AddOne"
	System__Jwt__DeleteOne       base.OperationMutationPath = "System/Jwt/DeleteOne"
	System__Jwt__UpdateOne       base.OperationMutationPath = "System/Jwt/UpdateOne"
	System__Log__ChangeOpen      base.OperationMutationPath = "System/Log/ChangeOpen"
	System__Log__CreateLog       base.OperationMutationPath = "System/Log/CreateLog"
	System__Log__DeleteLog       base.OperationMutationPath = "System/Log/DeleteLog"
	System__Log__DeleteOne       base.OperationMutationPath = "System/Log/DeleteOne"
	System__Menu__CreateOne      base.OperationMutationPath = "System/Menu/CreateOne"
	System__Menu__DeleteMany     base.OperationMutationPath = "System/Menu/DeleteMany"
	System__Menu__DeleteOne      base.OperationMutationPath = "System/Menu/DeleteOne"
	System__Menu__UpdateOne      base.OperationMutationPath = "System/Menu/UpdateOne"
	System__Operation__GetMany   base.OperationMutationPath = "System/Operation/GetMany"
	System__Perm__CreateMany     base.OperationMutationPath = "System/Perm/CreateMany"
	System__Role__AddOne         base.OperationMutationPath = "System/Role/AddOne"
	System__Role__BindMenus      base.OperationMutationPath = "System/Role/BindMenus"
	System__Role__BindRoleApis   base.OperationMutationPath = "System/Role/BindRoleApis"
	System__Role__DeleteMany     base.OperationMutationPath = "System/Role/DeleteMany"
	System__Role__DeleteOne      base.OperationMutationPath = "System/Role/DeleteOne"
	System__Role__SyncDeleteRole base.OperationMutationPath = "System/Role/SyncDeleteRole"
	System__Role__SyncRole       base.OperationMutationPath = "System/Role/SyncRole"
	System__Role__UnBindMenus    base.OperationMutationPath = "System/Role/UnBindMenus"
	System__Role__UpdateOne      base.OperationMutationPath = "System/Role/UpdateOne"
	System__Sub__CreateSub       base.OperationMutationPath = "System/Sub/CreateSub"
	System__User__ConnectRole    base.OperationMutationPath = "System/User/ConnectRole"
	System__User__CreateOne      base.OperationMutationPath = "System/User/CreateOne"
	System__User__DeleteOne      base.OperationMutationPath = "System/User/DeleteOne"
	System__User__DisconnectRole base.OperationMutationPath = "System/User/DisconnectRole"
	System__User__UpdateOne      base.OperationMutationPath = "System/User/UpdateOne"
	Proxy__bindmenu              base.OperationMutationPath = "proxy/bindmenu"
	Proxy__login                 base.OperationMutationPath = "proxy/login"
	Proxy__menuTree              base.OperationMutationPath = "proxy/menuTree"
	Proxy__perm                  base.OperationMutationPath = "proxy/perm"
	Proxy__unBindMenu            base.OperationMutationPath = "proxy/unBindMenu"
)
