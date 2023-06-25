import { createClient } from "./client";

export default createClient({
  baseURL: ""
});

function isNotEmpty(arg: any) {
  return arg !== undefined && arg !== null && arg !== "";
}

export type ConvertOptions<Fields> = {
  containsFields?: Fields[];
  equalsFields?: Fields[];
};

export function convertPageQuery(
  query: PageQuery,
  options?: ConvertOptions<keyof typeof query>
): FireboomQuery {
  const { pageNum, pageSize, ...rest } = query;
  const ret: FireboomQuery = {
    skip: (pageNum - 1) * pageSize,
    take: pageSize || 10,
    query: {}
  };
  if (Object.keys(rest).length) {
    if (options?.containsFields) {
      for (const field of options.containsFields) {
        if (isNotEmpty(query[field])) {
          (ret.query as Record<string, FireboomQueryValue>)[field] = {
            contains: query[field]
          };
        }
      }
    }
    if (options?.equalsFields) {
      for (const field of options.equalsFields) {
        if (isNotEmpty(query[field])) {
          (ret.query as Record<string, FireboomQueryValue>)[field] = {
            equals: query[field]
          };
        }
      }
    }
  }
  if (!Object.keys((ret as Record<string, any>)["query"])) {
    delete ret["query"];
  }
  return ret;
}
