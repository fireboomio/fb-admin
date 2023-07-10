import { http } from "@/utils/http";

// 新增文章
export const createOne = (
  title: string,
  username: string,
  content?: string,
  poster?: string,
  publishedAt?: string
) => {
  return http.request("post", `/operations/Post/CreateOne`, {
    data: {
      username,
      content,
      poster,
      publishedAt,
      title
    }
  });
};
