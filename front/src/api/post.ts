import axios from "@/utils/http";

// 新增文章
export const createOne = (
  title: string,
  username: string,
  content?: string,
  poster?: string,
  publishedAt?: string,
  userId?: number,
  cateId?: number,
) => {
  return axios.post<any>(`/operations/Post/CreateOne`, {
    data: {
      username,
      content,
      poster,
      publishedAt,
      title,
      userId,
      cateId,
    }
  });
};

/**
 * 文章的模糊查询
 */
export const getPostLike = (data: object) => {
  return axios.get<any>("/operations/Post/GetLikeList", {
    params: data,
  });
}
/**
 * 按照类别搜索文章
 */
export const getPostByCategory = (data: string) => {
  return axios.get<any>("/operations/Post/GetPostByCate", {
    params: {
      equals: data,
    }
  })
}

/**
 * 获取所有文章类别以及Id
 */
export const getCategory = () => {
  return axios.get<any>("/operations/Post/GetCategory");
}

/**
 * 文章更新
 */
export const updatePost = (data: object) => {
  return axios.post("/operations/Post/UpdateOne", {
    data
  });
}