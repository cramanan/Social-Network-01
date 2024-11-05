import { useState } from "react";

type QueryParams = {
    limit: number;
    offset: number;
};

export default function useQueryParams(limit?: number) {
    const [params, setParams] = useState<QueryParams>({
        limit: limit ?? 20,
        offset: 0,
    });

    const next = () =>
        setParams({ ...params, offset: params.offset + params.limit });

    const previous = () => {
        if (params.offset - params.limit >= 0)
            setParams({ ...params, offset: params.offset - params.limit });
    };

    const nextWithFunc = (fn: () => boolean) => {
        if (fn()) next();
    };

    return {
        limit: params.limit,
        offset: params.offset,
        setParams,
        next,
        previous,
        nextWithFunc,
    };
}
