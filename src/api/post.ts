import axios from "@/utils/http";

// 新增文章
export const createOne = (
  title: string,
  username: string,
  content?: string,
  poster?: string,
  publishedAt?: string
) => {
  return axios.post<any>(`/operations/Post/CreateOne`, {
    data: {
      username,
      content,
      poster,
      publishedAt,
      title
    }
  });
};
