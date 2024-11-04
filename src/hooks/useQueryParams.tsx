import { useState } from "react";

type QueryParams = {
    limit: number;
    offset: number;
};

export default function useQueryParams(defaultValue?: QueryParams) {
    const [params, setParams] = useState<QueryParams>(
        defaultValue ?? {
            limit: 20,
            offset: 0,
        }
    );

    const next = () =>
        setParams({ ...params, offset: params.offset + params.limit });

    const previous = () => {
        if (params.offset - params.limit >= 0)
            setParams({ ...params, offset: params.offset - params.limit });
    };

    return {
        limit: params.limit,
        offset: params.offset,
        setParams,
        next,
        previous,
    };
}
