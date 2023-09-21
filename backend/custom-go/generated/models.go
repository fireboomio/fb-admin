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
	Data []Casdoor__GetSMSProviderResponseData_data `json:"data"`
}
type Casdoor__GetSMSProviderResponseData_data struct {
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
	Data Casdoor__LoginResponseData_data `json:"data,omitempty"`
}
type Casdoor__LoginResponseData_data struct {
	Data    Casdoor__LoginResponseData_data_data `json:"data,omitempty"`
	Msg     string                               `json:"msg,omitempty"`
	Success bool                                 `json:"success,omitempty"`
}
type Casdoor__LoginResponseData_data_data struct {
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
	Data Casdoor__UpdateSMSProviderResponseData_data `json:"data,omitempty"`
}
type Casdoor__UpdateSMSProviderResponseData_data struct {
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
	Data Post__CreateCategoryResponseData_data `json:"data,omitempty"`
}
type Post__CreateCategoryResponseData_data struct {
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
	Author        string                                         `json:"author"`
	Case_category Post__CreateOneResponseData_data_case_category `json:"case_category,omitempty"`
	Id            int64                                          `json:"id"`
	Poster        string                                         `json:"poster,omitempty"`
	Published_at  string                                         `json:"published_at,omitempty"`
	Title         string                                         `json:"title"`
}
type Post__CreateOneResponseData_data_case_category struct {
	Id int64 `json:"id"`
}
type Post__DeleteCategoryInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__DeleteCategoryInternalInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__DeleteCategoryResponseData struct {
	Data Post__DeleteCategoryResponseData_data `json:"data,omitempty"`
}
type Post__DeleteCategoryResponseData_data struct {
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
	Data []Post__GetCategoryResponseData_data `json:"data"`
}
type Post__GetCategoryResponseData_data struct {
	Case_post []Post__GetCategoryResponseData_data_case_post `json:"case_post,omitempty"`
	Id        int64                                          `json:"id"`
	Name      string                                         `json:"name"`
}
type Post__GetCategoryResponseData_data_case_post struct {
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
	Data  []Post__GetLikeListResponseData_data `json:"data"`
	Data1 Post__GetLikeListResponseData_data1  `json:"data1,omitempty"`
}
type Post__GetLikeListResponseData_data struct {
	Author        string                                           `json:"author"`
	Case_category Post__GetLikeListResponseData_data_case_category `json:"case_category,omitempty"`
	Content       string                                           `json:"content,omitempty"`
	Id            int64                                            `json:"id"`
	Poster        string                                           `json:"poster,omitempty"`
	Published_at  string                                           `json:"published_at,omitempty"`
	Title         string                                           `json:"title"`
}
type Post__GetLikeListResponseData_data1 struct {
	Author        string                                            `json:"author"`
	Case_category Post__GetLikeListResponseData_data1_case_category `json:"case_category,omitempty"`
	Content       string                                            `json:"content,omitempty"`
	Id            int64                                             `json:"id"`
	Poster        string                                            `json:"poster,omitempty"`
	Published_at  string                                            `json:"published_at,omitempty"`
	Title         string                                            `json:"title"`
}
type Post__GetLikeListResponseData_data1_case_category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetLikeListResponseData_data_case_category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetListInput struct {
	Query *Main_case_postWhereInput `json:"query,omitempty"`
	Skip  int64                     `json:"skip,omitempty"`
	Take  int64                     `json:"take,omitempty"`
}
type Post__GetListInternalInput struct {
	Query *Main_case_postWhereInput `json:"query,omitempty"`
	Skip  int64                     `json:"skip,omitempty"`
	Take  int64                     `json:"take,omitempty"`
}
type Post__GetListResponseData struct {
	Data  []Post__GetListResponseData_data `json:"data"`
	Total int64                            `json:"total"`
}
type Post__GetListResponseData_data struct {
	Author        string                                       `json:"author"`
	Case_category Post__GetListResponseData_data_case_category `json:"case_category,omitempty"`
	Content       string                                       `json:"content,omitempty"`
	Id            int64                                        `json:"id"`
	Poster        string                                       `json:"poster,omitempty"`
	Published_at  string                                       `json:"published_at,omitempty"`
	Title         string                                       `json:"title"`
}
type Post__GetListResponseData_data_case_category struct {
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
	Admin_user    Post__GetOneResponseData_data_admin_user    `json:"admin_user,omitempty"`
	Author        string                                      `json:"author"`
	Case_category Post__GetOneResponseData_data_case_category `json:"case_category,omitempty"`
	Content       string                                      `json:"content,omitempty"`
	Id            int64                                       `json:"id"`
	Poster        string                                      `json:"poster,omitempty"`
	Published_at  string                                      `json:"published_at,omitempty"`
	Title         string                                      `json:"title"`
}
type Post__GetOneResponseData_data_admin_user struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetOneResponseData_data_case_category struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetPostByAuthorInput struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Post__GetPostByAuthorInternalInput struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
type Post__GetPostByAuthorResponseData struct {
	Data Post__GetPostByAuthorResponseData_data `json:"data,omitempty"`
}
type Post__GetPostByAuthorResponseData_data struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
type Post__GetPostByCateInput struct {
	Equals int64 `json:"equals"`
}
type Post__GetPostByCateInternalInput struct {
	Equals int64 `json:"equals"`
}
type Post__GetPostByCateResponseData struct {
	Data []Post__GetPostByCateResponseData_data `json:"data"`
}
type Post__GetPostByCateResponseData_data struct {
	Case_post []Post__GetPostByCateResponseData_data_case_post `json:"case_post,omitempty"`
}
type Post__GetPostByCateResponseData_data_case_post struct {
	Author       string `json:"author"`
	Cate         int64  `json:"cate,omitempty"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title"`
}
type Post__UpdateCategoryInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__UpdateCategoryInternalInput struct {
	Id int64 `json:"id,omitempty"`
}
type Post__UpdateCategoryResponseData struct {
	Data []Post__UpdateCategoryResponseData_data `json:"data"`
}
type Post__UpdateCategoryResponseData_data struct {
	Case_post []Post__UpdateCategoryResponseData_data_case_post `json:"case_post,omitempty"`
}
type Post__UpdateCategoryResponseData_data_case_post struct {
	Cate         int64  `json:"cate,omitempty"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title"`
}
type Post__UpdateOneInput struct {
	Author  string `json:"author,omitempty"`
	Cate    int64  `json:"cate,omitempty"`
	Content string `json:"content,omitempty"`
	Id      int64  `json:"id,omitempty"`
	Poster  string `json:"poster,omitempty"`
	Title   string `json:"title,omitempty"`
	UserId  int64  `json:"userId,omitempty"`
}
type Post__UpdateOneInternalInput struct {
	Author       string `json:"author,omitempty"`
	Cate         int64  `json:"cate,omitempty"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title,omitempty"`
	UserId       int64  `json:"userId,omitempty"`
}
type Post__UpdateOneResponseData struct {
	Main_updateOnecase_post Post__UpdateOneResponseData_main_updateOnecase_post `json:"main_updateOnecase_post,omitempty"`
}
type Post__UpdateOneResponseData_main_updateOnecase_post struct {
	Auth    int64  `json:"auth,omitempty"`
	Author  string `json:"author"`
	Cate    int64  `json:"cate,omitempty"`
	Content string `json:"content,omitempty"`
	Id      int64  `json:"id"`
	Poster  string `json:"poster,omitempty"`
	Title   string `json:"title"`
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
type System__GetMenusInput struct {
}
type System__GetMenusInternalInput struct {
	Pid int64 `json:"pid,omitempty"`
}
type System__GetMenusResponseData struct {
	Data []System__GetMenusResponseData_data `json:"data"`
}
type System__GetMenusResponseData_data struct {
	Children []System__GetMenusResponseData_data_children `json:"children"`
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
type System__Jwt__AddOneInput struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id,omitempty"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__AddOneInternalInput struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id,omitempty"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__AddOneResponseData struct {
	Main_createManyadmin_token System__Jwt__AddOneResponseData_main_createManyadmin_token `json:"main_createManyadmin_token,omitempty"`
}
type System__Jwt__AddOneResponseData_main_createManyadmin_token struct {
	Count int64 `json:"count"`
}
type System__Jwt__CheckBannedInput struct {
	Token string `json:"token,omitempty"`
}
type System__Jwt__CheckBannedInternalInput struct {
	Token string `json:"token,omitempty"`
}
type System__Jwt__CheckBannedResponseData struct {
	Data []System__Jwt__CheckBannedResponseData_data `json:"data"`
}
type System__Jwt__CheckBannedResponseData_data struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__DeleteOneInput struct {
	Id int64 `json:"id,omitempty"`
}
type System__Jwt__DeleteOneInternalInput struct {
	Id int64 `json:"id,omitempty"`
}
type System__Jwt__DeleteOneResponseData struct {
	Data System__Jwt__DeleteOneResponseData_data `json:"data,omitempty"`
}
type System__Jwt__DeleteOneResponseData_data struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__GetManyInput struct {
	Skip int64 `json:"skip,omitempty"`
	Take int64 `json:"take,omitempty"`
}
type System__Jwt__GetManyInternalInput struct {
	Skip int64 `json:"skip,omitempty"`
	Take int64 `json:"take,omitempty"`
}
type System__Jwt__GetManyResponseData struct {
	Main_findManyadmin_token []System__Jwt__GetManyResponseData_main_findManyadmin_token `json:"main_findManyadmin_token"`
}
type System__Jwt__GetManyResponseData_main_findManyadmin_token struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__UpdateOneInput struct {
	Banned      bool   `json:"banned,omitempty"`
	Equals      string `json:"equals"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__UpdateOneInternalInput struct {
	Banned      bool   `json:"banned,omitempty"`
	Equals      string `json:"equals"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type System__Jwt__UpdateOneResponseData struct {
	Main_updateManyadmin_token System__Jwt__UpdateOneResponseData_main_updateManyadmin_token `json:"main_updateManyadmin_token,omitempty"`
}
type System__Jwt__UpdateOneResponseData_main_updateManyadmin_token struct {
	Count int64 `json:"count"`
}
type System__Log__ChangeOpenInput struct {
	Set bool `json:"set,omitempty"`
}
type System__Log__ChangeOpenInternalInput struct {
	Set bool `json:"set,omitempty"`
}
type System__Log__ChangeOpenResponseData struct {
	Data System__Log__ChangeOpenResponseData_data `json:"data,omitempty"`
}
type System__Log__ChangeOpenResponseData_data struct {
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
	Data System__Log__CreateLogResponseData_data `json:"data,omitempty"`
}
type System__Log__CreateLogResponseData_data struct {
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
	Main_deleteManyadmin_apilog System__Log__DeleteLogResponseData_main_deleteManyadmin_apilog `json:"main_deleteManyadmin_apilog,omitempty"`
}
type System__Log__DeleteLogResponseData_main_deleteManyadmin_apilog struct {
	Count int64 `json:"count"`
}
type System__Log__DeleteOneInput struct {
	Id int64 `json:"id"`
}
type System__Log__DeleteOneInternalInput struct {
	Id int64 `json:"id"`
}
type System__Log__DeleteOneResponseData struct {
	Data System__Log__DeleteOneResponseData_data `json:"data,omitempty"`
}
type System__Log__DeleteOneResponseData_data struct {
	Id int64 `json:"id"`
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
	Data System__Log__GetIsOpenResponseData_data `json:"data,omitempty"`
}
type System__Log__GetIsOpenResponseData_data struct {
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
	Main_findManyadmin_apilog   []System__Log__GetLikeLogResponseData_main_findManyadmin_apilog `json:"main_findManyadmin_apilog"`
	Main_findUniqueadmin_apilog System__Log__GetLikeLogResponseData_main_findUniqueadmin_apilog `json:"main_findUniqueadmin_apilog,omitempty"`
}
type System__Log__GetLikeLogResponseData_main_findManyadmin_apilog struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserName  string `json:"userName"`
}
type System__Log__GetLikeLogResponseData_main_findUniqueadmin_apilog struct {
	Id        int64  `json:"id"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserName  string `json:"userName"`
}
type System__Log__GetLogInput struct {
	OrderBy []*Main_admin_apilogOrderByWithRelationInput `json:"orderBy,omitempty"`
	Skip    int64                                        `json:"skip,omitempty"`
	Take    int64                                        `json:"take,omitempty"`
}
type System__Log__GetLogInternalInput struct {
	OrderBy []*Main_admin_apilogOrderByWithRelationInput `json:"orderBy,omitempty"`
	Skip    int64                                        `json:"skip,omitempty"`
	Take    int64                                        `json:"take,omitempty"`
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
	ApiId    string `json:"apiId,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Label    string `json:"label"`
	Level    int64  `json:"level"`
	MenuType string `json:"menuType"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path"`
	Perms    string `json:"perms,omitempty"`
	Sort     int64  `json:"sort"`
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
	Main_deleteOnecase_menu System__Menu__DeleteOneResponseData_main_deleteOnecase_menu `json:"main_deleteOnecase_menu,omitempty"`
}
type System__Menu__DeleteOneResponseData_main_deleteOnecase_menu struct {
	Api_id      string `json:"api_id,omitempty"`
	Create_time string `json:"create_time,omitempty"`
	Icon        string `json:"icon,omitempty"`
}
type System__Menu__GetApiListInput struct {
}
type System__Menu__GetApiListInternalInput struct {
}
type System__Menu__GetApiListResponseData struct {
	Data []System__Menu__GetApiListResponseData_data `json:"data,omitempty"`
}
type System__Menu__GetApiListResponseData_data struct {
	Path string `json:"path,omitempty"`
}
type System__Menu__GetApisByMenusInput struct {
	MenuIds []int64 `json:"menuIds"`
}
type System__Menu__GetApisByMenusInternalInput struct {
	MenuIds []int64 `json:"menuIds"`
}
type System__Menu__GetApisByMenusResponseData struct {
	Data []string `json:"data"`
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
	OrderBy []*Main_case_menuOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_case_menuWhereInput                 `json:"query,omitempty"`
	Skip    int64                                     `json:"skip,omitempty"`
	Take    int64                                     `json:"take,omitempty"`
}
type System__Menu__GetListInternalInput struct {
	OrderBy []*Main_case_menuOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_case_menuWhereInput                 `json:"query,omitempty"`
	Skip    int64                                     `json:"skip,omitempty"`
	Take    int64                                     `json:"take,omitempty"`
}
type System__Menu__GetListResponseData struct {
	Data  []System__Menu__GetListResponseData_data `json:"data"`
	Total int64                                    `json:"total"`
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
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label,omitempty"`
	Level    int64  `json:"level,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort,omitempty"`
}
type System__Menu__UpdateOneInternalInput struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label,omitempty"`
	Level    int64  `json:"level,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort,omitempty"`
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
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label,omitempty"`
	Level    int64  `json:"level,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort,omitempty"`
}
type System__Operation__GetManyInternalInput struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label,omitempty"`
	Level    int64  `json:"level,omitempty"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort,omitempty"`
}
type System__Operation__GetManyResponseData struct {
	Data System__Operation__GetManyResponseData_data `json:"data,omitempty"`
}
type System__Operation__GetManyResponseData_data struct {
	Icon     string `json:"icon,omitempty"`
	Id       int64  `json:"id"`
	Label    string `json:"label"`
	Level    int64  `json:"level"`
	ParentId int64  `json:"parentId,omitempty"`
	Path     string `json:"path,omitempty"`
	Sort     int64  `json:"sort"`
}
type System__Perm__CreateManyInput struct {
	Data []*Main_admin_permissionCreateManyInput `json:"data"`
}
type System__Perm__CreateManyInternalInput struct {
	Data []*Main_admin_permissionCreateManyInput `json:"data"`
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
	Data []string `json:"data"`
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
	Data []*Main_admin_menu_roleCreateManyInput `json:"data"`
}
type System__Role__BindMenusInternalInput struct {
	Data []*Main_admin_menu_roleCreateManyInput `json:"data"`
}
type System__Role__BindMenusResponseData struct {
	Data System__Role__BindMenusResponseData_data `json:"data,omitempty"`
}
type System__Role__BindMenusResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Role__BindRoleApisInput struct {
	OperationPaths []string `json:"operationPaths"`
	RoleCodes      []string `json:"roleCodes"`
}
type System__Role__BindRoleApisInternalInput struct {
	OperationPaths []string `json:"operationPaths"`
	RoleCodes      []string `json:"roleCodes"`
}
type System__Role__BindRoleApisResponseData struct {
	System_bindRoleApis_post []string `json:"system_bindRoleApis_post,omitempty"`
}
type System__Role__DeleteManyInput struct {
	Ids []int64 `json:"ids,omitempty"`
}
type System__Role__DeleteManyInternalInput struct {
	Ids []int64 `json:"ids,omitempty"`
}
type System__Role__DeleteManyResponseData struct {
	Data  System__Role__DeleteManyResponseData_data  `json:"data,omitempty"`
	Data1 System__Role__DeleteManyResponseData_data1 `json:"data1,omitempty"`
	Data2 System__Role__DeleteManyResponseData_data2 `json:"data2,omitempty"`
}
type System__Role__DeleteManyResponseData_data struct {
	Count int64 `json:"count"`
}
type System__Role__DeleteManyResponseData_data1 struct {
	Count int64 `json:"count"`
}
type System__Role__DeleteManyResponseData_data2 struct {
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
	OrderBy []*Main_admin_roleOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_admin_roleWhereInput                 `json:"query,omitempty"`
	Skip    int64                                      `json:"skip,omitempty"`
	Take    int64                                      `json:"take,omitempty"`
}
type System__Role__GetListInternalInput struct {
	OrderBy []*Main_admin_roleOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_admin_roleWhereInput                 `json:"query,omitempty"`
	Skip    int64                                      `json:"skip,omitempty"`
	Take    int64                                      `json:"take,omitempty"`
}
type System__Role__GetListResponseData struct {
	Data  []System__Role__GetListResponseData_data `json:"data"`
	Total int64                                    `json:"total"`
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
type System__Role__GetRoleByIdInput struct {
	Id int64 `json:"id"`
}
type System__Role__GetRoleByIdInternalInput struct {
	Id int64 `json:"id"`
}
type System__Role__GetRoleByIdResponseData struct {
	Data System__Role__GetRoleByIdResponseData_data `json:"data,omitempty"`
}
type System__Role__GetRoleByIdResponseData_data struct {
	Code string `json:"code"`
}
type System__Role__GetRoleMenuIdInput struct {
	RoleId int64 `json:"roleId"`
}
type System__Role__GetRoleMenuIdInternalInput struct {
	RoleId int64 `json:"roleId"`
}
type System__Role__GetRoleMenuIdResponseData struct {
	Data []int64 `json:"data"`
}
type System__Role__SyncDeleteRoleInput struct {
	DataNames []string `json:"dataNames"`
}
type System__Role__SyncDeleteRoleInternalInput struct {
	DataNames []string `json:"dataNames"`
}
type System__Role__SyncDeleteRoleResponseData struct {
	System_batchDeleteRoles_delete bool `json:"system_batchDeleteRoles_delete,omitempty"`
}
type System__Role__SyncRoleInput struct {
	Data []*System_Role_input_object `json:"data"`
}
type System__Role__SyncRoleInternalInput struct {
	Data []*System_Role_input_object `json:"data"`
}
type System__Role__SyncRoleResponseData struct {
	Data []System__Role__SyncRoleResponseData_data `json:"data,omitempty"`
}
type System__Role__SyncRoleResponseData_data struct {
	DataName string `json:"dataName,omitempty"`
	Succeed  bool   `json:"succeed,omitempty"`
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
	Main_deleteManyadmin_menu_role System__Role__UnBindMenusResponseData_main_deleteManyadmin_menu_role `json:"main_deleteManyadmin_menu_role,omitempty"`
}
type System__Role__UnBindMenusResponseData_main_deleteManyadmin_menu_role struct {
	Count int64 `json:"count"`
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
	Main_createOnecase_sub System__Sub__CreateSubResponseData_main_createOnecase_sub `json:"main_createOnecase_sub,omitempty"`
}
type System__Sub__CreateSubResponseData_main_createOnecase_sub struct {
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
	Password     string `json:"password"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone"`
}
type System__User__CreateOneInternalInput struct {
	CountryCode  string `json:"countryCode,omitempty"`
	Name         string `json:"name"`
	Password     string `json:"password"`
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
	Main_deleteOneadmin_user System__User__DeleteOneResponseData_main_deleteOneadmin_user `json:"main_deleteOneadmin_user,omitempty"`
}
type System__User__DeleteOneResponseData_main_deleteOneadmin_user struct {
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
	Data System__User__GetAllListResponseData_data `json:"data"`
}
type System__User__GetAllListResponseData_data struct {
	Count System__User__GetAllListResponseData_data__count `json:"_count,omitempty"`
}
type System__User__GetAllListResponseData_data__count struct {
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
	Data  []System__User__GetLikeUserResponseData_data `json:"data"`
	Data1 System__User__GetLikeUserResponseData_data1  `json:"data1,omitempty"`
}
type System__User__GetLikeUserResponseData_data struct {
	Avatar    string                                               `json:"avatar,omitempty"`
	CreatedAt string                                               `json:"createdAt,omitempty"`
	Id        int64                                                `json:"id"`
	Name      string                                               `json:"name"`
	Phone     string                                               `json:"phone,omitempty"`
	User_role System__User__GetLikeUserResponseData_data_user_role `json:"user_role"`
}
type System__User__GetLikeUserResponseData_data1 struct {
	Avatar    string `json:"avatar,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Phone     string `json:"phone,omitempty"`
}
type System__User__GetLikeUserResponseData_data_user_role struct {
	Data []System__User__GetLikeUserResponseData_data_user_role_data `json:"data"`
}
type System__User__GetLikeUserResponseData_data_user_role_data struct {
	Role    System__User__GetLikeUserResponseData_data_user_role_data_role `json:"role"`
	Role_id int64                                                          `json:"role_id"`
}
type System__User__GetLikeUserResponseData_data_user_role_data_role struct {
	Main_findManyadmin_role []System__User__GetLikeUserResponseData_data_user_role_data_role_main_findManyadmin_role `json:"main_findManyadmin_role"`
}
type System__User__GetLikeUserResponseData_data_user_role_data_role_main_findManyadmin_role struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__GetListInput struct {
	OrderBy []*Main_admin_userOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_admin_userWhereInput                 `json:"query,omitempty"`
	Skip    int64                                      `json:"skip"`
	Take    int64                                      `json:"take"`
}
type System__User__GetListInternalInput struct {
	OrderBy []*Main_admin_userOrderByWithRelationInput `json:"orderBy,omitempty"`
	Query   *Main_admin_userWhereInput                 `json:"query,omitempty"`
	RoleId  int64                                      `json:"roleId,omitempty"`
	Skip    int64                                      `json:"skip"`
	Take    int64                                      `json:"take"`
	UserId  int64                                      `json:"userId,omitempty"`
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
	Main_findManyadmin_role_user []System__User__GetListResponseData_data__join_main_findManyadmin_role_user `json:"main_findManyadmin_role_user"`
}
type System__User__GetListResponseData_data__join_main_findManyadmin_role_user struct {
	Join    System__User__GetListResponseData_data__join_main_findManyadmin_role_user__join `json:"_join"`
	Role_id int64                                                                           `json:"role_id"`
	User_id int64                                                                           `json:"user_id"`
}
type System__User__GetListResponseData_data__join_main_findManyadmin_role_user__join struct {
	Main_findManyadmin_role []System__User__GetListResponseData_data__join_main_findManyadmin_role_user__join_main_findManyadmin_role `json:"main_findManyadmin_role"`
}
type System__User__GetListResponseData_data__join_main_findManyadmin_role_user__join_main_findManyadmin_role struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__GetOneInput struct {
	Name  string `json:"name,omitempty"`
	Phone string `json:"phone,omitempty"`
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
	Roles   []string `json:"roles"`
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
	Data []System__User__GetRoleUsersResponseData_data `json:"data"`
}
type System__User__GetRoleUsersResponseData_data struct {
	Code   string                                           `json:"code"`
	Id     int64                                            `json:"id"`
	Remark string                                           `json:"remark"`
	Role   System__User__GetRoleUsersResponseData_data_role `json:"role"`
}
type System__User__GetRoleUsersResponseData_data_role struct {
	Main_findManyadmin_role_user []System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user `json:"main_findManyadmin_role_user"`
}
type System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user struct {
	Join    System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user__join `json:"_join"`
	Role_id int64                                                                               `json:"role_id"`
	User_id int64                                                                               `json:"user_id"`
}
type System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user__join struct {
	User []System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user__join_user `json:"user"`
}
type System__User__GetRoleUsersResponseData_data_role_main_findManyadmin_role_user__join_user struct {
	Avatar     string `json:"avatar,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Id         int64  `json:"id"`
	Name       string `json:"name"`
}
type System__User__GetUserByUserIdInput struct {
	UserId string `json:"userId,omitempty"`
}
type System__User__GetUserByUserIdInternalInput struct {
	UserId string `json:"userId,omitempty"`
}
type System__User__GetUserByUserIdResponseData struct {
	Data []System__User__GetUserByUserIdResponseData_data `json:"data"`
}
type System__User__GetUserByUserIdResponseData_data struct {
	Avatar  string `json:"avatar,omitempty"`
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone,omitempty"`
	User_id string `json:"user_id,omitempty"`
}
type System__User__GetUserInfoInput struct {
	Equals string `json:"equals,omitempty"`
}
type System__User__GetUserInfoInternalInput struct {
	Equals string `json:"equals,omitempty"`
}
type System__User__GetUserInfoResponseData struct {
	Data []System__User__GetUserInfoResponseData_data `json:"data"`
}
type System__User__GetUserInfoResponseData_data struct {
	Avatar       string `json:"avatar,omitempty"`
	Country_code string `json:"country_code,omitempty"`
	Created_at   string `json:"created_at,omitempty"`
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Phone        string `json:"phone,omitempty"`
	User_id      string `json:"user_id,omitempty"`
}
type System__User__GetUserRoleInput struct {
	UserId int64 `json:"userId"`
}
type System__User__GetUserRoleInternalInput struct {
	RoleId int64 `json:"roleId,omitempty"`
	UserId int64 `json:"userId"`
}
type System__User__GetUserRoleResponseData struct {
	Data []System__User__GetUserRoleResponseData_data `json:"data"`
}
type System__User__GetUserRoleResponseData_data struct {
	Code   string `json:"code"`
	Id     int64  `json:"id"`
	Remark string `json:"remark"`
}
type System__User__UpdateOneInput struct {
	Avatar   string `json:"avatar,omitempty"`
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
	Phone    string `json:"phone,omitempty"`
	UserId   string `json:"userId"`
}
type System__User__UpdateOneInternalInput struct {
	Avatar   string `json:"avatar,omitempty"`
	Id       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password"`
	Phone    string `json:"phone,omitempty"`
	UserId   string `json:"userId"`
}
type System__User__UpdateOneResponseData struct {
	Casdoor System__User__UpdateOneResponseData_casdoor `json:"casdoor,omitempty"`
	Data    System__User__UpdateOneResponseData_data    `json:"data,omitempty"`
}
type System__User__UpdateOneResponseData_casdoor struct {
	Msg    string `json:"msg,omitempty"`
	Status int64  `json:"status,omitempty"`
}
type System__User__UpdateOneResponseData_data struct {
	Avatar string `json:"avatar,omitempty"`
}
type Casdoor_addUser_post_input_object struct {
	CountryCode  string `json:"countryCode,omitempty"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	PasswordType string `json:"passwordType"`
	Phone        string `json:"phone"`
}
type Casdoor_login_post_input_object struct {
	Code        string `json:"code"`
	CountryCode string `json:"countryCode,omitempty"`
	LoginType   string `json:"loginType"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Username    string `json:"username"`
}
type Casdoor_refreshToken_post_input_object struct {
	Refresh_token string `json:"refresh_token"`
}
type Casdoor_sendCode_post_input_object struct {
	CountryCode string `json:"countryCode,omitempty"`
	Dest        string `json:"dest"`
}
type Casdoor_updateProvider_post_input_object struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	SignName     string `json:"signName"`
	TemplateCode string `json:"templateCode"`
}
type Casdoor_updateUser_post_input_object struct {
	CountryCode  string `json:"countryCode,omitempty"`
	Name         string `json:"name"`
	Password     string `json:"password"`
	PasswordType string `json:"passwordType,omitempty"`
	Phone        string `json:"phone"`
	UserId       string `json:"userId"`
}
type Main_Admin_userRelationFilter struct {
	Is    *Main_admin_userWhereInput `json:"is,omitempty"`
	IsNot *Main_admin_userWhereInput `json:"isNot,omitempty"`
}
type Main_BoolFieldUpdateOperationsInput struct {
	Set bool `json:"set,omitempty"`
}
type Main_BoolFilter struct {
	Equals bool                   `json:"equals,omitempty"`
	Not    *Main_NestedBoolFilter `json:"not,omitempty"`
}
type Main_BoolNullableFilter struct {
	Equals bool                           `json:"equals,omitempty"`
	Not    *Main_NestedBoolNullableFilter `json:"not,omitempty"`
}
type Main_Case_categoryRelationFilter struct {
	Is    *Main_case_categoryWhereInput `json:"is,omitempty"`
	IsNot *Main_case_categoryWhereInput `json:"isNot,omitempty"`
}
type Main_Case_postListRelationFilter struct {
	Every *Main_case_postWhereInput `json:"every,omitempty"`
	None  *Main_case_postWhereInput `json:"none,omitempty"`
	Some  *Main_case_postWhereInput `json:"some,omitempty"`
}
type Main_DateTimeFieldUpdateOperationsInput struct {
	Set string `json:"set,omitempty"`
}
type Main_DateTimeFilter struct {
	Equals string                     `json:"equals,omitempty"`
	Gt     string                     `json:"gt,omitempty"`
	Gte    string                     `json:"gte,omitempty"`
	In     []string                   `json:"in,omitempty"`
	Lt     string                     `json:"lt,omitempty"`
	Lte    string                     `json:"lte,omitempty"`
	Not    *Main_NestedDateTimeFilter `json:"not,omitempty"`
	NotIn  []string                   `json:"notIn,omitempty"`
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
type Main_IntFieldUpdateOperationsInput struct {
	Decrement int64 `json:"decrement,omitempty"`
	Divide    int64 `json:"divide,omitempty"`
	Increment int64 `json:"increment,omitempty"`
	Multiply  int64 `json:"multiply,omitempty"`
	Set       int64 `json:"set,omitempty"`
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
type Main_NestedBoolFilter struct {
	Equals bool                   `json:"equals,omitempty"`
	Not    *Main_NestedBoolFilter `json:"not,omitempty"`
}
type Main_NestedBoolNullableFilter struct {
	Equals bool                           `json:"equals,omitempty"`
	Not    *Main_NestedBoolNullableFilter `json:"not,omitempty"`
}
type Main_NestedDateTimeFilter struct {
	Equals string                     `json:"equals,omitempty"`
	Gt     string                     `json:"gt,omitempty"`
	Gte    string                     `json:"gte,omitempty"`
	In     []string                   `json:"in,omitempty"`
	Lt     string                     `json:"lt,omitempty"`
	Lte    string                     `json:"lte,omitempty"`
	Not    *Main_NestedDateTimeFilter `json:"not,omitempty"`
	NotIn  []string                   `json:"notIn,omitempty"`
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
type Main_NullableBoolFieldUpdateOperationsInput struct {
	Set bool `json:"set,omitempty"`
}
type Main_NullableDateTimeFieldUpdateOperationsInput struct {
	Set string `json:"set,omitempty"`
}
type Main_NullableIntFieldUpdateOperationsInput struct {
	Decrement int64 `json:"decrement,omitempty"`
	Divide    int64 `json:"divide,omitempty"`
	Increment int64 `json:"increment,omitempty"`
	Multiply  int64 `json:"multiply,omitempty"`
	Set       int64 `json:"set,omitempty"`
}
type Main_NullableStringFieldUpdateOperationsInput struct {
	Set string `json:"set,omitempty"`
}
type Main_StringFieldUpdateOperationsInput struct {
	Set string `json:"set,omitempty"`
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
type Main_admin_apilogCreateInput struct {
	Code      string `json:"code"`
	CreatedAt string `json:"createdAt,omitempty"`
	DeletedAt string `json:"deletedAt,omitempty"`
	Ip        string `json:"ip"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	Ua        string `json:"ua"`
	UpdatedAt string `json:"updatedAt"`
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
}
type Main_admin_apilogOrderByWithRelationInput struct {
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
type Main_admin_apilogWhereInput struct {
	AND       *Main_admin_apilogWhereInput   `json:"AND,omitempty"`
	NOT       *Main_admin_apilogWhereInput   `json:"NOT,omitempty"`
	OR        []*Main_admin_apilogWhereInput `json:"OR,omitempty"`
	Code      *Main_StringFilter             `json:"code,omitempty"`
	CreatedAt *Main_DateTimeFilter           `json:"createdAt,omitempty"`
	DeletedAt *Main_DateTimeNullableFilter   `json:"deletedAt,omitempty"`
	Id        *Main_IntFilter                `json:"id,omitempty"`
	Ip        *Main_StringFilter             `json:"ip,omitempty"`
	Method    *Main_StringFilter             `json:"method,omitempty"`
	Path      *Main_StringFilter             `json:"path,omitempty"`
	Ua        *Main_StringFilter             `json:"ua,omitempty"`
	UpdatedAt *Main_DateTimeFilter           `json:"updatedAt,omitempty"`
	UserId    *Main_StringFilter             `json:"userId,omitempty"`
	UserName  *Main_StringFilter             `json:"userName,omitempty"`
}
type Main_admin_apilogWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_admin_dicUpdateInput struct {
	CreatedAt *Main_DateTimeFieldUpdateOperationsInput         `json:"createdAt,omitempty"`
	DeletedAt *Main_NullableDateTimeFieldUpdateOperationsInput `json:"deletedAt,omitempty"`
	IsOpen    *Main_BoolFieldUpdateOperationsInput             `json:"isOpen,omitempty"`
	UpdatedAt *Main_DateTimeFieldUpdateOperationsInput         `json:"updatedAt,omitempty"`
}
type Main_admin_dicWhereInput struct {
	AND       *Main_admin_dicWhereInput    `json:"AND,omitempty"`
	NOT       *Main_admin_dicWhereInput    `json:"NOT,omitempty"`
	OR        []*Main_admin_dicWhereInput  `json:"OR,omitempty"`
	CreatedAt *Main_DateTimeFilter         `json:"createdAt,omitempty"`
	DeletedAt *Main_DateTimeNullableFilter `json:"deletedAt,omitempty"`
	Id        *Main_IntFilter              `json:"id,omitempty"`
	IsOpen    *Main_BoolFilter             `json:"isOpen,omitempty"`
	UpdatedAt *Main_DateTimeFilter         `json:"updatedAt,omitempty"`
}
type Main_admin_dicWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_admin_menu_roleCreateManyInput struct {
	Id      int64 `json:"id,omitempty"`
	Menu_id int64 `json:"menu_id"`
	Role_id int64 `json:"role_id"`
}
type Main_admin_menu_roleWhereInput struct {
	AND     *Main_admin_menu_roleWhereInput   `json:"AND,omitempty"`
	NOT     *Main_admin_menu_roleWhereInput   `json:"NOT,omitempty"`
	OR      []*Main_admin_menu_roleWhereInput `json:"OR,omitempty"`
	Id      *Main_IntFilter                   `json:"id,omitempty"`
	Menu_id *Main_IntFilter                   `json:"menu_id,omitempty"`
	Role_id *Main_IntFilter                   `json:"role_id,omitempty"`
}
type Main_admin_permissionCreateManyInput struct {
	CreatedAt string `json:"createdAt,omitempty"`
	Enabled   int64  `json:"enabled,omitempty"`
	Id        int64  `json:"id"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}
type Main_admin_roleCreateInput struct {
	Code   string `json:"code"`
	Remark string `json:"remark"`
}
type Main_admin_roleOrderByWithRelationInput struct {
	Code   Main_SortOrder `json:"code,omitempty"`
	Id     Main_SortOrder `json:"id,omitempty"`
	Remark Main_SortOrder `json:"remark,omitempty"`
}
type Main_admin_roleUpdateInput struct {
	Code   *Main_StringFieldUpdateOperationsInput `json:"code,omitempty"`
	Remark *Main_StringFieldUpdateOperationsInput `json:"remark,omitempty"`
}
type Main_admin_roleWhereInput struct {
	AND    *Main_admin_roleWhereInput   `json:"AND,omitempty"`
	NOT    *Main_admin_roleWhereInput   `json:"NOT,omitempty"`
	OR     []*Main_admin_roleWhereInput `json:"OR,omitempty"`
	Code   *Main_StringFilter           `json:"code,omitempty"`
	Id     *Main_IntFilter              `json:"id,omitempty"`
	Remark *Main_StringFilter           `json:"remark,omitempty"`
}
type Main_admin_roleWhereUniqueInput struct {
	Code string `json:"code,omitempty"`
	Id   int64  `json:"id,omitempty"`
}
type Main_admin_role_userCreateInput struct {
	Role_id int64 `json:"role_id"`
	User_id int64 `json:"user_id"`
}
type Main_admin_role_userRole_idUser_idCompoundUniqueInput struct {
	Role_id int64 `json:"role_id"`
	User_id int64 `json:"user_id"`
}
type Main_admin_role_userWhereInput struct {
	AND     *Main_admin_role_userWhereInput   `json:"AND,omitempty"`
	NOT     *Main_admin_role_userWhereInput   `json:"NOT,omitempty"`
	OR      []*Main_admin_role_userWhereInput `json:"OR,omitempty"`
	Id      *Main_IntFilter                   `json:"id,omitempty"`
	Role_id *Main_IntFilter                   `json:"role_id,omitempty"`
	User_id *Main_IntFilter                   `json:"user_id,omitempty"`
}
type Main_admin_role_userWhereUniqueInput struct {
	Id              int64                                                  `json:"id,omitempty"`
	Role_id_user_id *Main_admin_role_userRole_idUser_idCompoundUniqueInput `json:"role_id_user_id,omitempty"`
}
type Main_admin_tokenCreateManyInput struct {
	Banned      bool   `json:"banned,omitempty"`
	Expire_time string `json:"expire_time,omitempty"`
	Flush_time  string `json:"flush_time,omitempty"`
	Id          int64  `json:"id,omitempty"`
	Token       string `json:"token,omitempty"`
	User_id     string `json:"user_id,omitempty"`
}
type Main_admin_tokenUpdateManyMutationInput struct {
	Banned      *Main_NullableBoolFieldUpdateOperationsInput     `json:"banned,omitempty"`
	Expire_time *Main_NullableDateTimeFieldUpdateOperationsInput `json:"expire_time,omitempty"`
	Flush_time  *Main_NullableDateTimeFieldUpdateOperationsInput `json:"flush_time,omitempty"`
	Token       *Main_NullableStringFieldUpdateOperationsInput   `json:"token,omitempty"`
	User_id     *Main_NullableStringFieldUpdateOperationsInput   `json:"user_id,omitempty"`
}
type Main_admin_tokenWhereInput struct {
	AND         *Main_admin_tokenWhereInput   `json:"AND,omitempty"`
	NOT         *Main_admin_tokenWhereInput   `json:"NOT,omitempty"`
	OR          []*Main_admin_tokenWhereInput `json:"OR,omitempty"`
	Banned      *Main_BoolNullableFilter      `json:"banned,omitempty"`
	Expire_time *Main_DateTimeNullableFilter  `json:"expire_time,omitempty"`
	Flush_time  *Main_DateTimeNullableFilter  `json:"flush_time,omitempty"`
	Id          *Main_IntFilter               `json:"id,omitempty"`
	Token       *Main_StringNullableFilter    `json:"token,omitempty"`
	User_id     *Main_StringNullableFilter    `json:"user_id,omitempty"`
}
type Main_admin_tokenWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_admin_userCreateNestedOneWithoutCase_postInput struct {
	Connect         *Main_admin_userWhereUniqueInput                     `json:"connect,omitempty"`
	ConnectOrCreate *Main_admin_userCreateOrConnectWithoutCase_postInput `json:"connectOrCreate,omitempty"`
	Create          *Main_admin_userCreateWithoutCase_postInput          `json:"create,omitempty"`
}
type Main_admin_userCreateOrConnectWithoutCase_postInput struct {
	Create *Main_admin_userCreateWithoutCase_postInput `json:"create"`
	Where  *Main_admin_userWhereUniqueInput            `json:"where"`
}
type Main_admin_userCreateWithoutCase_postInput struct {
	Avatar        string `json:"avatar,omitempty"`
	Country_code  string `json:"country_code,omitempty"`
	Created_at    string `json:"created_at,omitempty"`
	Name          string `json:"name"`
	Password      string `json:"password,omitempty"`
	Password_salt string `json:"password_salt,omitempty"`
	Password_type string `json:"password_type,omitempty"`
	Phone         string `json:"phone,omitempty"`
	User_id       string `json:"user_id,omitempty"`
	Wx_resp       string `json:"wx_resp,omitempty"`
}
type Main_admin_userOrderByWithRelationInput struct {
	Avatar        Main_SortOrder                               `json:"avatar,omitempty"`
	Case_post     *Main_case_postOrderByRelationAggregateInput `json:"case_post,omitempty"`
	Country_code  Main_SortOrder                               `json:"country_code,omitempty"`
	Created_at    Main_SortOrder                               `json:"created_at,omitempty"`
	Id            Main_SortOrder                               `json:"id,omitempty"`
	Name          Main_SortOrder                               `json:"name,omitempty"`
	Password      Main_SortOrder                               `json:"password,omitempty"`
	Password_salt Main_SortOrder                               `json:"password_salt,omitempty"`
	Password_type Main_SortOrder                               `json:"password_type,omitempty"`
	Phone         Main_SortOrder                               `json:"phone,omitempty"`
	User_id       Main_SortOrder                               `json:"user_id,omitempty"`
	Wx_resp       Main_SortOrder                               `json:"wx_resp,omitempty"`
}
type Main_admin_userUpdateInput struct {
	Avatar        *Main_NullableStringFieldUpdateOperationsInput        `json:"avatar,omitempty"`
	Case_post     *Main_case_postUpdateManyWithoutAdmin_userNestedInput `json:"case_post,omitempty"`
	Country_code  *Main_NullableStringFieldUpdateOperationsInput        `json:"country_code,omitempty"`
	Created_at    *Main_NullableDateTimeFieldUpdateOperationsInput      `json:"created_at,omitempty"`
	Name          *Main_StringFieldUpdateOperationsInput                `json:"name,omitempty"`
	Password      *Main_NullableStringFieldUpdateOperationsInput        `json:"password,omitempty"`
	Password_salt *Main_NullableStringFieldUpdateOperationsInput        `json:"password_salt,omitempty"`
	Password_type *Main_NullableStringFieldUpdateOperationsInput        `json:"password_type,omitempty"`
	Phone         *Main_NullableStringFieldUpdateOperationsInput        `json:"phone,omitempty"`
	User_id       *Main_NullableStringFieldUpdateOperationsInput        `json:"user_id,omitempty"`
	Wx_resp       *Main_NullableStringFieldUpdateOperationsInput        `json:"wx_resp,omitempty"`
}
type Main_admin_userUpdateOneWithoutCase_postNestedInput struct {
	Connect         *Main_admin_userWhereUniqueInput                     `json:"connect,omitempty"`
	ConnectOrCreate *Main_admin_userCreateOrConnectWithoutCase_postInput `json:"connectOrCreate,omitempty"`
	Create          *Main_admin_userCreateWithoutCase_postInput          `json:"create,omitempty"`
	Delete          bool                                                 `json:"delete,omitempty"`
	Disconnect      bool                                                 `json:"disconnect,omitempty"`
	Update          *Main_admin_userUpdateWithoutCase_postInput          `json:"update,omitempty"`
	Upsert          *Main_admin_userUpsertWithoutCase_postInput          `json:"upsert,omitempty"`
}
type Main_admin_userUpdateWithoutCase_postInput struct {
	Avatar        *Main_NullableStringFieldUpdateOperationsInput   `json:"avatar,omitempty"`
	Country_code  *Main_NullableStringFieldUpdateOperationsInput   `json:"country_code,omitempty"`
	Created_at    *Main_NullableDateTimeFieldUpdateOperationsInput `json:"created_at,omitempty"`
	Name          *Main_StringFieldUpdateOperationsInput           `json:"name,omitempty"`
	Password      *Main_NullableStringFieldUpdateOperationsInput   `json:"password,omitempty"`
	Password_salt *Main_NullableStringFieldUpdateOperationsInput   `json:"password_salt,omitempty"`
	Password_type *Main_NullableStringFieldUpdateOperationsInput   `json:"password_type,omitempty"`
	Phone         *Main_NullableStringFieldUpdateOperationsInput   `json:"phone,omitempty"`
	User_id       *Main_NullableStringFieldUpdateOperationsInput   `json:"user_id,omitempty"`
	Wx_resp       *Main_NullableStringFieldUpdateOperationsInput   `json:"wx_resp,omitempty"`
}
type Main_admin_userUpsertWithoutCase_postInput struct {
	Create *Main_admin_userCreateWithoutCase_postInput `json:"create"`
	Update *Main_admin_userUpdateWithoutCase_postInput `json:"update"`
}
type Main_admin_userWhereInput struct {
	AND           *Main_admin_userWhereInput        `json:"AND,omitempty"`
	NOT           *Main_admin_userWhereInput        `json:"NOT,omitempty"`
	OR            []*Main_admin_userWhereInput      `json:"OR,omitempty"`
	Avatar        *Main_StringNullableFilter        `json:"avatar,omitempty"`
	Case_post     *Main_Case_postListRelationFilter `json:"case_post,omitempty"`
	Country_code  *Main_StringNullableFilter        `json:"country_code,omitempty"`
	Created_at    *Main_DateTimeNullableFilter      `json:"created_at,omitempty"`
	Id            *Main_IntFilter                   `json:"id,omitempty"`
	Name          *Main_StringFilter                `json:"name,omitempty"`
	Password      *Main_StringNullableFilter        `json:"password,omitempty"`
	Password_salt *Main_StringNullableFilter        `json:"password_salt,omitempty"`
	Password_type *Main_StringNullableFilter        `json:"password_type,omitempty"`
	Phone         *Main_StringNullableFilter        `json:"phone,omitempty"`
	User_id       *Main_StringNullableFilter        `json:"user_id,omitempty"`
	Wx_resp       *Main_StringNullableFilter        `json:"wx_resp,omitempty"`
}
type Main_admin_userWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_case_categoryCreateInput struct {
	Case_post *Main_case_postCreateNestedManyWithoutCase_categoryInput `json:"case_post,omitempty"`
	Name      string                                                   `json:"name"`
}
type Main_case_categoryCreateNestedOneWithoutCase_postInput struct {
	Connect         *Main_case_categoryWhereUniqueInput                     `json:"connect,omitempty"`
	ConnectOrCreate *Main_case_categoryCreateOrConnectWithoutCase_postInput `json:"connectOrCreate,omitempty"`
	Create          *Main_case_categoryCreateWithoutCase_postInput          `json:"create,omitempty"`
}
type Main_case_categoryCreateOrConnectWithoutCase_postInput struct {
	Create *Main_case_categoryCreateWithoutCase_postInput `json:"create"`
	Where  *Main_case_categoryWhereUniqueInput            `json:"where"`
}
type Main_case_categoryCreateWithoutCase_postInput struct {
	Name string `json:"name"`
}
type Main_case_categoryOrderByWithRelationInput struct {
	Case_post *Main_case_postOrderByRelationAggregateInput `json:"case_post,omitempty"`
	Id        Main_SortOrder                               `json:"id,omitempty"`
	Name      Main_SortOrder                               `json:"name,omitempty"`
}
type Main_case_categoryUpdateInput struct {
	Case_post *Main_case_postUpdateManyWithoutCase_categoryNestedInput `json:"case_post,omitempty"`
	Name      *Main_StringFieldUpdateOperationsInput                   `json:"name,omitempty"`
}
type Main_case_categoryUpdateOneWithoutCase_postNestedInput struct {
	Connect         *Main_case_categoryWhereUniqueInput                     `json:"connect,omitempty"`
	ConnectOrCreate *Main_case_categoryCreateOrConnectWithoutCase_postInput `json:"connectOrCreate,omitempty"`
	Create          *Main_case_categoryCreateWithoutCase_postInput          `json:"create,omitempty"`
	Delete          bool                                                    `json:"delete,omitempty"`
	Disconnect      bool                                                    `json:"disconnect,omitempty"`
	Update          *Main_case_categoryUpdateWithoutCase_postInput          `json:"update,omitempty"`
	Upsert          *Main_case_categoryUpsertWithoutCase_postInput          `json:"upsert,omitempty"`
}
type Main_case_categoryUpdateWithoutCase_postInput struct {
	Name *Main_StringFieldUpdateOperationsInput `json:"name,omitempty"`
}
type Main_case_categoryUpsertWithoutCase_postInput struct {
	Create *Main_case_categoryCreateWithoutCase_postInput `json:"create"`
	Update *Main_case_categoryUpdateWithoutCase_postInput `json:"update"`
}
type Main_case_categoryWhereInput struct {
	AND       *Main_case_categoryWhereInput     `json:"AND,omitempty"`
	NOT       *Main_case_categoryWhereInput     `json:"NOT,omitempty"`
	OR        []*Main_case_categoryWhereInput   `json:"OR,omitempty"`
	Case_post *Main_Case_postListRelationFilter `json:"case_post,omitempty"`
	Id        *Main_IntFilter                   `json:"id,omitempty"`
	Name      *Main_StringFilter                `json:"name,omitempty"`
}
type Main_case_categoryWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_case_menuCreateInput struct {
	Api_id      string `json:"api_id,omitempty"`
	Create_time string `json:"create_time,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Is_bottom   int64  `json:"is_bottom,omitempty"`
	Label       string `json:"label"`
	Level       int64  `json:"level"`
	Menu_type   string `json:"menu_type,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	Path        string `json:"path,omitempty"`
	Perms       string `json:"perms,omitempty"`
	Sort        int64  `json:"sort"`
}
type Main_case_menuOrderByWithRelationInput struct {
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
type Main_case_menuUpdateInput struct {
	Api_id      *Main_NullableStringFieldUpdateOperationsInput   `json:"api_id,omitempty"`
	Create_time *Main_NullableDateTimeFieldUpdateOperationsInput `json:"create_time,omitempty"`
	Icon        *Main_NullableStringFieldUpdateOperationsInput   `json:"icon,omitempty"`
	Is_bottom   *Main_IntFieldUpdateOperationsInput              `json:"is_bottom,omitempty"`
	Label       *Main_StringFieldUpdateOperationsInput           `json:"label,omitempty"`
	Level       *Main_IntFieldUpdateOperationsInput              `json:"level,omitempty"`
	Menu_type   *Main_NullableStringFieldUpdateOperationsInput   `json:"menu_type,omitempty"`
	Name        *Main_NullableStringFieldUpdateOperationsInput   `json:"name,omitempty"`
	ParentId    *Main_NullableIntFieldUpdateOperationsInput      `json:"parentId,omitempty"`
	Path        *Main_NullableStringFieldUpdateOperationsInput   `json:"path,omitempty"`
	Perms       *Main_NullableStringFieldUpdateOperationsInput   `json:"perms,omitempty"`
	Sort        *Main_IntFieldUpdateOperationsInput              `json:"sort,omitempty"`
}
type Main_case_menuWhereInput struct {
	AND         *Main_case_menuWhereInput    `json:"AND,omitempty"`
	NOT         *Main_case_menuWhereInput    `json:"NOT,omitempty"`
	OR          []*Main_case_menuWhereInput  `json:"OR,omitempty"`
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
type Main_case_menuWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_case_postCreateInput struct {
	Admin_user    *Main_admin_userCreateNestedOneWithoutCase_postInput    `json:"admin_user,omitempty"`
	Author        string                                                  `json:"author"`
	Case_category *Main_case_categoryCreateNestedOneWithoutCase_postInput `json:"case_category,omitempty"`
	Content       string                                                  `json:"content,omitempty"`
	Poster        string                                                  `json:"poster,omitempty"`
	Published_at  string                                                  `json:"published_at,omitempty"`
	Title         string                                                  `json:"title"`
}
type Main_case_postCreateManyAdmin_userInput struct {
	Author       string `json:"author"`
	Cate         int64  `json:"cate,omitempty"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title"`
}
type Main_case_postCreateManyAdmin_userInputEnvelope struct {
	Data           []*Main_case_postCreateManyAdmin_userInput `json:"data"`
	SkipDuplicates bool                                       `json:"skipDuplicates,omitempty"`
}
type Main_case_postCreateManyCase_categoryInput struct {
	Auth         int64  `json:"auth,omitempty"`
	Author       string `json:"author"`
	Content      string `json:"content,omitempty"`
	Id           int64  `json:"id,omitempty"`
	Poster       string `json:"poster,omitempty"`
	Published_at string `json:"published_at,omitempty"`
	Title        string `json:"title"`
}
type Main_case_postCreateManyCase_categoryInputEnvelope struct {
	Data           []*Main_case_postCreateManyCase_categoryInput `json:"data"`
	SkipDuplicates bool                                          `json:"skipDuplicates,omitempty"`
}
type Main_case_postCreateNestedManyWithoutCase_categoryInput struct {
	Connect         *Main_case_postWhereUniqueInput                         `json:"connect,omitempty"`
	ConnectOrCreate *Main_case_postCreateOrConnectWithoutCase_categoryInput `json:"connectOrCreate,omitempty"`
	Create          *Main_case_postCreateWithoutCase_categoryInput          `json:"create,omitempty"`
	CreateMany      *Main_case_postCreateManyCase_categoryInputEnvelope     `json:"createMany,omitempty"`
}
type Main_case_postCreateOrConnectWithoutAdmin_userInput struct {
	Create *Main_case_postCreateWithoutAdmin_userInput `json:"create"`
	Where  *Main_case_postWhereUniqueInput             `json:"where"`
}
type Main_case_postCreateOrConnectWithoutCase_categoryInput struct {
	Create *Main_case_postCreateWithoutCase_categoryInput `json:"create"`
	Where  *Main_case_postWhereUniqueInput                `json:"where"`
}
type Main_case_postCreateWithoutAdmin_userInput struct {
	Author        string                                                  `json:"author"`
	Case_category *Main_case_categoryCreateNestedOneWithoutCase_postInput `json:"case_category,omitempty"`
	Content       string                                                  `json:"content,omitempty"`
	Poster        string                                                  `json:"poster,omitempty"`
	Published_at  string                                                  `json:"published_at,omitempty"`
	Title         string                                                  `json:"title"`
}
type Main_case_postCreateWithoutCase_categoryInput struct {
	Admin_user   *Main_admin_userCreateNestedOneWithoutCase_postInput `json:"admin_user,omitempty"`
	Author       string                                               `json:"author"`
	Content      string                                               `json:"content,omitempty"`
	Poster       string                                               `json:"poster,omitempty"`
	Published_at string                                               `json:"published_at,omitempty"`
	Title        string                                               `json:"title"`
}
type Main_case_postOrderByRelationAggregateInput struct {
	Count Main_SortOrder `json:"_count,omitempty"`
}
type Main_case_postOrderByWithRelationInput struct {
	Admin_user    *Main_admin_userOrderByWithRelationInput    `json:"admin_user,omitempty"`
	Auth          Main_SortOrder                              `json:"auth,omitempty"`
	Author        Main_SortOrder                              `json:"author,omitempty"`
	Case_category *Main_case_categoryOrderByWithRelationInput `json:"case_category,omitempty"`
	Cate          Main_SortOrder                              `json:"cate,omitempty"`
	Content       Main_SortOrder                              `json:"content,omitempty"`
	Id            Main_SortOrder                              `json:"id,omitempty"`
	Poster        Main_SortOrder                              `json:"poster,omitempty"`
	Published_at  Main_SortOrder                              `json:"published_at,omitempty"`
	Title         Main_SortOrder                              `json:"title,omitempty"`
}
type Main_case_postScalarWhereInput struct {
	AND          *Main_case_postScalarWhereInput   `json:"AND,omitempty"`
	NOT          *Main_case_postScalarWhereInput   `json:"NOT,omitempty"`
	OR           []*Main_case_postScalarWhereInput `json:"OR,omitempty"`
	Auth         *Main_IntNullableFilter           `json:"auth,omitempty"`
	Author       *Main_StringFilter                `json:"author,omitempty"`
	Cate         *Main_IntNullableFilter           `json:"cate,omitempty"`
	Content      *Main_StringNullableFilter        `json:"content,omitempty"`
	Id           *Main_IntFilter                   `json:"id,omitempty"`
	Poster       *Main_StringNullableFilter        `json:"poster,omitempty"`
	Published_at *Main_DateTimeNullableFilter      `json:"published_at,omitempty"`
	Title        *Main_StringFilter                `json:"title,omitempty"`
}
type Main_case_postUpdateInput struct {
	Admin_user    *Main_admin_userUpdateOneWithoutCase_postNestedInput    `json:"admin_user,omitempty"`
	Author        *Main_StringFieldUpdateOperationsInput                  `json:"author,omitempty"`
	Case_category *Main_case_categoryUpdateOneWithoutCase_postNestedInput `json:"case_category,omitempty"`
	Content       *Main_NullableStringFieldUpdateOperationsInput          `json:"content,omitempty"`
	Poster        *Main_NullableStringFieldUpdateOperationsInput          `json:"poster,omitempty"`
	Published_at  *Main_NullableDateTimeFieldUpdateOperationsInput        `json:"published_at,omitempty"`
	Title         *Main_StringFieldUpdateOperationsInput                  `json:"title,omitempty"`
}
type Main_case_postUpdateManyMutationInput struct {
	Author       *Main_StringFieldUpdateOperationsInput           `json:"author,omitempty"`
	Content      *Main_NullableStringFieldUpdateOperationsInput   `json:"content,omitempty"`
	Poster       *Main_NullableStringFieldUpdateOperationsInput   `json:"poster,omitempty"`
	Published_at *Main_NullableDateTimeFieldUpdateOperationsInput `json:"published_at,omitempty"`
	Title        *Main_StringFieldUpdateOperationsInput           `json:"title,omitempty"`
}
type Main_case_postUpdateManyWithWhereWithoutAdmin_userInput struct {
	Data  *Main_case_postUpdateManyMutationInput `json:"data"`
	Where *Main_case_postScalarWhereInput        `json:"where"`
}
type Main_case_postUpdateManyWithWhereWithoutCase_categoryInput struct {
	Data  *Main_case_postUpdateManyMutationInput `json:"data"`
	Where *Main_case_postScalarWhereInput        `json:"where"`
}
type Main_case_postUpdateManyWithoutAdmin_userNestedInput struct {
	Connect         *Main_case_postWhereUniqueInput                            `json:"connect,omitempty"`
	ConnectOrCreate *Main_case_postCreateOrConnectWithoutAdmin_userInput       `json:"connectOrCreate,omitempty"`
	Create          *Main_case_postCreateWithoutAdmin_userInput                `json:"create,omitempty"`
	CreateMany      *Main_case_postCreateManyAdmin_userInputEnvelope           `json:"createMany,omitempty"`
	Delete          *Main_case_postWhereUniqueInput                            `json:"delete,omitempty"`
	DeleteMany      *Main_case_postScalarWhereInput                            `json:"deleteMany,omitempty"`
	Disconnect      *Main_case_postWhereUniqueInput                            `json:"disconnect,omitempty"`
	Set             *Main_case_postWhereUniqueInput                            `json:"set,omitempty"`
	Update          *Main_case_postUpdateWithWhereUniqueWithoutAdmin_userInput `json:"update,omitempty"`
	UpdateMany      *Main_case_postUpdateManyWithWhereWithoutAdmin_userInput   `json:"updateMany,omitempty"`
	Upsert          *Main_case_postUpsertWithWhereUniqueWithoutAdmin_userInput `json:"upsert,omitempty"`
}
type Main_case_postUpdateManyWithoutCase_categoryNestedInput struct {
	Connect         *Main_case_postWhereUniqueInput                               `json:"connect,omitempty"`
	ConnectOrCreate *Main_case_postCreateOrConnectWithoutCase_categoryInput       `json:"connectOrCreate,omitempty"`
	Create          *Main_case_postCreateWithoutCase_categoryInput                `json:"create,omitempty"`
	CreateMany      *Main_case_postCreateManyCase_categoryInputEnvelope           `json:"createMany,omitempty"`
	Delete          *Main_case_postWhereUniqueInput                               `json:"delete,omitempty"`
	DeleteMany      *Main_case_postScalarWhereInput                               `json:"deleteMany,omitempty"`
	Disconnect      *Main_case_postWhereUniqueInput                               `json:"disconnect,omitempty"`
	Set             *Main_case_postWhereUniqueInput                               `json:"set,omitempty"`
	Update          *Main_case_postUpdateWithWhereUniqueWithoutCase_categoryInput `json:"update,omitempty"`
	UpdateMany      *Main_case_postUpdateManyWithWhereWithoutCase_categoryInput   `json:"updateMany,omitempty"`
	Upsert          *Main_case_postUpsertWithWhereUniqueWithoutCase_categoryInput `json:"upsert,omitempty"`
}
type Main_case_postUpdateWithWhereUniqueWithoutAdmin_userInput struct {
	Data  *Main_case_postUpdateWithoutAdmin_userInput `json:"data"`
	Where *Main_case_postWhereUniqueInput             `json:"where"`
}
type Main_case_postUpdateWithWhereUniqueWithoutCase_categoryInput struct {
	Data  *Main_case_postUpdateWithoutCase_categoryInput `json:"data"`
	Where *Main_case_postWhereUniqueInput                `json:"where"`
}
type Main_case_postUpdateWithoutAdmin_userInput struct {
	Author        *Main_StringFieldUpdateOperationsInput                  `json:"author,omitempty"`
	Case_category *Main_case_categoryUpdateOneWithoutCase_postNestedInput `json:"case_category,omitempty"`
	Content       *Main_NullableStringFieldUpdateOperationsInput          `json:"content,omitempty"`
	Poster        *Main_NullableStringFieldUpdateOperationsInput          `json:"poster,omitempty"`
	Published_at  *Main_NullableDateTimeFieldUpdateOperationsInput        `json:"published_at,omitempty"`
	Title         *Main_StringFieldUpdateOperationsInput                  `json:"title,omitempty"`
}
type Main_case_postUpdateWithoutCase_categoryInput struct {
	Admin_user   *Main_admin_userUpdateOneWithoutCase_postNestedInput `json:"admin_user,omitempty"`
	Author       *Main_StringFieldUpdateOperationsInput               `json:"author,omitempty"`
	Content      *Main_NullableStringFieldUpdateOperationsInput       `json:"content,omitempty"`
	Poster       *Main_NullableStringFieldUpdateOperationsInput       `json:"poster,omitempty"`
	Published_at *Main_NullableDateTimeFieldUpdateOperationsInput     `json:"published_at,omitempty"`
	Title        *Main_StringFieldUpdateOperationsInput               `json:"title,omitempty"`
}
type Main_case_postUpsertWithWhereUniqueWithoutAdmin_userInput struct {
	Create *Main_case_postCreateWithoutAdmin_userInput `json:"create"`
	Update *Main_case_postUpdateWithoutAdmin_userInput `json:"update"`
	Where  *Main_case_postWhereUniqueInput             `json:"where"`
}
type Main_case_postUpsertWithWhereUniqueWithoutCase_categoryInput struct {
	Create *Main_case_postCreateWithoutCase_categoryInput `json:"create"`
	Update *Main_case_postUpdateWithoutCase_categoryInput `json:"update"`
	Where  *Main_case_postWhereUniqueInput                `json:"where"`
}
type Main_case_postWhereInput struct {
	AND           *Main_case_postWhereInput         `json:"AND,omitempty"`
	NOT           *Main_case_postWhereInput         `json:"NOT,omitempty"`
	OR            []*Main_case_postWhereInput       `json:"OR,omitempty"`
	Admin_user    *Main_Admin_userRelationFilter    `json:"admin_user,omitempty"`
	Auth          *Main_IntNullableFilter           `json:"auth,omitempty"`
	Author        *Main_StringFilter                `json:"author,omitempty"`
	Case_category *Main_Case_categoryRelationFilter `json:"case_category,omitempty"`
	Cate          *Main_IntNullableFilter           `json:"cate,omitempty"`
	Content       *Main_StringNullableFilter        `json:"content,omitempty"`
	Id            *Main_IntFilter                   `json:"id,omitempty"`
	Poster        *Main_StringNullableFilter        `json:"poster,omitempty"`
	Published_at  *Main_DateTimeNullableFilter      `json:"published_at,omitempty"`
	Title         *Main_StringFilter                `json:"title,omitempty"`
}
type Main_case_postWhereUniqueInput struct {
	Id int64 `json:"id,omitempty"`
}
type Main_case_subCreateInput struct {
	Create_role string `json:"create_role"`
	CreatedAt   string `json:"createdAt,omitempty"`
	DeletedAt   string `json:"deletedAt,omitempty"`
	Message     string `json:"message"`
	Target_role string `json:"target_role"`
	Type        string `json:"type"`
	UpdatedAt   string `json:"updatedAt"`
}
type Main_case_subWhereInput struct {
	AND         *Main_case_subWhereInput     `json:"AND,omitempty"`
	NOT         *Main_case_subWhereInput     `json:"NOT,omitempty"`
	OR          []*Main_case_subWhereInput   `json:"OR,omitempty"`
	Create_role *Main_StringFilter           `json:"create_role,omitempty"`
	CreatedAt   *Main_DateTimeFilter         `json:"createdAt,omitempty"`
	DeletedAt   *Main_DateTimeNullableFilter `json:"deletedAt,omitempty"`
	Id          *Main_IntFilter              `json:"id,omitempty"`
	Message     *Main_StringFilter           `json:"message,omitempty"`
	Target_role *Main_StringFilter           `json:"target_role,omitempty"`
	Type        *Main_StringFilter           `json:"type,omitempty"`
	UpdatedAt   *Main_DateTimeFilter         `json:"updatedAt,omitempty"`
}
type System_Role_input_object struct {
	Code       string `json:"code,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	DeleteTime string `json:"deleteTime,omitempty"`
	Remark     string `json:"remark,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}
type System_bindRoleApis_post_input_object struct {
	OperationPaths []string                               `json:"operationPaths"`
	RbacType       System_bindRoleApis_post_rbacType_enum `json:"rbacType,omitempty"`
	RoleCodes      []string                               `json:"roleCodes"`
}

type Main_SortOrder string

const (
	Main_SortOrder_asc  Main_SortOrder = "asc"
	Main_SortOrder_desc Main_SortOrder = "desc"
)

type System_bindRoleApis_post_rbacType_enum string

const (
	System_bindRoleApis_post_rbacType_enum_denyMatchAll    System_bindRoleApis_post_rbacType_enum = "denyMatchAll"
	System_bindRoleApis_post_rbacType_enum_denyMatchAny    System_bindRoleApis_post_rbacType_enum = "denyMatchAny"
	System_bindRoleApis_post_rbacType_enum_requireMatchAll System_bindRoleApis_post_rbacType_enum = "requireMatchAll"
	System_bindRoleApis_post_rbacType_enum_requireMatchAny System_bindRoleApis_post_rbacType_enum = "requireMatchAny"
)
