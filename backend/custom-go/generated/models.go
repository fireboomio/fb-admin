package generated

type Casdoor__GetRolesByIdInput struct {
	UserId string `json:"userId"`
}
type Casdoor__GetRolesByIdInternalInput struct {
	Id     int64  `json:"id,omitempty"`
	RoleId int64  `json:"roleId,omitempty"`
	UserId string `json:"userId"`
}
type Casdoor__GetRolesByIdResponseData struct {
	Data []string `json:"data,omitempty"`
}
type Casdoor__GetSMSProviderInput struct {
}
type Casdoor__GetSMSProviderInternalInput struct {
}
type Casdoor__GetSMSProviderResponseData struct {
	Main_findManyprovider []Casdoor__GetSMSProviderResponseData_main_findManyprovider `json:"main_findManyprovider"`
}
type Casdoor__GetSMSProviderResponseData_main_findManyprovider struct {
	Client_id     string `json:"client_id,omitempty"`
	Client_secret string `json:"client_secret,omitempty"`
	Created_time  string `json:"created_time,omitempty"`
	Name          string `json:"name"`
	Owner         string `json:"owner"`
	Sign_name     string `json:"sign_name,omitempty"`
	Template_code string `json:"template_code,omitempty"`
	Type          string `json:"type,omitempty"`
}
type Casdoor__LoginInput struct {
	Code        string `json:"code,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	LoginType   string `json:"loginType,omitempty"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Username    string `json:"username,omitempty"`
}
type Casdoor__LoginInternalInput struct {
	Code        string `json:"code,omitempty"`
	CountryCode string `json:"countryCode,omitempty"`
	LoginType   string `json:"loginType,omitempty"`
	Password    string `json:"password,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Username    string `json:"username,omitempty"`
}
type Casdoor__LoginResponseData struct {
	Casdoor_apiLogin Casdoor__LoginResponseData_casdoor_apiLogin `json:"casdoor_apiLogin,omitempty"`
}
type Casdoor__LoginResponseData_casdoor_apiLogin struct {
	Data    Casdoor__LoginResponseData_casdoor_apiLogin_data `json:"data,omitempty"`
	Msg     string                                           `json:"msg,omitempty"`
	Success bool                                             `json:"success,omitempty"`
}
type Casdoor__LoginResponseData_casdoor_apiLogin_data struct {
	AccessToken  string `json:"accessToken,omitempty"`
	ExpireIn     int64  `json:"expireIn,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
type Casdoor__RefreshTokenInput struct {
	RefreshToken string `json:"refreshToken"`
}
type Casdoor__RefreshTokenInternalInput struct {
	RefreshToken string `json:"refreshToken"`
}
type Casdoor__RefreshTokenResponseData struct {
	Data Casdoor__RefreshTokenResponseData_data `json:"data,omitempty"`
}
type Casdoor__RefreshTokenResponseData_data struct {
	Data    Casdoor__RefreshTokenResponseData_data_data `json:"data,omitempty"`
	Msg     string                                      `json:"msg,omitempty"`
	Success bool                                        `json:"success,omitempty"`
}
type Casdoor__RefreshTokenResponseData_data_data struct {
	AccessToken  string `json:"accessToken,omitempty"`
	ExpireIn     int64  `json:"expireIn,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
type Casdoor__SendCodeInput struct {
	CountryCode string `json:"countryCode,omitempty"`
	Dest        string `json:"dest"`
}
type Casdoor__SendCodeInternalInput struct {
	CountryCode string `json:"countryCode,omitempty"`
	Dest        string `json:"dest"`
}
type Casdoor__SendCodeResponseData struct {
	Data Casdoor__SendCodeResponseData_data `json:"data,omitempty"`
}
type Casdoor__SendCodeResponseData_data struct {
	Msg string `json:"msg,omitempty"`
}
type Casdoor__UpdateSMSProviderInput struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	SignName     string `json:"signName"`
	TemplateCode string `json:"templateCode"`
}
type Casdoor__UpdateSMSProviderInternalInput struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	SignName     string `json:"signName"`
	TemplateCode string `json:"templateCode"`
}
type Casdoor__UpdateSMSProviderResponseData struct {
	Casdoor_apiUpdateProvider Casdoor__UpdateSMSProviderResponseData_casdoor_apiUpdateProvider `json:"casdoor_apiUpdateProvider,omitempty"`
}
type Casdoor__UpdateSMSProviderResponseData_casdoor_apiUpdateProvider struct {
	Msg    string `json:"msg,omitempty"`
	Status int64  `json:"status,omitempty"`
}
type Post__CreateCategoryInput struct {
	Name string `json:"name,omitempty"`
}
type Post__CreateCategoryInternalInput struct {
	Name string `json:"name,omitempty"`
}
type Post__CreateCategoryResponseData struct {
	Main_createOneCategory Post__CreateCategoryResponseData_main_createOneCategory `json:"main_createOneCategory,omitempty"`
}
type Post__CreateCategoryResponseData_main_createOneCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__CreateOneInput struct {
	CateId      int64  `json:"cateId"`
	Content     string `json:"content,omitempty"`
	Poster      string `json:"poster,omitempty"`
	PublishedAt string `json:"publishedAt,omitempty"`
	Title       string `json:"title"`
	UserId      int64  `json:"userId"`
	Username    string `json:"username"`
}
type Post__CreateOneInternalInput struct {
	CateId      int64  `json:"cateId"`
	Content     string `json:"content,omitempty"`
	Poster      string `json:"poster,omitempty"`
	PublishedAt string `json:"publishedAt,omitempty"`
	Title       string `json:"title"`
	UserId      int64  `json:"userId"`
	Username    string `json:"username"`
}
type Post__CreateOneResponseData struct {
	Data Post__CreateOneResponseData_data `json:"data,omitempty"`
}
type Post__CreateOneResponseData_data struct {
	Category     Post__CreateOneResponseData_data_Category `json:"Category,omitempty"`
	Author       string                                    `json:"author"`
	Id           int64                                     `json:"id"`
	Poster       string                                    `json:"poster,omitempty"`
	Published_at string                                    `json:"published_at,omitempty"`
	Title        string                                    `json:"title"`
}
type Post__CreateOneResponseData_data_Category struct {
	Id int64 `json:"id"`
}
type Post__DeleteCategoryInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__DeleteCategoryInternalInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__DeleteCategoryResponseData struct {
	Main_deleteOneCategory Post__DeleteCategoryResponseData_main_deleteOneCategory `json:"main_deleteOneCategory,omitempty"`
}
type Post__DeleteCategoryResponseData_main_deleteOneCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__DeleteManyInput struct {
	Ids []int64 `json:"ids"`
}
type Post__DeleteManyInternalInput struct {
	Ids []int64 `json:"ids"`
}
type Post__DeleteManyResponseData struct {
	Data Post__DeleteManyResponseData_data `json:"data,omitempty"`
}
type Post__DeleteManyResponseData_data struct {
	Count int64 `json:"count"`
}
type Post__DeleteOneInput struct {
	Id int64 `json:"id"`
}
type Post__DeleteOneInternalInput struct {
	Id int64 `json:"id"`
}
type Post__DeleteOneResponseData struct {
	Data Post__DeleteOneResponseData_data `json:"data,omitempty"`
}
type Post__DeleteOneResponseData_data struct {
	Id int64 `json:"id"`
}
type Post__GetCategoryInput struct {
	Skip int64 `json:"skip,omitempty"`
	Take int64 `json:"take,omitempty"`
}
type Post__GetCategoryInternalInput struct {
	Skip int64 `json:"skip,omitempty"`
	Take int64 `json:"take,omitempty"`
}
type Post__GetCategoryResponseData struct {
	Main_findManyCategory []Post__GetCategoryResponseData_main_findManyCategory `json:"main_findManyCategory"`
}
type Post__GetCategoryResponseData_main_findManyCategory struct {
	Id   int64                                                      `json:"id"`
	Name string                                                     `json:"name"`
	Post []Post__GetCategoryResponseData_main_findManyCategory_post `json:"post,omitempty"`
}
type Post__GetCategoryResponseData_main_findManyCategory_post struct {
	Author       string `json:"author"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title"`
}
type Post__GetLikeListInput struct {
	Author  string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
	Id      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
}
type Post__GetLikeListInternalInput struct {
	Author  string `json:"author,omitempty"`
	Content string `json:"content,omitempty"`
	Id      int64  `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
}
type Post__GetLikeListResponseData struct {
	Main_findManypost   []Post__GetLikeListResponseData_main_findManypost `json:"main_findManypost"`
	Main_findUniquepost Post__GetLikeListResponseData_main_findUniquepost `json:"main_findUniquepost,omitempty"`
}
type Post__GetLikeListResponseData_main_findManypost struct {
	Category     Post__GetLikeListResponseData_main_findManypost_Category `json:"Category,omitempty"`
	Author       string                                                   `json:"author"`
	Content      string                                                   `json:"content,omitempty"`
	Id           int64                                                    `json:"id"`
	Poster       string                                                   `json:"poster,omitempty"`
	Published_at string                                                   `json:"published_at,omitempty"`
	Title        string                                                   `json:"title"`
}
type Post__GetLikeListResponseData_main_findManypost_Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetLikeListResponseData_main_findUniquepost struct {
	Category     Post__GetLikeListResponseData_main_findUniquepost_Category `json:"Category,omitempty"`
	Author       string                                                     `json:"author"`
	Content      string                                                     `json:"content,omitempty"`
	Id           int64                                                      `json:"id"`
	Poster       string                                                     `json:"poster,omitempty"`
	Published_at string                                                     `json:"published_at,omitempty"`
	Title        string                                                     `json:"title"`
}
type Post__GetLikeListResponseData_main_findUniquepost_Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetListInput struct {
	Query *Main_postWhereInput `json:"query,omitempty"`
	Skip  int64                `json:"skip,omitempty"`
	Take  int64                `json:"take,omitempty"`
}
type Post__GetListInternalInput struct {
	Query *Main_postWhereInput `json:"query,omitempty"`
	Skip  int64                `json:"skip,omitempty"`
	Take  int64                `json:"take,omitempty"`
}
type Post__GetListResponseData struct {
	Data  []Post__GetListResponseData_data `json:"data"`
	Total int64                            `json:"total,omitempty"`
}
type Post__GetListResponseData_data struct {
	Category     Post__GetListResponseData_data_Category `json:"Category,omitempty"`
	Author       string                                  `json:"author"`
	Content      string                                  `json:"content,omitempty"`
	Id           int64                                   `json:"id"`
	Poster       string                                  `json:"poster,omitempty"`
	Published_at string                                  `json:"published_at,omitempty"`
	Title        string                                  `json:"title"`
}
type Post__GetListResponseData_data_Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetOneInput struct {
	Id int64 `json:"id"`
}
type Post__GetOneInternalInput struct {
	Id int64 `json:"id"`
}
type Post__GetOneResponseData struct {
	Data Post__GetOneResponseData_data `json:"data,omitempty"`
}
type Post__GetOneResponseData_data struct {
	Category     Post__GetOneResponseData_data_Category `json:"Category,omitempty"`
	Author       string                                 `json:"author"`
	Content      string                                 `json:"content,omitempty"`
	Id           int64                                  `json:"id"`
	Poster       string                                 `json:"poster,omitempty"`
	Published_at string                                 `json:"published_at,omitempty"`
	Title        string                                 `json:"title"`
	User         Post__GetOneResponseData_data_user     `json:"user,omitempty"`
}
type Post__GetOneResponseData_data_Category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetOneResponseData_data_user struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__UpdateCategoryInput struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Post__UpdateCategoryInternalInput struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Post__UpdateCategoryResponseData struct {
	Main_updateOneCategory Post__UpdateCategoryResponseData_main_updateOneCategory `json:"main_updateOneCategory,omitempty"`
}
type Post__UpdateCategoryResponseData_main_updateOneCategory struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__UpdateOneInput struct {
	Author  string `json:"author,omitempty"`
	Cate    int64  `json:"cate,omitempty"`
	Content string `json:"content"`
	Id      int64  `json:"id"`
	Poster  string `json:"poster,omitempty"`
	Title   string `json:"title,omitempty"`
	UserId  int64  `json:"userId,omitempty"`
}
type Post__UpdateOneInternalInput struct {
	Author  string `json:"author,omitempty"`
	Cate    int64  `json:"cate,omitempty"`
	Content string `json:"content"`
	Id      int64  `json:"id"`
	Poster  string `json:"poster,omitempty"`
	Title   string `json:"title,omitempty"`
	UserId  int64  `json:"userId,omitempty"`
}
type Post__UpdateOneResponseData struct {
	Data Post__UpdateOneResponseData_data `json:"data,omitempty"`
}
type Post__UpdateOneResponseData_data struct {
	Author       string                                `json:"author"`
	Content      string                                `json:"content,omitempty"`
	Id           int64                                 `json:"id"`
	Poster       string                                `json:"poster,omitempty"`
	Published_at string                                `json:"published_at,omitempty"`
	Title        string                                `json:"title"`
	User         Post__UpdateOneResponseData_data_user `json:"user,omitempty"`
}
type Post__UpdateOneResponseData_data_user struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Statistics__MonthlySalesInput struct {
}
type Statistics__MonthlySalesInternalInput struct {
}
type Statistics__MonthlySalesResponseData struct {
	Data []Statistics__MonthlySalesResponseData_data `json:"data,omitempty"`
}
type Statistics__MonthlySalesResponseData_data struct {
	Months     string  `json:"months,omitempty"`
	TotalSales float64 `json:"totalSales,omitempty"`
}
type Statistics__QueryRawInput struct {
	Query string `json:"query"`
}
type Statistics__QueryRawInternalInput struct {
	Query string `json:"query"`
}
type Statistics__QueryRawResponseData struct {
	Main_queryRaw any `json:"main_queryRaw,omitempty"`
}
type Statistics__SaleTypePercentInput struct {
}
type Statistics__SaleTypePercentInternalInput struct {
}
type Statistics__SaleTypePercentResponseData struct {
	Data []Statistics__SaleTypePercentResponseData_data `json:"data,omitempty"`
}
type Statistics__SaleTypePercentResponseData_data struct {
	TotalSales float64 `json:"totalSales,omitempty"`
	TypeId     int64   `json:"typeId,omitempty"`
	TypeName   string  `json:"typeName,omitempty"`
}
type Statistics__SalesTop10Input struct {
}
type Statistics__SalesTop10InternalInput struct {
}
type Statistics__SalesTop10ResponseData struct {
	Data []Statistics__SalesTop10ResponseData_data `json:"data,omitempty"`
}
type Statistics__SalesTop10ResponseData_data struct {
	ShopName   string  `json:"shopName,omitempty"`
	TotalSales float64 `json:"totalSales,omitempty"`
}
type Statistics__VisitTrendingInput struct {
}
type Statistics__VisitTrendingInternalInput struct {
}
type Statistics__VisitTrendingResponseData struct {
	Data []Statistics__VisitTrendingResponseData_data `json:"data,omitempty"`
}
type Statistics__VisitTrendingResponseData_data struct {
	Count int64  `json:"count,omitempty"`
	Days  string `json:"days,omitempty"`
}
type System__GetMenusInput struct {
	Pid int64 `json:"pid,omitempty"`
}
type System__GetMenusInternalInput struct {
	Pid int64 `json:"pid,omitempty"`
}
type System__GetMenusResponseData struct {
	Data []System__GetMenusResponseData_data `json:"data"`
}
type System__GetMenusResponseData_data struct {
	Children []System__GetMenusResponseData_data_children `json:"children,omitempty"`
	Id       int64                                        `json:"id"`
	Label    string                                       `json:"label"`
	Name     string                                       `json:"name,omitempty"`
	Path     string                                       `json:"path,omitempty"`
	Sort     int64                                        `json:"sort"`
}
type System__GetMenusResponseData_data_children struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
	Name  string `json:"name,omitempty"`
	Path  string `json:"path,omitempty"`
	Sort  int64  `json:"sort"`
}
type System__Log__ChangeOpenInput struct {
	Set bool `json:"set,omitempty"`
}
type System__Log__ChangeOpenInternalInput struct {
	Set bool `json:"set,omitempty"`
}
type System__Log__ChangeOpenResponseData struct {
	Main_updateOnedic System__Log__ChangeOpenResponseData_main_updateOnedic `json:"main_updateOnedic,omitempty"`
}
type System__Log__ChangeOpenResponseData_main_updateOnedic struct {
	Id     int64 `json:"id"`
	IsOpen bool  `json:"isOpen"`
}
type System__Log__CreateLogInput struct {
	Code      string `json:"code,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	Ua        string `json:"ua,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UserId    string `json:"userId,omitempty"`
	UserName  string `json:"userName,omitempty"`
}
type System__Log__CreateLogInternalInput struct {
	Code      string `json:"code,omitempty"`
	Ip        string `json:"ip,omitempty"`
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	Ua        string `json:"ua,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UserId    string `json:"userId,omitempty"`
	UserName  string `json:"userName,omitempty"`
}
type System__Log__CreateLogResponseData struct {
	Main_createOneapilog System__Log__CreateLogResponseData_main_createOneapilog `json:"main_createOneapilog,omitempty"`
}
type System__Log__CreateLogResponseData_main_createOneapilog struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
}
type System__Log__DeleteLogInput struct {
	Equals []int64 `json:"equals,omitempty"`
}
type System__Log__DeleteLogInternalInput struct {
	Equals []int64 `json:"equals,omitempty"`
}
type System__Log__DeleteLogResponseData struct {
	Main_deleteManyapilog System__Log__DeleteLogResponseData_main_deleteManyapilog `json:"main_deleteManyapilog,omitempty"`
}
type System__Log__DeleteLogResponseData_main_deleteManyapilog struct {
	Count int64 `json:"count"`
}
type System__Log__GetAllLogInput struct {
}
type System__Log__GetAllLogInternalInput struct {
}
type System__Log__GetAllLogResponseData struct {
	LogNum System__Log__GetAllLogResponseData_logNum `json:"logNum"`
}
type System__Log__GetAllLogResponseData_logNum struct {
	Count System__Log__GetAllLogResponseData_logNum__count `json:"_count,omitempty"`
}
type System__Log__GetAllLogResponseData_logNum__count struct {
	Id int64 `json:"id"`
}
type System__Log__GetIsOpenInput struct {
}
type System__Log__GetIsOpenInternalInput struct {
}
type System__Log__GetIsOpenResponseData struct {
	Main_findFirstdic System__Log__GetIsOpenResponseData_main_findFirstdic `json:"main_findFirstdic,omitempty"`
}
type System__Log__GetIsOpenResponseData_main_findFirstdic struct {
	Id     int64 `json:"id"`
	IsOpen bool  `json:"isOpen"`
}
type System__Log__GetLikeLogInput struct {
	Id     int64  `json:"id,omitempty"`
	Ip     string `json:"ip,omitempty"`
	Method string `json:"method,omitempty"`
	Name   string `json:"name,omitempty"`
	Path   string `json:"path,omitempty"`
}
type System__Log__GetLikeLogInternalInput struct {
	Id     int64  `json:"id,omitempty"`
	Ip     string `json:"ip,omitempty"`
	Method string `json:"method,omitempty"`
	Name   string `json:"name,omitempty"`
	Path   string `json:"path,omitempty"`
}
type System__Log__GetLikeLogResponseData struct {
	Main_findManyapilog   []System__Log__GetLikeLogResponseData_main_findManyapilog `json:"main_findManyapilog"`
	Main_findUniqueapilog System__Log__GetLikeLogResponseData_main_findUniqueapilog `json:"main_findUniqueapilog,omitempty"`
}
type System__Log__GetLikeLogResponseData_main_findManyapilog struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserName  string `json:"userName"`
}
type System__Log__GetLikeLogResponseData_main_findUniqueapilog struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserName  string `json:"userName"`
}
type System__Log__GetLogInput struct {
	OrderBy []*Main_apilogOrderByWithRelationInput `json:"orderBy,omitempty"`
	Skip    int64                                  `json:"skip,omitempty"`
	Take    int64                                  `json:"take,omitempty"`
}
type System__Log__GetLogInternalInput struct {
	OrderBy []*Main_apilogOrderByWithRelationInput `json:"orderBy,omitempty"`
	Skip    int64                                  `json:"skip,omitempty"`
	Take    int64                                  `json:"take,omitempty"`
}
type System__Log__GetLogResponseData struct {
	Data []System__Log__GetLogResponseData_data `json:"data"`
}
type System__Log__GetLogResponseData_data struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
}
type System__Menu__CreateOneInput struct {
	ApiId       string `json:"apiId,omitempty"`
	Create_time string `json:"create_time"`
	Icon        string `json:"icon,omitempty"`
	Label       string `json:"label"`
	Level       int64  `json:"level"`
	MenuType    string `json:"menuType"`
	ParentId    int64  `json:"parentId,omitempty"`
	Path        string `json:"path"`
	Perms       string `json:"perms,omitempty"`
	Sort        int64  `json:"sort"`
}
type System__Menu__CreateOneInternalInput struct {
	ApiId       string `json:"apiId,omitempty"`
	Create_time string `json:"create_time"`
	Icon        string `json:"icon,omitempty"`
	Label       string `json:"label"`
	Level       int64  `json:"level"`
	MenuType    string `json:"menuType"`
	ParentId    int64  `json:"parentId,omitempty"`
	Path        string `json:"path"`
	Perms       string `json:"perms,omitempty"`
	Sort        int64  `json:"sort"`
}
type System__Menu__CreateOneResponseData struct {
	Data System__Menu__CreateOneResponseData_data `json:"data,omitempty"`
}
type System__Menu__CreateOneResponseData_data struct {
	Id int64 `json:"id"`
}
type System__Menu__DeleteManyInput struct {
	Ids []int64 `json:"ids"`
}
type System__Menu__DeleteManyInternalInput struct {
	Ids []int64 `json:"ids"`
}
type System__Menu__DeleteManyResponseData struct {
	Data System__Menu__DeleteManyResponseData_data `json:"data,omitempty"`
}
type System__Menu__DeleteManyResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Menu__DeleteOneInput struct {
	Id int64 `json:"id"`
}
type System__Menu__DeleteOneInternalInput struct {
	Id int64 `json:"id"`
}
type System__Menu__DeleteOneResponseData struct {
	Data System__Menu__DeleteOneResponseData_data `json:"data,omitempty"`
}
type System__Menu__DeleteOneResponseData_data struct {
	Id int64 `json:"id"`
}
type System__Menu__GetApiListInput struct {
}
type System__Menu__GetApiListInternalInput struct {
}
type System__Menu__GetApiListResponseData struct {
	Data []System__Menu__GetApiListResponseData_data `json:"data,omitempty"`
}
type System__Menu__GetApiListResponseData_data struct {
	Id   int64  `json:"id,omitempty"`
	Path string `json:"path,omitempty"`
}
type System__Menu__GetApisByMenusInput struct {
	MenuIds []int64 `json:"menuIds"`
}
type System__Menu__GetApisByMenusInternalInput struct {
	MenuIds []int64 `json:"menuIds"`
}
type System__Menu__GetApisByMenusResponseData struct {
	Data []string `json:"data,omitempty"`
}
type System__Menu__GetChildrenMenusInput struct {
	Pid int64 `json:"pid"`
}
type System__Menu__GetChildrenMenusInternalInput struct {
	Pid int64 `json:"pid"`
}
type System__Menu__GetChildrenMenusResponseData struct {
	Data []System__Menu__GetChildrenMenusResponseData_data `json:"data"`
}
type System__Menu__GetChildrenMenusResponseData_data struct {
	Id    int64  `json:"id"`
	Label string `json:"label"`
	Level int64  `json:"level"`
	Name  string `json:"name,omitempty"`
	Path  string `json:"path,omitempty"`
	Sort  int64  `json:"sort"`
}
type System__Menu__GetListInput struct {
	OrderBy []*Main_menuOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_menuWhereInput                 `json:"query,omitempty"`
	Skip    int64                                `json:"skip,omitempty"`
	Take    int64                                `json:"take,omitempty"`
}
type System__Menu__GetListInternalInput struct {
	OrderBy []*Main_menuOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_menuWhereInput                 `json:"query,omitempty"`
	Skip    int64                                `json:"skip,omitempty"`
	Take    int64                                `json:"take,omitempty"`
}
type System__Menu__GetListResponseData struct {
	Data  []System__Menu__GetListResponseData_data `json:"data"`
	Total int64                                    `json:"total,omitempty"`
}
type System__Menu__GetListResponseData_data struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Level    int64  `json:"level"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort"`
}
type System__Menu__GetManyInput struct {
}
type System__Menu__GetManyInternalInput struct {
}
type System__Menu__GetManyResponseData struct {
	Data []System__Menu__GetManyResponseData_data `json:"data"`
}
type System__Menu__GetManyResponseData_data struct {
	Api_id      string `json:"api_id,omitempty"`
	Create_time string `json:"create_time,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Id          int64  `json:"id"`
	Is_bottom   int64  `json:"is_bottom"`
	Label       string `json:"label"`
	Level       int64  `json:"level"`
	Menu_type   string `json:"menu_type,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	Path        string `json:"path,omitempty"`
	Perms       string `json:"perms,omitempty"`
	Sort        int64  `json:"sort"`
}
type System__Menu__GetMenuByLevelOrPidInput struct {
	Level int64 `json:"level,omitempty"`
	Pid   int64 `json:"pid,omitempty"`
}
type System__Menu__GetMenuByLevelOrPidInternalInput struct {
	Level int64 `json:"level,omitempty"`
	Pid   int64 `json:"pid,omitempty"`
}
type System__Menu__GetMenuByLevelOrPidResponseData struct {
	Data []System__Menu__GetMenuByLevelOrPidResponseData_data `json:"data"`
}
type System__Menu__GetMenuByLevelOrPidResponseData_data struct {
	Id        int64  `json:"id"`
	Is_bottom int64  `json:"is_bottom"`
	Label     string `json:"label"`
}
type System__Menu__GetOneInput struct {
	Id int64 `json:"id"`
}
type System__Menu__GetOneInternalInput struct {
	Id int64 `json:"id"`
}
type System__Menu__GetOneResponseData struct {
	Data System__Menu__GetOneResponseData_data `json:"data,omitempty"`
}
type System__Menu__GetOneResponseData_data struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Level    int64  `json:"level"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort"`
}
type System__Menu__UpdateOneInput struct {
	Icon  string `json:"icon,omitempty"`
	Id    int64  `json:"id"`
	Label string `json:"label,omitempty"`
	Level int64  `json:"level,omitempty"`
	Path  string `json:"path,omitempty"`
	Sort  int64  `json:"sort,omitempty"`
}
type System__Menu__UpdateOneInternalInput struct {
	Icon  string `json:"icon,omitempty"`
	Id    int64  `json:"id"`
	Label string `json:"label,omitempty"`
	Level int64  `json:"level,omitempty"`
	Path  string `json:"path,omitempty"`
	Sort  int64  `json:"sort,omitempty"`
}
type System__Menu__UpdateOneResponseData struct {
	Data System__Menu__UpdateOneResponseData_data `json:"data,omitempty"`
}
type System__Menu__UpdateOneResponseData_data struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Level    int64  `json:"level"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort"`
}
type System__Operation__GetManyInput struct {
}
type System__Operation__GetManyInternalInput struct {
}
type System__Operation__GetManyResponseData struct {
	Data []System__Operation__GetManyResponseData_data `json:"data,omitempty"`
}
type System__Operation__GetManyResponseData_data struct {
	CreateTime    string `json:"createTime,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
	Id            int64  `json:"id,omitempty"`
	LiveQuery     bool   `json:"liveQuery,omitempty"`
	Method        string `json:"method,omitempty"`
	OperationType string `json:"operationType,omitempty"`
	Title         string `json:"title,omitempty"`
}
type System__Perm__CreateManyInput struct {
	Data []*Main_permissionCreateManyInput `json:"data"`
}
type System__Perm__CreateManyInternalInput struct {
	Data []*Main_permissionCreateManyInput `json:"data"`
}
type System__Perm__CreateManyResponseData struct {
	Data System__Perm__CreateManyResponseData_data `json:"data,omitempty"`
}
type System__Perm__CreateManyResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Perm__GetBindPermsInput struct {
}
type System__Perm__GetBindPermsInternalInput struct {
}
type System__Perm__GetBindPermsResponseData struct {
	Data []System__Perm__GetBindPermsResponseData_data `json:"data"`
}
type System__Perm__GetBindPermsResponseData_data struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Enabled   int64  `json:"enabled"`
	Method    string `json:"method"`
	Path      string `json:"path"`
}
type System__Perm__GetRolePermsInput struct {
	Code []string `json:"code"`
}
type System__Perm__GetRolePermsInternalInput struct {
	Code   []string `json:"code"`
	MenuId int64    `json:"menuId,omitempty"`
	RoleId int64    `json:"roleId,omitempty"`
}
type System__Perm__GetRolePermsResponseData struct {
	Data []System__Perm__GetRolePermsResponseData_data `json:"data"`
}
type System__Perm__GetRolePermsResponseData_data struct {
	Join System__Perm__GetRolePermsResponseData_data__join `json:"_join"`
	Id   int64                                             `json:"id"`
}
type System__Perm__GetRolePermsResponseData_data__join struct {
	Data []string `json:"data,omitempty"`
}
type System__Role__AddOneInput struct {
	Code   string `json:"code"`
	Remark string `json:"remark"`
}
type System__Role__AddOneInternalInput struct {
	Code   string `json:"code"`
	Remark string `json:"remark"`
}
type System__Role__AddOneResponseData struct {
	Data System__Role__AddOneResponseData_data `json:"data,omitempty"`
}
type System__Role__AddOneResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__Role__BindMenusInput struct {
	Data []*Main_menu_roleCreateManyInput `json:"data"`
}
type System__Role__BindMenusInternalInput struct {
	Data []*Main_menu_roleCreateManyInput `json:"data"`
}
type System__Role__BindMenusResponseData struct {
	Data System__Role__BindMenusResponseData_data `json:"data,omitempty"`
}
type System__Role__BindMenusResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Role__BindRoleApisInput struct {
	AllRoles []string `json:"allRoles"`
	Apis     []int64  `json:"apis"`
	RoleCode string   `json:"roleCode"`
}
type System__Role__BindRoleApisInternalInput struct {
	AllRoles []string `json:"allRoles"`
	Apis     []int64  `json:"apis"`
	RoleCode string   `json:"roleCode"`
}
type System__Role__BindRoleApisResponseData struct {
	System_bindRoleApis System__Role__BindRoleApisResponseData_system_bindRoleApis `json:"system_bindRoleApis,omitempty"`
}
type System__Role__BindRoleApisResponseData_system_bindRoleApis struct {
	Count int64 `json:"count,omitempty"`
}
type System__Role__DeleteManyInput struct {
	Ids []int64 `json:"ids,omitempty"`
}
type System__Role__DeleteManyInternalInput struct {
	Ids []int64 `json:"ids,omitempty"`
}
type System__Role__DeleteManyResponseData struct {
	Data System__Role__DeleteManyResponseData_data `json:"data,omitempty"`
}
type System__Role__DeleteManyResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Role__DeleteOneInput struct {
	Code string `json:"code"`
}
type System__Role__DeleteOneInternalInput struct {
	Code string `json:"code"`
}
type System__Role__DeleteOneResponseData struct {
	Data System__Role__DeleteOneResponseData_data `json:"data,omitempty"`
}
type System__Role__DeleteOneResponseData_data struct {
	Id int64 `json:"id"`
}
type System__Role__GetListInput struct {
	OrderBy []*Main_roleOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_roleWhereInput                 `json:"query,omitempty"`
	Skip    int64                                `json:"skip,omitempty"`
	Take    int64                                `json:"take,omitempty"`
}
type System__Role__GetListInternalInput struct {
	OrderBy []*Main_roleOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_roleWhereInput                 `json:"query,omitempty"`
	Skip    int64                                `json:"skip,omitempty"`
	Take    int64                                `json:"take,omitempty"`
}
type System__Role__GetListResponseData struct {
	Data  []System__Role__GetListResponseData_data `json:"data"`
	Total int64                                    `json:"total,omitempty"`
}
type System__Role__GetListResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__Role__GetManyInput struct {
}
type System__Role__GetManyInternalInput struct {
}
type System__Role__GetManyResponseData struct {
	Data []System__Role__GetManyResponseData_data `json:"data"`
}
type System__Role__GetManyResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__Role__GetOneInput struct {
	Id int64 `json:"id"`
}
type System__Role__GetOneInternalInput struct {
	Id int64 `json:"id"`
}
type System__Role__GetOneResponseData struct {
	Data System__Role__GetOneResponseData_data `json:"data,omitempty"`
}
type System__Role__GetOneResponseData_data struct {
	Code string `json:"code"`
}
type System__Role__GetRoleMenuIdInput struct {
	RoleId int64 `json:"roleId"`
}
type System__Role__GetRoleMenuIdInternalInput struct {
	RoleId int64 `json:"roleId"`
}
type System__Role__GetRoleMenuIdResponseData struct {
	Data []int64 `json:"data,omitempty"`
}
type System__Role__UnBindMenusInput struct {
	MenuIds []int64 `json:"menuIds"`
	RoleId  int64   `json:"roleId"`
}
type System__Role__UnBindMenusInternalInput struct {
	MenuIds []int64 `json:"menuIds"`
	RoleId  int64   `json:"roleId"`
}
type System__Role__UnBindMenusResponseData struct {
	Main_deleteManymenu_role System__Role__UnBindMenusResponseData_main_deleteManymenu_role `json:"main_deleteManymenu_role,omitempty"`
}
type System__Role__UnBindMenusResponseData_main_deleteManymenu_role struct {
	Count int64 `json:"count"`
}
type System__Role__UnBindRoleApisInput struct {
	Apis     []int64 `json:"apis"`
	RoleCode string  `json:"roleCode"`
}
type System__Role__UnBindRoleApisInternalInput struct {
	Apis     []int64 `json:"apis"`
	RoleCode string  `json:"roleCode"`
}
type System__Role__UnBindRoleApisResponseData struct {
	System_unBindRoleApis System__Role__UnBindRoleApisResponseData_system_unBindRoleApis `json:"system_unBindRoleApis,omitempty"`
}
type System__Role__UnBindRoleApisResponseData_system_unBindRoleApis struct {
	Count int64 `json:"count,omitempty"`
}
type System__Role__UpdateOneInput struct {
	Id     int64  `json:"id"`
	Remark string `json:"remark,omitempty"`
}
type System__Role__UpdateOneInternalInput struct {
	Id     int64  `json:"id"`
	Remark string `json:"remark,omitempty"`
}
type System__Role__UpdateOneResponseData struct {
	Data System__Role__UpdateOneResponseData_data `json:"data,omitempty"`
}
type System__Role__UpdateOneResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__Sub__CreateSubInput struct {
	Create_role string `json:"create_role,omitempty"`
	Message     string `json:"message,omitempty"`
	Target_role string `json:"target_role,omitempty"`
	Type        string `json:"type,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
type System__Sub__CreateSubInternalInput struct {
	Create_role string `json:"create_role,omitempty"`
	Message     string `json:"message,omitempty"`
	Target_role string `json:"target_role,omitempty"`
	Type        string `json:"type,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}
type System__Sub__CreateSubResponseData struct {
	Main_createOnesub System__Sub__CreateSubResponseData_main_createOnesub `json:"main_createOnesub,omitempty"`
}
type System__Sub__CreateSubResponseData_main_createOnesub struct {
	Create_role string `json:"create_role"`
	Id          int64  `json:"id"`
	Message     string `json:"message"`
	Target_role string `json:"target_role"`
	UpdatedAt   string `json:"updatedAt"`
}
type System__Sub__SSEInput struct {
	Roles string `json:"roles"`
}
type System__Sub__SSEInternalInput struct {
	Roles string `json:"roles"`
}
type System__Sub__SSEResponseData struct {
	Data []System__Sub__SSEResponseData_data `json:"data"`
}
type System__Sub__SSEResponseData_data struct {
	Create_role string `json:"create_role"`
	Id          int64  `json:"id"`
	Message     string `json:"message"`
	Target_role string `json:"target_role"`
	Type        string `json:"type"`
	UpdatedAt   string `json:"updatedAt"`
}
type System__User__ConnectRoleInput struct {
	RoleId int64 `json:"roleId"`
	UserId int64 `json:"userId"`
}
type System__User__ConnectRoleInternalInput struct {
	RoleId int64 `json:"roleId"`
	UserId int64 `json:"userId"`
}
type System__User__ConnectRoleResponseData struct {
	Data System__User__ConnectRoleResponseData_data `json:"data,omitempty"`
}
type System__User__ConnectRoleResponseData_data struct {
	Id int64 `json:"id"`
}
type System__User__CountUsersInput struct {
}
type System__User__CountUsersInternalInput struct {
}
type System__User__CountUsersResponseData struct {
	Data System__User__CountUsersResponseData_data `json:"data"`
}
type System__User__CountUsersResponseData_data struct {
	Count System__User__CountUsersResponseData_data_count `json:"count,omitempty"`
}
type System__User__CountUsersResponseData_data_count struct {
	Id int64 `json:"id"`
}
type System__User__CreateOneInput struct {
	CountryCode  string `json:"countryCode,omitempty"`
	Name         string `json:"name"`
	Password     string `json:"password,omitempty"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone"`
}
type System__User__CreateOneInternalInput struct {
	CountryCode  string `json:"countryCode,omitempty"`
	Name         string `json:"name"`
	Password     string `json:"password,omitempty"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone"`
}
type System__User__CreateOneResponseData struct {
	Data System__User__CreateOneResponseData_data `json:"data,omitempty"`
}
type System__User__CreateOneResponseData_data struct {
	Msg    string `json:"msg,omitempty"`
	Status int64  `json:"status,omitempty"`
}
type System__User__DeleteOneInput struct {
	Id int64 `json:"id,omitempty"`
}
type System__User__DeleteOneInternalInput struct {
	Id int64 `json:"id,omitempty"`
}
type System__User__DeleteOneResponseData struct {
	Main_deleteOneuser System__User__DeleteOneResponseData_main_deleteOneuser `json:"main_deleteOneuser,omitempty"`
}
type System__User__DeleteOneResponseData_main_deleteOneuser struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	Password_salt string `json:"password_salt,omitempty"`
	Password_type string `json:"password_type,omitempty"`
	Phone         string `json:"phone,omitempty"`
	User_id       string `json:"user_id,omitempty"`
}
type System__User__DisconnectRoleInput struct {
	RoleId int64 `json:"roleId"`
	UserId int64 `json:"userId,omitempty"`
}
type System__User__DisconnectRoleInternalInput struct {
	RoleId int64 `json:"roleId"`
	UserId int64 `json:"userId,omitempty"`
}
type System__User__DisconnectRoleResponseData struct {
	Data System__User__DisconnectRoleResponseData_data `json:"data,omitempty"`
}
type System__User__DisconnectRoleResponseData_data struct {
	Id int64 `json:"id"`
}
type System__User__GetAllListInput struct {
}
type System__User__GetAllListInternalInput struct {
}
type System__User__GetAllListResponseData struct {
	Main_aggregateuser System__User__GetAllListResponseData_main_aggregateuser `json:"main_aggregateuser"`
}
type System__User__GetAllListResponseData_main_aggregateuser struct {
	Count System__User__GetAllListResponseData_main_aggregateuser__count `json:"_count,omitempty"`
}
type System__User__GetAllListResponseData_main_aggregateuser__count struct {
	Id int64 `json:"id"`
}
type System__User__GetLikeUserInput struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
}
type System__User__GetLikeUserInternalInput struct {
	Id     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	RoleId int64  `json:"roleId,omitempty"`
	UserId int64  `json:"userId,omitempty"`
}
type System__User__GetLikeUserResponseData struct {
	Main_findManyuser   []System__User__GetLikeUserResponseData_main_findManyuser `json:"main_findManyuser"`
	Main_findUniqueuser System__User__GetLikeUserResponseData_main_findUniqueuser `json:"main_findUniqueuser,omitempty"`
}
type System__User__GetLikeUserResponseData_main_findManyuser struct {
	Avatar    string                                                            `json:"avatar,omitempty"`
	CreatedAt string                                                            `json:"createdAt,omitempty"`
	Id        int64                                                             `json:"id"`
	Name      string                                                            `json:"name"`
	Phone     string                                                            `json:"phone,omitempty"`
	User_role System__User__GetLikeUserResponseData_main_findManyuser_user_role `json:"user_role"`
}
type System__User__GetLikeUserResponseData_main_findManyuser_user_role struct {
	Data []System__User__GetLikeUserResponseData_main_findManyuser_user_role_data `json:"data"`
}
type System__User__GetLikeUserResponseData_main_findManyuser_user_role_data struct {
	Role    System__User__GetLikeUserResponseData_main_findManyuser_user_role_data_role `json:"role"`
	Role_id int64                                                                       `json:"role_id"`
}
type System__User__GetLikeUserResponseData_main_findManyuser_user_role_data_role struct {
	Main_findManyrole []System__User__GetLikeUserResponseData_main_findManyuser_user_role_data_role_main_findManyrole `json:"main_findManyrole"`
}
type System__User__GetLikeUserResponseData_main_findManyuser_user_role_data_role_main_findManyrole struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__GetLikeUserResponseData_main_findUniqueuser struct {
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
}
type System__User__GetListInput struct {
	OrderBy []*Main_userOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_userWhereInput                 `json:"query,omitempty"`
	RoleId  int64                                `json:"roleId,omitempty"`
	Skip    int64                                `json:"skip"`
	Take    int64                                `json:"take"`
	UserId  int64                                `json:"userId,omitempty"`
}
type System__User__GetListInternalInput struct {
	OrderBy []*Main_userOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_userWhereInput                 `json:"query,omitempty"`
	RoleId  int64                                `json:"roleId,omitempty"`
	Skip    int64                                `json:"skip"`
	Take    int64                                `json:"take"`
	UserId  int64                                `json:"userId,omitempty"`
}
type System__User__GetListResponseData struct {
	Data []System__User__GetListResponseData_data `json:"data"`
}
type System__User__GetListResponseData_data struct {
	Join      System__User__GetListResponseData_data__join `json:"_join"`
	Avatar    string                                       `json:"avatar,omitempty"`
	CreatedAt string                                       `json:"createdAt,omitempty"`
	Id        int64                                        `json:"id"`
	Name      string                                       `json:"name"`
	Phone     string                                       `json:"phone,omitempty"`
	User_id   string                                       `json:"user_id,omitempty"`
}
type System__User__GetListResponseData_data__join struct {
	Main_findManyrole_user []System__User__GetListResponseData_data__join_main_findManyrole_user `json:"main_findManyrole_user"`
}
type System__User__GetListResponseData_data__join_main_findManyrole_user struct {
	Join    System__User__GetListResponseData_data__join_main_findManyrole_user__join `json:"_join"`
	Role_id int64                                                                     `json:"role_id"`
	User_id int64                                                                     `json:"user_id"`
}
type System__User__GetListResponseData_data__join_main_findManyrole_user__join struct {
	Main_findManyrole []System__User__GetListResponseData_data__join_main_findManyrole_user__join_main_findManyrole `json:"main_findManyrole"`
}
type System__User__GetListResponseData_data__join_main_findManyrole_user__join_main_findManyrole struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__GetOneInput struct {
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	RoleId int64  `json:"roleId,omitempty"`
	UserId int64  `json:"userId,omitempty"`
}
type System__User__GetOneInternalInput struct {
	Name   string `json:"name,omitempty"`
	Phone  string `json:"phone,omitempty"`
	RoleId int64  `json:"roleId,omitempty"`
	UserId int64  `json:"userId,omitempty"`
}
type System__User__GetOneResponseData struct {
	Data System__User__GetOneResponseData_data `json:"data,omitempty"`
}
type System__User__GetOneResponseData_data struct {
	Avatar  string   `json:"avatar,omitempty"`
	Id      int64    `json:"id"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	User_id string   `json:"user_id,omitempty"`
}
type System__User__GetRoleUsersInput struct {
	Code string `json:"code"`
}
type System__User__GetRoleUsersInternalInput struct {
	Code   string `json:"code"`
	RoleId int64  `json:"roleId,omitempty"`
	UserId int64  `json:"userId,omitempty"`
}
type System__User__GetRoleUsersResponseData struct {
	Main_findManyrole []System__User__GetRoleUsersResponseData_main_findManyrole `json:"main_findManyrole"`
}
type System__User__GetRoleUsersResponseData_main_findManyrole struct {
	Join   System__User__GetRoleUsersResponseData_main_findManyrole__join `json:"_join"`
	Code   string                                                         `json:"code"`
	Id     int64                                                          `json:"id"`
	Remark string                                                         `json:"remark"`
}
type System__User__GetRoleUsersResponseData_main_findManyrole__join struct {
	Main_findManyrole_user []System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user `json:"main_findManyrole_user"`
}
type System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user struct {
	Join    System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user__join `json:"_join"`
	Role_id int64                                                                                       `json:"role_id"`
	User_id int64                                                                                       `json:"user_id"`
}
type System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user__join struct {
	Main_findManyuser []System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user__join_main_findManyuser `json:"main_findManyuser"`
}
type System__User__GetRoleUsersResponseData_main_findManyrole__join_main_findManyrole_user__join_main_findManyuser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type System__User__GetUserByUserIdInput struct {
	UserId string `json:"userId,omitempty"`
}
type System__User__GetUserByUserIdInternalInput struct {
	UserId string `json:"userId,omitempty"`
}
type System__User__GetUserByUserIdResponseData struct {
	Main_findManyuser []System__User__GetUserByUserIdResponseData_main_findManyuser `json:"main_findManyuser"`
}
type System__User__GetUserByUserIdResponseData_main_findManyuser struct {
	Avatar  string `json:"avatar,omitempty"`
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone,omitempty"`
	User_id string `json:"user_id,omitempty"`
}
type System__User__GetUserRoleInput struct {
	UserId int64 `json:"userId"`
}
type System__User__GetUserRoleInternalInput struct {
	RoleId int64 `json:"roleId,omitempty"`
	UserId int64 `json:"userId"`
}
type System__User__GetUserRoleResponseData struct {
	Data []System__User__GetUserRoleResponseData_data `json:"data,omitempty"`
}
type System__User__GetUserRoleResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__UpdateOneInput struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Set          string `json:"set,omitempty"`
	UserId       string `json:"userId"`
}
type System__User__UpdateOneInternalInput struct {
	Id           int64  `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Password     string `json:"password,omitempty"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Set          string `json:"set,omitempty"`
	UserId       string `json:"userId"`
}
type System__User__UpdateOneResponseData struct {
	Casdoor_apiUpdateUser System__User__UpdateOneResponseData_casdoor_apiUpdateUser `json:"casdoor_apiUpdateUser,omitempty"`
	Main_updateOneuser    System__User__UpdateOneResponseData_main_updateOneuser    `json:"main_updateOneuser,omitempty"`
}
type System__User__UpdateOneResponseData_casdoor_apiUpdateUser struct {
	Msg    string `json:"msg,omitempty"`
	Status int64  `json:"status,omitempty"`
}
type System__User__UpdateOneResponseData_main_updateOneuser struct {
	Avatar string `json:"avatar,omitempty"`
}
type Main_CategoryRelationFilter struct {
	Is    *Main_CategoryWhereInput `json:"is,omitempty"`
	IsNot *Main_CategoryWhereInput `json:"isNot,omitempty"`
}
type Main_CategoryWhereInput struct {
	AND  *Main_CategoryWhereInput     `json:"AND,omitempty"`
	NOT  *Main_CategoryWhereInput     `json:"NOT,omitempty"`
	OR   []*Main_CategoryWhereInput   `json:"OR,omitempty"`
	Id   *Main_IntFilter              `json:"id,omitempty"`
	Name *Main_StringFilter           `json:"name,omitempty"`
	Post *Main_PostListRelationFilter `json:"post,omitempty"`
}
type Main_DateTimeNullableFilter struct {
	Equals string                             `json:"equals,omitempty"`
	Gt     string                             `json:"gt,omitempty"`
	Gte    string                             `json:"gte,omitempty"`
	In     []string                           `json:"in,omitempty"`
	Lt     string                             `json:"lt,omitempty"`
	Lte    string                             `json:"lte,omitempty"`
	Not    *Main_NestedDateTimeNullableFilter `json:"not,omitempty"`
	NotIn  []string                           `json:"notIn,omitempty"`
}
type Main_IntFilter struct {
	Equals int64                 `json:"equals,omitempty"`
	Gt     int64                 `json:"gt,omitempty"`
	Gte    int64                 `json:"gte,omitempty"`
	In     []int64               `json:"in,omitempty"`
	Lt     int64                 `json:"lt,omitempty"`
	Lte    int64                 `json:"lte,omitempty"`
	Not    *Main_NestedIntFilter `json:"not,omitempty"`
	NotIn  []int64               `json:"notIn,omitempty"`
}
type Main_IntNullableFilter struct {
	Equals int64                         `json:"equals,omitempty"`
	Gt     int64                         `json:"gt,omitempty"`
	Gte    int64                         `json:"gte,omitempty"`
	In     []int64                       `json:"in,omitempty"`
	Lt     int64                         `json:"lt,omitempty"`
	Lte    int64                         `json:"lte,omitempty"`
	Not    *Main_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                       `json:"notIn,omitempty"`
}
type Main_NestedDateTimeNullableFilter struct {
	Equals string                             `json:"equals,omitempty"`
	Gt     string                             `json:"gt,omitempty"`
	Gte    string                             `json:"gte,omitempty"`
	In     []string                           `json:"in,omitempty"`
	Lt     string                             `json:"lt,omitempty"`
	Lte    string                             `json:"lte,omitempty"`
	Not    *Main_NestedDateTimeNullableFilter `json:"not,omitempty"`
	NotIn  []string                           `json:"notIn,omitempty"`
}
type Main_NestedIntFilter struct {
	Equals int64                 `json:"equals,omitempty"`
	Gt     int64                 `json:"gt,omitempty"`
	Gte    int64                 `json:"gte,omitempty"`
	In     []int64               `json:"in,omitempty"`
	Lt     int64                 `json:"lt,omitempty"`
	Lte    int64                 `json:"lte,omitempty"`
	Not    *Main_NestedIntFilter `json:"not,omitempty"`
	NotIn  []int64               `json:"notIn,omitempty"`
}
type Main_NestedIntNullableFilter struct {
	Equals int64                         `json:"equals,omitempty"`
	Gt     int64                         `json:"gt,omitempty"`
	Gte    int64                         `json:"gte,omitempty"`
	In     []int64                       `json:"in,omitempty"`
	Lt     int64                         `json:"lt,omitempty"`
	Lte    int64                         `json:"lte,omitempty"`
	Not    *Main_NestedIntNullableFilter `json:"not,omitempty"`
	NotIn  []int64                       `json:"notIn,omitempty"`
}
type Main_NestedStringFilter struct {
	Contains   string                   `json:"contains,omitempty"`
	EndsWith   string                   `json:"endsWith,omitempty"`
	Equals     string                   `json:"equals,omitempty"`
	Gt         string                   `json:"gt,omitempty"`
	Gte        string                   `json:"gte,omitempty"`
	In         []string                 `json:"in,omitempty"`
	Lt         string                   `json:"lt,omitempty"`
	Lte        string                   `json:"lte,omitempty"`
	Not        *Main_NestedStringFilter `json:"not,omitempty"`
	NotIn      []string                 `json:"notIn,omitempty"`
	StartsWith string                   `json:"startsWith,omitempty"`
}
type Main_NestedStringNullableFilter struct {
	Contains   string                           `json:"contains,omitempty"`
	EndsWith   string                           `json:"endsWith,omitempty"`
	Equals     string                           `json:"equals,omitempty"`
	Gt         string                           `json:"gt,omitempty"`
	Gte        string                           `json:"gte,omitempty"`
	In         []string                         `json:"in,omitempty"`
	Lt         string                           `json:"lt,omitempty"`
	Lte        string                           `json:"lte,omitempty"`
	Not        *Main_NestedStringNullableFilter `json:"not,omitempty"`
	NotIn      []string                         `json:"notIn,omitempty"`
	StartsWith string                           `json:"startsWith,omitempty"`
}
type Main_PostListRelationFilter struct {
	Every *Main_postWhereInput `json:"every,omitempty"`
	None  *Main_postWhereInput `json:"none,omitempty"`
	Some  *Main_postWhereInput `json:"some,omitempty"`
}
type Main_StringFilter struct {
	Contains   string                   `json:"contains,omitempty"`
	EndsWith   string                   `json:"endsWith,omitempty"`
	Equals     string                   `json:"equals,omitempty"`
	Gt         string                   `json:"gt,omitempty"`
	Gte        string                   `json:"gte,omitempty"`
	In         []string                 `json:"in,omitempty"`
	Lt         string                   `json:"lt,omitempty"`
	Lte        string                   `json:"lte,omitempty"`
	Not        *Main_NestedStringFilter `json:"not,omitempty"`
	NotIn      []string                 `json:"notIn,omitempty"`
	StartsWith string                   `json:"startsWith,omitempty"`
}
type Main_StringNullableFilter struct {
	Contains   string                           `json:"contains,omitempty"`
	EndsWith   string                           `json:"endsWith,omitempty"`
	Equals     string                           `json:"equals,omitempty"`
	Gt         string                           `json:"gt,omitempty"`
	Gte        string                           `json:"gte,omitempty"`
	In         []string                         `json:"in,omitempty"`
	Lt         string                           `json:"lt,omitempty"`
	Lte        string                           `json:"lte,omitempty"`
	Not        *Main_NestedStringNullableFilter `json:"not,omitempty"`
	NotIn      []string                         `json:"notIn,omitempty"`
	StartsWith string                           `json:"startsWith,omitempty"`
}
type Main_UserRelationFilter struct {
	Is    *Main_userWhereInput `json:"is,omitempty"`
	IsNot *Main_userWhereInput `json:"isNot,omitempty"`
}
type Main_apilogOrderByWithRelationInput struct {
	Code      Main_SortOrder `json:"code,omitempty"`
	CreatedAt Main_SortOrder `json:"createdAt,omitempty"`
	DeletedAt Main_SortOrder `json:"deletedAt,omitempty"`
	Id        Main_SortOrder `json:"id,omitempty"`
	Ip        Main_SortOrder `json:"ip,omitempty"`
	Method    Main_SortOrder `json:"method,omitempty"`
	Path      Main_SortOrder `json:"path,omitempty"`
	Ua        Main_SortOrder `json:"ua,omitempty"`
	UpdatedAt Main_SortOrder `json:"updatedAt,omitempty"`
	UserId    Main_SortOrder `json:"userId,omitempty"`
	UserName  Main_SortOrder `json:"userName,omitempty"`
}
type Main_menuOrderByWithRelationInput struct {
	Api_id      Main_SortOrder `json:"api_id,omitempty"`
	Create_time Main_SortOrder `json:"create_time,omitempty"`
	Icon        Main_SortOrder `json:"icon,omitempty"`
	Id          Main_SortOrder `json:"id,omitempty"`
	Is_bottom   Main_SortOrder `json:"is_bottom,omitempty"`
	Label       Main_SortOrder `json:"label,omitempty"`
	Level       Main_SortOrder `json:"level,omitempty"`
	Menu_type   Main_SortOrder `json:"menu_type,omitempty"`
	Name        Main_SortOrder `json:"name,omitempty"`
	ParentId    Main_SortOrder `json:"parentId,omitempty"`
	Path        Main_SortOrder `json:"path,omitempty"`
	Perms       Main_SortOrder `json:"perms,omitempty"`
	Sort        Main_SortOrder `json:"sort,omitempty"`
}
type Main_menuWhereInput struct {
	AND         *Main_menuWhereInput         `json:"AND,omitempty"`
	NOT         *Main_menuWhereInput         `json:"NOT,omitempty"`
	OR          []*Main_menuWhereInput       `json:"OR,omitempty"`
	Api_id      *Main_StringNullableFilter   `json:"api_id,omitempty"`
	Create_time *Main_DateTimeNullableFilter `json:"create_time,omitempty"`
	Icon        *Main_StringNullableFilter   `json:"icon,omitempty"`
	Id          *Main_IntFilter              `json:"id,omitempty"`
	Is_bottom   *Main_IntFilter              `json:"is_bottom,omitempty"`
	Label       *Main_StringFilter           `json:"label,omitempty"`
	Level       *Main_IntFilter              `json:"level,omitempty"`
	Menu_type   *Main_StringNullableFilter   `json:"menu_type,omitempty"`
	Name        *Main_StringNullableFilter   `json:"name,omitempty"`
	ParentId    *Main_IntNullableFilter      `json:"parentId,omitempty"`
	Path        *Main_StringNullableFilter   `json:"path,omitempty"`
	Perms       *Main_StringNullableFilter   `json:"perms,omitempty"`
	Sort        *Main_IntFilter              `json:"sort,omitempty"`
}
type Main_menu_roleCreateManyInput struct {
	Id      int64 `json:"id,omitempty"`
	Menu_id int64 `json:"menu_id,omitempty"`
	Role_id int64 `json:"role_id,omitempty"`
}
type Main_permissionCreateManyInput struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Enabled   int64  `json:"enabled,omitempty"`
	Id        string `json:"id,omitempty"`
	Method    string `json:"method,omitempty"`
	Path      string `json:"path,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
type Main_postOrderByRelationAggregateInput struct {
	Count Main_SortOrder `json:"_count,omitempty"`
}
type Main_postWhereInput struct {
	AND          *Main_postWhereInput         `json:"AND,omitempty"`
	Category     *Main_CategoryRelationFilter `json:"Category,omitempty"`
	NOT          *Main_postWhereInput         `json:"NOT,omitempty"`
	OR           []*Main_postWhereInput       `json:"OR,omitempty"`
	Auth         *Main_IntNullableFilter      `json:"auth,omitempty"`
	Author       *Main_StringFilter           `json:"author,omitempty"`
	Cate         *Main_IntNullableFilter      `json:"cate,omitempty"`
	Content      *Main_StringNullableFilter   `json:"content,omitempty"`
	Id           *Main_IntFilter              `json:"id,omitempty"`
	Poster       *Main_StringNullableFilter   `json:"poster,omitempty"`
	Published_at *Main_DateTimeNullableFilter `json:"published_at,omitempty"`
	Title        *Main_StringFilter           `json:"title,omitempty"`
	User         *Main_UserRelationFilter     `json:"user,omitempty"`
}
type Main_roleOrderByWithRelationInput struct {
	Code   Main_SortOrder `json:"code,omitempty"`
	Id     Main_SortOrder `json:"id,omitempty"`
	Remark Main_SortOrder `json:"remark,omitempty"`
}
type Main_roleWhereInput struct {
	AND    *Main_roleWhereInput   `json:"AND,omitempty"`
	NOT    *Main_roleWhereInput   `json:"NOT,omitempty"`
	OR     []*Main_roleWhereInput `json:"OR,omitempty"`
	Code   *Main_StringFilter     `json:"code,omitempty"`
	Id     *Main_IntFilter        `json:"id,omitempty"`
	Remark *Main_StringFilter     `json:"remark,omitempty"`
}
type Main_userOrderByWithRelationInput struct {
	Avatar        Main_SortOrder                          `json:"avatar,omitempty"`
	Country_code  Main_SortOrder                          `json:"country_code,omitempty"`
	Created_at    Main_SortOrder                          `json:"created_at,omitempty"`
	Id            Main_SortOrder                          `json:"id,omitempty"`
	Name          Main_SortOrder                          `json:"name,omitempty"`
	Password      Main_SortOrder                          `json:"password,omitempty"`
	Password_salt Main_SortOrder                          `json:"password_salt,omitempty"`
	Password_type Main_SortOrder                          `json:"password_type,omitempty"`
	Phone         Main_SortOrder                          `json:"phone,omitempty"`
	Post          *Main_postOrderByRelationAggregateInput `json:"post,omitempty"`
	User_id       Main_SortOrder                          `json:"user_id,omitempty"`
}
type Main_userWhereInput struct {
	AND           *Main_userWhereInput         `json:"AND,omitempty"`
	NOT           *Main_userWhereInput         `json:"NOT,omitempty"`
	OR            []*Main_userWhereInput       `json:"OR,omitempty"`
	Avatar        *Main_StringNullableFilter   `json:"avatar,omitempty"`
	Country_code  *Main_StringNullableFilter   `json:"country_code,omitempty"`
	Created_at    *Main_DateTimeNullableFilter `json:"created_at,omitempty"`
	Id            *Main_IntFilter              `json:"id,omitempty"`
	Name          *Main_StringFilter           `json:"name,omitempty"`
	Password      *Main_StringNullableFilter   `json:"password,omitempty"`
	Password_salt *Main_StringNullableFilter   `json:"password_salt,omitempty"`
	Password_type *Main_StringNullableFilter   `json:"password_type,omitempty"`
	Phone         *Main_StringNullableFilter   `json:"phone,omitempty"`
	Post          *Main_PostListRelationFilter `json:"post,omitempty"`
	User_id       *Main_StringNullableFilter   `json:"user_id,omitempty"`
}

type Main_SortOrder string

const (
	Main_SortOrder_asc  Main_SortOrder = "asc"
	Main_SortOrder_desc Main_SortOrder = "desc"
)
